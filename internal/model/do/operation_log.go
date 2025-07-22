// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OperationLog is the golang structure of table operation_log for DAO operations like Where/Data.
type OperationLog struct {
	g.Meta        `orm:"table:operation_log, do:true"`
	Id            interface{} // 日志ID
	UserId        interface{} // 操作用户ID
	Username      interface{} // 操作用户名
	Operation     interface{} // 操作类型
	Module        interface{} // 操作模块
	Description   interface{} // 操作描述
	RequestMethod interface{} // 请求方法
	RequestUrl    interface{} // 请求URL
	RequestParams interface{} // 请求参数
	ResponseData  interface{} // 响应数据
	IpAddress     interface{} // IP地址
	UserAgent     interface{} // 用户代理
	Status        interface{} // 操作状态 1-成功 2-失败
	ErrorMessage  interface{} // 错误信息
	ExecutionTime interface{} // 执行时间(毫秒)
	CreateTime    *gtime.Time // 创建时间
}
