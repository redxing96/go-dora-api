/*
 * @Description: 删除角色
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 20:26:16
 * @LastEditTime: 2025-06-25 21:33:15
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/sys_role/v1/delete.go
 */
package v1

import "github.com/gogf/gf/v2/frame/g"

type RoleDeleteReq struct {
	g.Meta `path:"/v1/role/delete" method:"post" tags:"角色管理" summary:"删除角色" security:"api_key"`
	Ids    []int `v:"required|integer" json:"ids" dc:"角色ID"`
}

type RoleDeleteRes struct {
	Ids []int `json:"ids" dc:"角色ID"`
}
