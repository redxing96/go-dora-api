/*
 * @Description: 角色基础逻辑
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 10:24:28
 * @LastEditTime: 2025-06-25 22:13:49
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/sys_role/base.go
 */
package sys_role

import "go-dora-api/internal/service"

type sSysRole struct{}

func New() *sSysRole {
	return &sSysRole{}
}

func init() {
	service.RegisterSysRole(New())
}
