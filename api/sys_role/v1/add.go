/*
 * @Description: 添加角色
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 21:20:52
 * @LastEditTime: 2025-06-25 22:34:30
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/sys_role/v1/add.go
 */
package v1

import "github.com/gogf/gf/v2/frame/g"

type RoleAddReq struct {
	g.Meta `path:"/v1/role/add" method:"post" tags:"角色管理" summary:"添加角色" security:"api_key"`
	Name   string `v:"required|length:1,100" d:"" json:"name" dc:"角色名称"`
	Desc   string `d:"" json:"desc" dc:"备注"`
	Status int    `v:"required|integer" d:"1" json:"status" dc:"状态 1-正常 2-禁用"`
}

type RoleAddRes struct {
	Id int `json:"id" dc:"角色ID"`
}
