// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DictionaryDao is the data access object for the table dictionary.
type DictionaryDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns DictionaryColumns // columns contains all the column names of Table for convenient usage.
}

// DictionaryColumns defines and stores column names for the table dictionary.
type DictionaryColumns struct {
	Id         string //
	Type       string // 类型
	Name       string // 名称
	Code       string // 编码
	Desc       string // 描述
	IsEdit     string // 是否可编辑 1-是 2-否
	Status     string // 状态 1-正常 2-禁用
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// dictionaryColumns holds the columns for the table dictionary.
var dictionaryColumns = DictionaryColumns{
	Id:         "id",
	Type:       "type",
	Name:       "name",
	Code:       "code",
	Desc:       "desc",
	IsEdit:     "is_edit",
	Status:     "status",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewDictionaryDao creates and returns a new DAO object for table data access.
func NewDictionaryDao() *DictionaryDao {
	return &DictionaryDao{
		group:   "default",
		table:   "dictionary",
		columns: dictionaryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DictionaryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DictionaryDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DictionaryDao) Columns() DictionaryColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DictionaryDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DictionaryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *DictionaryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
