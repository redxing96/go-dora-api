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
	ICrontab interface {
		InitData(ctx context.Context) error
		Start(ctx context.Context, name string)
		Stop(ctx context.Context, name string)
		Remove(ctx context.Context, names []string)
		StopAll(ctx context.Context)
		Search(ctx context.Context, name string) *model.CrontabSearchOutput
		All(ctx context.Context) (output []model.CrontabSearchOutput)
	}
)

var (
	localCrontab ICrontab
)

func Crontab() ICrontab {
	if localCrontab == nil {
		panic("implement not found for interface ICrontab, forgot register?")
	}
	return localCrontab
}

func RegisterCrontab(i ICrontab) {
	localCrontab = i
}
