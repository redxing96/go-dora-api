/*
 * @Description: 删除角色
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 22:22:32
 * @LastEditTime: 2025-07-01 20:09:27
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/sys_role/delete.go
 */
package sys_role

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/errors/gerror"
)

// 删除角色
func (s *sSysRole) Delete(ctx context.Context, in *model.SysRoleDeleteInput) (out *model.SysRoleDeleteOutput, err error) {
	// 根据传入的角色ID列表，删除对应的角色
	result, err := dao.SysRole.Ctx(ctx).WhereIn(dao.SysRole.Columns().Id, in.Ids).Delete()
	if err != nil {
		// 如果删除失败，记录错误日志，并返回错误信息
		logger.SystemLogger.Errorf("删除角色失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "delete_role_failed", 500))
		return
	}
	// 获取受影响的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		// 如果受影响的行数为0，记录错误日志，并返回错误信息
		logger.SystemLogger.Errorf("删除角色失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "delete_role_failed", 500))
		return
	}

	// 删除角色菜单
	dao.SysRoleMenu.Ctx(ctx).WhereIn(dao.SysRoleMenu.Columns().RoleId, in.Ids).Delete()

	// 构造返回结果
	out = &model.SysRoleDeleteOutput{
		Ids: in.Ids,
	}
	return
}
