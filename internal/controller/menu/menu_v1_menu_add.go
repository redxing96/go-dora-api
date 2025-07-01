/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 12:06:47
 * @LastEditTime: 2025-07-01 16:20:20
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/menu/menu_v1_menu_add.go
 */
package menu

import (
	"context"

	v1 "go-dora-api/api/menu/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// MenuAdd 函数用于添加菜单
func (c *ControllerV1) MenuAdd(ctx context.Context, req *v1.MenuAddReq) (res *v1.MenuAddRes, err error) {
	// 调用 service.SysMenu().Add 方法添加菜单
	result, err := service.SysMenu().Add(ctx, &model.MenuAddInput{
		Pid:           req.Pid,           // 父级ID
		Name:          req.Name,          // 名称
		Type:          req.Type,          // 类型
		Path:          req.Path,          // 路径
		Sort:          req.Sort,          // 排序
		Component:     req.Component,     // 组件
		Title:         req.Title,         // 标题
		Icon:          req.Icon,          // 图标
		IconSvg:       req.IconSvg,       // 图标svg
		IsHidden:      req.IsHidden,      // 是否隐藏
		IsKeepAlive:   req.IsKeepAlive,   // 是否缓存
		ActiveMenu:    req.ActiveMenu,    // 激活菜单
		IsLargeScreen: req.IsLargeScreen, // 是否大屏幕
		Link:          req.Link,          // 是否链接
		Remark:        req.Remark,        // 备注
		Status:        req.Status,        // 状态
	})
	if err != nil {
		return nil, err
	}

	// 创建 MenuAddResponse 对象
	r := &model.MenuAddResponse{
		Id:   result.Id,
		Name: req.Name,
	}
	// 创建 MenuAddRes 对象
	resp := v1.MenuAddRes(r)

	return &resp, nil
}
