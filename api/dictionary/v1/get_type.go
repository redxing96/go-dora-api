/*
 * @Description: 获取字典类型
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 13:26:38
 * @LastEditTime: 2025-07-02 13:26:55
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/dictionary/v1/get_type.go
 */
package v1

import "github.com/gogf/gf/v2/frame/g"

type DictionaryGetTypeReq struct {
	g.Meta `path:"/v1/dictionary/type/all" method:"get" tags:"字典管理" summary:"字典类型" security:"api_key"`
}

type DictionaryGetTypeRes []string
