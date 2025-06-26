/*
 * @Description: jwt
 * @Author: redxing96@163.com
 * @Date: 2025-06-24 13:48:55
 * @LastEditTime: 2025-06-24 20:41:30
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/utility/jwt/jwt.go
 */
package jwt

import (
	"context"
	"sync"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
)

type sJWT struct{}

var (
	JWT         = &sJWT{}
	blackTokens = sync.Map{}
)

// secret 获取jwt密钥
func (s *sJWT) secret() string {
	return g.Cfg().MustGet(context.TODO(), "jwt.secret").String()
}

// ttl 获取jwt过期时间
func (s *sJWT) ttl() time.Duration {
	return time.Duration(g.Cfg().MustGet(context.TODO(), "jwt.ttl").Int()) * time.Second
}

// refreshTTL 获取jwt刷新时间
func (s *sJWT) refreshTTL() time.Duration {
	return time.Duration(g.Cfg().MustGet(context.TODO(), "jwt.refresh_ttl").Int()) * time.Minute
}

// Builder函数用于生成一个JWT字符串
func (s *sJWT) Builder(ctx context.Context, payload map[string]any) (string, error) {
	// 创建一个空的jwt.MapClaims对象
	claims := jwt.MapClaims{}
	// 遍历payload中的键值对，将它们添加到claims中
	for k, v := range payload {
		claims[k] = v
	}
	// 设置过期时间为当前时间加上ttl()
	claims["exp"] = time.Now().Add(s.ttl()).Unix()

	// 使用HS256算法和claims生成一个新的JWT对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用secret()函数返回的密钥对JWT进行签名，并返回签名后的JWT字符串
	return token.SignedString([]byte(s.secret()))
}

// 验证 token，并返回 payload map
// Auth 方法用于验证 token 是否有效
func (s *sJWT) Auth(ctx context.Context, tokenStr string) (map[string]any, error) {
	// 检查 token 是否在黑名单中
	if s.isBlacklisted(tokenStr) {
		return nil, gerror.New("token is blacklisted")
	}

	// 解析 token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// 返回 secret
		return []byte(s.secret()), nil
	})
	if err != nil || !token.Valid {
		// 如果解析失败或者 token 无效，则返回错误
		return nil, gerror.Wrap(err, "token invalid")
	}

	// 将 token.Claims 转换为 jwt.MapClaims 类型
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		// 如果转换失败，则返回错误
		return nil, gerror.New("invalid claims type")
	}

	// 过期检查：如果过期，则立即加入黑名单
	expUnix := gconv.Int64(claims["exp"])
	if time.Now().Unix() > expUnix {
		s.Invalidate(tokenStr)
		return nil, gerror.New("token expired and now blacklisted")
	}

	// 返回 token.Claims
	return claims, nil
}

// 刷新 token（旧的加入黑名单）
// Refresh函数用于刷新JWT令牌
func (s *sJWT) Refresh(ctx context.Context, oldToken string) (string, error) {
	// 使用Auth函数验证旧令牌
	payload, err := s.Auth(ctx, oldToken)
	if err != nil {
		// 如果验证失败，返回错误
		return "", err
	}
	// 使旧令牌失效
	s.Invalidate(oldToken)
	// 使用Builder函数生成新的令牌
	return s.Builder(ctx, payload)
}

// 加入黑名单
// Invalidate方法用于使指定的token失效
func (s *sJWT) Invalidate(token string) {
	// 将token存储到blackTokens中，并设置失效时间为当前时间加上refreshTTL
	blackTokens.Store(token, time.Now().Add(s.refreshTTL()))
}

// 是否在黑名单
// 判断token是否在黑名单中
func (s *sJWT) isBlacklisted(token string) bool {
	// 从黑名单中加载token
	v, ok := blackTokens.Load(token)
	if !ok {
		// 如果token不在黑名单中，返回false
		return false
	}
	// 如果token在黑名单中，判断token是否过期
	if expireAt, ok := v.(time.Time); ok && expireAt.After(time.Now()) {
		// 如果token未过期，返回true
		return true
	}
	// 如果token已过期，从黑名单中删除token
	blackTokens.Delete(token)
	// 返回false
	return false
}

// 从请求中获取token
func (s *sJWT) GetTokenFromRequest(ctx context.Context) string {
	// 从上下文中获取请求
	r := g.RequestFromCtx(ctx)
	// 从请求头中获取Authorization字段，即token
	return r.Header.Get("Authorization") // Bearer xxx
}

// 解析JWT并获取指定字段的值
func (s *sJWT) GetPayloadField(tokenStr, field string) (interface{}, error) {
	// 解析JWT
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// 返回密钥
		return []byte(s.secret()), nil
	})
	// 如果解析失败或JWT无效，则返回错误
	if err != nil || !token.Valid {
		return nil, err
	}

	// 将JWT的claims转换为MapClaims类型
	claims, ok := token.Claims.(jwt.MapClaims)
	// 如果转换失败，则返回错误
	if !ok {
		return nil, gerror.New("invalid claims")
	}

	return claims[field], nil
}
