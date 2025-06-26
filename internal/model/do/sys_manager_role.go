// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysManagerRole is the golang structure of table sys_manager_role for DAO operations like Where/Data.
type SysManagerRole struct {
	g.Meta    `orm:"table:sys_manager_role, do:true"`
	Id        interface{} // ID
	ManagerId interface{} // 用户ID
	RoleId    interface{} // 角色ID
}
