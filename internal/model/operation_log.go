/*
 * @Description: 操作日志模型
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 10:00:00
 * @LastEditTime: 2025-07-02 10:00:00
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/model/operation_log.go
 */
package model

import "go-dora-api/internal/model/entity"

// 操作日志创建输入
type OperationLogCreateInput struct {
	UserId        int64  `json:"userId"        description:"操作用户ID"`
	Username      string `json:"username"      description:"操作用户名"`
	Operation     string `json:"operation"     description:"操作类型"`
	Module        string `json:"module"        description:"操作模块"`
	Description   string `json:"description"   description:"操作描述"`
	RequestMethod string `json:"requestMethod" description:"请求方法"`
	RequestUrl    string `json:"requestUrl"    description:"请求URL"`
	RequestParams string `json:"requestParams" description:"请求参数"`
	ResponseData  string `json:"responseData"  description:"响应数据"`
	IpAddress     string `json:"ipAddress"     description:"IP地址"`
	UserAgent     string `json:"userAgent"     description:"用户代理"`
	Status        int    `json:"status"        description:"操作状态 1-成功 2-失败"`
	ErrorMessage  string `json:"errorMessage"  description:"错误信息"`
	ExecutionTime int    `json:"executionTime" description:"执行时间(毫秒)"`
}

// 操作日志创建输出
type OperationLogCreateOutput struct {
	Id int64 `json:"id" description:"日志ID"`
}

// 操作日志查询输入
type OperationLogListInput struct {
	Page      int    `json:"page"       description:"页码"`
	PageSize  int    `json:"pageSize"   description:"每页数量"`
	UserId    int64  `json:"userId"     description:"操作用户ID"`
	Username  string `json:"username"   description:"操作用户名"`
	Operation string `json:"operation"  description:"操作类型"`
	Module    string `json:"module"     description:"操作模块"`
	Status    int    `json:"status"     description:"操作状态"`
	StartTime string `json:"startTime"  description:"开始时间"`
	EndTime   string `json:"endTime"    description:"结束时间"`
	Keyword   string `json:"keyword"    description:"关键词搜索"`
}

// 操作日志查询输出
type OperationLogListOutput struct {
	List  []*entity.OperationLog `json:"list"  description:"日志列表"`
	Total int                    `json:"total" description:"总数"`
}

// 操作日志详情输出
type OperationLogDetailOutput struct {
	*entity.OperationLog
}

// 操作日志删除输入
type OperationLogDeleteInput struct {
	Ids []int64 `json:"ids" description:"日志ID列表"`
}

// 操作日志删除输出
type OperationLogDeleteOutput struct {
	Ids []int64 `json:"ids" description:"删除的日志ID列表"`
}
