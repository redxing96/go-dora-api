/*
 * @Description: 删除角色
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 21:40:44
 * @LastEditTime: 2025-06-25 22:34:01
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/sys_role/sys_role_v1_role_delete.go
 */
package sys_role

import (
	"context"

	v1 "go-dora-api/api/sys_role/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// RoleDelete 删除角色
func (c *ControllerV1) RoleDelete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error) {
	// 调用service.SysRole().Delete方法删除角色
	result, err := service.SysRole().Delete(ctx, &model.SysRoleDeleteInput{
		Ids: req.Ids,
	})
	if err != nil {
		// 如果删除失败，返回错误
		return nil, err
	}

	// 构造返回结果
	resp := v1.RoleDeleteRes{
		Ids: result.Ids,
	}

	// 返回结果
	return &resp, nil
}
