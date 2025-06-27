/*
 * @Description: WebSocket服务器增强功能 - 提供连接池、消息队列、监控指标等高级功能
 * @Author: redxing96@163.com
 * @Date: 2025-06-27 11:05:00
 * @LastEditTime: 2025-06-27 11:11:38
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/utility/websocket/websocket_enhanced.go
 */
package websocket

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// 增强功能常量
const (
	// 消息队列大小
	messageQueueSize = 10000

	// 连接池大小
	connectionPoolSize = 1000

	// 消息处理超时时间
	messageProcessTimeout = 5 * time.Second

	// 批量发送大小
	batchSendSize = 100

	// 监控指标更新间隔
	metricsUpdateInterval = 10 * time.Second
)

// EnhancedWsServer 增强版WebSocket服务器
type EnhancedWsServer struct {
	*WsServer                      // 继承基础WebSocket服务器
	messageQueue   *gqueue.Queue   // 消息队列
	connectionPool *ConnectionPool // 连接池
	metrics        *Metrics        // 监控指标
	rateLimiter    *RateLimiter    // 速率限制器
	messageHandler MessageHandler  // 消息处理器
	ctx            context.Context
	cancel         context.CancelFunc
}

// ConnectionPool 连接池
type ConnectionPool struct {
	pool    *gmap.StrAnyMap
	maxSize int
	mu      sync.RWMutex
}

// Metrics 监控指标
type Metrics struct {
	// 连接相关指标
	ConnectionAttempts int64 // 连接尝试次数
	ConnectionSuccess  int64 // 连接成功次数
	ConnectionFailures int64 // 连接失败次数
	ConnectionDrops    int64 // 连接断开次数

	// 消息相关指标
	MessagesReceived int64         // 接收消息数
	MessagesSent     int64         // 发送消息数
	MessagesDropped  int64         // 丢弃消息数
	MessageLatency   time.Duration // 消息延迟

	// 性能相关指标
	AverageResponseTime time.Duration // 平均响应时间
	PeakConnections     int64         // 峰值连接数
	CurrentLoad         float64       // 当前负载

	// 错误相关指标
	ErrorCount    int64     // 错误计数
	LastErrorTime time.Time // 最后错误时间
	LastError     error     // 最后错误

	mu sync.RWMutex
}

// RateLimiter 速率限制器
type RateLimiter struct {
	limits       *gmap.StrAnyMap // clientId -> *TokenBucket
	defaultLimit int64           // 默认限制
	mu           sync.RWMutex
}

// TokenBucket 令牌桶算法实现
type TokenBucket struct {
	tokens     int64     // 当前令牌数
	capacity   int64     // 桶容量
	rate       int64     // 令牌产生速率
	lastRefill time.Time // 上次补充时间
	mu         sync.Mutex
}

// MessageHandler 消息处理器接口
type MessageHandler interface {
	HandleMessage(ctx context.Context, clientId string, message []byte) error
	HandleError(ctx context.Context, clientId string, err error)
}

// DefaultMessageHandler 默认消息处理器
type DefaultMessageHandler struct {
	server *EnhancedWsServer
}

// NewEnhancedWsServer 创建增强版WebSocket服务器
func NewEnhancedWsServer() *EnhancedWsServer {
	ctx, cancel := context.WithCancel(context.Background())

	enhanced := &EnhancedWsServer{
		WsServer:       NewWsServer(),
		messageQueue:   gqueue.New(messageQueueSize),
		connectionPool: NewConnectionPool(connectionPoolSize),
		metrics:        NewMetrics(),
		rateLimiter:    NewRateLimiter(100), // 默认每秒100个请求
		messageHandler: &DefaultMessageHandler{},
		ctx:            ctx,
		cancel:         cancel,
	}

	// 设置消息处理器
	enhanced.messageHandler = &DefaultMessageHandler{server: enhanced}

	// 启动后台协程
	go enhanced.messageProcessor()
	go enhanced.metricsCollector()

	return enhanced
}

// NewConnectionPool 创建连接池
func NewConnectionPool(maxSize int) *ConnectionPool {
	return &ConnectionPool{
		pool:    gmap.NewStrAnyMap(true),
		maxSize: maxSize,
	}
}

// NewMetrics 创建监控指标
func NewMetrics() *Metrics {
	return &Metrics{
		LastErrorTime: time.Now(),
	}
}

// NewRateLimiter 创建速率限制器
func NewRateLimiter(defaultLimit int64) *RateLimiter {
	return &RateLimiter{
		limits:       gmap.NewStrAnyMap(true),
		defaultLimit: defaultLimit,
	}
}

// NewTokenBucket 创建令牌桶
func NewTokenBucket(capacity, rate int64) *TokenBucket {
	return &TokenBucket{
		tokens:     capacity,
		capacity:   capacity,
		rate:       rate,
		lastRefill: time.Now(),
	}
}

// HandleWsConnection 增强版连接处理
func (e *EnhancedWsServer) HandleWsConnection(w http.ResponseWriter, r *http.Request) {
	// 更新连接尝试指标
	atomic.AddInt64(&e.metrics.ConnectionAttempts, 1)

	// 检查连接池是否已满
	if e.connectionPool.IsFull() {
		atomic.AddInt64(&e.metrics.ConnectionFailures, 1)
		http.Error(w, "服务器连接数已达上限", http.StatusServiceUnavailable)
		return
	}

	// 调用基础连接处理
	e.WsServer.HandleWsConnection(w, r)

	// 更新连接成功指标
	atomic.AddInt64(&e.metrics.ConnectionSuccess, 1)

	// 更新峰值连接数
	currentConnections := atomic.LoadInt64(&e.WsServer.stats.ActiveConnections)
	if currentConnections > atomic.LoadInt64(&e.metrics.PeakConnections) {
		atomic.StoreInt64(&e.metrics.PeakConnections, currentConnections)
	}
}

// SendMessageWithRetry 带重试的消息发送
func (e *EnhancedWsServer) SendMessageWithRetry(ctx context.Context, clientId string, message []byte, maxRetries int) error {
	for i := 0; i <= maxRetries; i++ {
		if e.WsServer.SendToClient(ctx, clientId, message) {
			atomic.AddInt64(&e.metrics.MessagesSent, 1)
			return nil
		}

		if i < maxRetries {
			time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
		}
	}

	atomic.AddInt64(&e.metrics.MessagesDropped, 1)
	return fmt.Errorf("消息发送失败，已重试 %d 次", maxRetries)
}

// SendBatchMessage 批量发送消息
func (e *EnhancedWsServer) SendBatchMessage(ctx context.Context, clientIds []string, message []byte) (int, error) {
	if len(clientIds) == 0 {
		return 0, nil
	}

	successCount := 0
	errorCount := 0

	// 分批处理
	for i := 0; i < len(clientIds); i += batchSendSize {
		end := i + batchSendSize
		if end > len(clientIds) {
			end = len(clientIds)
		}

		batch := clientIds[i:end]

		var wg sync.WaitGroup
		results := make(chan bool, len(batch))

		for _, clientId := range batch {
			wg.Add(1)
			go func(id string) {
				defer wg.Done()
				success := e.WsServer.SendToClient(ctx, id, message)
				results <- success
			}(clientId)
		}

		wg.Wait()
		close(results)

		for success := range results {
			if success {
				successCount++
			} else {
				errorCount++
			}
		}
	}

	atomic.AddInt64(&e.metrics.MessagesSent, int64(successCount))
	atomic.AddInt64(&e.metrics.MessagesDropped, int64(errorCount))

	return successCount, nil
}

// RateLimitCheck 速率限制检查
func (e *EnhancedWsServer) RateLimitCheck(clientId string) bool {
	return e.rateLimiter.Allow(clientId)
}

// SetRateLimit 设置客户端速率限制
func (e *EnhancedWsServer) SetRateLimit(clientId string, limit int64) {
	e.rateLimiter.SetLimit(clientId, limit)
}

// GetMetrics 获取监控指标
func (e *EnhancedWsServer) GetMetrics() *Metrics {
	e.metrics.mu.RLock()
	defer e.metrics.mu.RUnlock()

	metrics := &Metrics{
		ConnectionAttempts:  atomic.LoadInt64(&e.metrics.ConnectionAttempts),
		ConnectionSuccess:   atomic.LoadInt64(&e.metrics.ConnectionSuccess),
		ConnectionFailures:  atomic.LoadInt64(&e.metrics.ConnectionFailures),
		ConnectionDrops:     atomic.LoadInt64(&e.metrics.ConnectionDrops),
		MessagesReceived:    atomic.LoadInt64(&e.metrics.MessagesReceived),
		MessagesSent:        atomic.LoadInt64(&e.metrics.MessagesSent),
		MessagesDropped:     atomic.LoadInt64(&e.metrics.MessagesDropped),
		MessageLatency:      e.metrics.MessageLatency,
		AverageResponseTime: e.metrics.AverageResponseTime,
		PeakConnections:     atomic.LoadInt64(&e.metrics.PeakConnections),
		CurrentLoad:         e.metrics.CurrentLoad,
		ErrorCount:          atomic.LoadInt64(&e.metrics.ErrorCount),
		LastErrorTime:       e.metrics.LastErrorTime,
		LastError:           e.metrics.LastError,
	}

	return metrics
}

// messageProcessor 消息处理器
func (e *EnhancedWsServer) messageProcessor() {
	defer func() {
		if err := recover(); err != nil {
			g.Log().Errorf(e.ctx, "消息处理器panic: %v", err)
		}
	}()

	for {
		select {
		case <-e.ctx.Done():
			return
		default:
		}

		// 从队列中获取消息
		if item := e.messageQueue.Pop(); item != nil {
			msg := item.(*QueuedMessage)

			// 检查速率限制
			if !e.RateLimitCheck(msg.ClientId) {
				atomic.AddInt64(&e.metrics.MessagesDropped, 1)
				continue
			}

			// 处理消息
			ctx, cancel := context.WithTimeout(e.ctx, messageProcessTimeout)
			err := e.messageHandler.HandleMessage(ctx, msg.ClientId, msg.Data)
			cancel()

			if err != nil {
				e.messageHandler.HandleError(ctx, msg.ClientId, err)
			}
		} else {
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// metricsCollector 指标收集器
func (e *EnhancedWsServer) metricsCollector() {
	ticker := time.NewTicker(metricsUpdateInterval)
	defer ticker.Stop()

	for {
		select {
		case <-e.ctx.Done():
			return
		case <-ticker.C:
			e.updateMetrics()
		}
	}
}

// updateMetrics 更新监控指标
func (e *EnhancedWsServer) updateMetrics() {
	e.metrics.mu.Lock()
	defer e.metrics.mu.Unlock()

	// 计算当前负载
	currentConnections := atomic.LoadInt64(&e.WsServer.stats.ActiveConnections)
	maxConnections := int64(connectionPoolSize)
	e.metrics.CurrentLoad = float64(currentConnections) / float64(maxConnections)

	// 计算平均响应时间（这里简化处理）
	e.metrics.AverageResponseTime = time.Millisecond * 50
}

// QueuedMessage 队列消息
type QueuedMessage struct {
	ClientId string
	Data     []byte
	Time     time.Time
}

// IsFull 检查连接池是否已满
func (cp *ConnectionPool) IsFull() bool {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return cp.pool.Size() >= cp.maxSize
}

// Add 添加连接到连接池
func (cp *ConnectionPool) Add(clientId string, client interface{}) bool {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	if cp.pool.Size() >= cp.maxSize {
		return false
	}

	cp.pool.Set(clientId, client)
	return true
}

// Remove 从连接池移除连接
func (cp *ConnectionPool) Remove(clientId string) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.pool.Remove(clientId)
}

// Allow 检查是否允许请求（令牌桶算法）
func (rl *RateLimiter) Allow(clientId string) bool {
	rl.mu.RLock()
	bucket := rl.limits.Get(clientId)
	rl.mu.RUnlock()

	if bucket == nil {
		// 创建新的令牌桶
		rl.mu.Lock()
		bucket = NewTokenBucket(rl.defaultLimit, rl.defaultLimit)
		rl.limits.Set(clientId, bucket)
		rl.mu.Unlock()
	}

	return bucket.(*TokenBucket).Take()
}

// SetLimit 设置客户端速率限制
func (rl *RateLimiter) SetLimit(clientId string, limit int64) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	bucket := NewTokenBucket(limit, limit)
	rl.limits.Set(clientId, bucket)
}

// Take 从令牌桶中取一个令牌
func (tb *TokenBucket) Take() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	// 补充令牌
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)
	tokensToAdd := int64(elapsed.Seconds() * float64(tb.rate))

	if tokensToAdd > 0 {
		tb.tokens = gconv.Int64(gconv.Int64(tb.tokens) + tokensToAdd)
		if tb.tokens > tb.capacity {
			tb.tokens = tb.capacity
		}

		tb.lastRefill = now
	}

	// 检查是否有可用令牌
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}

	return false
}

// HandleMessage 默认消息处理器实现
func (dmh *DefaultMessageHandler) HandleMessage(ctx context.Context, clientId string, message []byte) error {
	// 更新接收消息计数
	atomic.AddInt64(&dmh.server.metrics.MessagesReceived, 1)

	// 调用基础消息处理
	dmh.server.WsServer.handleMessage(ctx, clientId, message)
	return nil
}

// HandleError 默认错误处理器实现
func (dmh *DefaultMessageHandler) HandleError(ctx context.Context, clientId string, err error) {
	// 更新错误计数
	atomic.AddInt64(&dmh.server.metrics.ErrorCount, 1)

	dmh.server.metrics.mu.Lock()
	dmh.server.metrics.LastError = err
	dmh.server.metrics.LastErrorTime = time.Now()
	dmh.server.metrics.mu.Unlock()

	g.Log().Errorf(ctx, "客户端 %s 消息处理错误: %v", clientId, err)
}

// Close 关闭增强版服务器
func (e *EnhancedWsServer) Close() {
	e.cancel()
	e.WsServer.Close()

	g.Log().Info(context.Background(), "增强版WebSocket服务器已关闭")
}

// GetHealthStatus 获取健康状态
func (e *EnhancedWsServer) GetHealthStatus() map[string]interface{} {
	metrics := e.GetMetrics()

	status := map[string]interface{}{
		"status":             "healthy",
		"active_connections": atomic.LoadInt64(&e.WsServer.stats.ActiveConnections),
		"total_connections":  atomic.LoadInt64(&e.WsServer.stats.TotalConnections),
		"message_queue_size": e.messageQueue.Size(),
		"current_load":       metrics.CurrentLoad,
		"error_count":        metrics.ErrorCount,
		"last_error_time":    metrics.LastErrorTime,
		"uptime":             time.Since(e.metrics.LastErrorTime),
	}

	// 检查健康状态
	if metrics.CurrentLoad > 0.9 {
		status["status"] = "overloaded"
	}

	if metrics.ErrorCount > 100 {
		status["status"] = "error"
	}

	return status
}

// BroadcastWithFilter 带过滤器的广播
func (e *EnhancedWsServer) BroadcastWithFilter(ctx context.Context, message []byte, filter func(clientId string) bool) (int, error) {
	if len(message) == 0 {
		return 0, ErrInvalidMessage
	}

	count := 0
	e.WsServer.Clients.RLockFunc(func(m map[string]interface{}) {
		for clientId, v := range m {
			if !filter(clientId) {
				continue
			}

			client := v.(*WsClient)
			if !client.IsClosing && !client.IsClosed {
				select {
				case client.send <- message:
					count++
				default:
					go e.WsServer.closeClient(ctx, client.ClientId, false)
				}
			}
		}
	})

	atomic.AddInt64(&e.metrics.MessagesSent, int64(count))
	return count, nil
}

// GetConnectionInfo 获取连接详细信息
func (e *EnhancedWsServer) GetConnectionInfo(clientId string) map[string]interface{} {
	client := e.WsServer.Clients.Get(clientId)
	if client == nil {
		return nil
	}

	wsClient := client.(*WsClient)
	wsClient.mu.RLock()
	defer wsClient.mu.RUnlock()

	info := map[string]interface{}{
		"client_id":       clientId,
		"uid":             e.WsServer.GetUidByClientId(clientId),
		"is_closing":      wsClient.IsClosing,
		"is_closed":       wsClient.IsClosed,
		"last_activity":   wsClient.lastActivity,
		"groups":          wsClient.Groups.Slice(),
		"session":         gconv.Map(wsClient.Session),
		"send_buffer":     len(wsClient.send),
		"reconnect_count": atomic.LoadInt32(&wsClient.reconnectCount),
	}

	return info
}
