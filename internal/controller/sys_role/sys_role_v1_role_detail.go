/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-07-01 19:58:04
 * @LastEditTime: 2025-07-01 20:05:23
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/sys_role/sys_role_v1_role_detail.go
 */
package sys_role

import (
	"context"

	v1 "go-dora-api/api/sys_role/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// 根据角色ID获取角色详情
func (c *ControllerV1) RoleDetail(ctx context.Context, req *v1.RoleDetailReq) (res *v1.RoleDetailRes, err error) {
	// 调用服务层获取角色信息
	result, err := service.SysRole().All(ctx, &model.SysRoleAllInput{
		ManagerId: req.Id,
	})
	// 如果获取失败或者结果为空，则返回错误
	if err != nil || len(result) <= 0 {
		return nil, err
	}

	// 创建角色详情结构体
	r := new(model.SysRoleDetailRes)
	// 设置角色ID
	r.Id = result[0].Id
	// 设置角色名称
	r.RoleName = result[0].RoleName
	// 设置角色编码
	r.RoleCode = result[0].RoleCode
	// 设置角色描述
	r.RoleDesc = result[0].RoleDesc
	// 设置角色状态
	r.Status = result[0].Status
	// 设置创建时间
	r.CreateTime = result[0].CreateTime

	// 获取管理菜单
	menuList, err := service.SysRole().GetMenu(ctx, []int64{int64(req.Id)}, []int{1, 2, 3, 4})
	if err != nil {
		return nil, err
	}

	// 转换为菜单项
	menuItems := service.SysMenu().ConvertToMenuItems(menuList)

	// 构建菜单树
	menuTree := service.SysMenu().BuildMenuTree(menuItems, 0)

	// 排序菜单项
	service.SysMenu().SortMenuItems(menuTree)

	// 设置菜单
	r.Menus = menuTree

	// 构建返回结果
	resp := v1.RoleDetailRes(r)

	return &resp, nil
}
