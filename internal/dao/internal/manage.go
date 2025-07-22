// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ManageDao is the data access object for the table manage.
type ManageDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of the current DAO.
	columns ManageColumns // columns contains all the column names of Table for convenient usage.
}

// ManageColumns defines and stores column names for the table manage.
type ManageColumns struct {
	Id         string //
	Account    string // 用户账号
	Password   string // 用户密码
	Email      string // 电子邮箱
	Phone      string // 电话号码
	Avatar     string // 头像
	Status     string // 状态 0-初始化 1-正常 2-冻结
	IsSuper    string // 是否超管 1-是
	IsDelete   string // 是否删除 1-是 2-否
	CreateTime string // 创建时间
	UpdateTime string // 修改时间
}

// manageColumns holds the columns for the table manage.
var manageColumns = ManageColumns{
	Id:         "id",
	Account:    "account",
	Password:   "password",
	Email:      "email",
	Phone:      "phone",
	Avatar:     "avatar",
	Status:     "status",
	IsSuper:    "is_super",
	IsDelete:   "is_delete",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewManageDao creates and returns a new DAO object for table data access.
func NewManageDao() *ManageDao {
	return &ManageDao{
		group:   "default",
		table:   "manage",
		columns: manageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ManageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ManageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ManageDao) Columns() ManageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ManageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ManageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ManageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
