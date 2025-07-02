/*
 * @Description: 获取所有字典
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 10:01:17
 * @LastEditTime: 2025-07-02 10:01:20
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/dictionary/v1/get_all.go
 */
package v1

import (
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type DictionaryGetAllReq struct {
	g.Meta   `path:"/v1/dictionary/get_all" method:"get" tags:"字典管理" summary:"获取所有字典" security:"api_key"`
	Type     string `json:"type" dc:"字典类型"`
	LikeName string `json:"like_name" dc:"模糊查询名称"`
	LikeCode string `json:"like_code" dc:"模糊查询编码"`
	Status   int    `json:"status" dc:"状态 1-正常 2-禁用"`
}

type DictionaryGetAllRes []*model.DictionaryDetailOutput
