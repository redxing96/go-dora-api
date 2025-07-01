// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Manage is the golang structure for table manage.
type Manage struct {
	Id         int         `json:"id"         orm:"id"          description:""`                   //
	Account    string      `json:"account"    orm:"account"     description:"用户账号"`               // 用户账号
	Password   string      `json:"password"   orm:"password"    description:"用户密码"`               // 用户密码
	Email      string      `json:"email"      orm:"email"       description:"电子邮箱"`               // 电子邮箱
	Phone      string      `json:"phone"      orm:"phone"       description:"电话号码"`               // 电话号码
	Avatar     string      `json:"avatar"     orm:"avatar"      description:"头像"`                 // 头像
	Status     int         `json:"status"     orm:"status"      description:"状态 0-初始化 1-正常 2-冻结"` // 状态 0-初始化 1-正常 2-冻结
	IsSuper    int         `json:"isSuper"    orm:"is_super"    description:"是否超管 1-是"`           // 是否超管 1-是
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`               // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:"修改时间"`               // 修改时间
}
