/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-29 12:20:12
 * @LastEditTime: 2025-06-29 12:22:14
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/file/v1/file_update.go
 */
package v1

import (
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/v1/file/upload" tags:"文件" method:"post" mime:"multipart/form-data" summary:"上传文件" security:"api_key"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传的文件"`
}

type FileUploadRes *model.FileUploadRes
