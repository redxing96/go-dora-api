// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// 中间件函数，用于验证客户端的JWT
		MiddlewareClientJWT(r *ghttp.Request)
		// 中间件函数，用于验证管理端JWT
		MiddlewareManageJWT(r *ghttp.Request)
		// 全局异常捕获
		MiddlewareErrorHandler(r *ghttp.Request)
		// 跨域中间件
		MiddlewareCORS(r *ghttp.Request)
		// 访问日志中间件
		MiddlewareAccessLog(r *ghttp.Request)
		// 全局响应中间件
		MiddlewareHandlerResponse(r *ghttp.Request)
		// MiddlewarePermission 函数用于检查用户是否有权限访问某个路径
		MiddlewarePermission(r *ghttp.Request)
		// 限流中间件
		// RateLimitMiddleware 函数用于限制请求的速率
		RateLimitMiddleware(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
