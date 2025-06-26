/*
 * @Description: 更新角色
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 21:40:44
 * @LastEditTime: 2025-06-25 22:35:39
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/sys_role/sys_role_v1_role_update.go
 */
package sys_role

import (
	"context"

	v1 "go-dora-api/api/sys_role/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// RoleUpdate 更新角色信息
func (c *ControllerV1) RoleUpdate(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.RoleUpdateRes, err error) {
	// 调用service.SysRole().Update方法更新角色信息
	result, err := service.SysRole().Update(ctx, &model.SysRoleUpdateInput{
		Id:       req.Id,     // 角色ID
		RoleName: req.Name,   // 角色名称
		RoleDesc: req.Desc,   // 角色描述
		Status:   req.Status, // 角色状态
	})
	if err != nil {
		// 如果更新失败，返回错误
		return nil, err
	}

	// 构造返回结果
	resp := v1.RoleUpdateRes{
		Id: result.Id, // 角色ID
	}

	// 返回结果
	return &resp, nil
}
