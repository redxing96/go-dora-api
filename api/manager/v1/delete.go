package v1

import "github.com/gogf/gf/v2/frame/g"

type ManagerDeleteReq struct {
	g.Meta `path:"/v1/manager/delete" method:"get" tags:"管理员管理" summary:"删除管理员" security:"api_key"`
	Ids    []int `v:"required" json:"ids" dc:"管理员ID"`
}

type ManagerDeleteRes struct {
	Ids []int `json:"ids" dc:"管理员ID"`
}
