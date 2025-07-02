/*
 * @Description: 根据角色ID获取角色详情
 * @Author: redxing96@163.com
 * @Date: 2025-07-01 21:07:34
 * @LastEditTime: 2025-07-01 21:08:50
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/sys_role/detail.go
 */
package sys_role

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"
	"go-dora-api/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sSysRole) Detail(ctx context.Context, in *model.SysRoleDetailInput) (out *model.SysRoleDetailRes, err error) {
	query := dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().Id, in.Id)
	var role []*entity.SysRole
	err = query.Scan(&role)
	if err != nil {
		logger.SystemLogger.Errorf("获取角色详情失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_role_detail_failed", 500))
		return
	}
	if len(role) <= 0 {
		err = gerror.NewCode(common_return.ErrorCode("", "role_not_found", 404))
		logger.SystemLogger.Errorf("角色不存在: %v", err)
		return
	}

	out = &model.SysRoleDetailRes{
		Id:         role[0].Id,
		RoleName:   role[0].RoleName,
		RoleCode:   role[0].RoleCode,
		RoleDesc:   role[0].RoleDesc,
		Status:     role[0].Status,
		CreateTime: role[0].CreateTime,
	}

	return out, nil
}
