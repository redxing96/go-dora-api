// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Dictionary is the golang structure of table dictionary for DAO operations like Where/Data.
type Dictionary struct {
	g.Meta     `orm:"table:dictionary, do:true"`
	Id         interface{} //
	Type       interface{} // 类型
	Name       interface{} // 名称
	Code       interface{} // 编码
	Desc       interface{} // 描述
	IsEdit     interface{} // 是否可编辑 1-是 2-否
	Status     interface{} // 状态 1-正常 2-禁用
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
}
