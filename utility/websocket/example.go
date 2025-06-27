/*
 * @Description: WebSocket服务器使用示例 - 展示基础版和增强版的使用方法
 * @Author: redxing96@163.com
 * @Date: 2025-06-27 11:10:00
 * @LastEditTime: 2025-06-27 11:10:00
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/utility/websocket/example.go
 */
package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// ExampleBasicServer 基础版WebSocket服务器使用示例
func ExampleBasicServer() {
	// 创建基础版WebSocket服务器
	server := NewWsServer()
	defer server.Close()

	// 设置HTTP路由
	http.HandleFunc("/ws", server.HandleWsConnection)

	// 启动HTTP服务器
	go func() {
		g.Log().Info(context.Background(), "基础版WebSocket服务器启动在 :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			g.Log().Error(context.Background(), "HTTP服务器启动失败:", err)
		}
	}()

	// 模拟消息发送
	go func() {
		time.Sleep(2 * time.Second)

		// 发送广播消息
		message := Message{
			Type: "broadcast",
			Data: "Hello, World!",
			Time: time.Now().Unix(),
		}

		if data, err := json.Marshal(message); err == nil {
			server.SendToAll(context.Background(), data)
		}

		// 获取服务器统计信息
		stats := server.GetStats()
		g.Log().Infof(context.Background(), "服务器统计: 活跃连接=%d, 总连接=%d, 总消息=%d",
			stats.ActiveConnections, stats.TotalConnections, stats.TotalMessages)
	}()

	// 保持服务器运行
	select {}
}

// ExampleEnhancedServer 增强版WebSocket服务器使用示例
func ExampleEnhancedServer() {
	// 创建增强版WebSocket服务器
	server := NewEnhancedWsServer()
	defer server.Close()

	// 设置HTTP路由
	http.HandleFunc("/ws/enhanced", server.HandleWsConnection)

	// 添加健康检查端点
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		status := server.GetHealthStatus()
		if data, err := json.Marshal(status); err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		}
	})

	// 启动HTTP服务器
	go func() {
		g.Log().Info(context.Background(), "增强版WebSocket服务器启动在 :8081")
		if err := http.ListenAndServe(":8081", nil); err != nil {
			g.Log().Error(context.Background(), "HTTP服务器启动失败:", err)
		}
	}()

	// 模拟高级功能使用
	go func() {
		time.Sleep(2 * time.Second)

		// 设置速率限制
		server.SetRateLimit("client1", 10) // 每秒10个请求

		// 发送带重试的消息
		message := Message{
			Type: "notification",
			Data: "Important notification",
			Time: time.Now().Unix(),
		}

		if data, err := json.Marshal(message); err == nil {
			server.SendMessageWithRetry(context.Background(), "client1", data, 3)
		}

		// 批量发送消息
		clientIds := []string{"client1", "client2", "client3"}
		batchMessage := Message{
			Type: "batch_notification",
			Data: "Batch message",
			Time: time.Now().Unix(),
		}

		if data, err := json.Marshal(batchMessage); err == nil {
			count, err := server.SendBatchMessage(context.Background(), clientIds, data)
			g.Log().Infof(context.Background(), "批量发送结果: 成功=%d, 错误=%v", count, err)
		}

		// 获取监控指标
		metrics := server.GetMetrics()
		g.Log().Infof(context.Background(), "监控指标: 连接成功=%d, 消息发送=%d, 当前负载=%.2f",
			metrics.ConnectionSuccess, metrics.MessagesSent, metrics.CurrentLoad)
	}()

	// 保持服务器运行
	select {}
}

// ExampleGroupChat 群聊功能示例
func ExampleGroupChat() {
	server := NewEnhancedWsServer()
	defer server.Close()

	http.HandleFunc("/ws/chat", server.HandleWsConnection)

	go func() {
		g.Log().Info(context.Background(), "群聊WebSocket服务器启动在 :8082")
		if err := http.ListenAndServe(":8082", nil); err != nil {
			g.Log().Error(context.Background(), "HTTP服务器启动失败:", err)
		}
	}()

	// 模拟群聊功能
	go func() {
		time.Sleep(2 * time.Second)

		// 创建聊天室
		roomId := "room1"

		// 模拟用户加入聊天室
		users := []string{"user1", "user2", "user3"}
		for _, uid := range users {
			// 这里需要先有客户端连接，然后绑定UID
			// 在实际应用中，客户端连接后会自动绑定UID
			g.Log().Infof(context.Background(), "用户 %s 加入聊天室 %s", uid, roomId)
		}

		// 发送群聊消息
		chatMessage := Message{
			Type:  "chat",
			Data:  "大家好！",
			From:  "user1",
			Group: roomId,
			Time:  time.Now().Unix(),
		}

		if data, err := json.Marshal(chatMessage); err == nil {
			count, err := server.SendToGroup(context.Background(), roomId, data)
			g.Log().Infof(context.Background(), "群聊消息发送: 成功=%d, 错误=%v", count, err)
		}

		// 获取聊天室信息
		clientCount := server.GetClientIdCountByGroup(roomId)
		uidCount := server.GetUidCountByGroup(roomId)
		g.Log().Infof(context.Background(), "聊天室 %s: 客户端数=%d, 用户数=%d", roomId, clientCount, uidCount)
	}()

	select {}
}

// ExampleCustomMessageHandler 自定义消息处理器示例
func ExampleCustomMessageHandler() {
	server := NewEnhancedWsServer()
	defer server.Close()

	// 设置自定义消息处理器
	server.messageHandler = &CustomMessageHandler{server: server}

	http.HandleFunc("/ws/custom", server.HandleWsConnection)

	go func() {
		g.Log().Info(context.Background(), "自定义处理器WebSocket服务器启动在 :8083")
		if err := http.ListenAndServe(":8083", nil); err != nil {
			g.Log().Error(context.Background(), "HTTP服务器启动失败:", err)
		}
	}()

	select {}
}

// CustomMessageHandler 自定义消息处理器
type CustomMessageHandler struct {
	server *EnhancedWsServer
}

// HandleMessage 自定义消息处理逻辑
func (cmh *CustomMessageHandler) HandleMessage(ctx context.Context, clientId string, message []byte) error {
	var msg Message
	if err := json.Unmarshal(message, &msg); err != nil {
		return err
	}

	// 自定义消息处理逻辑
	switch msg.Type {
	case "custom_event":
		// 处理自定义事件
		g.Log().Infof(ctx, "收到自定义事件: %s", msg.Data)

		// 发送响应
		response := Message{
			Type: "custom_response",
			Data: "Event processed successfully",
			Time: time.Now().Unix(),
		}

		if data, err := json.Marshal(response); err == nil {
			cmh.server.SendToClient(ctx, clientId, data)
		}

	case "echo":
		// 回声功能
		echoMsg := Message{
			Type: "echo_response",
			Data: msg.Data,
			Time: time.Now().Unix(),
		}

		if data, err := json.Marshal(echoMsg); err == nil {
			cmh.server.SendToClient(ctx, clientId, data)
		}

	default:
		// 调用默认处理器
		cmh.server.WsServer.handleMessage(ctx, clientId, message)
		return nil
	}

	return nil
}

// HandleError 自定义错误处理逻辑
func (cmh *CustomMessageHandler) HandleError(ctx context.Context, clientId string, err error) {
	g.Log().Errorf(ctx, "自定义错误处理 - 客户端 %s: %v", clientId, err)

	// 发送错误通知给客户端
	errorMsg := Message{
		Type: "error",
		Data: fmt.Sprintf("处理消息时发生错误: %v", err),
		Time: time.Now().Unix(),
	}

	if data, err := json.Marshal(errorMsg); err == nil {
		cmh.server.SendToClient(ctx, clientId, data)
	}
}

// ExamplePerformanceTest 性能测试示例
func ExamplePerformanceTest() {
	server := NewEnhancedWsServer()
	defer server.Close()

	http.HandleFunc("/ws/perf", server.HandleWsConnection)

	go func() {
		g.Log().Info(context.Background(), "性能测试WebSocket服务器启动在 :8084")
		if err := http.ListenAndServe(":8084", nil); err != nil {
			g.Log().Error(context.Background(), "HTTP服务器启动失败:", err)
		}
	}()

	// 性能监控
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				metrics := server.GetMetrics()
				stats := server.GetStats()

				g.Log().Infof(context.Background(),
					"性能监控 - 活跃连接: %d, 消息/秒: %.2f, 负载: %.2f%%, 错误率: %.2f%%",
					stats.ActiveConnections,
					float64(metrics.MessagesSent)/5.0, // 假设5秒间隔
					metrics.CurrentLoad*100,
					float64(metrics.ErrorCount)/float64(metrics.MessagesReceived)*100,
				)
			}
		}
	}()

	select {}
}

// ExampleLoadBalancing 负载均衡示例
func ExampleLoadBalancing() {
	// 创建多个WebSocket服务器实例
	servers := make([]*EnhancedWsServer, 3)

	for i := 0; i < 3; i++ {
		servers[i] = NewEnhancedWsServer()
		defer servers[i].Close()

		port := 8090 + i
		pattern := fmt.Sprintf("/ws/lb%d", i)

		http.HandleFunc(pattern, servers[i].HandleWsConnection)

		g.Log().Infof(context.Background(), "负载均衡服务器 %d 启动在 :%d", i, port)
	}

	// 启动HTTP服务器
	go func() {
		if err := http.ListenAndServe(":8090", nil); err != nil {
			g.Log().Error(context.Background(), "HTTP服务器启动失败:", err)
		}
	}()

	// 负载监控
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				for i, server := range servers {
					metrics := server.GetMetrics()
					stats := server.GetStats()

					g.Log().Infof(context.Background(),
						"服务器 %d - 连接: %d, 负载: %.2f%%, 消息: %d",
						i, stats.ActiveConnections, metrics.CurrentLoad*100, metrics.MessagesSent,
					)
				}
			}
		}
	}()

	select {}
}
