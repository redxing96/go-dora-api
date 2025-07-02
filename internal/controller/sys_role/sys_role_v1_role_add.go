/*
 * @Description: 添加角色
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 22:34:34
 * @LastEditTime: 2025-07-01 20:55:59
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/sys_role/sys_role_v1_role_add.go
 */
package sys_role

import (
	"context"

	v1 "go-dora-api/api/sys_role/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

func (c *ControllerV1) RoleAdd(ctx context.Context, req *v1.RoleAddReq) (res *v1.RoleAddRes, err error) {
	// 调用 service.SysRole().Add 方法添加角色
	result, err := service.SysRole().Add(ctx, &model.SysRoleAddInput{
		RoleName: req.RoleName, // 角色名称
		RoleDesc: req.RoleDesc, // 角色描述
		RoleCode: req.RoleCode, // 角色编码
		Status:   req.Status,   // 角色状态
		MenuIds:  req.MenuIds,  // 菜单IDs
	})
	if err != nil {
		// 如果添加角色失败，则返回错误
		return nil, err
	}

	// 构造返回结果
	resp := v1.RoleAddRes{
		Id: result.Id, // 角色ID
	}

	// 返回结果
	return &resp, nil
}
