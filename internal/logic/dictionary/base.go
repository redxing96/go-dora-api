/*
 * @Description: 字典基础逻辑
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 09:32:19
 * @LastEditTime: 2025-07-02 09:54:16
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/dictionary/base.go
 */
package dictionary

import "go-dora-api/internal/service"

type sDictionary struct{}

func New() *sDictionary {
	return &sDictionary{}
}

func init() {
	service.RegisterDictionary(New())
}
