// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure for table sys_menu.
type SysMenu struct {
	Id            int64       `json:"id"            orm:"id"              description:"菜单ID"`                // 菜单ID
	Pid           int64       `json:"pid"           orm:"pid"             description:"父菜单ID"`               // 父菜单ID
	Type          int         `json:"type"          orm:"type"            description:"权限类型(1菜单,2接口,3按钮)"`   // 权限类型(1菜单,2接口,3按钮)
	Path          string      `json:"path"          orm:"path"            description:"路由地址"`                // 路由地址
	Sort          int         `json:"sort"          orm:"sort"            description:"排序"`                  // 排序
	Component     string      `json:"component"     orm:"component"       description:"组件路径"`                // 组件路径
	Title         string      `json:"title"         orm:"title"           description:"菜单标题"`                // 菜单标题
	Icon          string      `json:"icon"          orm:"icon"            description:"图标类名"`                // 图标类名
	Hidden        int         `json:"hidden"        orm:"hidden"          description:"是否隐藏(2-否,1-是)"`       // 是否隐藏(2-否,1-是)
	KeepAlive     int         `json:"keepAlive"     orm:"keep_alive"      description:"是否缓存(2-否,1-是)"`       // 是否缓存(2-否,1-是)
	ActiveMenu    string      `json:"activeMenu"    orm:"active_menu"     description:"激活菜单的path"`           // 激活菜单的path
	AlwaysShow    int         `json:"alwaysShow"    orm:"always_show"     description:"是否总是显示为父菜单(2-否,1-是)"` // 是否总是显示为父菜单(2-否,1-是)
	IsLargeScreen int         `json:"isLargeScreen" orm:"is_large_screen" description:"是否仅在大屏显示(2-否,1-是)"`   // 是否仅在大屏显示(2-否,1-是)
	IsFirstLevel  int         `json:"isFirstLevel"  orm:"is_first_level"  description:"是否是一级导航(2-否,1-是)"`    // 是否是一级导航(2-否,1-是)
	IsSecondLevel int         `json:"isSecondLevel" orm:"is_second_level" description:"是否是二级导航(2-否,1-是)"`    // 是否是二级导航(2-否,1-是)
	NoRedirect    int         `json:"noRedirect"    orm:"no_redirect"     description:"是否禁止重定向(2-否,1-是)"`    // 是否禁止重定向(2-否,1-是)
	IsLink        string      `json:"isLink"        orm:"is_link"         description:"是否是外部链接(2-否,1-是)"`    // 是否是外部链接(2-否,1-是)
	Status        int         `json:"status"        orm:"status"          description:"状态 1-正常 2-禁用"`        // 状态 1-正常 2-禁用
	CreateTime    *gtime.Time `json:"createTime"    orm:"create_time"     description:"创建时间"`                // 创建时间
	UpdateTime    *gtime.Time `json:"updateTime"    orm:"update_time"     description:"更新时间"`                // 更新时间
	Remark        string      `json:"remark"        orm:"remark"          description:"备注"`                  // 备注
}
