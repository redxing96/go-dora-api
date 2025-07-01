// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"go-dora-api/internal/model"
	"go-dora-api/internal/model/entity"
)

type (
	IMedia interface {
		// UploadFile 上传文件（支持秒传）
		UploadFile(ctx context.Context, in *model.MediaUploadInput) (res *model.MediaUploadOutput, err error)
		// GetMediaByMd5 根据MD5哈希查找文件
		GetMediaByMd5(ctx context.Context, md5Hash string) (res *entity.Media, err error)
		// GetMediaById 根据ID获取媒体文件信息
		GetMediaById(ctx context.Context, id int) (res *entity.Media, err error)
		// UpdateViewCount 更新文件访问统计
		UpdateViewCount(ctx context.Context, id int) error
		// UpdateDownloadCount 更新文件下载统计
		UpdateDownloadCount(ctx context.Context, id int) error
		// DeleteMedia 删除媒体文件
		DeleteMedia(ctx context.Context, id int) error
		// GetMediaList 获取媒体文件列表
		GetMediaList(ctx context.Context, in *model.MediaListInput) (res []*entity.Media, total int, err error)
	}
)

var (
	localMedia IMedia
)

func Media() IMedia {
	if localMedia == nil {
		panic("implement not found for interface IMedia, forgot register?")
	}
	return localMedia
}

func RegisterMedia(i IMedia) {
	localMedia = i
}
