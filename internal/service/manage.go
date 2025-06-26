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
	IManage interface {
		// 根据账户获取管理员详情
		GetManageDetailToAccount(ctx context.Context, in *model.GetManageDetailInput) (res *model.GetManageDetailOutput, err error)
	}
)

var (
	localManage IManage
)

func Manage() IManage {
	if localManage == nil {
		panic("implement not found for interface IManage, forgot register?")
	}
	return localManage
}

func RegisterManage(i IManage) {
	localManage = i
}
