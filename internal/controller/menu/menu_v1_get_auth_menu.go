/*
 * @Description: 获取菜单权限接口（根据token获取）
 * @Author: redxing96@163.com
 * @Date: 2025-06-26 17:50:44
 * @LastEditTime: 2025-06-26 18:00:55
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/menu/menu_v1_get_auth_menu.go
 */
package menu

import (
	"context"

	"github.com/spf13/cast"

	v1 "go-dora-api/api/menu/v1"
	"go-dora-api/internal/service"
)

// GetAuthMenu 获取权限菜单
func (c *ControllerV1) GetAuthMenu(ctx context.Context, req *v1.GetAuthMenuReq) (res *v1.GetAuthMenuRes, err error) {

	// 获取管理菜单
	menuList, err := service.SysMenu().GetManageMenu(ctx, cast.ToInt64(ctx.Value("manager_id")), 1)
	if err != nil {
		return nil, err
	}

	// 转换为菜单项
	menuItems := service.SysMenu().ConvertToMenuItems(menuList)

	// 构建菜单树
	menuTree := service.SysMenu().BuildMenuTree(menuItems, 0)

	// 排序菜单项
	service.SysMenu().SortMenuItems(menuTree)

	// 构建响应
	resp := v1.GetAuthMenuRes{
		MenuList: menuTree,
	}

	return &resp, nil
}
