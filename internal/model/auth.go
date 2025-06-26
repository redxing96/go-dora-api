/*
 * @Description: 管理端登录输出
 * @Author: redxing96@163.com
 * @Date: 2025-06-24 18:44:40
 * @LastEditTime: 2025-06-24 20:28:05
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/model/auth.go
 */
package model

type LoginOutput struct {
	Token    string `json:"token" dc:"token"`
	Exp      int64  `json:"exp" dc:"过期时间"`
	ManageId int    `json:"manage_id" dc:"管理员ID"`
	Account  string `json:"account" dc:"管理员账号"`
	IsSuper  int    `json:"is_super" dc:"是否超管"`
	Status   int    `json:"status" dc:"管理员状态"`
}

type ManageLoginInput struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type ManageLoginOutput struct {
	Token string `json:"token"`
}
