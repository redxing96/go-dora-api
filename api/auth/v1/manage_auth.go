/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-24 18:55:15
 * @LastEditTime: 2025-06-24 18:55:18
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/auth/v1/manage_auth.go
 */
package v1

import (
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/v1/auth/login" method:"post" tags:"身份验证" summary:"登录接口"`
	Account  string `v:"required|length:4,30" d:"test" json:"account" dc:"账号"`
	Password string `v:"required|length:6,128" d:"e10adc3949ba59abbe56e057f20f883e" json:"password" dc:"密码"`
}

type LoginRes *model.LoginOutput
