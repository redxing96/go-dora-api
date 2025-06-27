/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-27 11:31:32
 * @LastEditTime: 2025-06-27 11:41:52
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/websocket/websocket.go
 */
package websocket

import (
	"go-dora-api/internal/service"
	"go-dora-api/utility/websocket"
	"net/http"
)

type sWebSocket struct {
	server *websocket.EnhancedWsServer
}

// 创建一个新的sWebSocket实例
func New() *sWebSocket {
	// 创建一个新的websocket服务器
	server := websocket.NewEnhancedWsServer()
	// 返回一个新的sWebSocket实例，包含创建的websocket服务器
	return &sWebSocket{server: server}
}

func init() {
	service.RegisterWebSocket(New())
}

// 处理WebSocket连接
func (s *sWebSocket) HandleWsConnection(w http.ResponseWriter, r *http.Request) {
	// 调用server的HandleWsConnection方法处理WebSocket连接
	s.server.HandleWsConnection(w, r)
}

// 获取sWebSocket的server
func (s *sWebSocket) GetServer() *websocket.EnhancedWsServer {
	// 返回sWebSocket的server
	return s.server
}
