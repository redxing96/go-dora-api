// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Media is the golang structure of table media for DAO operations like Where/Data.
type Media struct {
	g.Meta        `orm:"table:media, do:true"`
	Id            interface{} // ID
	FileName      interface{} // 文件名
	OriginalName  interface{} // 原始文件名
	FilePath      interface{} // 文件路径
	FileUrl       interface{} // 文件访问URL
	FileSize      interface{} // 文件大小(字节)
	FileType      interface{} // 文件类型
	MimeType      interface{} // MIME类型
	Md5Hash       interface{} // 文件MD5哈希
	Sha1Hash      interface{} // 文件SHA1哈希
	UploaderId    interface{} // 上传者ID
	UploaderType  interface{} // 上传者类型
	Status        interface{} // 状态 0-待处理 1-正常 2-删除
	IsPublic      interface{} // 是否公开 0-私有 1-公开
	DownloadCount interface{} // 下载次数
	ViewCount     interface{} // 查看次数
	CreateTime    *gtime.Time // 创建时间
	UpdateTime    *gtime.Time // 修改时间
}
