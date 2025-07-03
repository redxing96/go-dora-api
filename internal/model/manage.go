/*
 * @Description: 管理员模型
 * @Author: redxing96@163.com
 * @Date: 2025-06-24 20:05:04
 * @LastEditTime: 2025-07-02 17:21:12
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/model/manage.go
 */
package model

import (
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
	Id         int         `json:"id" dc:"管理员ID"`
	Account    string      `json:"account" dc:"管理员账号"`
	Status     int         `json:"status" dc:"管理员状态 0-初始化 1-正常 2-冻结"`
	IsSuper    int         `json:"is_super" dc:"是否超管 1-是"`
	Email      string      `json:"email" dc:"电子邮箱"`
	Phone      string      `json:"phone" dc:"电话号码"`
	Avatar     string      `json:"avatar" dc:"头像"`
	CreateTime *gtime.Time `json:"create_time" dc:"创建时间"`
	UpdateTime *gtime.Time `json:"update_time" dc:"更新时间"`
	RoleIds    []int       `json:"role_ids" dc:"角色IDs"`
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

type ManagerAddInput struct {
	Account  string `json:"account" dc:"管理员账号"`
	Password string `json:"password" dc:"管理员密码"`
	Email    string `json:"email" dc:"电子邮箱"`
	Phone    string `json:"phone" dc:"电话号码"`
	Avatar   string `json:"avatar" dc:"头像"`
	Status   int    `json:"status" dc:"状态 0-初始化 1-正常 2-冻结"`
	IsSuper  int    `json:"is_super" dc:"是否超级管理员 1-是 2-否"`
	RoleIds  []int  `json:"role_ids" dc:"角色IDs"`
}

type ManagerAddOutput struct {
	Id      int    `json:"id" dc:"管理员ID"`
	Account string `json:"account" dc:"管理员账号"`
}

type ManageListRes struct {
	BaseOutput
	List []*GetManageDetailOutput `json:"list" dc:"管理员列表"`
}

type ManageListInput struct {
	BaseInput
}
