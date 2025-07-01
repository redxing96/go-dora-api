package sys_role

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model/entity"
	"go-dora-api/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sSysRole) GetMenu(ctx context.Context, roleIDs []int64, menuType []int) (menuList []*entity.SysMenu, err error) {
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
