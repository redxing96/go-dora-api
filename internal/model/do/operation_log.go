/*
 * @Description: 操作日志DO模型
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 10:00:00
 * @LastEditTime: 2025-07-02 10:00:00
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/model/do/operation_log.go
 */
package do

// 操作日志DO模型
type OperationLog struct {
	Id            interface{} `json:"id"            description:"日志ID"`
	UserId        interface{} `json:"userId"        description:"操作用户ID"`
	Username      interface{} `json:"username"      description:"操作用户名"`
	Operation     interface{} `json:"operation"     description:"操作类型"`
	Module        interface{} `json:"module"        description:"操作模块"`
	Description   interface{} `json:"description"   description:"操作描述"`
	RequestMethod interface{} `json:"requestMethod" description:"请求方法"`
	RequestUrl    interface{} `json:"requestUrl"    description:"请求URL"`
	RequestParams interface{} `json:"requestParams" description:"请求参数"`
	ResponseData  interface{} `json:"responseData"  description:"响应数据"`
	IpAddress     interface{} `json:"ipAddress"     description:"IP地址"`
	UserAgent     interface{} `json:"userAgent"     description:"用户代理"`
	Status        interface{} `json:"status"        description:"操作状态 1-成功 2-失败"`
	ErrorMessage  interface{} `json:"errorMessage"  description:"错误信息"`
	ExecutionTime interface{} `json:"executionTime" description:"执行时间(毫秒)"`
	CreateTime    interface{} `json:"createTime"    description:"创建时间"`
}
