/*
 * @Description: 管理员退出登录
 * @Author: redxing96@163.com
 * @Date: 2025-06-26 16:35:59
 * @LastEditTime: 2025-06-26 17:25:02
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/auth/v1/manage_logout.go
 */
package v1

import "github.com/gogf/gf/v2/frame/g"

type LogoutReq struct {
	g.Meta `path:"/v1/auth/logout" method:"post" tags:"身份验证" summary:"退出登录接口" security:"api_key"`
}

type LogoutRes struct{}
