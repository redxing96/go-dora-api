/*
 * @Description: 添加路由权限
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 11:23:37
 * @LastEditTime: 2025-07-01 17:05:36
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/menu/v1/add.go
 */
package v1

import (
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type MenuAddReq struct {
	g.Meta        `path:"/v1/menu/add" method:"post" tags:"路由权限管理" summary:"添加路由权限" security:"api_key"`
	Pid           int    `v:"required|integer" d:"0" json:"pid" dc:"父级ID"`
	Name          string `v:"required|length:1,100" d:"" json:"name" dc:"名称"`
	Type          int    `v:"required|integer" d:"0" json:"type" dc:"类型 1菜单,2接口,3按钮,4目录"`
	Path          string `d:"" json:"path" dc:"路径"`
	Sort          int    `v:"required|integer" d:"0" json:"sort" dc:"排序"`
	Component     string `d:"" json:"component" dc:"组件"`
	Title         string `v:"required|length:1,100" d:"" json:"title" dc:"标题"`
	Icon          string `d:"" json:"icon" dc:"图标"`
	IconSvg       string `d:"" json:"icon_svg" dc:"图标svg"`
	IsHidden      int    `v:"required|integer" d:"0" json:"is_hidden" dc:"是否隐藏 2:否 1:是"`
	IsKeepAlive   int    `d:"0" json:"is_keep_alive" dc:"是否缓存 2:否 1:是"`
	ActiveMenu    string `d:"" json:"active_menu" dc:"激活菜单"`
	IsLargeScreen int    `d:"0" json:"is_large_screen" dc:"是否大屏 2:否 1:是"`
	Link          string `d:"0" json:"link" dc:"外部连接"`
	Remark        string `d:"" json:"remark" dc:"备注"`
	Status        int    `v:"required|integer" d:"0" json:"status" dc:"状态 1-正常 2-禁用"`
}

type MenuAddRes *model.MenuAddResponse
