// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysManagerRole is the golang structure for table sys_manager_role.
type SysManagerRole struct {
	Id        int64 `json:"id"        orm:"id"         description:"ID"`   // ID
	ManagerId int64 `json:"managerId" orm:"manager_id" description:"用户ID"` // 用户ID
	RoleId    int64 `json:"roleId"    orm:"role_id"    description:"角色ID"` // 角色ID
}
