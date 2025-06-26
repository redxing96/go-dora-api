/*
 * @Description: 更新角色
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 21:39:59
 * @LastEditTime: 2025-06-25 21:40:07
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/sys_role/v1/update.go
 */
package v1

import "github.com/gogf/gf/v2/frame/g"

type RoleUpdateReq struct {
	g.Meta `path:"/v1/role/update" method:"post" tags:"角色管理" summary:"更新角色" security:"api_key"`
	Id     int    `v:"required|integer" d:"0" json:"id" dc:"ID"`
	Name   string `v:"required|length:1,100" d:"" json:"name" dc:"角色名称"`
	Desc   string `d:"" json:"desc" dc:"备注"`
	Status int    `v:"required|integer" d:"1" json:"status" dc:"状态 1-正常 2-禁用"`
}

type RoleUpdateRes struct {
	Id int `json:"id" dc:"角色ID"`
}
