// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure of table sys_menu for DAO operations like Where/Data.
type SysMenu struct {
	g.Meta        `orm:"table:sys_menu, do:true"`
	Id            interface{} // 菜单ID
	Pid           interface{} // 父菜单ID
	Type          interface{} // 权限类型(1菜单,2接口,3按钮)
	Path          interface{} // 路由地址
	Sort          interface{} // 排序
	Component     interface{} // 组件路径
	Title         interface{} // 菜单标题
	Icon          interface{} // 图标类名
	Hidden        interface{} // 是否隐藏(2-否,1-是)
	KeepAlive     interface{} // 是否缓存(2-否,1-是)
	ActiveMenu    interface{} // 激活菜单的path
	AlwaysShow    interface{} // 是否总是显示为父菜单(2-否,1-是)
	IsLargeScreen interface{} // 是否仅在大屏显示(2-否,1-是)
	IsFirstLevel  interface{} // 是否是一级导航(2-否,1-是)
	IsSecondLevel interface{} // 是否是二级导航(2-否,1-是)
	NoRedirect    interface{} // 是否禁止重定向(2-否,1-是)
	IsLink        interface{} // 是否是外部链接(2-否,1-是)
	Status        interface{} // 状态 1-正常 2-禁用
	CreateTime    *gtime.Time // 创建时间
	UpdateTime    *gtime.Time // 更新时间
	Remark        interface{} // 备注
}
