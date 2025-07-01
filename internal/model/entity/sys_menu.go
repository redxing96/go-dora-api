// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure for table sys_menu.
type SysMenu struct {
	Id            int64       `json:"id"            orm:"id"              description:"菜单ID"`                  // 菜单ID
	Pid           int64       `json:"pid"           orm:"pid"             description:"父菜单ID"`                 // 父菜单ID
	Type          int         `json:"type"          orm:"type"            description:"权限类型(1菜单,2接口,3按钮,4目录)"` // 权限类型(1菜单,2接口,3按钮,4目录)
	Path          string      `json:"path"          orm:"path"            description:"路由地址"`                  // 路由地址
	Sort          int         `json:"sort"          orm:"sort"            description:"排序"`                    // 排序
	Component     string      `json:"component"     orm:"component"       description:"组件路径"`                  // 组件路径
	Name          string      `json:"name"          orm:"name"            description:"菜单名称"`                  // 菜单名称
	Title         string      `json:"title"         orm:"title"           description:"菜单标题"`                  // 菜单标题
	Icon          string      `json:"icon"          orm:"icon"            description:"图标类名"`                  // 图标类名
	IconSvg       string      `json:"iconSvg"       orm:"icon_svg"        description:"svg图标"`                 // svg图标
	IsHidden      int         `json:"isHidden"      orm:"is_hidden"       description:"是否隐藏(2-否,1-是)"`         // 是否隐藏(2-否,1-是)
	IsKeepAlive   int         `json:"isKeepAlive"   orm:"is_keep_alive"   description:"是否缓存(2-否,1-是)"`         // 是否缓存(2-否,1-是)
	ActiveMenu    string      `json:"activeMenu"    orm:"active_menu"     description:"激活菜单的path"`             // 激活菜单的path
	IsLargeScreen int         `json:"isLargeScreen" orm:"is_large_screen" description:"是否仅在大屏显示(2-否,1-是)"`     // 是否仅在大屏显示(2-否,1-是)
	Link          string      `json:"link"          orm:"link"            description:"外部链接"`                  // 外部链接
	Status        int         `json:"status"        orm:"status"          description:"状态(1正常,2禁用)"`           // 状态(1正常,2禁用)
	CreateTime    *gtime.Time `json:"createTime"    orm:"create_time"     description:"创建时间"`                  // 创建时间
	UpdateTime    *gtime.Time `json:"updateTime"    orm:"update_time"     description:"更新时间"`                  // 更新时间
	Remark        string      `json:"remark"        orm:"remark"          description:"备注"`                    // 备注
}
