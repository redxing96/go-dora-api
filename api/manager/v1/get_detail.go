/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-29 11:23:09
 * @LastEditTime: 2025-06-29 11:36:52
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/manager/v1/get_detail.go
 */
package v1

import (
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ManagerDetailReq struct {
	g.Meta `path:"/v1/manager/detail" method:"get" tags:"管理员管理" summary:"管理员详情" security:"api_key"`
	Id     int `v:"required" d:"1" json:"id" dc:"管理员ID"`
}

type ManagerDetailRes *model.ManagerDetailRes
