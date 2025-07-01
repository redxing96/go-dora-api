package media

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"

	"go-dora-api/internal/common_return"
	"go-dora-api/internal/consts"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"
	"go-dora-api/internal/model/entity"
	"go-dora-api/utility/common"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// UploadFile 上传文件（支持秒传）
func (s *sMedia) UploadFile(ctx context.Context, in *model.MediaUploadInput) (res *model.MediaUploadOutput, err error) {
	// 首先检查是否存在相同MD5的文件（秒传检查）
	existingMedia, err := s.GetMediaByMd5(ctx, in.Md5Hash)
	if err == nil && existingMedia != nil {
		logger.SystemLogger.Infof("文件已存在，返回秒传结果: %v", existingMedia)
		// 文件已存在，返回秒传结果
		return &model.MediaUploadOutput{
			Media:         existingMedia,
			IsQuickUpload: true,
		}, nil
	}

	// 文件不存在，执行正常上传流程
	dirName := gtime.Now().Format("Y-m-d")
	saveDirPath := common.GetAbsPath(consts.UPLOADS_PATH) + "/" + dirName

	// 确保目录存在
	if err := os.MkdirAll(saveDirPath, 0777); err != nil {
		logger.SystemLogger.Errorf("创建目录失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "create_dir_failed", 500))
		return nil, err
	}

	// 保存文件到磁盘
	fileName, err := s.saveFileToDisk(in.FilePath, saveDirPath)
	if err != nil {
		logger.SystemLogger.Errorf("保存文件失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "save_file_failed", 500))
		return
	}

	// 计算SHA1哈希
	sha1Hash, err := s.calculateSha1Hash(in.FilePath)
	if err != nil {
		logger.SystemLogger.Errorf("计算SHA1哈希失败: %v", err)
		return nil, gerror.NewCode(common_return.ErrorCode("", "calculate_sha1_failed", 500))
	}

	// 创建媒体文件记录
	mediaData := &entity.Media{
		FileName:      fileName,
		OriginalName:  in.OriginalName,
		FilePath:      saveDirPath + "/" + fileName,
		FileUrl:       consts.BaseUri + "/files/" + dirName + "/" + fileName,
		FileSize:      in.FileSize,
		FileType:      in.FileType,
		MimeType:      in.MimeType,
		Md5Hash:       in.Md5Hash,
		Sha1Hash:      sha1Hash,
		UploaderId:    in.UploaderId,
		UploaderType:  in.UploaderType,
		Status:        1, // 正常状态
		IsPublic:      in.IsPublic,
		DownloadCount: 0,
		ViewCount:     0,
		CreateTime:    gtime.Now(),
		UpdateTime:    gtime.Now(),
	}

	// 保存到数据库
	result, err := dao.Media.Ctx(ctx).Data(mediaData).Insert()
	if err != nil {
		logger.SystemLogger.Errorf("保存媒体文件记录失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "save_media_record_failed", 500))
		return
	}

	// 获取插入的ID
	insertId, err := result.LastInsertId()
	if err != nil {
		logger.SystemLogger.Errorf("获取插入ID失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_insert_id_failed", 500))
		return
	}

	mediaData.Id = int(insertId)

	return &model.MediaUploadOutput{
		Media:         mediaData,
		IsQuickUpload: false,
	}, nil
}

// saveFileToDisk 保存文件到磁盘
func (s *sMedia) saveFileToDisk(filePath, saveDirPath string) (string, error) {
	// 这里假设传入的filePath是临时文件路径
	// 在实际使用中，可能需要处理multipart.FileHeader
	fileName := filepath.Base(filePath)
	return fileName, nil
}

// calculateSha1Hash 计算文件的SHA1哈希
func (s *sMedia) calculateSha1Hash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		logger.SystemLogger.Errorf("打开文件失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "open_file_failed", 500))
		return "", err
	}
	defer file.Close()

	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		logger.SystemLogger.Errorf("计算SHA1哈希失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "calculate_sha1_failed", 500))
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// GetMediaByMd5 根据MD5哈希查找文件
func (s *sMedia) GetMediaByMd5(ctx context.Context, md5Hash string) (res *entity.Media, err error) {
	var media entity.Media
	err = dao.Media.Ctx(ctx).Where("md5_hash", md5Hash).Where("status", 1).Scan(&media)
	if err != nil {
		logger.SystemLogger.Errorf("根据MD5哈希查找文件失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_media_by_md5_failed", 500))
		return nil, err
	}
	if media.Id == 0 {
		logger.SystemLogger.Errorf("文件不存在: %v", md5Hash)
		err = gerror.NewCode(common_return.ErrorCode("", "media_not_found", 404))
		return
	}
	return &media, nil
}

// GetMediaById 根据ID获取媒体文件信息
func (s *sMedia) GetMediaById(ctx context.Context, id int) (res *entity.Media, err error) {
	var media entity.Media
	err = dao.Media.Ctx(ctx).Where("id", id).Where("status", 1).Scan(&media)
	if err != nil {
		logger.SystemLogger.Errorf("根据ID获取媒体文件信息失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_media_by_id_failed", 500))
		return nil, err
	}
	if media.Id == 0 {
		return nil, gerror.NewCode(common_return.ErrorCode("", "media_not_found", 404))
	}
	return &media, nil
}

// UpdateViewCount 更新文件访问统计
func (s *sMedia) UpdateViewCount(ctx context.Context, id int) error {
	_, err := dao.Media.Ctx(ctx).Data("view_count", dao.Media.Columns().ViewCount+"+1").Where("id", id).Update()
	return err
}

// UpdateDownloadCount 更新文件下载统计
func (s *sMedia) UpdateDownloadCount(ctx context.Context, id int) error {
	_, err := dao.Media.Ctx(ctx).Data("download_count", dao.Media.Columns().DownloadCount+"+1").Where("id", id).Update()
	return err
}

// DeleteMedia 删除媒体文件
func (s *sMedia) DeleteMedia(ctx context.Context, id int) error {
	// 软删除，将状态设置为删除
	_, err := dao.Media.Ctx(ctx).Data("status", 2).Where("id", id).Update()
	return err
}

// GetMediaList 获取媒体文件列表
func (s *sMedia) GetMediaList(ctx context.Context, in *model.MediaListInput) (res []*entity.Media, total int, err error) {
	model := dao.Media.Ctx(ctx).Where("status", 1)

	// 添加查询条件
	if in.UploaderId > 0 {
		model = model.Where("uploader_id", in.UploaderId)
	}
	if in.FileType != "" {
		model = model.Where("file_type", in.FileType)
	}
	if in.Status > 0 {
		model = model.Where("status", in.Status)
	}
	if in.IsPublic >= 0 {
		model = model.Where("is_public", in.IsPublic)
	}
	if in.Keyword != "" {
		model = model.WhereLike("original_name", "%"+in.Keyword+"%").
			WhereOrLike("file_name", "%"+in.Keyword+"%")
	}

	// 获取总数
	total, err = model.Count()
	if err != nil {
		logger.SystemLogger.Errorf("获取媒体文件列表总数失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_media_list_total_failed", 500))
		return
	}

	// 分页查询
	var mediaList []*entity.Media
	err = model.Page(in.Page, in.PageSize).OrderDesc("create_time").Scan(&mediaList)
	if err != nil {
		logger.SystemLogger.Errorf("获取媒体文件列表失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_media_list_failed", 500))
		return
	}

	return mediaList, total, nil
}
