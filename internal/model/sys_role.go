/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 21:58:14
 * @LastEditTime: 2025-07-01 21:54:49
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/model/sys_role.go
 */
package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type SysRoleListInput struct {
	BaseInput
}

type SysRoleListRes struct {
	BaseOutput
	List []*SysRoleItemOutput `json:"list" dc:"列表"`
}

type SysRoleItemOutput struct {
	Id         int64       `json:"id" dc:"角色ID"`
	RoleName   string      `json:"role_name" dc:"角色名称"`
	RoleCode   string      `json:"role_code" dc:"角色编码"`
	RoleDesc   string      `json:"role_desc" dc:"角色描述"`
	Status     int         `json:"status" dc:"状态"`
	CreateTime *gtime.Time `json:"create_time" dc:"创建时间"`
}

type SysRoleAddInput struct {
	RoleName string  `json:"role_name" dc:"角色名称"`
	RoleDesc string  `json:"role_desc" dc:"角色描述"`
	RoleCode string  `json:"role_code" dc:"角色编码"`
	Status   int     `json:"status" dc:"状态"`
	MenuIds  []int64 `json:"menu_ids" dc:"菜单IDs"`
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
	Id       int     `json:"id" dc:"角色ID"`
	RoleName string  `json:"role_name" dc:"角色名称"`
	RoleCode string  `json:"role_code" dc:"角色编码"`
	MenuIds  []int64 `json:"menu_ids" dc:"菜单IDs"`
	RoleDesc string  `json:"role_desc" dc:"角色描述"`
	Status   int     `json:"status" dc:"状态"`
}

type SysRoleUpdateOutput struct {
	Id int `json:"id" dc:"角色ID"`
}

type SysRoleAllInput struct {
	Status    int `json:"status" dc:"状态 1-正常 2-禁用"`
	ManagerId int `json:"manager_id" dc:"管理员ID"`
}

type SysRoleAllRes struct {
	List []*SysRoleItemOutput `json:"list" dc:"角色列表"`
}

type SysRoleDetailRes struct {
	Id         int64       `json:"id" dc:"角色ID"`
	RoleName   string      `json:"role_name" dc:"角色名称"`
	RoleCode   string      `json:"role_code" dc:"角色编码"`
	RoleDesc   string      `json:"role_desc" dc:"角色描述"`
	Status     int         `json:"status" dc:"状态"`
	CreateTime *gtime.Time `json:"create_time" dc:"创建时间"`
	MenuIds    []int64     `json:"menu_ids" dc:"菜单IDs"`
}

type SysRoleDetailInput struct {
	Id int `json:"id" dc:"角色ID"`
}
