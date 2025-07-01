/*
 * @Description: 限流中间件
 * @Author: redxing96@163.com
 * @Date: 2025-07-01 09:39:09
 * @LastEditTime: 2025-07-01 09:44:34
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/middleware/rate_limit_middleware.go
 */
package middleware

import (
	"go-dora-api/internal/common_return"
	"net/http"
	"sync"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"golang.org/x/time/rate"
)

// 限流器
var limiterMap = make(map[string]*rate.Limiter)

// 互斥锁
var mu sync.Mutex

// 获取限流器
// 根据IP地址获取限流器
func getLimiter(ip string) *rate.Limiter {
	// 加锁
	mu.Lock()
	// 在函数结束时解锁
	defer mu.Unlock()

	// 从限流器映射中获取IP地址对应的限流器
	limiter, ok := limiterMap[ip]
	// 如果没有找到对应的限流器
	if !ok {
		// 创建一个新的限流器
		limiter = rate.NewLimiter(rate.Limit(10), 10)
		// 将新的限流器添加到限流器映射中
		limiterMap[ip] = limiter
	}
	// 返回限流器
	return limiter
}

// 限流中间件
// RateLimitMiddleware 函数用于限制请求的速率
func (s *sMiddleware) RateLimitMiddleware(r *ghttp.Request) {
	// 获取客户端的IP地址
	ip := r.GetClientIp()
	// 获取IP地址对应的限流器
	limiter := getLimiter(ip)

	// 如果限流器不允许请求，则返回错误
	if !limiter.Allow() {
		// 创建错误信息
		err := gerror.NewCode(common_return.ErrorCode("", "rate_limit_exceeded", 429))
		// 设置错误信息
		r.SetError(err)
		// 返回429状态码
		r.Response.WriteStatus(http.StatusTooManyRequests)
		return
	}

	// 继续执行下一个中间件
	r.Middleware.Next()
}
