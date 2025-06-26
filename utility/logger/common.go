/*
 * @Description: 日志组件公共接口
 * @Author: redxing96@163.com
 * @Date: 2025-06-23 16:40:55
 * @LastEditTime: 2025-06-23 16:41:04
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/utility/logger/common.go
 */
package logger

import (
	"os"
	"path/filepath"
)

// 日志组件接口
type ILogger interface {
	SetLevel(level string) error
	Level() string
	Debug(msg string, fields ...IField)
	Info(msg string, fields ...IField)
	Warn(msg string, fields ...IField)
	Error(msg string, fields ...IField)
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Rotate() error
	Sync() error
}

// 字段组件接口
type IField interface {
	GetKey() string
	GetValue() interface{}
}

// 判断是否是合法的文件路径
func IsValidFile(fp string) bool {
	stat, err := os.Stat(fp)
	if err == nil {
		return !stat.IsDir()
	}

	if err := os.MkdirAll(filepath.Dir(fp), 0744); err != nil {
		return false
	}

	var d []byte
	if err := os.WriteFile(fp, d, 0644); err == nil {
		os.Remove(fp)
		return true
	}
	return false
}
