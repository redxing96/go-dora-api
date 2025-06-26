/*
 * @Description: 菜单工具函数
 * @Author: redxing96@163.com
 * @Date: 2025-06-26 17:52:53
 * @LastEditTime: 2025-06-26 17:54:00
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/sys_menu/tools.go
 */
package sys_menu

import (
	"go-dora-api/internal/model"
	"go-dora-api/internal/model/entity"
)

// 将SysMenu列表转换为MenuItem列表
func (s *sSysMenu) ConvertToMenuItems(sysMenus []*entity.SysMenu) []*model.MenuItem {
	menuItems := make([]*model.MenuItem, 0, len(sysMenus))
	for _, sysMenu := range sysMenus {
		menuItem := &model.MenuItem{
			Id:        sysMenu.Id,
			Pid:       sysMenu.Pid,
			Path:      sysMenu.Path,
			Name:      sysMenu.Title,
			Sort:      sysMenu.Sort,
			Status:    sysMenu.Status,
			Component: sysMenu.Component,
			Meta: model.MenuItemMeta{
				Title:         sysMenu.Title,
				Icon:          sysMenu.Icon,
				Hidden:        sysMenu.Hidden == 1,
				KeepAlive:     sysMenu.KeepAlive == 1,
				ActiveMenu:    sysMenu.ActiveMenu,
				AlwaysShow:    sysMenu.AlwaysShow == 1,
				IsLargeScreen: sysMenu.IsLargeScreen == 1,
				IsFirstLevel:  sysMenu.IsFirstLevel == 1,
				IsSecondLevel: sysMenu.IsSecondLevel == 1,
				NoRedirect:    sysMenu.NoRedirect == 1,
				IsLink:        sysMenu.IsLink,
			},
		}
		menuItems = append(menuItems, menuItem)
	}
	return menuItems
}

// 构建菜单树
func (s *sSysMenu) BuildMenuTree(menus []*model.MenuItem, parentId int64) []*model.MenuItem {
	// 定义菜单树
	var menuTree []*model.MenuItem
	// 遍历菜单
	for _, menu := range menus {
		// 如果菜单的父id等于传入的父id
		if menu.Pid == parentId {
			// 递归构建子菜单
			children := s.BuildMenuTree(menus, menu.Id)
			// 如果子菜单不为空
			if len(children) > 0 {
				// 将子菜单赋值给当前菜单的Children字段
				menu.Children = children
			}
			// 将当前菜单添加到菜单树中
			menuTree = append(menuTree, menu)
		}
	}
	// 按sort排序
	s.SortMenuItems(menuTree)
	// 返回菜单树
	return menuTree
}

// 对菜单项进行排序
func (s *sSysMenu) SortMenuItems(menus []*model.MenuItem) {
	// 如果菜单项数量小于等于1，则无需排序
	if len(menus) <= 1 {
		return
	}
	// 外层循环控制比较的轮数
	for i := 0; i < len(menus)-1; i++ {
		// 内层循环控制每一轮的比较次数
		for j := 0; j < len(menus)-i-1; j++ {
			// 如果当前菜单项的排序值大于下一个菜单项的排序值，则交换位置
			if menus[j].Sort > menus[j+1].Sort {
				menus[j], menus[j+1] = menus[j+1], menus[j]
			}
		}
	}
}
