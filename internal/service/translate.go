// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ITranslate interface {
		// 翻译
		Translate(ctx context.Context, content string, args ...any) string
		// 获取默认语言
		GetDefaultLanguage() string
	}
)

var (
	localTranslate ITranslate
)

func Translate() ITranslate {
	if localTranslate == nil {
		panic("implement not found for interface ITranslate, forgot register?")
	}
	return localTranslate
}

func RegisterTranslate(i ITranslate) {
	localTranslate = i
}
