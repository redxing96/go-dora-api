/*
 * @Description: 获取所有角色
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 22:37:38
 * @LastEditTime: 2025-06-25 22:38:58
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/sys_role/v1/all.go
 */
package v1

import (
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type RoleAllReq struct {
	g.Meta `path:"/v1/role/all" method:"get" tags:"角色管理" summary:"获取所有角色" security:"api_key"`
	Status int `json:"status" dc:"状态 1-正常 2-禁用"`
}

type RoleAllRes *model.SysRoleAllRes
