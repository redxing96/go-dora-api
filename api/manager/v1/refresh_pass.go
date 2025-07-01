/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-29 12:02:09
 * @LastEditTime: 2025-06-29 12:13:54
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/manager/v1/refresh_pass.go
 */
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ManagerRefreshPassReq struct {
	g.Meta   `path:"/v1/manager/refresh/pass" method:"post" tags:"管理员管理" summary:"刷新管理员密码(默认密码:Dora.123456)" security:"api_key"`
	Id       int    `d:"" json:"id" dc:"需要重置密码的管理员ID"`
	Password string `d:"" json:"password" dc:"当前登录的管理员密码"`
}

type ManagerRefreshPassRes struct {
	Id int `json:"id" dc:"管理员ID"`
}
