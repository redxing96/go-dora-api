/*
 * @Description: 角色列表
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 21:28:59
 * @LastEditTime: 2025-06-25 22:14:38
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/sys_role/v1/list.go
 */
package v1

import (
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type RoleListReq struct {
	g.Meta   `path:"/v1/role/list" method:"get" tags:"角色管理" summary:"角色列表" security:"api_key"`
	Page     int    `d:"1" json:"page" dc:"当前页"`
	PageSize int    `d:"20" json:"page_size" dc:"每页条数"`
	Search   string `d:"" json:"search" dc:"搜索关键字"`
}

type RoleListRes *model.SysRoleListRes
