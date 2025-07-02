/*
 * @Description: 添加管理员
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 10:05:56
 * @LastEditTime: 2025-07-02 10:07:06
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/manager/v1/add.go
 */
package v1

import "github.com/gogf/gf/v2/frame/g"

type ManagerAddReq struct {
	g.Meta   `path:"/v1/manager/add" method:"post" tags:"管理员管理" summary:"添加管理员" security:"api_key"`
	Account  string `json:"account" dc:"账号"`
	Password string `json:"password" dc:"密码"`
	Email    string `json:"email" dc:"邮箱"`
	Phone    string `json:"phone" dc:"手机号"`
	Avatar   string `json:"avatar" dc:"头像"`
	Status   int    `json:"status" dc:"状态 0-初始化 1-正常 2-冻结"`
	IsSuper  int    `json:"is_super" dc:"是否超级管理员 1-是 2-否"`
	RoleIds  []int  `json:"role_ids" dc:"角色IDs"`
}

type ManagerAddRes struct {
	Id      int    `json:"id" dc:"管理员ID"`
	Account string `json:"account" dc:"账号"`
}
