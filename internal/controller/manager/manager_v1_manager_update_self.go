/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-29 11:50:54
 * @LastEditTime: 2025-06-29 12:01:46
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/manager/manager_v1_manager_update_self.go
 */
package manager

import (
	"context"

	"github.com/spf13/cast"

	v1 "go-dora-api/api/manager/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// ManagerUpdateSelf 更新当前登录用户的信息
func (c *ControllerV1) ManagerUpdateSelf(ctx context.Context, req *v1.ManagerUpdateSelfReq) (res *v1.ManagerUpdateSelfRes, err error) {
	// 从上下文中获取当前登录用户的ID
	managerId := cast.ToInt(ctx.Value("manager_id"))
	// 调用service层的方法，更新用户信息
	result, err := service.Manage().Update(ctx, &model.UpdateInput{
		Id:       managerId,
		Account:  req.Account,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
		Avatar:   req.Avatar,
	})
	if err != nil {
		// 如果更新失败，返回错误
		return
	}

	// 根据管理员ID获取管理员详细信息
	resultDetail, err := service.Manage().GetManageDetailToId(ctx, &model.GetManageDetailInput{
		Id: managerId,
	})
	if err != nil {
		return
	}

	var roleIds []int
	// 根据管理员ID获取管理员角色信息
	role, err := service.SysRole().All(ctx, &model.SysRoleAllInput{
		ManagerId: managerId,
	})
	if err != nil {
		return
	}

	for _, item := range role {
		roleIds = append(roleIds, int(item.Id))
	}

	// 构造返回结果
	r := &model.ManagerDetailRes{
		Id:         result.Id,
		Account:    result.Account,
		Email:      result.Email,
		Phone:      result.Phone,
		Avatar:     result.Avatar,
		Status:     resultDetail.Status,
		IsSuper:    resultDetail.IsSuper,
		CreateTime: resultDetail.CreateTime,
		RoleIds:    roleIds,
		UpdateTime: resultDetail.UpdateTime,
	}

	// 构造返回的响应
	resp := v1.ManagerUpdateSelfRes(r)

	// 返回结果和错误
	return &resp, nil
}
