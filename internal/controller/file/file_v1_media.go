/*
 * @Description: 媒体文件控制器
 * @Author: redxing96@163.com
 * @Date: 2025-06-29 13:00:00
 * @LastEditTime: 2025-06-29 13:48:22
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/file/file_v1_media.go
 */
package file

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strings"

	v1 "go-dora-api/api/file/v1"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/consts"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
	"go-dora-api/utility/common"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/spf13/cast"
)

// MediaUpload 媒体文件上传（支持秒传）
func (c *ControllerV1) MediaUpload(ctx context.Context, req *v1.MediaUploadReq) (res *v1.MediaUploadRes, err error) {

	managerId := cast.ToInt(ctx.Value("manager_id"))

	// 验证文件
	if req.File == nil {
		err = gerror.NewCode(common_return.ErrorCode("", "file_required", 400))
		return
	}

	// 获取文件信息
	fileHeader := req.File
	originalName := fileHeader.Filename
	fileSize := fileHeader.Size

	// 检查文件大小限制（例如：100MB）
	maxFileSize := int64(100 * 1024 * 1024)
	if fileSize > maxFileSize {
		err = gerror.NewCode(common_return.ErrorCode("", "file_too_large", 400))
		return
	}

	// 获取文件扩展名和MIME类型
	fileExt := strings.ToLower(filepath.Ext(originalName))
	mimeType := fileHeader.Header.Get("Content-Type")

	// 检查文件类型是否允许
	allowedTypes := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".mp4", ".avi", ".mov", ".wmv", ".flv", ".mp3", ".wav", ".flac", ".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".txt", ".zip", ".rar", ".7z"}
	isAllowed := false
	for _, allowedType := range allowedTypes {
		if fileExt == allowedType {
			isAllowed = true
			break
		}
	}
	if !isAllowed {
		err = gerror.NewCode(common_return.ErrorCode("", "file_type_not_allowed", 400))
		return
	}

	// 保存文件到临时目录
	dirName := gtime.Now().Format("Y-m-d")
	saveDirPath := common.GetAbsPath(consts.UPLOADS_PATH) + "/" + dirName
	if err := os.MkdirAll(saveDirPath, 0777); err != nil {
		logger.SystemLogger.Errorf("创建目录失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "create_dir_failed", 500))
		return nil, err
	}

	// 生成唯一文件名
	fileName, err := req.File.Save(saveDirPath, true)
	if err != nil {
		logger.SystemLogger.Errorf("保存文件失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "save_file_failed", 500))
		return
	}

	// 计算MD5哈希
	md5Hash, err := c.calculateFileMd5(saveDirPath + "/" + fileName)
	if err != nil {
		logger.SystemLogger.Errorf("计算MD5哈希失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "calculate_md5_failed", 500))
		return
	}

	// 构建上传输入参数
	uploadInput := &model.MediaUploadInput{
		FileName:     fileName,
		OriginalName: originalName,
		FilePath:     saveDirPath + "/" + fileName,
		FileUrl:      consts.BaseUri + "/files/" + dirName + "/" + fileName,
		FileSize:     fileSize,
		FileType:     fileExt,
		MimeType:     mimeType,
		Md5Hash:      md5Hash,
		Sha1Hash:     "", // 将在服务层计算
		UploaderId:   managerId,
		UploaderType: "manager",
		IsPublic:     req.IsPublic,
	}

	// 调用服务层上传文件
	uploadOutput, err := service.Media().UploadFile(ctx, uploadInput)
	if err != nil {
		return nil, err
	}

	return &v1.MediaUploadRes{
		Media:         uploadOutput.Media,
		IsQuickUpload: uploadOutput.IsQuickUpload,
	}, nil
}

// MediaList 获取媒体文件列表
func (c *ControllerV1) MediaList(ctx context.Context, req *v1.MediaListReq) (res *v1.MediaListRes, err error) {
	// 构建查询参数
	listInput := &model.MediaListInput{
		Page:       req.Page,
		PageSize:   req.PageSize,
		UploaderId: req.UploaderId,
		FileType:   req.FileType,
		Status:     req.Status,
		IsPublic:   req.IsPublic,
		Keyword:    req.Keyword,
	}

	// 调用服务层获取列表
	mediaList, total, err := service.Media().GetMediaList(ctx, listInput)
	if err != nil {
		return nil, err
	}

	return &v1.MediaListRes{
		List:  mediaList,
		Total: total,
		Page:  req.Page,
		Size:  req.PageSize,
	}, nil
}

// MediaDetail 获取媒体文件详情
func (c *ControllerV1) MediaDetail(ctx context.Context, req *v1.MediaDetailReq) (res *v1.MediaDetailRes, err error) {
	// 调用服务层获取详情
	media, err := service.Media().GetMediaById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	resp := v1.MediaDetailRes(media)

	return &resp, nil
}

// MediaDelete 删除媒体文件
func (c *ControllerV1) MediaDelete(ctx context.Context, req *v1.MediaDeleteReq) (res *v1.MediaDeleteRes, err error) {
	// 调用服务层删除文件
	err = service.Media().DeleteMedia(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.MediaDeleteRes{
		Success: true,
	}, nil
}

// MediaView 记录媒体文件查看
func (c *ControllerV1) MediaView(ctx context.Context, req *v1.MediaViewReq) (res *v1.MediaViewRes, err error) {
	// 调用服务层更新查看次数
	err = service.Media().UpdateViewCount(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.MediaViewRes{
		Success: true,
	}, nil
}

// MediaDownload 记录媒体文件下载
func (c *ControllerV1) MediaDownload(ctx context.Context, req *v1.MediaDownloadReq) (res *v1.MediaDownloadRes, err error) {
	// 调用服务层更新下载次数
	err = service.Media().UpdateDownloadCount(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.MediaDownloadRes{
		Success: true,
	}, nil
}

// MediaCheckMd5 检查文件MD5是否存在
func (c *ControllerV1) MediaCheckMd5(ctx context.Context, req *v1.MediaCheckMd5Req) (res *v1.MediaCheckMd5Res, err error) {
	// 调用服务层检查MD5
	media, err := service.Media().GetMediaByMd5(ctx, req.Md5Hash)
	if err != nil {
		// 如果文件不存在，返回false
		if gerror.Code(err).Code() == common_return.ErrorCode("", "media_not_found", 404).Code() {
			return &v1.MediaCheckMd5Res{
				Exists: false,
				Media:  nil,
			}, nil
		}
		return nil, err
	}

	return &v1.MediaCheckMd5Res{
		Exists: true,
		Media:  media,
	}, nil
}

// calculateFileMd5 计算文件MD5哈希
func (c *ControllerV1) calculateFileMd5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
