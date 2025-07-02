// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Dictionary is the golang structure for table dictionary.
type Dictionary struct {
	Id         int         `json:"id"         orm:"id"          description:""`              //
	Type       string      `json:"type"       orm:"type"        description:"类型"`            // 类型
	Name       string      `json:"name"       orm:"name"        description:"名称"`            // 名称
	Code       string      `json:"code"       orm:"code"        description:"编码"`            // 编码
	Desc       string      `json:"desc"       orm:"desc"        description:"描述"`            // 描述
	IsEdit     int         `json:"isEdit"     orm:"is_edit"     description:"是否可编辑 1-是 2-否"` // 是否可编辑 1-是 2-否
	Status     int         `json:"status"     orm:"status"      description:"状态 1-正常 2-禁用"`  // 状态 1-正常 2-禁用
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`          // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:"更新时间"`          // 更新时间
}
