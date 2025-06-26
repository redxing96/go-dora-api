// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysUserRole is the golang structure for table sys_user_role.
type SysUserRole struct {
	Id     int64 `json:"id"     orm:"id"      description:"ID"`   // ID
	UserId int64 `json:"userId" orm:"user_id" description:"用户ID"` // 用户ID
	RoleId int64 `json:"roleId" orm:"role_id" description:"角色ID"` // 角色ID
}
