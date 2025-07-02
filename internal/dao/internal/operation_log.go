// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OperationLogDao is the data access object for the table operation_log.
type OperationLogDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns OperationLogColumns // columns contains all the column names of Table for convenient usage.
}

// OperationLogColumns defines and stores column names for the table operation_log.
type OperationLogColumns struct {
	Id            string //
	UserId        string // 操作用户ID
	Username      string // 操作用户名
	Operation     string // 操作类型
	Module        string // 操作模块
	Description   string // 操作描述
	RequestMethod string // 请求方法
	RequestUrl    string // 请求URL
	RequestParams string // 请求参数
	ResponseData  string // 响应数据
	IpAddress     string // IP地址
	UserAgent     string // 用户代理
	Status        string // 操作状态 1-成功 2-失败
	ErrorMessage  string // 错误信息
	ExecutionTime string // 执行时间(毫秒)
	CreateTime    string // 创建时间
}

// operationLogColumns holds the columns for the table operation_log.
var operationLogColumns = OperationLogColumns{
	Id:            "id",
	UserId:        "user_id",
	Username:      "username",
	Operation:     "operation",
	Module:        "module",
	Description:   "description",
	RequestMethod: "request_method",
	RequestUrl:    "request_url",
	RequestParams: "request_params",
	ResponseData:  "response_data",
	IpAddress:     "ip_address",
	UserAgent:     "user_agent",
	Status:        "status",
	ErrorMessage:  "error_message",
	ExecutionTime: "execution_time",
	CreateTime:    "create_time",
}

// NewOperationLogDao creates and returns a new DAO object for table data access.
func NewOperationLogDao() *OperationLogDao {
	return &OperationLogDao{
		group:   "default",
		table:   "operation_log",
		columns: operationLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OperationLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OperationLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OperationLogDao) Columns() OperationLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OperationLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OperationLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *OperationLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
