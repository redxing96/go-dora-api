/*
 * @Description: 字典更新
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 09:21:50
 * @LastEditTime: 2025-07-02 09:26:51
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/dictionary/v1/update.go
 */
package v1

import "github.com/gogf/gf/v2/frame/g"

type DictionaryUpdateReq struct {
	g.Meta `path:"/v1/dictionary/update" method:"post" tags:"字典管理" summary:"更新字典" security:"api_key"`
	Id     int    `json:"id" dc:"字典ID"`
	Name   string `json:"name" dc:"字典名称"`
	Code   string `json:"code" dc:"字典编码"`
	Type   string `json:"type" dc:"字典类型"`
	Desc   string `json:"desc" dc:"字典描述"`
	Status int    `json:"status" dc:"状态 1-正常 2-禁用"`
}

type DictionaryUpdateRes struct {
	Id   int    `json:"id" dc:"字典ID"`
	Name string `json:"name" dc:"字典名称"`
}
