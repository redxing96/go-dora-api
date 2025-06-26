/*
 * @Description: 菜单基础逻辑
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 10:24:28
 * @LastEditTime: 2025-06-25 10:59:13
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/sys_menu/base.go
 */
package sys_menu

import "go-dora-api/internal/service"

type sSysMenu struct{}

func New() *sSysMenu {
	return &sSysMenu{}
}

func init() {
	service.RegisterSysMenu(New())
}
