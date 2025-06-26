/*
 * @Description: 管理员退出登录
 * @Author: redxing96@163.com
 * @Date: 2025-06-26 17:17:46
 * @LastEditTime: 2025-06-26 17:47:31
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/auth/auth_v1_logout.go
 */
package auth

import (
	"context"

	"github.com/spf13/cast"

	v1 "go-dora-api/api/auth/v1"
	"go-dora-api/internal/service"
	"go-dora-api/utility/jwt"
)

// Logout函数用于用户登出
func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	// 从上下文中获取manage_token
	token := ctx.Value("manage_token")

	// 使用JWT库将token失效
	jwt.JWT.Invalidate(cast.ToString(token))

	// 移除用户权限
	service.PermissionCache().RemoveManagerPermissions(ctx, cast.ToInt64(ctx.Value("manager_id")))

	// 创建一个空的LogoutRes对象
	resp := v1.LogoutRes{}

	// 返回LogoutRes对象和nil错误
	return &resp, nil
}
