/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 10:57:51
 * @LastEditTime: 2025-06-26 17:55:28
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/menu/menu_v1_menu_all_list.go
 */
package menu

import (
	"context"

	v1 "go-dora-api/api/menu/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// MenuAllList 获取所有菜单列表
func (c *ControllerV1) MenuAllList(ctx context.Context, req *v1.MenuAllListReq) (res *v1.MenuAllListRes, err error) {
	// 创建获取所有菜单的输入参数
	search := new(model.GetAllSysMenuInput)
	search.Type = 1

	// 调用服务获取所有菜单列表
	list, _, err := service.SysMenu().GetAll(ctx, search)
	if err != nil {
		return nil, err
	}

	// 转换为MenuItem结构
	menuItems := service.SysMenu().ConvertToMenuItems(list)

	// 构建菜单树
	menuTree := service.SysMenu().BuildMenuTree(menuItems, 0)

	// 对菜单树进行排序
	service.SysMenu().SortMenuItems(menuTree)

	// 创建返回结果
	r := new(model.MenuAllResponse)
	r.List = menuTree

	// 创建返回结果
	resp := v1.MenuAllListRes(r)

	return &resp, nil
}
