/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-26 17:49:48
 * @LastEditTime: 2025-06-26 17:56:25
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/menu/v1/get_auth_menu.go
 */
package v1

import (
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type GetAuthMenuReq struct {
	g.Meta `path:"/v1/menu/get_auth_menu" method:"get" tags:"路由权限管理" summary:"获取菜单权限接口（根据token获取）" security:"api_key"`
}

type GetAuthMenuRes struct {
	MenuList []*model.MenuItem `json:"menu_list"`
}
