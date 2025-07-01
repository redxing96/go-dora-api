// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Id         int64       `json:"id"         orm:"id"          description:"角色ID"`        // 角色ID
	RoleName   string      `json:"roleName"   orm:"role_name"   description:"角色名称"`        // 角色名称
	RoleCode   string      `json:"roleCode"   orm:"role_code"   description:"角色编码"`        // 角色编码
	RoleDesc   string      `json:"roleDesc"   orm:"role_desc"   description:"角色描述"`        // 角色描述
	Status     int         `json:"status"     orm:"status"      description:"状态(1正常,0禁用)"` // 状态(1正常,0禁用)
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`        // 创建时间
}
