/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-29 11:42:32
 * @LastEditTime: 2025-06-29 12:14:09
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/manager/v1/update.go
 */
package v1

import (
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ManagerUpdateReq struct {
	g.Meta   `path:"/v1/manager/update/" method:"post" tags:"管理员管理" summary:"修改详情信息" security:"api_key"`
	Id       int    `d:"" json:"id" dc:"管理员ID"`
	Account  string `d:"" json:"account" dc:"管理员账号"`
	Password string `d:"" json:"password" dc:"管理员密码"`
	Email    string `d:"" json:"email" dc:"电子邮箱"`
	Phone    string `d:"" json:"phone" dc:"电话号码"`
	Avatar   string `d:"" json:"avatar" dc:"头像"`
}

type ManagerUpdateRes *model.ManagerDetailRes
