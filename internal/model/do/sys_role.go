// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure of table sys_role for DAO operations like Where/Data.
type SysRole struct {
	g.Meta     `orm:"table:sys_role, do:true"`
	Id         interface{} // 角色ID
	RoleName   interface{} // 角色名称
	RoleCode   interface{} // 角色编码
	RoleDesc   interface{} // 角色描述
	Status     interface{} // 状态(1正常,0禁用)
	CreateTime *gtime.Time // 创建时间
}
