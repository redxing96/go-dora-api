/*
 * @Description: 中间件
 * @Author: redxing96@163.com
 * @Date: 2025-06-23 16:37:26
 * @LastEditTime: 2025-07-01 18:24:07
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/middleware/middleware.go
 */
package middleware

import (
	"context"
	"fmt"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/consts"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/service"
	"net/http"
	"strings"

	logger2 "go-dora-api/utility/logger"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"

	common_return2 "go-dora-api/utility/common_return"
)

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

// 全局异常捕获
func (s *sMiddleware) MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		// 捕获全局异常
		logger.SystemLogger.Errorf("捕获全局错误: %s", err)
	}
}

// 跨域中间件
func (s *sMiddleware) MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// 访问日志中间件
func (s *sMiddleware) MiddlewareAccessLog(r *ghttp.Request) {
	// 获取开始请求的时间
	startTime := gtime.TimestampMilli()
	r.Middleware.Next()

	var (
		scheme       = "http"
		proto        = r.Header.Get("X-Forwarded-Proto")
		endTime      = gtime.TimestampMilli()
		err          = r.GetError()
		stackContent = ""
	)

	if r.TLS != nil || gstr.Equal(proto, "https") {
		scheme = "https"
	}
	if stack := gerror.Stack(err); stack != "" {
		stackContent = stack
	}

	// 日志写入文件
	logger.SystemLogger.Info("route",
		logger2.LogField{Key: "status", Value: r.Response.Status},
		logger2.LogField{Key: "method", Value: r.Method},
		logger2.LogField{Key: "scheme", Value: scheme},
		logger2.LogField{Key: "host", Value: r.Host},
		logger2.LogField{Key: "url", Value: r.URL.String()},
		logger2.LogField{Key: "proto", Value: r.Proto},
		logger2.LogField{Key: "duration", Value: float64(endTime-startTime) / 1000},
		logger2.LogField{Key: "ip", Value: r.GetClientIp()},
		logger2.LogField{Key: "referer", Value: r.Referer()},
		logger2.LogField{Key: "userAgent", Value: r.UserAgent()},
		logger2.LogField{Key: "stack", Value: stackContent},
	)

	// ====== 操作日志写入数据库 ======
	// 获取用户信息（如有登录）
	// userId := int64(0)
	// username := "游客"
	// if v := r.Session.MustGet("userId"); v != nil {
	// 	userId = v.Int64()
	// }
	// if v := r.Session.MustGet("username"); v != nil {
	// 	username = v.String()
	// }

	// // 操作类型
	// operation := r.Method
	// // 操作模块（可用路由前缀或首段）
	// module := ""
	// urlPath := r.URL.Path
	// if urlPath != "" {
	// 	parts := gstr.SplitAndTrim(urlPath, "/")
	// 	if len(parts) > 1 {
	// 		module = parts[1]
	// 	} else if len(parts) == 1 {
	// 		module = parts[0]
	// 	}
	// }
	// // 操作描述
	// description := r.Method + " " + r.URL.String()
	// // 请求参数
	// params := r.GetMap()
	// // 响应内容
	// response := r.Response.BufferString()
	// // 状态
	// status := 1
	// if r.Response.Status >= 400 {
	// 	status = 2
	// }
	// // 错误信息
	// errMsg := ""
	// if err != nil {
	// 	errMsg = err.Error()
	// }
	// // 执行时间
	// executionTime := int(endTime - startTime)

	// // 组装请求信息
	// requestInfo := map[string]interface{}{
	// 	"method":    r.Method,
	// 	"url":       r.URL.String(),
	// 	"params":    params,
	// 	"response":  response,
	// 	"ip":        r.GetClientIp(),
	// 	"userAgent": r.UserAgent(),
	// }

	// 写入操作日志表
	// _ = service.OperationLog().RecordOperation(
	// 	r.GetCtx(),
	// 	userId,
	// 	username,
	// 	operation,
	// 	module,
	// 	description,
	// 	status,
	// 	errMsg,
	// 	executionTime,
	// 	requestInfo,
	// )
}

// 全局响应中间件
func (s *sMiddleware) MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()
	// 有自定义缓冲区内容，然后退出当前处理程序。
	if r.Response.BufferLength() > 0 {
		return
	}
	// 获取错误码
	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
		lang = r.Header.Get("lang")
	)

	if strings.ToUpper(lang) == "ZH-CN" {
		lang = consts.LANG_CN
	}

	//多语言切换
	switch lang {
	case consts.LANG_CN, consts.LANG_EN:
	default:
		lang = service.Translate().GetDefaultLanguage()
	}

	// 错误处理
	if err != nil {
		msg = err.Error()
		if code == gcode.CodeNil {
			code = common_return.Error("", "internal_error")
			msg = code.Message()
		}
	} else {
		// 状态码处理
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			switch r.Response.Status {
			case http.StatusNotFound:
				code = common_return.Error("", "not_found")
				msg = code.Message()
			case http.StatusForbidden:
				code = common_return.Error("", "not_authorized")
				msg = code.Message()
			default:
				code = common_return.Error("", "unknown_error")
				msg = code.Message()
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = common_return.Success(res)
			msg = code.Message()
		}
	}

	// 错误处理
	var errorDetail interface{}
	errorKey := ""
	var errorCode int
	errorMsg := msg

	// 错误信息处理
	errorKey, errorCode, _ = common_return.SplitErrorMessage(msg)
	if bussDetail, ok := code.Detail().(common_return2.BussDetail); ok {
		errorDetail = bussDetail.Detail
	}

	// 错误信息处理
	if errorCode != -1 {
		ctx := gi18n.WithLanguage(context.TODO(), lang)
		errorText := gi18n.Translate(ctx, errorKey)
		_, errorCode, errorMsg = common_return.SplitErrorMessage(errorText)
		if bussDetail, ok := code.Detail().(common_return2.BussDetail); ok {
			errorMsg = fmt.Sprintf(errorMsg, bussDetail.Args...)
			errorDetail = bussDetail.Detail
		}
	}

	if errorCode == -1 {
		errorCode = code.Code()
	}
	if errorDetail == nil {
		errorDetail = ""
	}
	fmt.Println("响应编码：", errorCode)
	// 返回响应
	r.Response.WriteJson(common_return.DefaultJsonResponse(errorCode, errorMsg, errorDetail))
}
