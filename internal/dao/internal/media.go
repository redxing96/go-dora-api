// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MediaDao is the data access object for the table media.
type MediaDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns MediaColumns // columns contains all the column names of Table for convenient usage.
}

// MediaColumns defines and stores column names for the table media.
type MediaColumns struct {
	Id            string // ID
	FileName      string // 文件名
	OriginalName  string // 原始文件名
	FilePath      string // 文件路径
	FileUrl       string // 文件访问URL
	FileSize      string // 文件大小(字节)
	FileType      string // 文件类型
	MimeType      string // MIME类型
	Md5Hash       string // 文件MD5哈希
	Sha1Hash      string // 文件SHA1哈希
	UploaderId    string // 上传者ID
	UploaderType  string // 上传者类型
	Status        string // 状态 0-待处理 1-正常 2-删除
	IsPublic      string // 是否公开 0-私有 1-公开
	DownloadCount string // 下载次数
	ViewCount     string // 查看次数
	CreateTime    string // 创建时间
	UpdateTime    string // 修改时间
}

// mediaColumns holds the columns for the table media.
var mediaColumns = MediaColumns{
	Id:            "id",
	FileName:      "file_name",
	OriginalName:  "original_name",
	FilePath:      "file_path",
	FileUrl:       "file_url",
	FileSize:      "file_size",
	FileType:      "file_type",
	MimeType:      "mime_type",
	Md5Hash:       "md5_hash",
	Sha1Hash:      "sha1_hash",
	UploaderId:    "uploader_id",
	UploaderType:  "uploader_type",
	Status:        "status",
	IsPublic:      "is_public",
	DownloadCount: "download_count",
	ViewCount:     "view_count",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewMediaDao creates and returns a new DAO object for table data access.
func NewMediaDao() *MediaDao {
	return &MediaDao{
		group:   "default",
		table:   "media",
		columns: mediaColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MediaDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MediaDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MediaDao) Columns() MediaColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MediaDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MediaDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *MediaDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
