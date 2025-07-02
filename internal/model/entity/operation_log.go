/*
 * @Description: 操作日志实体
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 10:00:00
 * @LastEditTime: 2025-07-02 10:00:00
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/model/entity/operation_log.go
 */
package entity

import "github.com/gogf/gf/v2/os/gtime"

// 操作日志实体
type OperationLog struct {
	Id            int64       `json:"id"            description:"日志ID"`
	UserId        int64       `json:"userId"        description:"操作用户ID"`
	Username      string      `json:"username"      description:"操作用户名"`
	Operation     string      `json:"operation"     description:"操作类型"`
	Module        string      `json:"module"        description:"操作模块"`
	Description   string      `json:"description"   description:"操作描述"`
	RequestMethod string      `json:"requestMethod" description:"请求方法"`
	RequestUrl    string      `json:"requestUrl"    description:"请求URL"`
	RequestParams string      `json:"requestParams" description:"请求参数"`
	ResponseData  string      `json:"responseData"  description:"响应数据"`
	IpAddress     string      `json:"ipAddress"     description:"IP地址"`
	UserAgent     string      `json:"userAgent"     description:"用户代理"`
	Status        int         `json:"status"        description:"操作状态 1-成功 2-失败"`
	ErrorMessage  string      `json:"errorMessage"  description:"错误信息"`
	ExecutionTime int         `json:"executionTime" description:"执行时间(毫秒)"`
	CreateTime    *gtime.Time `json:"createTime"    description:"创建时间"`
}
