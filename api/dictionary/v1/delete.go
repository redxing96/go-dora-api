/*
 * @Description: 字典删除
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 09:22:27
 * @LastEditTime: 2025-07-02 09:25:12
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/dictionary/v1/delete.go
 */
package v1

import "github.com/gogf/gf/v2/frame/g"

type DictionaryDeleteReq struct {
	g.Meta `path:"/v1/dictionary/delete" method:"post" tags:"字典管理" summary:"删除字典" security:"api_key"`
	Ids    []int `json:"ids" dc:"字典IDs"`
}

type DictionaryDeleteRes struct {
	Ids []int `json:"ids" dc:"字典IDs"`
}
