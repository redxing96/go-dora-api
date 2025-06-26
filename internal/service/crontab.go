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
		// 初始化数据
		InitData(ctx context.Context) error
		// 启动sCrontab的定时任务
		Start(ctx context.Context, name string)
		// 停止定时任务
		Stop(ctx context.Context, name string)
		// Remove函数用于从sCrontab中移除指定的定时任务
		Remove(ctx context.Context, names []string)
		// 停止所有定时任务
		StopAll(ctx context.Context)
		// 根据传入的name参数，在s.cronObj中搜索对应的entry，如果找不到，则返回nil，否则返回一个model.CrontabSearchOutput结构体，包含entry的Name、Status、IsSingleton和RegisterTime属性
		Search(ctx context.Context, name string) *model.CrontabSearchOutput
		// All函数用于获取所有定时任务的信息
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
