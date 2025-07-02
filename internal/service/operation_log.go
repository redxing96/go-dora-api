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
	IOperationLog interface {
		// 创建操作日志
		Create(ctx context.Context, in *model.OperationLogCreateInput) (out *model.OperationLogCreateOutput, err error)
		// 获取操作日志列表
		GetList(ctx context.Context, in *model.OperationLogListInput) (out *model.OperationLogListOutput, err error)
		// 获取操作日志详情
		GetDetail(ctx context.Context, id int64) (out *model.OperationLogDetailOutput, err error)
		// 删除操作日志
		Delete(ctx context.Context, in *model.OperationLogDeleteInput) (out *model.OperationLogDeleteOutput, err error)
		// 记录操作日志（便捷方法）
		RecordOperation(ctx context.Context, userId int64, username string, operation string, module string, description string, status int, errMsg string, executionTime int, requestInfo map[string]interface{}) error
	}
)

var (
	localOperationLog IOperationLog
)

func OperationLog() IOperationLog {
	if localOperationLog == nil {
		panic("implement not found for interface IOperationLog, forgot register?")
	}
	return localOperationLog
}

func RegisterOperationLog(i IOperationLog) {
	localOperationLog = i
}
