package sys_role

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

// 根据角色ID获取菜单ID
func (s *sSysRole) GetRoleMenuIDs(ctx context.Context, roleIDs []int64) (menuIds []int64, err error) {
	// 定义角色菜单列表和菜单ID列表
	var (
		roleMenuList []*entity.SysRoleMenu
	)
	// 根据角色ID查询角色菜单列表
	err = dao.SysRoleMenu.Ctx(ctx).WhereIn(dao.SysRoleMenu.Columns().RoleId, roleIDs).Scan(&roleMenuList)
	if err != nil {
		// 记录错误日志
		logger.SystemLogger.Errorf("获取角色菜单失败: %v", err)
		// 返回错误信息
		err = gerror.NewCode(common_return.ErrorCode("", "get_role_menu_failed", 500))
		return
	}
	// 遍历角色菜单列表，获取菜单ID
	for _, v := range roleMenuList {
		menuIds = append(menuIds, v.MenuId)
	}

	return
}
