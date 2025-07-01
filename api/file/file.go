// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package file

import (
	"context"

	"go-dora-api/api/file/v1"
)

type IFileV1 interface {
	FileUpload(ctx context.Context, req *v1.FileUploadReq) (res *v1.FileUploadRes, err error)
	MediaUpload(ctx context.Context, req *v1.MediaUploadReq) (res *v1.MediaUploadRes, err error)
	MediaList(ctx context.Context, req *v1.MediaListReq) (res *v1.MediaListRes, err error)
	MediaDetail(ctx context.Context, req *v1.MediaDetailReq) (res *v1.MediaDetailRes, err error)
	MediaDelete(ctx context.Context, req *v1.MediaDeleteReq) (res *v1.MediaDeleteRes, err error)
	MediaView(ctx context.Context, req *v1.MediaViewReq) (res *v1.MediaViewRes, err error)
	MediaDownload(ctx context.Context, req *v1.MediaDownloadReq) (res *v1.MediaDownloadRes, err error)
	MediaCheckMd5(ctx context.Context, req *v1.MediaCheckMd5Req) (res *v1.MediaCheckMd5Res, err error)
}
