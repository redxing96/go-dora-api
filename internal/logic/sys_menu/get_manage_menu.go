/*
 * @Description: 获取管理员菜单
 * @Author: redxing96@163.com
 * @Date: 2025-06-26 13:10:32
 * @LastEditTime: 2025-07-01 12:03:34
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/sys_menu/get_manage_menu.go
 */
package sys_menu

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model/entity"
	"go-dora-api/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

// 获取管理员菜单
func (s *sSysMenu) GetManageMenu(ctx context.Context, managerID int64, menuType []int) (menuList []*entity.SysMenu, err error) {
	// 根据用户ID获取到所有的角色IDs
	roleIDs, err := service.SysRole().GetManagerRoleIDs(ctx, managerID)
	if err != nil {
		logger.SystemLogger.Errorf("获取管理员角色失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_manager_role_failed", 500))
		return
	}
	// 根据角色IDs获取到所有的菜单IDs
	menuIDs, err := service.SysRole().GetRoleMenuIDs(ctx, roleIDs)
	if err != nil {
		logger.SystemLogger.Errorf("获取角色菜单失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_role_menu_failed", 500))
		return
	}
	// 获取菜单
	query := dao.SysMenu.Ctx(ctx).WhereIn(dao.SysMenu.Columns().Id, menuIDs)

	if len(menuType) > 0 {
		query = query.WhereIn(dao.SysMenu.Columns().Type, menuType)
	}

	err = query.Scan(&menuList)
	if err != nil {
		logger.SystemLogger.Errorf("获取菜单失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_menu_failed", 500))
		return
	}

	return
}
