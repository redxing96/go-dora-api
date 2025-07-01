// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Media is the golang structure for table media.
type Media struct {
	Id            int         `json:"id"            orm:"id"             description:"ID"`                 // ID
	FileName      string      `json:"fileName"      orm:"file_name"      description:"文件名"`                // 文件名
	OriginalName  string      `json:"originalName"  orm:"original_name"  description:"原始文件名"`              // 原始文件名
	FilePath      string      `json:"filePath"      orm:"file_path"      description:"文件路径"`               // 文件路径
	FileUrl       string      `json:"fileUrl"       orm:"file_url"       description:"文件访问URL"`            // 文件访问URL
	FileSize      int64       `json:"fileSize"      orm:"file_size"      description:"文件大小(字节)"`           // 文件大小(字节)
	FileType      string      `json:"fileType"      orm:"file_type"      description:"文件类型"`               // 文件类型
	MimeType      string      `json:"mimeType"      orm:"mime_type"      description:"MIME类型"`             // MIME类型
	Md5Hash       string      `json:"md5Hash"       orm:"md5_hash"       description:"文件MD5哈希"`            // 文件MD5哈希
	Sha1Hash      string      `json:"sha1Hash"      orm:"sha1_hash"      description:"文件SHA1哈希"`           // 文件SHA1哈希
	UploaderId    int         `json:"uploaderId"    orm:"uploader_id"    description:"上传者ID"`              // 上传者ID
	UploaderType  string      `json:"uploaderType"  orm:"uploader_type"  description:"上传者类型"`              // 上传者类型
	Status        int         `json:"status"        orm:"status"         description:"状态 0-待处理 1-正常 2-删除"` // 状态 0-待处理 1-正常 2-删除
	IsPublic      int         `json:"isPublic"      orm:"is_public"      description:"是否公开 0-私有 1-公开"`     // 是否公开 0-私有 1-公开
	DownloadCount int         `json:"downloadCount" orm:"download_count" description:"下载次数"`               // 下载次数
	ViewCount     int         `json:"viewCount"     orm:"view_count"     description:"查看次数"`               // 查看次数
	CreateTime    *gtime.Time `json:"createTime"    orm:"create_time"    description:"创建时间"`               // 创建时间
	UpdateTime    *gtime.Time `json:"updateTime"    orm:"update_time"    description:"修改时间"`               // 修改时间
}
