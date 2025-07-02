/*
 * @Description: 字典列表
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 09:21:14
 * @LastEditTime: 2025-07-02 09:57:51
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/dictionary/v1/list.go
 */
package v1

import (
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type DictionaryListReq struct {
	g.Meta   `path:"/v1/dictionary/list" method:"get" tags:"字典管理" summary:"字典列表" security:"api_key"`
	Page     int    `json:"page" dc:"页码" d:"1"`
	PageSize int    `json:"page_size" dc:"每页数量" d:"20"`
	Search   string `json:"search" dc:"搜索关键词"`
	Type     string `json:"type" dc:"字典类型"`
	LikeName string `json:"like_name" dc:"模糊查询名称"`
	LikeCode string `json:"like_code" dc:"模糊查询编码"`
	Status   int    `json:"status" dc:"状态 1-正常 2-禁用"`
}

type DictionaryListRes *model.DictionaryListRes
