/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-29 12:29:45
 * @LastEditTime: 2025-06-29 12:30:21
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/media/base.go
 */
package media

import "go-dora-api/internal/service"

type sMedia struct{}

func New() *sMedia {
	return &sMedia{}
}

func init() {
	service.RegisterMedia(New())
}
