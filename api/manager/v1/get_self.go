/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-29 11:23:19
 * @LastEditTime: 2025-06-29 11:25:59
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/manager/v1/get_self.go
 */
package v1

import (
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ManagerSelfReq struct {
	g.Meta `path:"/v1/manager/self" method:"get" tags:"管理员管理" summary:"获取自己的详情信息" security:"api_key"`
}

type ManagerSelfRes *model.ManagerDetailRes
