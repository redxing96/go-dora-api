/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-29 12:02:52
 * @LastEditTime: 2025-06-29 12:12:20
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/manager/manager_v1_manager_refresh_pass.go
 */
package manager

import (
	"context"

	v1 "go-dora-api/api/manager/v1"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
	"go-dora-api/utility/common"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/spf13/cast"
)

// 刷新管理员密码
func (c *ControllerV1) ManagerRefreshPass(ctx context.Context, req *v1.ManagerRefreshPassReq) (res *v1.ManagerRefreshPassRes, err error) {
	// 获取管理员ID
	managerId := cast.ToInt(ctx.Value("manager_id"))

	if managerId == req.Id {
		err = gerror.NewCode(common_return.ErrorCode("", "refresh_pass_failed", 500))
		logger.SystemLogger.Errorf("刷新管理员密码失败: %v", err)
		return
	}

	// 根据管理员ID获取管理员详情
	manager, err := service.Manage().GetManageDetailToId(ctx, &model.GetManageDetailInput{
		Id: managerId,
	})
	if err != nil {
		return
	}

	// 比较密码是否正确
	if !common.CompareHashAndPassword(manager.Password, req.Password) {
		// 密码错误，返回错误信息
		err = gerror.NewCode(common_return.ErrorCode("", "refresh_pass_failed", 500))
		logger.SystemLogger.Errorf("刷新管理员密码失败: %v", err)
		return
	}

	// 获取默认密码
	password, _ := g.Cfg().Get(ctx, "system.default_password", "Dora.123456")

	// 更新管理员密码
	result, err := service.Manage().Update(ctx, &model.UpdateInput{
		Id:       req.Id,
		Password: common.GenerateFromPassword(password.String()),
	})
	if err != nil {
		return
	}

	// 构造返回结果
	resp := v1.ManagerRefreshPassRes{
		Id: result.Id,
	}

	return &resp, nil
}
