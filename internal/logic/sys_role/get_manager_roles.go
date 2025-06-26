/*
 * @Description: 获取管理员角色
 * @Author: redxing96@163.com
 * @Date: 2025-06-26 13:15:46
 * @LastEditTime: 2025-06-26 13:30:19
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/sys_role/get_manager_roles.go
 */
package sys_role

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

// 获取管理员角色ID
func (s *sSysRole) GetManagerRoleIDs(ctx context.Context, managerID int64) (roleIDs []int64, err error) {
	// 定义一个切片，用于存储管理员角色
	var (
		list []*entity.SysManagerRole
	)
	// 从数据库中查询管理员角色
	err = dao.SysManagerRole.Ctx(ctx).Where(dao.SysManagerRole.Columns().ManagerId, managerID).Scan(&list)
	// 如果查询失败，则记录错误日志，并返回错误信息
	if err != nil {
		logger.SystemLogger.Errorf("获取管理员角色失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_manager_role_failed", 500))
		return
	}
	// 遍历查询结果，将角色ID添加到切片中
	for _, v := range list {
		roleIDs = append(roleIDs, v.RoleId)
	}

	return
}
