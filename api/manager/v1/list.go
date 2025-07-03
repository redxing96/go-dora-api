package v1

import (
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ManageListReq struct {
	g.Meta   `path:"/v1/manager/list" method:"get" tags:"管理员管理" summary:"管理员列表" security:"api_key"`
	Page     int    `d:"1" json:"page" dc:"当前页"`
	PageSize int    `d:"20" json:"page_size" dc:"每页条数"`
	Search   string `d:"" json:"search" dc:"搜索关键字"`
}

type ManageListRes *model.ManageListRes
