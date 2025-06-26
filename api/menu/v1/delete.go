/*
 * @Description: 删除路由权限
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 20:26:16
 * @LastEditTime: 2025-06-25 20:39:34
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/menu/v1/delete.go
 */
package v1

import "github.com/gogf/gf/v2/frame/g"

type MenuDeleteReq struct {
	g.Meta `path:"/v1/menu/delete" method:"post" tags:"路由权限管理" summary:"删除路由权限" security:"api_key"`
	Ids    []int `v:"required|integer" json:"ids" dc:"路由权限ID"`
}

type MenuDeleteRes struct {
	Ids []int `json:"ids" dc:"路由权限ID"`
}
