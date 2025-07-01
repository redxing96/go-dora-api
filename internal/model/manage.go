package model

import (
	"go-dora-api/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

type GetManageDetailInput struct {
	Id      int    `json:"id" dc:"管理员ID"`
	Account string `json:"account" dc:"管理员账号"`
}

type GetManageDetailOutput struct {
	Id         int         `json:"id" dc:"管理员ID"`
	Account    string      `json:"account" dc:"管理员账号"`
	Password   string      `json:"password" dc:"管理员密码"`
	Status     int         `json:"status" dc:"管理员状态 0-初始化 1-正常 2-冻结"`
	IsSuper    int         `json:"is_super" dc:"是否超管 1-是"`
	Email      string      `json:"email" dc:"电子邮箱"`
	Phone      string      `json:"phone" dc:"电话号码"`
	Avatar     string      `json:"avatar" dc:"头像"`
	CreateTime *gtime.Time `json:"create_time" dc:"创建时间"`
	UpdateTime *gtime.Time `json:"update_time" dc:"更新时间"`
}

type ManagerDetailRes struct {
	Id         int               `json:"id" dc:"管理员ID"`
	Account    string            `json:"account" dc:"管理员账号"`
	Status     int               `json:"status" dc:"管理员状态 0-初始化 1-正常 2-冻结"`
	IsSuper    int               `json:"is_super" dc:"是否超管 1-是"`
	Email      string            `json:"email" dc:"电子邮箱"`
	Phone      string            `json:"phone" dc:"电话号码"`
	Avatar     string            `json:"avatar" dc:"头像"`
	CreateTime *gtime.Time       `json:"create_time" dc:"创建时间"`
	UpdateTime *gtime.Time       `json:"update_time" dc:"更新时间"`
	Role       []*entity.SysRole `json:"role" dc:"角色"`
}

type UpdateInput struct {
	Id       int    `json:"id" dc:"管理员ID"`
	Account  string `json:"account" dc:"管理员账号"`
	Password string `json:"password" dc:"管理员密码"`
	Email    string `json:"email" dc:"电子邮箱"`
	Phone    string `json:"phone" dc:"电话号码"`
	Avatar   string `json:"avatar" dc:"头像"`
}

type UpdateOutput struct {
	Id      int    `json:"id" dc:"管理员ID"`
	Account string `json:"account" dc:"管理员账号"`
	Email   string `json:"email" dc:"电子邮箱"`
	Phone   string `json:"phone" dc:"电话号码"`
	Avatar  string `json:"avatar" dc:"头像"`
}
