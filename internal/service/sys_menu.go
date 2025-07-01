// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"go-dora-api/internal/model"
	"go-dora-api/internal/model/entity"
)

type (
	ISysMenu interface {
		// 添加菜单
		Add(ctx context.Context, in *model.MenuAddInput) (res *model.MenuAddOutput, err error)
		// 删除路由权限
		Delete(ctx context.Context, in *model.MenuDeleteInput) (res *model.MenuDeleteOutput, err error)
		// 获取所有系统菜单
		GetAll(ctx context.Context, in *model.GetAllSysMenuInput) (res []*entity.SysMenu, total int, err error)
		// 获取管理员菜单
		GetManageMenu(ctx context.Context, managerID int64, menuType []int) (menuList []*entity.SysMenu, err error)
		// 将SysMenu列表转换为MenuItem列表
		ConvertToMenuItems(sysMenus []*entity.SysMenu) []*model.MenuItem
		// 构建菜单树
		BuildMenuTree(menus []*model.MenuItem, parentId int64) []*model.MenuItem
		// 对菜单项进行排序
		SortMenuItems(menus []*model.MenuItem)
		// 更新路由权限
		Update(ctx context.Context, in *model.MenuUpdateInput) (out *model.MenuUpdateOutput, err error)
	}
)

var (
	localSysMenu ISysMenu
)

func SysMenu() ISysMenu {
	if localSysMenu == nil {
		panic("implement not found for interface ISysMenu, forgot register?")
	}
	return localSysMenu
}

func RegisterSysMenu(i ISysMenu) {
	localSysMenu = i
}
