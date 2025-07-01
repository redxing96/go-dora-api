/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-29 12:22:33
 * @LastEditTime: 2025-06-29 12:29:47
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/file/file_v1_file_upload.go
 */
package file

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"

	v1 "go-dora-api/api/file/v1"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/consts"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"
	"go-dora-api/utility/common"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

func (c *ControllerV1) FileUpload(ctx context.Context, req *v1.FileUploadReq) (res *v1.FileUploadRes, err error) {
	dirName := gtime.Now().Format("Y-m-d")
	saveDirPath := common.GetAbsPath(consts.UPLOADS_PATH) + "/" + dirName
	os.MkdirAll(saveDirPath, 0777)

	fileName, err := req.File.Save(saveDirPath, true)
	if err != nil {
		logger.SystemLogger.Errorf("保存文件失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "save_file_failed", 500))
		return
	}

	openedFile, err := req.File.Open()
	if err != nil {
		err = gerror.NewCode(common_return.ErrorCode("", "open_file_failed", 500))
		return
	}
	defer openedFile.Close()

	// 创建 MD5 哈希对象
	hash := md5.New()
	// 将文件内容复制到哈希对象
	if _, err = io.Copy(hash, openedFile); err != nil {
		err = gerror.NewCode(common_return.ErrorCode("", "copy_file_failed", 500))
		return
	}

	// 获取哈希结果
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	newData := new(model.FileUploadRes)
	newData.Url = consts.BaseUri + "/files/" + dirName + "/" + fileName
	newData.Path = saveDirPath + "/" + fileName
	newData.Name = fileName
	newData.Md5Hash = hashString
	resp := v1.FileUploadRes(newData)
	return &resp, nil
}
