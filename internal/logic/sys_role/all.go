/*
 * @Description: 获取所有角色
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 22:28:25
 * @LastEditTime: 2025-06-29 13:14:31
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/sys_role/all.go
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

// 获取所有角色
func (s *sSysRole) All(ctx context.Context, in *model.SysRoleAllInput) (out []*entity.SysRole, err error) {
	// 创建查询对象
	query := dao.SysRole.Ctx(ctx)
	// 如果传入的角色状态不为0，则添加查询条件
	if in.Status != 0 {
		query = query.Where(dao.SysRole.Columns().Status, in.Status)
	}
	if in.ManagerId != 0 {
		var roleList []*entity.SysManagerRole
		// 根据管理员ID查询管理员角色
		err = dao.SysManagerRole.Ctx(ctx).Where(dao.SysManagerRole.Columns().ManagerId, in.ManagerId).Scan(&roleList)
		if err != nil {
			// 记录错误日志
			logger.SystemLogger.Errorf("获取管理员角色失败: %v", err)
			// 返回错误信息
			err = gerror.NewCode(common_return.ErrorCode("", "get_manager_role_failed", 500))
			return
		}
		var roleIds []int64
		for _, role := range roleList {
			roleIds = append(roleIds, role.RoleId)
		}

		if len(roleIds) > 0 {
			// 如果查询到管理员角色，则添加查询条件
			query = query.WhereIn(dao.SysRole.Columns().Id, roleIds)
		}
	}
	// 执行查询，并将结果扫描到out中
	err = query.Scan(&out)
	// 如果查询失败，则记录错误日志，并返回错误信息
	if err != nil {
		logger.SystemLogger.Errorf("获取所有角色失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_role_all_failed", 500))
		return
	}
	// 返回查询结果
	return
}
