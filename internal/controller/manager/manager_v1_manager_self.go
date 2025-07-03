/*
 * @Description: 获取自己的详情信息
 * @Author: redxing96@163.com
 * @Date: 2025-06-29 11:28:25
 * @LastEditTime: 2025-06-29 11:50:45
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/manager/manager_v1_manager_self.go
 */
package manager

import (
	"context"

	"github.com/spf13/cast"

	v1 "go-dora-api/api/manager/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// ManagerSelf 方法用于获取当前管理员的详细信息
func (c *ControllerV1) ManagerSelf(ctx context.Context, req *v1.ManagerSelfReq) (res *v1.ManagerSelfRes, err error) {
	// 从上下文中获取管理员ID
	managerId := cast.ToInt(ctx.Value("manager_id"))

	// 根据管理员ID获取管理员详细信息
	result, err := service.Manage().GetManageDetailToId(ctx, &model.GetManageDetailInput{
		Id: managerId,
	})
	if err != nil {
		return
	}

	// 根据管理员ID获取管理员角色信息
	role, err := service.SysRole().All(ctx, &model.SysRoleAllInput{
		ManagerId: managerId,
	})
	if err != nil {
		return
	}

	// 构造管理员详细信息响应
	r := &model.ManagerDetailRes{
		Id:         result.Id,
		Account:    result.Account,
		Status:     result.Status,
		Email:      result.Email,
		Phone:      result.Phone,
		Avatar:     result.Avatar,
		IsSuper:    result.IsSuper,
		CreateTime: result.CreateTime,
		UpdateTime: result.UpdateTime,
	}

	for _, item := range role {
		r.RoleIds = append(r.RoleIds, int(item.Id))
	}

	// 将管理员详细信息转换为响应
	resp := v1.ManagerSelfRes(r)

	// 返回响应和错误信息
	return &resp, nil
}
