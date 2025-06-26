/*
 * @Description: 获取所有菜单
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 10:51:16
 * @LastEditTime: 2025-06-25 20:45:48
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/menu/v1/get_all.go
 */
package v1

import (
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type MenuAllListReq struct {
	g.Meta `path:"/v1/menu/all" method:"get" tags:"路由权限管理" summary:"获取全部菜单权限列表(全部菜单权限)" security:"api_key"`
	IsTree int `d:"0" json:"is_tree" dc:"是否树形结构 0:否 1:是"`
}

type MenuAllListRes *model.MenuAllResponse
