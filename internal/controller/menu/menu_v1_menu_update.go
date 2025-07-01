/*
 * @Description: 更新路由权限
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 21:40:44
 * @LastEditTime: 2025-07-01 16:20:31
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/menu/menu_v1_menu_update.go
 */
package menu

import (
	"context"

	v1 "go-dora-api/api/menu/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// MenuUpdate 更新菜单
func (c *ControllerV1) MenuUpdate(ctx context.Context, req *v1.MenuUpdateReq) (res *v1.MenuUpdateRes, err error) {
	// 调用service.SysMenu().Update方法更新菜单
	result, err := service.SysMenu().Update(ctx, &model.MenuUpdateInput{
		Id:            req.Id,            // 菜单ID
		Pid:           req.Pid,           // 父菜单ID
		Type:          req.Type,          // 菜单类型
		Path:          req.Path,          // 菜单路径
		Sort:          req.Sort,          // 菜单排序
		Component:     req.Component,     // 菜单组件
		Title:         req.Title,         // 菜单标题
		Icon:          req.Icon,          // 菜单图标
		IconSvg:       req.IconSvg,       // 菜单图标svg
		IsHidden:      req.IsHidden,      // 是否隐藏
		IsKeepAlive:   req.IsKeepAlive,   // 是否缓存
		ActiveMenu:    req.ActiveMenu,    // 是否激活菜单
		IsLargeScreen: req.IsLargeScreen, // 是否大屏幕
		Link:          req.Link,          // 是否链接
		Remark:        req.Remark,        // 备注
		Status:        req.Status,        // 状态
	})

	// 如果更新失败，返回错误
	if err != nil {
		return nil, err
	}

	// 返回更新后的菜单ID
	return &v1.MenuUpdateRes{
		Id:   result.Id,
		Name: req.Name,
	}, nil
}
