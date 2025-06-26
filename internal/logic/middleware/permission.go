/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-26 13:54:09
 * @LastEditTime: 2025-06-26 17:32:53
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/middleware/permission.go
 */
package middleware

import (
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// MiddlewarePermission 函数用于检查用户是否有权限访问某个路径
func (s *sMiddleware) MiddlewarePermission(r *ghttp.Request) {

	// 获取请求的路径
	path := r.Request.URL.Path

	// 检查用户是否有权限访问该路径
	if !service.PermissionCache().CheckManagerPermission(r.Context(), path) {
		logger.SystemLogger.Error("no permission")
		r.SetError(gerror.NewCode(common_return.ErrorCode("", "no_permission", 401)))
		r.Exit()
		return
	}

	// 如果有权限，则继续执行下一个中间件
	r.Middleware.Next()
}
