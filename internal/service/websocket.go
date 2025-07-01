// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"go-dora-api/utility/websocket"
	"net/http"
)

type (
	IWebSocket interface {
		// 处理WebSocket连接
		HandleWsConnection(w http.ResponseWriter, r *http.Request)
		// 获取sWebSocket的server
		GetServer() *websocket.EnhancedWsServer
	}
)

var (
	localWebSocket IWebSocket
)

func WebSocket() IWebSocket {
	if localWebSocket == nil {
		panic("implement not found for interface IWebSocket, forgot register?")
	}
	return localWebSocket
}

func RegisterWebSocket(i IWebSocket) {
	localWebSocket = i
}
