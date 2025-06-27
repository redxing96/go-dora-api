# WebSocket服务器优化文档

## 概述

这是一个高性能、高并发的WebSocket服务器实现，提供了基础版和增强版两种模式，支持实时通信、群组管理、用户绑定等功能。

## 主要特性

### 基础版特性
- ✅ 高性能WebSocket连接管理
- ✅ 客户端生命周期管理
- ✅ 用户ID绑定和组管理
- ✅ 消息广播和定向发送
- ✅ 会话数据管理
- ✅ 连接状态监控
- ✅ 自动清理机制
- ✅ 错误处理和恢复

### 增强版特性
- ✅ 连接池管理
- ✅ 消息队列处理
- ✅ 速率限制（令牌桶算法）
- ✅ 监控指标收集
- ✅ 健康状态检查
- ✅ 批量消息发送
- ✅ 消息重试机制
- ✅ 自定义消息处理器
- ✅ 负载均衡支持

## 架构设计

### 核心组件

1. **WsServer**: 基础WebSocket服务器
2. **EnhancedWsServer**: 增强版WebSocket服务器
3. **WsClient**: 客户端连接对象
4. **ConnectionPool**: 连接池管理
5. **RateLimiter**: 速率限制器
6. **Metrics**: 监控指标收集

### 数据结构

```go
// 客户端结构
type WsClient struct {
    Conn         *websocket.Conn    // WebSocket连接
    ClientId     string             // 客户端ID
    Uid          string             // 用户ID
    Groups       *gset.StrSet       // 所属组
    Session      g.Map              // 会话数据
    IsClosing    bool               // 关闭状态
    send         chan []byte        // 发送通道
    mu           sync.RWMutex       // 读写锁
    lastActivity time.Time          // 最后活动时间
    ctx          context.Context    // 上下文
    cancel       context.CancelFunc // 取消函数
}

// 服务器结构
type WsServer struct {
    Clients       *gmap.StrAnyMap    // 客户端映射
    UidBindings   *gmap.StrStrMap    // UID绑定
    UidClientsMap *gmap.StrAnyMap    // UID客户端映射
    Groups        *gmap.StrAnyMap    // 组管理
    stats         *ServerStats       // 统计信息
    cleanupTicker *time.Ticker       // 清理定时器
}
```

## 使用方法

### 基础版使用

```go
package main

import (
    "net/http"
    "github.com/your-project/utility/websocket"
)

func main() {
    // 创建WebSocket服务器
    server := websocket.NewWsServer()
    defer server.Close()

    // 设置路由
    http.HandleFunc("/ws", server.HandleWsConnection)

    // 启动服务器
    http.ListenAndServe(":8080", nil)
}
```

### 增强版使用

```go
package main

import (
    "net/http"
    "github.com/your-project/utility/websocket"
)

func main() {
    // 创建增强版WebSocket服务器
    server := websocket.NewEnhancedWsServer()
    defer server.Close()

    // 设置路由
    http.HandleFunc("/ws", server.HandleWsConnection)
    
    // 健康检查端点
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        status := server.GetHealthStatus()
        // 返回JSON格式的健康状态
    })

    // 启动服务器
    http.ListenAndServe(":8080", nil)
}
```

## API参考

### 连接管理

```go
// 处理WebSocket连接
func (s *WsServer) HandleWsConnection(w http.ResponseWriter, r *http.Request)

// 关闭客户端连接
func (s *WsServer) CloseClient(ctx context.Context, clientId string) error

// 检查客户端是否在线
func (s *WsServer) IsOnline(clientId string) bool
```

### 消息发送

```go
// 发送给所有客户端
func (s *WsServer) SendToAll(ctx context.Context, message []byte) error

// 发送给指定客户端
func (s *WsServer) SendToClient(ctx context.Context, clientId string, message []byte) bool

// 发送给指定用户
func (s *WsServer) SendToUid(ctx context.Context, uid string, message []byte) (int, error)

// 发送给指定组
func (s *WsServer) SendToGroup(ctx context.Context, groupId string, message []byte) (int, error)

// 带重试的消息发送（增强版）
func (e *EnhancedWsServer) SendMessageWithRetry(ctx context.Context, clientId string, message []byte, maxRetries int) error

// 批量发送消息（增强版）
func (e *EnhancedWsServer) SendBatchMessage(ctx context.Context, clientIds []string, message []byte) (int, error)
```

### 用户管理

```go
// 绑定用户ID
func (s *WsServer) BindUid(ctx context.Context, clientId string, uid string) error

// 解绑用户ID
func (s *WsServer) UnbindUid(ctx context.Context, clientId string, uid string) error

// 检查用户是否在线
func (s *WsServer) IsUidOnline(uid string) bool

// 获取用户的客户端列表
func (s *WsServer) GetClientIdByUid(uid string) []string
```

### 组管理

```go
// 加入组
func (s *WsServer) JoinGroup(ctx context.Context, clientId string, groupId string) error

// 离开组
func (s *WsServer) LeaveGroup(ctx context.Context, clientId string, groupId string) error

// 解散组
func (s *WsServer) Ungroup(ctx context.Context, groupId string) error

// 获取组内客户端数量
func (s *WsServer) GetClientIdCountByGroup(groupId string) int

// 获取组内客户端列表
func (s *WsServer) GetClientIdListByGroup(groupId string) []string
```

### 会话管理

```go
// 设置会话数据
func (s *WsServer) SetSession(ctx context.Context, clientId string, session g.Map) error

// 更新会话数据
func (s *WsServer) UpdateSession(ctx context.Context, clientId string, session g.Map) error

// 获取会话数据
func (s *WsServer) GetSession(clientId string) g.Map
```

### 监控和统计

```go
// 获取服务器统计信息
func (s *WsServer) GetStats() *ServerStats

// 获取监控指标（增强版）
func (e *EnhancedWsServer) GetMetrics() *Metrics

// 获取健康状态（增强版）
func (e *EnhancedWsServer) GetHealthStatus() map[string]interface{}

// 设置速率限制（增强版）
func (e *EnhancedWsServer) SetRateLimit(clientId string, limit int64)
```

## 消息格式

### 标准消息格式

```json
{
    "type": "message_type",
    "data": "message_data",
    "from": "sender_id",
    "to": "receiver_id",
    "group": "group_id",
    "time": 1640995200
}
```

### 预定义消息类型

- `ping`: 心跳消息
- `pong`: 心跳响应
- `bind_uid`: 绑定用户ID
- `join_group`: 加入组
- `leave_group`: 离开组
- `broadcast`: 广播消息
- `chat`: 聊天消息

## 性能优化

### 并发安全
- 使用读写锁保护共享数据
- 原子操作更新计数器
- 协程安全的连接管理

### 内存管理
- 连接池限制最大连接数
- 消息队列防止内存溢出
- 定期清理无效连接

### 网络优化
- 消息压缩支持
- 批量消息发送
- 连接复用

### 监控指标
- 连接数统计
- 消息吞吐量
- 错误率监控
- 响应时间统计

## 配置参数

### 基础配置

```go
const (
    writeWait = 10 * time.Second        // 写入超时
    pongWait = 60 * time.Second         // Pong超时
    pingPeriod = 54 * time.Second       // Ping间隔
    maxMessageSize = 512                // 最大消息大小
    sendBufferSize = 256                // 发送缓冲区大小
    cleanupInterval = 5 * time.Minute   // 清理间隔
)
```

### 增强版配置

```go
const (
    messageQueueSize = 10000            // 消息队列大小
    connectionPoolSize = 1000           // 连接池大小
    messageProcessTimeout = 5 * time.Second  // 消息处理超时
    batchSendSize = 100                 // 批量发送大小
    metricsUpdateInterval = 10 * time.Second // 指标更新间隔
)
```

## 错误处理

### 错误类型

```go
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
```

### 错误恢复

- 自动重连机制
- 消息重试机制
- 连接状态检查
- 异常恢复处理

## 部署建议

### 生产环境配置

1. **负载均衡**: 使用多个WebSocket服务器实例
2. **监控告警**: 设置连接数和错误率告警
3. **日志管理**: 配置结构化日志输出
4. **资源限制**: 设置合理的连接数和内存限制

### 性能调优

1. **连接池大小**: 根据服务器资源调整
2. **消息队列大小**: 根据业务需求设置
3. **清理间隔**: 平衡性能和资源使用
4. **速率限制**: 防止恶意攻击

## 示例代码

详细的示例代码请参考 `example.go` 文件，包括：

- 基础服务器使用
- 增强版服务器使用
- 群聊功能实现
- 自定义消息处理器
- 性能测试
- 负载均衡

## 注意事项

1. **内存使用**: 大量连接时注意内存消耗
2. **网络带宽**: 高并发时注意网络带宽使用
3. **错误处理**: 正确处理连接断开和异常情况
4. **安全考虑**: 生产环境需要添加认证和授权
5. **监控告警**: 建议添加完善的监控和告警机制

## 更新日志

### v1.0.0 (2025-06-27)
- 初始版本发布
- 基础WebSocket功能
- 用户和组管理
- 消息发送功能

### v1.1.0 (2025-06-27)
- 增强版功能发布
- 连接池和消息队列
- 监控指标收集
- 速率限制功能
- 性能优化 