// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"go-dora-api/internal/model"
)

type (
	IDictionary interface {
		// 向字典中添加一个新条目
		Add(ctx context.Context, in *model.DictionaryAddInput) (out *model.DictionaryAddOutput, err error)
		// 删除字典
		Delete(ctx context.Context, in *model.DictionaryDeleteInput) (out *model.DictionaryDeleteOutput, err error)
		// GetAll函数用于获取所有字典数据
		GetAll(ctx context.Context, in *model.DictionaryGetAllInput) (out []*model.DictionaryDetailOutput, err error)
		GetType(ctx context.Context) (out []string, err error)
		// List函数用于查询字典列表
		List(ctx context.Context, in *model.DictionaryListInput) (out []*model.DictionaryDetailOutput, total int, err error)
		// 更新字典
		Update(ctx context.Context, in *model.DictionaryUpdateInput) (out *model.DictionaryUpdateOutput, err error)
	}
)

var (
	localDictionary IDictionary
)

func Dictionary() IDictionary {
	if localDictionary == nil {
		panic("implement not found for interface IDictionary, forgot register?")
	}
	return localDictionary
}

func RegisterDictionary(i IDictionary) {
	localDictionary = i
}
