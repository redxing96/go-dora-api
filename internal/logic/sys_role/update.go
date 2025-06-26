/*
 * @Description: 更新角色
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 22:25:31
 * @LastEditTime: 2025-06-25 22:28:10
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/sys_role/update.go
 */
package sys_role

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"
	"go-dora-api/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
)

// 更新角色信息
func (s *sSysRole) Update(ctx context.Context, in *model.SysRoleUpdateInput) (out *model.SysRoleUpdateOutput, err error) {
	// 根据角色ID更新角色信息
	result, err := dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().Id, in.Id).Data(do.SysRole{
		RoleName: in.RoleName,
		RoleDesc: in.RoleDesc,
		Status:   in.Status,
	}).Update()
	if err != nil {
		// 记录错误日志
		logger.SystemLogger.Errorf("更新角色失败: %v", err)
		// 返回错误信息
		err = gerror.NewCode(common_return.ErrorCode("", "update_role_failed", 500))
		return
	}

	// 获取受影响的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		// 记录错误日志
		logger.SystemLogger.Errorf("更新角色失败: %v", err)
		// 返回错误信息
		err = gerror.NewCode(common_return.ErrorCode("", "update_role_failed", 500))
		return
	}

	// 构造返回结果
	out = &model.SysRoleUpdateOutput{
		Id: in.Id,
	}

	return
}
