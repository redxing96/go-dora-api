/*
 * @Description: 公共方法
 * @Author: redxing96@163.com
 * @Date: 2025-06-23 16:48:39
 * @LastEditTime: 2025-06-23 16:48:46
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/utility/common/common.go
 */
package common

import (
	"path/filepath"

	"github.com/gogf/gf/v2/os/gfile"
)

// 获取绝对路径
func GetAbsPath(relativePath string) string {
	return filepath.Join(gfile.SelfDir(), relativePath)
}
