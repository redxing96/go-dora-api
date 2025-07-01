/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-29 12:21:47
 * @LastEditTime: 2025-06-29 12:59:50
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/model/file.go
 */
package model

import "go-dora-api/internal/model/entity"

type FileUploadRes struct {
	Md5Hash string `json:"md5_hash" dc:"文件MD5"`
	Url     string `json:"url" dc:"文件URL"`
	Path    string `json:"path" dc:"文件路径"`
	Name    string `json:"name" dc:"文件名称"`
}

// MediaUploadInput 媒体文件上传输入
type MediaUploadInput struct {
	FileName     string `json:"fileName"     dc:"文件名"`
	OriginalName string `json:"originalName" dc:"原始文件名"`
	FilePath     string `json:"filePath"     dc:"文件路径"`
	FileUrl      string `json:"fileUrl"      dc:"文件访问URL"`
	FileSize     int64  `json:"fileSize"     dc:"文件大小"`
	FileType     string `json:"fileType"     dc:"文件类型"`
	MimeType     string `json:"mimeType"     dc:"MIME类型"`
	Md5Hash      string `json:"md5Hash"      dc:"MD5哈希"`
	Sha1Hash     string `json:"sha1Hash"     dc:"SHA1哈希"`
	UploaderId   int    `json:"uploaderId"   dc:"上传者ID"`
	UploaderType string `json:"uploaderType" dc:"上传者类型"`
	IsPublic     int    `json:"isPublic"     dc:"是否公开"`
}

// MediaUploadOutput 媒体文件上传输出
type MediaUploadOutput struct {
	Media         *entity.Media `json:"media"       dc:"媒体文件信息"`
	IsQuickUpload bool          `json:"isQuickUpload" dc:"是否秒传"`
}

// MediaListInput 媒体文件列表输入
type MediaListInput struct {
	Page       int    `json:"page"        dc:"页码"`
	PageSize   int    `json:"pageSize"    dc:"每页数量"`
	UploaderId int    `json:"uploaderId"  dc:"上传者ID"`
	FileType   string `json:"fileType"    dc:"文件类型"`
	Status     int    `json:"status"      dc:"状态"`
	IsPublic   int    `json:"isPublic"    dc:"是否公开"`
	Keyword    string `json:"keyword"     dc:"关键词搜索"`
}

// MediaListOutput 媒体文件列表输出
type MediaListOutput struct {
	List  []*entity.Media `json:"list"  dc:"媒体文件列表"`
	Total int             `json:"total" dc:"总数"`
	Page  int             `json:"page"  dc:"当前页"`
	Size  int             `json:"size"  dc:"每页数量"`
}
