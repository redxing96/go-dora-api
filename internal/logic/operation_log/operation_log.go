/*
 * @Description: 操作日志业务逻辑
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 10:00:00
 * @LastEditTime: 2025-07-02 13:37:02
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/operation_log/operation_log.go
 */
package operation_log

import (
	"context"
	"encoding/json"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"
	"go-dora-api/internal/model/do"
	"go-dora-api/internal/model/entity"
	"go-dora-api/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type sOperationLog struct{}

func init() {
	service.RegisterOperationLog(New())
}

func New() *sOperationLog {
	return &sOperationLog{}
}

// 创建操作日志
func (s *sOperationLog) Create(ctx context.Context, in *model.OperationLogCreateInput) (out *model.OperationLogCreateOutput, err error) {
	// 插入操作日志
	result, err := dao.OperationLog.Ctx(ctx).Data(do.OperationLog{
		UserId:        in.UserId,
		Username:      in.Username,
		Operation:     in.Operation,
		Module:        in.Module,
		Description:   in.Description,
		RequestMethod: in.RequestMethod,
		RequestUrl:    in.RequestUrl,
		RequestParams: in.RequestParams,
		ResponseData:  in.ResponseData,
		IpAddress:     in.IpAddress,
		UserAgent:     in.UserAgent,
		Status:        in.Status,
		ErrorMessage:  in.ErrorMessage,
		ExecutionTime: in.ExecutionTime,
		CreateTime:    gtime.Now(),
	}).InsertAndGetId()
	if err != nil {
		logger.SystemLogger.Errorf("创建操作日志失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "create_operation_log_failed", 500))
		return
	}

	out = &model.OperationLogCreateOutput{
		Id: result,
	}
	return
}

// 获取操作日志列表
func (s *sOperationLog) GetList(ctx context.Context, in *model.OperationLogListInput) (out []*entity.OperationLog, total int, err error) {
	model := dao.OperationLog.Ctx(ctx)

	// 添加查询条件
	if in.UserId > 0 {
		model = model.Where(dao.OperationLog.Columns().UserId, in.UserId)
	}
	if in.Username != "" {
		model = model.WhereLike(dao.OperationLog.Columns().Username, "%"+in.Username+"%")
	}
	if in.Operation != "" {
		model = model.Where(dao.OperationLog.Columns().Operation, in.Operation)
	}
	if in.Module != "" {
		model = model.Where(dao.OperationLog.Columns().Module, in.Module)
	}
	if in.Status > 0 {
		model = model.Where(dao.OperationLog.Columns().Status, in.Status)
	}
	if in.StartTime != "" {
		model = model.WhereGTE(dao.OperationLog.Columns().CreateTime, in.StartTime)
	}
	if in.EndTime != "" {
		model = model.WhereLTE(dao.OperationLog.Columns().CreateTime, in.EndTime)
	}
	if in.Keyword != "" {
		model = model.WhereLike(dao.OperationLog.Columns().Description, "%"+in.Keyword+"%").
			WhereOrLike(dao.OperationLog.Columns().RequestUrl, "%"+in.Keyword+"%")
	}

	// 获取总数
	total, err = model.Count()
	if err != nil {
		logger.SystemLogger.Errorf("获取操作日志列表总数失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_operation_log_total_failed", 500))
		return
	}

	// 分页查询
	err = model.Page(in.Page, in.PageSize).OrderDesc(dao.OperationLog.Columns().CreateTime).Scan(&out)
	if err != nil {
		logger.SystemLogger.Errorf("获取操作日志列表失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_operation_log_list_failed", 500))
		return
	}

	return
}

// 获取操作日志详情
func (s *sOperationLog) GetDetail(ctx context.Context, id int64) (out *model.OperationLogDetailOutput, err error) {
	var log model.OperationLogDetailOutput
	err = dao.OperationLog.Ctx(ctx).Where(dao.OperationLog.Columns().Id, id).Scan(&log)
	if err != nil {
		logger.SystemLogger.Errorf("获取操作日志详情失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_operation_log_detail_failed", 500))
		return
	}
	if log.Id == 0 {
		err = gerror.NewCode(common_return.ErrorCode("", "operation_log_not_found", 404))
		return
	}

	out = &log
	return
}

// 删除操作日志
func (s *sOperationLog) Delete(ctx context.Context, in *model.OperationLogDeleteInput) (out *model.OperationLogDeleteOutput, err error) {
	result, err := dao.OperationLog.Ctx(ctx).WhereIn(dao.OperationLog.Columns().Id, in.Ids).Delete()
	if err != nil {
		logger.SystemLogger.Errorf("删除操作日志失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "delete_operation_log_failed", 500))
		return
	}

	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		logger.SystemLogger.Errorf("删除操作日志失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "delete_operation_log_failed", 500))
		return
	}

	out = &model.OperationLogDeleteOutput{
		Ids: in.Ids,
	}
	return
}

// 记录操作日志（便捷方法）
func (s *sOperationLog) RecordOperation(ctx context.Context, userId int64, username, operation, module, description string, status int, errMsg string, executionTime int, requestInfo map[string]interface{}) error {
	// 处理请求信息
	requestMethod := ""
	requestUrl := ""
	requestParams := ""
	responseData := ""
	ipAddress := ""
	userAgent := ""

	if requestInfo != nil {
		if method, ok := requestInfo["method"].(string); ok {
			requestMethod = method
		}
		if url, ok := requestInfo["url"].(string); ok {
			requestUrl = url
		}
		if params, ok := requestInfo["params"]; ok {
			if paramsBytes, err := json.Marshal(params); err == nil {
				requestParams = string(paramsBytes)
			}
		}
		if response, ok := requestInfo["response"]; ok {
			if responseBytes, err := json.Marshal(response); err == nil {
				responseData = string(responseBytes)
			}
		}
		if ip, ok := requestInfo["ip"].(string); ok {
			ipAddress = ip
		}
		if ua, ok := requestInfo["userAgent"].(string); ok {
			userAgent = ua
		}
	}

	// 创建操作日志
	_, err := s.Create(ctx, &model.OperationLogCreateInput{
		UserId:        userId,
		Username:      username,
		Operation:     operation,
		Module:        module,
		Description:   description,
		RequestMethod: requestMethod,
		RequestUrl:    requestUrl,
		RequestParams: requestParams,
		ResponseData:  responseData,
		IpAddress:     ipAddress,
		UserAgent:     userAgent,
		Status:        status,
		ErrorMessage:  errMsg,
		ExecutionTime: executionTime,
	})

	return err
}
