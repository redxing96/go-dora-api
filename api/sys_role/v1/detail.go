/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-07-01 19:56:28
 * @LastEditTime: 2025-07-01 21:11:05
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/sys_role/v1/detail.go
 */
package v1

import (
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type RoleDetailReq struct {
	g.Meta `path:"/v1/role/detail" method:"get" tags:"角色管理" summary:"角色详情" security:"api_key"`
	Id     int `d:"0" json:"id" dc:"角色ID"`
}

type RoleDetailRes *model.SysRoleDetailRes
