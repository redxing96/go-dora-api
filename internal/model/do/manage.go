// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Manage is the golang structure of table manage for DAO operations like Where/Data.
type Manage struct {
	g.Meta     `orm:"table:manage, do:true"`
	Id         interface{} //
	Account    interface{} // 用户账号
	Password   interface{} // 用户密码
	Email      interface{} // 电子邮箱
	Phone      interface{} // 电话号码
	Avatar     interface{} // 头像
	Status     interface{} // 状态 0-初始化 1-正常 2-冻结
	IsSuper    interface{} // 是否超管 1-是
	IsDelete   interface{} // 是否删除 1-是 2-否
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 修改时间
}
