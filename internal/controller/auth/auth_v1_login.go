/*
 * @Description: 管理端登录控制器
 * @Author: redxing96@163.com
 * @Date: 2025-06-24 18:55:48
 * @LastEditTime: 2025-06-26 17:35:52
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/auth/auth_v1_login.go
 */
package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/spf13/cast"

	v1 "go-dora-api/api/auth/v1"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
	"go-dora-api/utility/common"
	"go-dora-api/utility/jwt"
)

// Login函数用于处理登录请求
func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	// 调用service.Manage().GetManageDetailToAccount函数获取管理员详情
	detail, err := service.Manage().GetManageDetailToAccount(ctx, &model.GetManageDetailInput{
		Account: req.Account,
	})
	if err != nil {
		err = gerror.NewCode(common_return.ErrorCode("", "account_and_password_error", 500))
		return nil, err
	}

	// 比较密码
	if !common.CompareHashAndPassword(detail.Password, req.Password) {
		err = gerror.NewCode(common_return.ErrorCode("", "account_and_password_error", 500))
		return
	}

	// 如果管理员状态为2，则返回错误
	if detail.Status == 2 {
		err = gerror.NewCode(common_return.ErrorCode("", "manager_frozen", 500))
		return
	}

	// 生成token
	token, err := jwt.JWT.Builder(ctx, map[string]any{
		"manager_id": detail.Id,
		"is_super":   detail.IsSuper,
	})
	if err != nil {
		return nil, err
	}

	// 获取token中的exp字段
	expTime, err := jwt.JWT.GetPayloadField(token, "exp")
	if err != nil {
		return nil, err
	}

	err = service.PermissionCache().LoadUserPermissions(ctx, int64(detail.Id))
	if err != nil {
		logger.SystemLogger.Errorf("load user permissions error: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "load_user_permissions_error", 500))
		return nil, err
	}

	// 构造登录输出
	r := &model.LoginOutput{
		Token:    token,
		Exp:      cast.ToInt64(expTime),
		ManageId: detail.Id,
		Account:  detail.Account,
		IsSuper:  detail.IsSuper,
		Status:   detail.Status,
	}

	// 构造登录响应
	resp := v1.LoginRes(r)

	return &resp, nil
}
