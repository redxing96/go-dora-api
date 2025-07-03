/*
 * @Description: 常量
 * @Author: redxing96@163.com
 * @Date: 2025-06-23 16:26:20
 * @LastEditTime: 2025-06-29 12:25:33
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/consts/consts.go
 */
package consts

var (
	//软件版本(打包时记得修改)
	APP_VERSION = "v1.0.0"

	//上传文件大小
	Max_UPLOAD_SIZE int64 = 500 * 1024 * 1024

	// 多语言
	LANG_CN = "zh-CN"
	LANG_EN = "en"

	// 客户端校验字段
	CLIENT_TOKEN_FIELD = "Authorization"
	// 管理端校验字段
	MANAGE_TOKEN_FIELD = "Authorization"

	// 分页默认配置
	DEFAULT_PAGE      = 20
	DEFAULT_PAGE_SIZE = 10

	// 上传文件路径
	UPLOADS_PATH = "resource/public/uploads/file"

	// 基础URL
	BaseUri = "http://192.168.10.180:4545"
)
