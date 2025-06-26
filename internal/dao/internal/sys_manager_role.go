// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysManagerRoleDao is the data access object for the table sys_manager_role.
type SysManagerRoleDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns SysManagerRoleColumns // columns contains all the column names of Table for convenient usage.
}

// SysManagerRoleColumns defines and stores column names for the table sys_manager_role.
type SysManagerRoleColumns struct {
	Id        string // ID
	ManagerId string // 用户ID
	RoleId    string // 角色ID
}

// sysManagerRoleColumns holds the columns for the table sys_manager_role.
var sysManagerRoleColumns = SysManagerRoleColumns{
	Id:        "id",
	ManagerId: "manager_id",
	RoleId:    "role_id",
}

// NewSysManagerRoleDao creates and returns a new DAO object for table data access.
func NewSysManagerRoleDao() *SysManagerRoleDao {
	return &SysManagerRoleDao{
		group:   "default",
		table:   "sys_manager_role",
		columns: sysManagerRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysManagerRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysManagerRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysManagerRoleDao) Columns() SysManagerRoleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysManagerRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysManagerRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysManagerRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
