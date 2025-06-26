/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-23 17:05:49
 * @LastEditTime: 2025-06-23 17:23:24
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/common_return/common_return.go
 */
package common_return

import (
	"go-dora-api/internal/consts"
	"go-dora-api/utility/common_return"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
)

// 全局Json返回结构体
type JsonResponse struct {
	Code    int         `json:"code"`    // 业务码
	Message string      `json:"message"` // 业务码消息
	Version string      `json:"version"` // 软件版本
	Data    interface{} `json:"data"`    // 数据
}

// 成功返回
func Success(data interface{}) gcode.Code {
	return common_return.New(200, data, "success")
}

// 错误返回
func Error(data interface{}, msg string, args ...interface{}) gcode.Code {
	return common_return.New(500, data, msg, args...)
}

// 错误码返回
func ErrorCode(data interface{}, msg string, code int, args ...interface{}) gcode.Code {
	// 调用error_code包中的New函数，传入错误码、错误数据、错误信息和可选参数，返回一个错误码
	return common_return.New(code, data, msg, args...)
}

// 分割错误信息
func SplitErrorMessage(msg string) (key string, code int, message string) {
	message = ""
	code = -1
	message = msg

	if msg == "" {
		return
	}
	count := strings.Count(msg, "|")
	if count < 2 {
		return
	}

	index1 := strings.Index(msg, "|")
	if index1 == 0 {
		return
	}
	errorKey := msg[:index1]

	if index1+1 >= len(msg) {
		return
	}
	msg = msg[index1+1:]
	index2 := strings.Index(msg, "|")
	num, err := strconv.Atoi(msg[:index2])
	if err != nil {
		return
	}

	return errorKey, num, msg[index2+1:]
}

// 默认全局Json返回结构体封装
func DefaultJsonResponse(code int, message string, data interface{}) JsonResponse {
	return JsonResponse{
		Code:    code,
		Message: message,
		Version: consts.APP_VERSION,
		Data:    data,
	}
}
