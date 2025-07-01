/*
 * @Description: 更新路由权限
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 21:36:08
 * @LastEditTime: 2025-07-01 18:17:58
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/menu/v1/update.go
 */
package v1

import "github.com/gogf/gf/v2/frame/g"

type MenuUpdateReq struct {
	g.Meta        `path:"/v1/menu/update" method:"post" tags:"路由权限管理" summary:"更新路由权限" security:"api_key"`
	Id            int    `v:"required|integer" d:"0" json:"id" dc:"ID"`
	Pid           int    `v:"required|integer" d:"0" json:"pid" dc:"父级ID"`
	Type          int    `v:"required|integer" d:"0" json:"type" dc:"类型 1:菜单 2:接口"`
	Path          string `d:"" json:"path" dc:"路径"`
	Sort          int    `v:"required|integer" d:"0" json:"sort" dc:"排序"`
	Component     string `d:"" json:"component" dc:"组件"`
	Title         string `v:"required|length:1,100" d:"" json:"title" dc:"标题"`
	Name          string `v:"required|length:1,100" d:"" json:"name" dc:"名称"`
	Icon          string `d:"" json:"icon" dc:"图标"`
	IconSvg       string `d:"" json:"icon_svg" dc:"图标svg"`
	IsHidden      int    `d:"2" json:"is_hidden" dc:"是否隐藏 2:否 1:是"`
	IsKeepAlive   int    `d:"2" json:"is_keep_alive" dc:"是否缓存 2:否 1:是"`
	ActiveMenu    string `d:"" json:"active_menu" dc:"激活菜单"`
	IsLargeScreen int    `d:"2" json:"is_large_screen" dc:"是否大屏 2:否 1:是"`
	Link          int    `d:"" json:"link" dc:"是否链接 2:否 1:是"`
	Remark        string `d:"" json:"remark" dc:"备注"`
	Status        int    `v:"required|integer" d:"0" json:"status" dc:"状态 1-正常 2-禁用"`
}

type MenuUpdateRes struct {
	Id   int    `json:"id" dc:"角色ID"`
	Name string `json:"name" dc:"名称"`
}
