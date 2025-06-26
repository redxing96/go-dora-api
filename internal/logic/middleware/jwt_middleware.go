/*
 * @Description: jwt中间件
 * @Author: redxing96@163.com
 * @Date: 2025-06-24 13:31:27
 * @LastEditTime: 2025-06-26 17:29:47
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/middleware/jwt_middleware.go
 */
package middleware

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/consts"
	"go-dora-api/internal/logger"
	"go-dora-api/utility/jwt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 中间件函数，用于验证客户端的JWT
func (s *sMiddleware) MiddlewareClientJWT(r *ghttp.Request) {
	// 获取请求头中的token
	token := r.GetHeader(consts.CLIENT_TOKEN_FIELD)
	if token == "" {
		// 如果token为空，则获取请求参数中是否包含
		token = r.GetQuery(consts.CLIENT_TOKEN_FIELD).String()
		if token == "" {
			// 如果token为空，则返回错误
			logger.SystemLogger.Error("client token is required")
			r.SetError(gerror.NewCode(common_return.ErrorCode("", "token_empty", 401)))
			r.Exit()
			return
		}
	}

	// 验证token
	data, err := jwt.JWT.Auth(context.TODO(), token)
	if err != nil {
		// 如果token无效，则返回错误
		logger.SystemLogger.Error("client token is invalid")
		r.SetError(gerror.NewCode(common_return.ErrorCode("", "token_invalid", 401)))
		r.Exit()
		return
	}

	// 将token存入上下文
	r.SetCtxVar("client_token", token)

	// 如果token中包含client_id，则将其存入上下文
	if id, ok := data["client_id"]; ok {
		r.SetCtxVar("client_id", id)
	}

	// 执行下一个中间件
	r.Middleware.Next()
}

// 中间件函数，用于验证管理端JWT
func (s *sMiddleware) MiddlewareManageJWT(r *ghttp.Request) {
	// 获取请求头中的token
	token := r.GetHeader(consts.MANAGE_TOKEN_FIELD)
	if token == "" {
		// 如果token为空，则获取请求参数中是否包含
		token = r.GetQuery(consts.MANAGE_TOKEN_FIELD).String()
		if token == "" {
			// 如果token为空，则返回错误
			logger.SystemLogger.Error("client token is required")
			r.SetError(gerror.NewCode(common_return.ErrorCode("", "token_empty", 401)))
			r.Exit()
			return
		}
	}

	// 验证token
	data, err := jwt.JWT.Auth(context.TODO(), token)
	if err != nil {
		// 如果token无效，则返回错误
		logger.SystemLogger.Error("manage token is invalid")
		r.SetError(gerror.NewCode(common_return.ErrorCode("", "token_invalid", 401)))
		r.Exit()
		return
	}

	// 将token存入上下文
	r.SetCtxVar("manage_token", token)

	// 如果token中包含manager_id，则将其存入上下文
	if id, ok := data["manager_id"]; ok {
		r.SetCtxVar("manager_id", id)
	}

	// 如果token中包含is_super，则将其存入上下文
	if isSuper, ok := data["is_super"]; ok {
		r.SetCtxVar("is_super", isSuper)
	}

	// 执行下一个中间件
	r.Middleware.Next()
}
