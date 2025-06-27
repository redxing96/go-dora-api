/*
 * @Description: WebSocket服务器实现 - 提供高性能、高并发的实时通信功能
 * @Author: redxing96@163.com
 * @Date: 2025-06-27 10:35:01
 * @LastEditTime: 2025-06-27 11:12:03
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/utility/websocket/websocket.go
 */
package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/gorilla/websocket"
)

// 常量定义 - 配置参数
const (
	// 允许写入的等待时间 - 防止客户端阻塞
	writeWait = 10 * time.Second

	// 读取下一个pong消息的超时时间 - 心跳检测
	pongWait = 60 * time.Second

	// 发送ping消息的时间间隔，必须小于pongWait - 保持连接活跃
	pingPeriod = (pongWait * 9) / 10

	// 最大消息大小（字节）- 防止内存溢出
	maxMessageSize = 512

	// 发送缓冲区大小 - 平衡内存使用和性能
	sendBufferSize = 256

	// 连接清理间隔 - 定期清理无效连接
	cleanupInterval = 5 * time.Minute

	// 最大重连次数 - 防止无限重连
	maxReconnectAttempts = 3

	// 重连延迟时间
	reconnectDelay = 1 * time.Second
)

// 特殊字符定义
var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// 错误类型定义
var (
	ErrClientNotFound   = fmt.Errorf("客户端不存在")
	ErrClientOffline    = fmt.Errorf("客户端已离线")
	ErrGroupNotFound    = fmt.Errorf("组不存在")
	ErrInvalidMessage   = fmt.Errorf("无效消息")
	ErrConnectionClosed = fmt.Errorf("连接已关闭")
	ErrBufferFull       = fmt.Errorf("发送缓冲区已满")
	ErrInvalidUID       = fmt.Errorf("无效的用户ID")
	ErrInvalidGroupID   = fmt.Errorf("无效的组ID")
	ErrInvalidClientID  = fmt.Errorf("无效的客户端ID")
)

// WebSocket升级器配置
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许跨域请求 - 生产环境应该根据实际需求配置
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	// 启用压缩
	EnableCompression: true,
}

// WsClient 表示一个WebSocket客户端连接
type WsClient struct {
	Conn           *websocket.Conn    // WebSocket连接对象
	ClientId       string             // 客户端唯一标识
	Uid            string             // 用户ID（可选）
	Groups         *gset.StrSet       // 客户端所在的组集合
	Session        g.Map              // 客户端会话数据
	IsClosing      bool               // 是否正在关闭
	IsClosed       bool               // 是否已关闭
	send           chan []byte        // 发送消息通道
	mu             sync.RWMutex       // 读写锁，保护客户端状态
	lastActivity   time.Time          // 最后活动时间
	reconnectCount int32              // 重连次数
	ctx            context.Context    // 上下文
	cancel         context.CancelFunc // 取消函数
}

// WsServer 表示WebSocket服务器
type WsServer struct {
	Clients       *gmap.StrAnyMap    // clientId -> *WsClient
	UidBindings   *gmap.StrStrMap    // clientId -> uid
	UidClientsMap *gmap.StrAnyMap    // uid -> *gset.StrSet
	Groups        *gmap.StrAnyMap    // groupId -> *gset.StrSet
	mu            sync.RWMutex       // 服务器级别的读写锁
	ctx           context.Context    // 服务器上下文
	cancel        context.CancelFunc // 取消函数
	stats         *ServerStats       // 服务器统计信息
	cleanupTicker *time.Ticker       // 清理定时器
}

// ServerStats 服务器统计信息
type ServerStats struct {
	TotalConnections  int64        // 总连接数
	ActiveConnections int64        // 活跃连接数
	TotalMessages     int64        // 总消息数
	TotalGroups       int64        // 总组数
	TotalUsers        int64        // 总用户数
	LastCleanupTime   time.Time    // 最后清理时间
	mu                sync.RWMutex // 统计信息锁
}

// Message 消息结构体
type Message struct {
	Type  string      `json:"type"`  // 消息类型
	Data  interface{} `json:"data"`  // 消息数据
	From  string      `json:"from"`  // 发送者
	To    string      `json:"to"`    // 接收者
	Group string      `json:"group"` // 组ID
	Time  int64       `json:"time"`  // 时间戳
}

// NewWsServer 创建一个新的WebSocket服务器实例
func NewWsServer() *WsServer {
	ctx, cancel := context.WithCancel(context.Background())

	server := &WsServer{
		Clients:       gmap.NewStrAnyMap(true),
		UidBindings:   gmap.NewStrStrMap(true),
		UidClientsMap: gmap.NewStrAnyMap(true),
		Groups:        gmap.NewStrAnyMap(true),
		ctx:           ctx,
		cancel:        cancel,
		stats: &ServerStats{
			LastCleanupTime: time.Now(),
		},
		cleanupTicker: time.NewTicker(cleanupInterval),
	}

	// 启动后台清理协程
	go server.cleanupRoutine()

	// 启动统计信息更新协程
	go server.statsRoutine()

	return server
}

// Close 关闭WebSocket服务器
func (s *WsServer) Close() {
	// 取消上下文
	s.cancel()

	// 停止清理定时器
	if s.cleanupTicker != nil {
		s.cleanupTicker.Stop()
	}

	// 关闭所有客户端连接
	s.Clients.RLockFunc(func(m map[string]interface{}) {
		for clientId := range m {
			s.closeClient(context.Background(), clientId, false)
		}
	})

	g.Log().Info(context.Background(), "WebSocket服务器已关闭")
}

// HandleWsConnection 处理WebSocket连接请求
func (s *WsServer) HandleWsConnection(w http.ResponseWriter, r *http.Request) {
	// 使用recover防止panic
	defer func() {
		if err := recover(); err != nil {
			g.Log().Errorf(r.Context(), "WebSocket连接处理panic: %v\n%s", err, debug.Stack())
		}
	}()

	// 检查服务器是否已关闭
	select {
	case <-s.ctx.Done():
		http.Error(w, "服务器已关闭", http.StatusServiceUnavailable)
		return
	default:
	}

	// 将HTTP连接升级为WebSocket连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		g.Log().Errorf(r.Context(), "WebSocket连接升级失败: %v", err)
		return
	}

	// 生成客户端ID
	clientId := generateClientId()

	// 创建客户端上下文
	clientCtx, clientCancel := context.WithCancel(s.ctx)

	// 创建客户端对象
	client := &WsClient{
		Conn:         conn,
		ClientId:     clientId,
		Groups:       gset.NewStrSet(),
		Session:      make(g.Map),
		send:         make(chan []byte, sendBufferSize),
		lastActivity: time.Now(),
		ctx:          clientCtx,
		cancel:       clientCancel,
	}

	// 将客户端添加到服务器
	s.mu.Lock()
	s.Clients.Set(clientId, client)
	s.mu.Unlock()

	// 更新统计信息
	atomic.AddInt64(&s.stats.TotalConnections, 1)
	atomic.AddInt64(&s.stats.ActiveConnections, 1)

	// 启动读写协程
	go client.writePump(s)
	go client.readPump(s)

	g.Log().Debugf(r.Context(), "新客户端连接: %s, 当前活跃连接数: %d", clientId, atomic.LoadInt64(&s.stats.ActiveConnections))
}

// writePump 客户端写入协程 - 负责向客户端发送消息
func (c *WsClient) writePump(server *WsServer) {
	// 使用recover防止panic
	defer func() {
		if err := recover(); err != nil {
			g.Log().Errorf(c.ctx, "客户端 %s 写入协程panic: %v\n%s", c.ClientId, err, debug.Stack())
		}
		server.closeClient(c.ctx, c.ClientId, true)
	}()

	// 创建ping定时器
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			// 上下文取消，退出协程
			return

		case message, ok := <-c.send:
			// 设置写超时
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))

			if !ok {
				// 通道关闭，发送关闭消息
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			// 获取消息写入器
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				g.Log().Errorf(c.ctx, "客户端 %s 获取写入器失败: %v", c.ClientId, err)
				return
			}

			// 写入消息
			if _, err := w.Write(message); err != nil {
				g.Log().Errorf(c.ctx, "客户端 %s 写入消息失败: %v", c.ClientId, err)
				w.Close()
				return
			}

			// 写入队列中的其他消息
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			// 关闭写入器
			if err := w.Close(); err != nil {
				g.Log().Errorf(c.ctx, "客户端 %s 关闭写入器失败: %v", c.ClientId, err)
				return
			}

			// 更新最后活动时间
			c.mu.Lock()
			c.lastActivity = time.Now()
			c.mu.Unlock()

		case <-ticker.C:
			// 发送ping消息
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				g.Log().Errorf(c.ctx, "客户端 %s 发送ping失败: %v", c.ClientId, err)
				return
			}
		}
	}
}

// readPump 客户端读取协程 - 负责从客户端读取消息
func (c *WsClient) readPump(server *WsServer) {
	// 使用recover防止panic
	defer func() {
		if err := recover(); err != nil {
			g.Log().Errorf(c.ctx, "客户端 %s 读取协程panic: %v\n%s", c.ClientId, err, debug.Stack())
		}
		server.closeClient(c.ctx, c.ClientId, true)
	}()

	// 设置连接参数
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		// 读取消息
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				g.Log().Errorf(c.ctx, "客户端 %s 读取消息错误: %v", c.ClientId, err)
			}
			break
		}

		// 更新最后活动时间
		c.mu.Lock()
		c.lastActivity = time.Now()
		c.mu.Unlock()

		// 更新消息统计
		atomic.AddInt64(&server.stats.TotalMessages, 1)

		// 处理消息
		go server.handleMessage(c.ctx, c.ClientId, message)
	}
}

// handleMessage 处理接收到的消息
func (s *WsServer) handleMessage(ctx context.Context, clientId string, message []byte) {
	defer func() {
		if err := recover(); err != nil {
			g.Log().Errorf(ctx, "处理消息panic: %v\n%s", err, debug.Stack())
		}
	}()

	// 解析消息
	var msg Message
	if err := json.Unmarshal(message, &msg); err != nil {
		g.Log().Errorf(ctx, "客户端 %s 发送的消息格式无效: %v", clientId, err)
		return
	}

	// 设置消息时间戳
	if msg.Time == 0 {
		msg.Time = time.Now().Unix()
	}

	// 根据消息类型处理
	switch msg.Type {
	case "ping":
		// 心跳消息
		response := Message{
			Type: "pong",
			Time: time.Now().Unix(),
		}
		if data, err := json.Marshal(response); err == nil {
			s.SendToClient(ctx, clientId, data)
		}

	case "bind_uid":
		// 绑定用户ID
		if uid, ok := msg.Data.(string); ok && uid != "" {
			s.BindUid(ctx, clientId, uid)
		}

	case "join_group":
		// 加入组
		if groupId, ok := msg.Data.(string); ok && groupId != "" {
			s.JoinGroup(ctx, clientId, groupId)
		}

	case "leave_group":
		// 离开组
		if groupId, ok := msg.Data.(string); ok && groupId != "" {
			s.LeaveGroup(ctx, clientId, groupId)
		}

	case "broadcast":
		// 广播消息
		if msg.Group != "" {
			s.SendToGroup(ctx, msg.Group, message)
		} else {
			s.SendToAll(ctx, message)
		}

	default:
		g.Log().Debugf(ctx, "收到来自客户端 %s 的消息: %s", clientId, string(message))
	}
}

// generateClientId 生成唯一的客户端ID
func generateClientId() string {
	return gconv.String(guid.S())
}

// SendToAll 向所有客户端发送消息
func (s *WsServer) SendToAll(ctx context.Context, message []byte) error {
	if len(message) == 0 {
		return ErrInvalidMessage
	}

	count := 0
	s.Clients.RLockFunc(func(m map[string]interface{}) {
		for _, v := range m {
			client := v.(*WsClient)
			if !client.IsClosing && !client.IsClosed {
				select {
				case client.send <- message:
					count++
				default:
					// 发送缓冲区已满，关闭连接
					go s.closeClient(ctx, client.ClientId, false)
				}
			}
		}
	})

	g.Log().Debugf(ctx, "向所有客户端发送消息，成功发送给 %d 个客户端", count)
	return nil
}

// SendToClient 向指定客户端发送消息
func (s *WsServer) SendToClient(ctx context.Context, clientId string, message []byte) bool {
	if clientId == "" || len(message) == 0 {
		return false
	}

	client := s.Clients.Get(clientId)
	if client == nil {
		return false
	}

	wsClient := client.(*WsClient)
	if wsClient.IsClosing || wsClient.IsClosed {
		return false
	}

	select {
	case wsClient.send <- message:
		return true
	default:
		// 发送缓冲区已满，关闭连接
		go s.closeClient(ctx, clientId, false)
		return false
	}
}

// CloseClient 关闭指定客户端连接
func (s *WsServer) CloseClient(ctx context.Context, clientId string) error {
	if clientId == "" {
		return ErrInvalidClientID
	}

	s.closeClient(ctx, clientId, false)
	return nil
}

// closeClient 关闭客户端连接的内部实现
func (s *WsServer) closeClient(ctx context.Context, clientId string, isReadError bool) {
	client := s.Clients.Get(clientId)
	if client == nil {
		return
	}

	wsClient := client.(*WsClient)

	// 设置关闭状态
	wsClient.mu.Lock()
	if wsClient.IsClosing || wsClient.IsClosed {
		wsClient.mu.Unlock()
		return
	}
	wsClient.IsClosing = true
	wsClient.mu.Unlock()

	// 取消客户端上下文
	wsClient.cancel()

	// 解绑UID
	if uid := s.UidBindings.Get(clientId); uid != "" {
		s.unbindUid(ctx, clientId, uid)
	}

	// 离开所有组
	wsClient.Groups.Iterator(func(groupId string) bool {
		s.leaveGroup(ctx, clientId, groupId)
		return true
	})

	// 关闭连接
	close(wsClient.send)
	wsClient.Conn.Close()

	// 标记为已关闭
	wsClient.mu.Lock()
	wsClient.IsClosed = true
	wsClient.mu.Unlock()

	// 从客户端列表中移除
	s.Clients.Remove(clientId)

	// 更新统计信息
	atomic.AddInt64(&s.stats.ActiveConnections, -1)

	if !isReadError {
		g.Log().Debugf(ctx, "客户端 %s 连接已关闭", clientId)
	}
}

// IsOnline 检查客户端是否在线
func (s *WsServer) IsOnline(clientId string) bool {
	if clientId == "" {
		return false
	}

	client := s.Clients.Get(clientId)
	if client == nil {
		return false
	}

	wsClient := client.(*WsClient)
	return !wsClient.IsClosing && !wsClient.IsClosed
}

// BindUid 绑定客户端到UID
func (s *WsServer) BindUid(ctx context.Context, clientId string, uid string) error {
	if clientId == "" {
		return ErrInvalidClientID
	}
	if uid == "" {
		return ErrInvalidUID
	}

	// 检查客户端是否在线
	if !s.IsOnline(clientId) {
		return ErrClientOffline
	}

	// 先解除之前的绑定
	oldUid := s.UidBindings.Get(clientId)
	if oldUid != "" && oldUid != uid {
		s.unbindUid(ctx, clientId, oldUid)
	}

	// 设置新的绑定
	s.UidBindings.Set(clientId, uid)
	s.UidClientsMap.GetOrSetFuncLock(uid, func() interface{} {
		return gset.NewStrSet()
	}).(*gset.StrSet).Add(clientId)

	g.Log().Debugf(ctx, "客户端 %s 已绑定到UID %s", clientId, uid)
	return nil
}

// UnbindUid 解除客户端与UID的绑定
func (s *WsServer) UnbindUid(ctx context.Context, clientId string, uid string) error {
	if clientId == "" {
		return ErrInvalidClientID
	}
	if uid == "" {
		return ErrInvalidUID
	}

	s.unbindUid(ctx, clientId, uid)
	return nil
}

// unbindUid 解除绑定的内部实现
func (s *WsServer) unbindUid(ctx context.Context, clientId string, uid string) {
	if s.UidBindings.Get(clientId) != uid {
		return
	}

	s.UidBindings.Remove(clientId)

	if clients := s.UidClientsMap.Get(uid); clients != nil {
		clients.(*gset.StrSet).Remove(clientId)
		if clients.(*gset.StrSet).Size() == 0 {
			s.UidClientsMap.Remove(uid)
		}
	}

	g.Log().Debugf(ctx, "客户端 %s 已解除与UID %s 的绑定", clientId, uid)
}

// IsUidOnline 检查UID是否在线
func (s *WsServer) IsUidOnline(uid string) bool {
	if uid == "" {
		return false
	}

	if clients := s.UidClientsMap.Get(uid); clients != nil {
		return clients.(*gset.StrSet).Size() > 0
	}
	return false
}

// GetClientIdByUid 获取UID对应的所有客户端ID
func (s *WsServer) GetClientIdByUid(uid string) []string {
	if uid == "" {
		return []string{}
	}

	if clients := s.UidClientsMap.Get(uid); clients != nil {
		return clients.(*gset.StrSet).Slice()
	}
	return []string{}
}

// GetUidByClientId 获取客户端ID对应的UID
func (s *WsServer) GetUidByClientId(clientId string) string {
	if clientId == "" {
		return ""
	}

	return s.UidBindings.Get(clientId)
}

// SendToUid 向指定用户发送消息
func (s *WsServer) SendToUid(ctx context.Context, uid string, message []byte) (int, error) {
	if uid == "" {
		return 0, ErrInvalidUID
	}
	if len(message) == 0 {
		return 0, ErrInvalidMessage
	}

	count := 0
	if clients := s.UidClientsMap.Get(uid); clients != nil {
		clients.(*gset.StrSet).Iterator(func(clientId string) bool {
			if s.SendToClient(ctx, clientId, message) {
				count++
			}
			return true
		})
	}

	return count, nil
}

// JoinGroup 将客户端加入组
func (s *WsServer) JoinGroup(ctx context.Context, clientId string, groupId string) error {
	if clientId == "" {
		return ErrInvalidClientID
	}
	if groupId == "" {
		return ErrInvalidGroupID
	}

	client := s.Clients.Get(clientId)
	if client == nil {
		return ErrClientNotFound
	}

	wsClient := client.(*WsClient)
	if wsClient.IsClosing || wsClient.IsClosed {
		return ErrClientOffline
	}

	// 将客户端加入组
	wsClient.mu.Lock()
	wsClient.Groups.Add(groupId)
	wsClient.mu.Unlock()

	// 将客户端添加到组
	s.Groups.GetOrSetFuncLock(groupId, func() interface{} {
		atomic.AddInt64(&s.stats.TotalGroups, 1)
		return gset.NewStrSet()
	}).(*gset.StrSet).Add(clientId)

	g.Log().Debugf(ctx, "客户端 %s 已加入组 %s", clientId, groupId)
	return nil
}

// LeaveGroup 将客户端从组中移除
func (s *WsServer) LeaveGroup(ctx context.Context, clientId string, groupId string) error {
	if clientId == "" {
		return ErrInvalidClientID
	}
	if groupId == "" {
		return ErrInvalidGroupID
	}

	s.leaveGroup(ctx, clientId, groupId)
	return nil
}

// leaveGroup 从组中移除客户端的内部实现
func (s *WsServer) leaveGroup(ctx context.Context, clientId string, groupId string) {
	client := s.Clients.Get(clientId)
	if client == nil {
		return
	}

	wsClient := client.(*WsClient)
	wsClient.mu.Lock()
	wsClient.Groups.Remove(groupId)
	wsClient.mu.Unlock()

	if group := s.Groups.Get(groupId); group != nil {
		group.(*gset.StrSet).Remove(clientId)
		if group.(*gset.StrSet).Size() == 0 {
			s.Groups.Remove(groupId)
			atomic.AddInt64(&s.stats.TotalGroups, -1)
		}
	}

	g.Log().Debugf(ctx, "客户端 %s 已离开组 %s", clientId, groupId)
}

// Ungroup 解散组
func (s *WsServer) Ungroup(ctx context.Context, groupId string) error {
	if groupId == "" {
		return ErrInvalidGroupID
	}

	if group := s.Groups.Get(groupId); group != nil {
		group.(*gset.StrSet).Iterator(func(clientId string) bool {
			if client := s.Clients.Get(clientId); client != nil {
				client.(*WsClient).mu.Lock()
				client.(*WsClient).Groups.Remove(groupId)
				client.(*WsClient).mu.Unlock()
			}
			return true
		})
		s.Groups.Remove(groupId)
		atomic.AddInt64(&s.stats.TotalGroups, -1)
	}

	g.Log().Debugf(ctx, "组 %s 已解散", groupId)
	return nil
}

// SendToGroup 向组内所有客户端发送消息
func (s *WsServer) SendToGroup(ctx context.Context, groupId string, message []byte) (int, error) {
	if groupId == "" {
		return 0, ErrInvalidGroupID
	}
	if len(message) == 0 {
		return 0, ErrInvalidMessage
	}

	count := 0
	if group := s.Groups.Get(groupId); group != nil {
		group.(*gset.StrSet).Iterator(func(clientId string) bool {
			if s.SendToClient(ctx, clientId, message) {
				count++
			}
			return true
		})
	}

	return count, nil
}

// GetClientIdCountByGroup 获取组内客户端数量
func (s *WsServer) GetClientIdCountByGroup(groupId string) int {
	if groupId == "" {
		return 0
	}

	if group := s.Groups.Get(groupId); group != nil {
		return group.(*gset.StrSet).Size()
	}
	return 0
}

// GetClientSessionsByGroup 获取组内所有客户端的会话信息
func (s *WsServer) GetClientSessionsByGroup(groupId string) []g.Map {
	if groupId == "" {
		return []g.Map{}
	}

	sessions := make([]g.Map, 0)
	if group := s.Groups.Get(groupId); group != nil {
		group.(*gset.StrSet).Iterator(func(clientId string) bool {
			if client := s.Clients.Get(clientId); client != nil {
				client.(*WsClient).mu.RLock()
				session := gconv.Map(client.(*WsClient).Session)
				session["client_id"] = clientId
				session["uid"] = s.GetUidByClientId(clientId)
				session["last_activity"] = client.(*WsClient).lastActivity
				client.(*WsClient).mu.RUnlock()
				sessions = append(sessions, session)
			}
			return true
		})
	}
	return sessions
}

// GetAllClientIdCount 获取所有客户端数量
func (s *WsServer) GetAllClientIdCount() int {
	return s.Clients.Size()
}

// GetAllClientSessions 获取所有客户端的会话信息
func (s *WsServer) GetAllClientSessions() []g.Map {
	sessions := make([]g.Map, 0)
	s.Clients.RLockFunc(func(m map[string]interface{}) {
		for clientId, v := range m {
			client := v.(*WsClient)
			client.mu.RLock()
			session := gconv.Map(client.Session)
			session["client_id"] = clientId
			session["uid"] = s.GetUidByClientId(clientId)
			session["last_activity"] = client.lastActivity
			session["is_closing"] = client.IsClosing
			session["is_closed"] = client.IsClosed
			client.mu.RUnlock()
			sessions = append(sessions, session)
		}
	})
	return sessions
}

// SetSession 设置客户端会话
func (s *WsServer) SetSession(ctx context.Context, clientId string, session g.Map) error {
	if clientId == "" {
		return ErrInvalidClientID
	}

	client := s.Clients.Get(clientId)
	if client == nil {
		return ErrClientNotFound
	}

	wsClient := client.(*WsClient)
	if wsClient.IsClosing || wsClient.IsClosed {
		return ErrClientOffline
	}

	wsClient.mu.Lock()
	wsClient.Session = session
	wsClient.mu.Unlock()

	return nil
}

// UpdateSession 更新客户端会话
func (s *WsServer) UpdateSession(ctx context.Context, clientId string, session g.Map) error {
	if clientId == "" {
		return ErrInvalidClientID
	}

	client := s.Clients.Get(clientId)
	if client == nil {
		return ErrClientNotFound
	}

	wsClient := client.(*WsClient)
	if wsClient.IsClosing || wsClient.IsClosed {
		return ErrClientOffline
	}

	wsClient.mu.Lock()
	for k, v := range session {
		wsClient.Session[k] = v
	}
	wsClient.mu.Unlock()

	return nil
}

// GetSession 获取客户端会话
func (s *WsServer) GetSession(clientId string) g.Map {
	if clientId == "" {
		return nil
	}

	client := s.Clients.Get(clientId)
	if client == nil {
		return nil
	}

	wsClient := client.(*WsClient)
	if wsClient.IsClosing || wsClient.IsClosed {
		return nil
	}

	wsClient.mu.RLock()
	session := gconv.Map(wsClient.Session)
	wsClient.mu.RUnlock()
	return session
}

// GetClientIdListByGroup 获取组内所有客户端ID列表
func (s *WsServer) GetClientIdListByGroup(groupId string) []string {
	if groupId == "" {
		return []string{}
	}

	if group := s.Groups.Get(groupId); group != nil {
		return group.(*gset.StrSet).Slice()
	}
	return []string{}
}

// GetAllClientIdList 获取所有客户端ID列表
func (s *WsServer) GetAllClientIdList() []string {
	clientIds := make([]string, 0)
	s.Clients.RLockFunc(func(m map[string]interface{}) {
		for clientId := range m {
			clientIds = append(clientIds, clientId)
		}
	})
	return clientIds
}

// GetUidListByGroup 获取组内所有UID列表
func (s *WsServer) GetUidListByGroup(groupId string) []string {
	if groupId == "" {
		return []string{}
	}

	uids := gset.NewStrSet()
	if group := s.Groups.Get(groupId); group != nil {
		group.(*gset.StrSet).Iterator(func(clientId string) bool {
			if uid := s.GetUidByClientId(clientId); uid != "" {
				uids.Add(uid)
			}
			return true
		})
	}
	return uids.Slice()
}

// GetUidCountByGroup 获取组内UID数量
func (s *WsServer) GetUidCountByGroup(groupId string) int {
	return len(s.GetUidListByGroup(groupId))
}

// GetAllUidList 获取所有UID列表
func (s *WsServer) GetAllUidList() []string {
	uids := gset.NewStrSet()
	s.Clients.RLockFunc(func(m map[string]interface{}) {
		for clientId := range m {
			if uid := s.GetUidByClientId(clientId); uid != "" {
				uids.Add(uid)
			}
		}
	})
	return uids.Slice()
}

// GetAllUidCount 获取所有UID数量
func (s *WsServer) GetAllUidCount() int {
	return len(s.GetAllUidList())
}

// GetAllGroupIdList 获取所有组ID列表
func (s *WsServer) GetAllGroupIdList() []string {
	groupIds := make([]string, 0)
	s.Groups.RLockFunc(func(m map[string]interface{}) {
		for groupId := range m {
			groupIds = append(groupIds, groupId)
		}
	})
	return groupIds
}

// GetStats 获取服务器统计信息
func (s *WsServer) GetStats() *ServerStats {
	s.stats.mu.RLock()
	defer s.stats.mu.RUnlock()

	stats := &ServerStats{
		TotalConnections:  atomic.LoadInt64(&s.stats.TotalConnections),
		ActiveConnections: atomic.LoadInt64(&s.stats.ActiveConnections),
		TotalMessages:     atomic.LoadInt64(&s.stats.TotalMessages),
		TotalGroups:       atomic.LoadInt64(&s.stats.TotalGroups),
		TotalUsers:        int64(len(s.GetAllUidList())),
		LastCleanupTime:   s.stats.LastCleanupTime,
	}

	return stats
}

// cleanupRoutine 定期清理协程
func (s *WsServer) cleanupRoutine() {
	defer func() {
		if err := recover(); err != nil {
			g.Log().Errorf(s.ctx, "清理协程panic: %v\n%s", err, debug.Stack())
		}
	}()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-s.cleanupTicker.C:
			s.cleanup()
		}
	}
}

// cleanup 执行清理操作
func (s *WsServer) cleanup() {
	defer func() {
		if err := recover(); err != nil {
			g.Log().Errorf(s.ctx, "清理操作panic: %v\n%s", err, debug.Stack())
		}
	}()

	// 清理超时的客户端连接
	timeout := time.Now().Add(-pongWait * 2)
	s.Clients.RLockFunc(func(m map[string]interface{}) {
		for clientId, v := range m {
			client := v.(*WsClient)
			client.mu.RLock()
			lastActivity := client.lastActivity
			client.mu.RUnlock()

			if lastActivity.Before(timeout) {
				g.Log().Debugf(s.ctx, "清理超时客户端: %s", clientId)
				go s.closeClient(s.ctx, clientId, false)
			}
		}
	})

	// 清理空的组
	s.Groups.RLockFunc(func(m map[string]interface{}) {
		for groupId, v := range m {
			if v.(*gset.StrSet).Size() == 0 {
				s.Groups.Remove(groupId)
				atomic.AddInt64(&s.stats.TotalGroups, -1)
			}
		}
	})

	// 更新清理时间
	s.stats.mu.Lock()
	s.stats.LastCleanupTime = time.Now()
	s.stats.mu.Unlock()

	g.Log().Debugf(s.ctx, "定期清理完成，当前活跃连接数: %d", atomic.LoadInt64(&s.stats.ActiveConnections))
}

// statsRoutine 统计信息更新协程
func (s *WsServer) statsRoutine() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			// 更新用户数量统计
			atomic.StoreInt64(&s.stats.TotalUsers, int64(len(s.GetAllUidList())))
		}
	}
}
