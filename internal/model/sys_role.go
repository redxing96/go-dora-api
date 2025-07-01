/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 21:58:14
 * @LastEditTime: 2025-07-01 20:02:04
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/model/sys_role.go
 */
package model

import (
	"go-dora-api/internal/model/entity"
)

type SysRoleListInput struct {
	BaseInput
}

type SysRoleListRes struct {
	BaseOutput
	List []*entity.SysRole
}

type SysRoleAddInput struct {
	RoleName string `json:"role_name" dc:"角色名称"`
	RoleDesc string `json:"role_desc" dc:"角色描述"`
	RoleCode string `json:"role_code" dc:"角色编码"`
	Status   int    `json:"status" dc:"状态"`
	MenuIds  []int  `json:"menu_ids" dc:"菜单IDs"`
}

type SysRoleAddOutput struct {
	Id int `json:"id" dc:"角色ID"`
}

type SysRoleDeleteInput struct {
	Ids []int `json:"ids" dc:"角色ID"`
}

type SysRoleDeleteOutput struct {
	Ids []int `json:"id" dc:"角色ID"`
}

type SysRoleUpdateInput struct {
	Id       int    `json:"id" dc:"角色ID"`
	RoleName string `json:"role_name" dc:"角色名称"`
	RoleCode string `json:"role_code" dc:"角色编码"`
	MenuIds  []int  `json:"menu_ids" dc:"菜单IDs"`
	RoleDesc string `json:"role_desc" dc:"角色描述"`
	Status   int    `json:"status" dc:"状态"`
}

type SysRoleUpdateOutput struct {
	Id int `json:"id" dc:"角色ID"`
}

type SysRoleAllInput struct {
	Status    int `json:"status" dc:"状态 1-正常 2-禁用"`
	ManagerId int `json:"manager_id" dc:"管理员ID"`
}

type SysRoleAllRes struct {
	List []*entity.SysRole `json:"list" dc:"角色列表"`
}

type SysRoleDetailRes struct {
	entity.SysRole
	Menus []*MenuItem `json:"menus" dc:"菜单列表"`
}
