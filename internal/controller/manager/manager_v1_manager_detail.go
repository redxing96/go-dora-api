/*
 * @Description: 管理员详情
 * @Author: redxing96@163.com
 * @Date: 2025-06-29 11:28:25
 * @LastEditTime: 2025-06-29 11:50:37
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/manager/manager_v1_manager_detail.go
 */
package manager

import (
	"context"

	v1 "go-dora-api/api/manager/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

func (c *ControllerV1) ManagerDetail(ctx context.Context, req *v1.ManagerDetailReq) (res *v1.ManagerDetailRes, err error) {
	// 从上下文中获取管理员ID
	managerId := req.Id

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
		IsSuper:    result.IsSuper,
		Email:      result.Email,
		Phone:      result.Phone,
		Avatar:     result.Avatar,
		CreateTime: result.CreateTime,
		UpdateTime: result.UpdateTime,
		Role:       role,
	}

	// 将管理员详细信息转换为响应
	resp := v1.ManagerDetailRes(r)

	// 返回响应和错误信息
	return &resp, nil
}
