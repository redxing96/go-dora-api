/*
 * @Description: 获取所有角色
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 22:39:48
 * @LastEditTime: 2025-07-01 20:47:31
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/sys_role/sys_role_v1_role_all.go
 */
package sys_role

import (
	"context"

	v1 "go-dora-api/api/sys_role/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// RoleAll 函数用于获取所有角色信息
func (c *ControllerV1) RoleAll(ctx context.Context, req *v1.RoleAllReq) (res *v1.RoleAllRes, err error) {
	// 调用 service.SysRole().All 函数获取所有角色信息
	result, err := service.SysRole().All(ctx, &model.SysRoleAllInput{
		Status: req.Status,
	})
	if err != nil {
		// 如果出现错误，则返回错误信息
		return nil, err
	}

	// 创建 SysRoleAllRes 结构体，并将获取到的角色信息赋值给 List 字段
	r := &model.SysRoleAllRes{}

	for _, item := range result {
		r.List = append(r.List, &model.SysRoleItemOutput{
			Id:         item.Id,
			RoleName:   item.RoleName,
			RoleCode:   item.RoleCode,
			RoleDesc:   item.RoleDesc,
			Status:     item.Status,
			CreateTime: item.CreateTime,
		})
	}

	// 将 SysRoleAllRes 结构体转换为 v1.RoleAllRes 结构体
	resp := v1.RoleAllRes(r)

	// 返回转换后的 v1.RoleAllRes 结构体和 nil 错误信息
	return &resp, nil
}
