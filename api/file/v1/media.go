/*
 * @Description: 媒体文件API接口定义
 * @Author: redxing96@163.com
 * @Date: 2025-06-29 13:00:00
 * @LastEditTime: 2025-06-29 13:10:49
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/api/file/v1/media.go
 */
package v1

import (
	"go-dora-api/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// MediaUploadReq 媒体文件上传请求
type MediaUploadReq struct {
	g.Meta   `path:"/v1/media/upload" tags:"媒体文件" method:"post" mime:"multipart/form-data" summary:"上传媒体文件" security:"api_key"`
	File     *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传的文件" v:"required#请选择要上传的文件"`
	IsPublic int               `json:"isPublic" dc:"是否公开 0-私有 1-公开" d:"0"`
}

// MediaUploadRes 媒体文件上传响应
type MediaUploadRes struct {
	Media         *entity.Media `json:"media" dc:"媒体文件信息"`
	IsQuickUpload bool          `json:"isQuickUpload" dc:"是否秒传"`
}

// MediaListReq 媒体文件列表请求
type MediaListReq struct {
	g.Meta     `path:"/v1/media/list" tags:"媒体文件" method:"get" summary:"获取媒体文件列表" security:"api_key"`
	Page       int    `json:"page" dc:"页码" d:"1"`
	PageSize   int    `json:"pageSize" dc:"每页数量" d:"10"`
	UploaderId int    `json:"uploaderId" dc:"上传者ID"`
	FileType   string `json:"fileType" dc:"文件类型"`
	Status     int    `json:"status" dc:"状态"`
	IsPublic   int    `json:"isPublic" dc:"是否公开"`
	Keyword    string `json:"keyword" dc:"关键词搜索"`
}

// MediaListRes 媒体文件列表响应
type MediaListRes struct {
	List  []*entity.Media `json:"list" dc:"媒体文件列表"`
	Total int             `json:"total" dc:"总数"`
	Page  int             `json:"page" dc:"当前页"`
	Size  int             `json:"size" dc:"每页数量"`
}

// MediaDetailReq 媒体文件详情请求
type MediaDetailReq struct {
	g.Meta `path:"/v1/media/detail" tags:"媒体文件" method:"get" summary:"获取媒体文件详情" security:"api_key"`
	Id     int `json:"id" dc:"媒体文件ID" v:"required#媒体文件ID不能为空"`
}

// MediaDetailRes 媒体文件详情响应
type MediaDetailRes *entity.Media

// MediaDeleteReq 媒体文件删除请求
type MediaDeleteReq struct {
	g.Meta `path:"/v1/media/delete" tags:"媒体文件" method:"delete" summary:"删除媒体文件" security:"api_key"`
	Id     int `json:"id" dc:"媒体文件ID" v:"required#媒体文件ID不能为空"`
}

// MediaDeleteRes 媒体文件删除响应
type MediaDeleteRes struct {
	Success bool `json:"success" dc:"是否成功"`
}

// MediaViewReq 媒体文件查看请求
type MediaViewReq struct {
	g.Meta `path:"/v1/media/view" tags:"媒体文件" method:"post" summary:"记录媒体文件查看" security:"api_key"`
	Id     int `json:"id" dc:"媒体文件ID" v:"required#媒体文件ID不能为空"`
}

// MediaViewRes 媒体文件查看响应
type MediaViewRes struct {
	Success bool `json:"success" dc:"是否成功"`
}

// MediaDownloadReq 媒体文件下载请求
type MediaDownloadReq struct {
	g.Meta `path:"/v1/media/download" tags:"媒体文件" method:"post" summary:"记录媒体文件下载" security:"api_key"`
	Id     int `json:"id" dc:"媒体文件ID" v:"required#媒体文件ID不能为空"`
}

// MediaDownloadRes 媒体文件下载响应
type MediaDownloadRes struct {
	Success bool `json:"success" dc:"是否成功"`
}

// MediaCheckMd5Req 检查MD5哈希请求
type MediaCheckMd5Req struct {
	g.Meta  `path:"/v1/media/check-md5" tags:"媒体文件" method:"post" summary:"检查文件MD5是否存在" security:"api_key"`
	Md5Hash string `json:"md5Hash" dc:"文件MD5哈希" v:"required#MD5哈希不能为空"`
}

// MediaCheckMd5Res 检查MD5哈希响应
type MediaCheckMd5Res struct {
	Exists bool          `json:"exists" dc:"文件是否存在"`
	Media  *entity.Media `json:"media" dc:"媒体文件信息（如果存在）"`
}
