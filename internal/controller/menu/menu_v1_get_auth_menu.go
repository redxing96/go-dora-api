/*
 * @Description: 获取菜单权限接口（根据token获取）
 * @Author: redxing96@163.com
 * @Date: 2025-06-26 17:50:44
 * @LastEditTime: 2025-07-01 12:04:18
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/menu/menu_v1_get_auth_menu.go
 */
package menu

import (
	"context"

	"github.com/spf13/cast"

	v1 "go-dora-api/api/menu/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// GetAuthMenu 获取权限菜单
func (c *ControllerV1) GetAuthMenu(ctx context.Context, req *v1.GetAuthMenuReq) (res *v1.GetAuthMenuRes, err error) {

	isSuper := cast.ToInt(ctx.Value("is_super"))

	var menuItems []*model.MenuItem
	var btnList []string

	if isSuper == 1 {
		// 超级管理员获取所有菜单
		menuList, _, err := service.SysMenu().GetAll(ctx, &model.GetAllSysMenuInput{
			Type:   []int{1, 4},
			Status: 1,
		})
		if err != nil {
			return nil, err
		}

		// 转换为菜单项
		menuItems = service.SysMenu().ConvertToMenuItems(menuList)

		btns, _, err := service.SysMenu().GetAll(ctx, &model.GetAllSysMenuInput{
			Type:   []int{3},
			Status: 1,
		})
		if err != nil {
			return nil, err
		}

		for _, item := range btns {
			btnList = append(btnList, item.Component)
		}
	} else {
		// 获取管理菜单
		menuList, err := service.SysMenu().GetManageMenu(ctx, &model.GetManageMenuInput{
			ManagerID: cast.ToInt64(ctx.Value("manager_id")),
			MenuType:  []int{1, 4},
			Status:    1,
		})
		if err != nil {
			return nil, err
		}

		// 转换为菜单项
		menuItems = service.SysMenu().ConvertToMenuItems(menuList)

		btns, err := service.SysMenu().GetManageMenu(ctx, &model.GetManageMenuInput{
			ManagerID: cast.ToInt64(ctx.Value("manager_id")),
			MenuType:  []int{3},
			Status:    1,
		})
		if err != nil {
			return nil, err
		}

		for _, item := range btns {
			btnList = append(btnList, item.Component)
		}
	}

	// 构建菜单树
	menuTree := service.SysMenu().BuildMenuTree(menuItems, 0)

	// 排序菜单项
	service.SysMenu().SortMenuItems(menuTree)

	// 构建响应
	resp := v1.GetAuthMenuRes{
		MenuList: menuTree,
		BtnList:  btnList,
	}

	return &resp, nil
}
