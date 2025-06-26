// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysMenuDao is the data access object for the table sys_menu.
type SysMenuDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns SysMenuColumns // columns contains all the column names of Table for convenient usage.
}

// SysMenuColumns defines and stores column names for the table sys_menu.
type SysMenuColumns struct {
	Id            string // 菜单ID
	Pid           string // 父菜单ID
	Type          string // 权限类型(1菜单,2接口,3按钮)
	Path          string // 路由地址
	Sort          string // 排序
	Component     string // 组件路径
	Title         string // 菜单标题
	Icon          string // 图标类名
	Hidden        string // 是否隐藏(2-否,1-是)
	KeepAlive     string // 是否缓存(2-否,1-是)
	ActiveMenu    string // 激活菜单的path
	AlwaysShow    string // 是否总是显示为父菜单(2-否,1-是)
	IsLargeScreen string // 是否仅在大屏显示(2-否,1-是)
	IsFirstLevel  string // 是否是一级导航(2-否,1-是)
	IsSecondLevel string // 是否是二级导航(2-否,1-是)
	NoRedirect    string // 是否禁止重定向(2-否,1-是)
	IsLink        string // 是否是外部链接(2-否,1-是)
	Status        string // 状态 1-正常 2-禁用
	CreateTime    string // 创建时间
	UpdateTime    string // 更新时间
	Remark        string // 备注
}

// sysMenuColumns holds the columns for the table sys_menu.
var sysMenuColumns = SysMenuColumns{
	Id:            "id",
	Pid:           "pid",
	Type:          "type",
	Path:          "path",
	Sort:          "sort",
	Component:     "component",
	Title:         "title",
	Icon:          "icon",
	Hidden:        "hidden",
	KeepAlive:     "keep_alive",
	ActiveMenu:    "active_menu",
	AlwaysShow:    "always_show",
	IsLargeScreen: "is_large_screen",
	IsFirstLevel:  "is_first_level",
	IsSecondLevel: "is_second_level",
	NoRedirect:    "no_redirect",
	IsLink:        "is_link",
	Status:        "status",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
	Remark:        "remark",
}

// NewSysMenuDao creates and returns a new DAO object for table data access.
func NewSysMenuDao() *SysMenuDao {
	return &SysMenuDao{
		group:   "default",
		table:   "sys_menu",
		columns: sysMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysMenuDao) Columns() SysMenuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
