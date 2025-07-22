// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OperationLog is the golang structure for table operation_log.
type OperationLog struct {
	Id            int64       `json:"id"            orm:"id"             description:"日志ID"`           // 日志ID
	UserId        int64       `json:"userId"        orm:"user_id"        description:"操作用户ID"`         // 操作用户ID
	Username      string      `json:"username"      orm:"username"       description:"操作用户名"`          // 操作用户名
	Operation     string      `json:"operation"     orm:"operation"      description:"操作类型"`           // 操作类型
	Module        string      `json:"module"        orm:"module"         description:"操作模块"`           // 操作模块
	Description   string      `json:"description"   orm:"description"    description:"操作描述"`           // 操作描述
	RequestMethod string      `json:"requestMethod" orm:"request_method" description:"请求方法"`           // 请求方法
	RequestUrl    string      `json:"requestUrl"    orm:"request_url"    description:"请求URL"`          // 请求URL
	RequestParams string      `json:"requestParams" orm:"request_params" description:"请求参数"`           // 请求参数
	ResponseData  string      `json:"responseData"  orm:"response_data"  description:"响应数据"`           // 响应数据
	IpAddress     string      `json:"ipAddress"     orm:"ip_address"     description:"IP地址"`           // IP地址
	UserAgent     string      `json:"userAgent"     orm:"user_agent"     description:"用户代理"`           // 用户代理
	Status        int         `json:"status"        orm:"status"         description:"操作状态 1-成功 2-失败"` // 操作状态 1-成功 2-失败
	ErrorMessage  string      `json:"errorMessage"  orm:"error_message"  description:"错误信息"`           // 错误信息
	ExecutionTime int         `json:"executionTime" orm:"execution_time" description:"执行时间(毫秒)"`       // 执行时间(毫秒)
	CreateTime    *gtime.Time `json:"createTime"    orm:"create_time"    description:"创建时间"`           // 创建时间
}
