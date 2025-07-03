/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 17:12:16
 * @LastEditTime: 2025-07-02 17:23:03
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/manager/manager_v1_manage_list.go
 */
package manager

import (
	"context"

	v1 "go-dora-api/api/manager/v1"
	"go-dora-api/internal/consts"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// ManageList 函数用于处理管理列表请求
func (c *ControllerV1) ManageList(ctx context.Context, req *v1.ManageListReq) (res *v1.ManageListRes, err error) {
	// 创建一个 ManageListInput 结构体
	search := new(model.ManageListInput)
	// 将请求中的页码赋值给 search 结构体
	search.Page = req.Page
	// 如果页码小于等于0，则将其设置为默认页码
	if search.Page <= 0 {
		search.Page = consts.DEFAULT_PAGE
	}
	// 将请求中的每页数量赋值给 search 结构体
	search.PageSize = req.PageSize
	// 如果每页数量小于等于0，则将其设置为默认每页数量
	if search.PageSize <= 0 {
		search.PageSize = consts.DEFAULT_PAGE_SIZE
	}
	// 将请求中的搜索条件赋值给 search 结构体
	search.Search = req.Search

	// 调用 service 中的 GetList 函数，获取管理列表
	list, total, err := service.Manage().GetList(ctx, search)
	// 如果出现错误，则返回错误
	if err != nil {
		return nil, err
	}

	// 创建一个 ManageListRes 结构体
	r := new(model.ManageListRes)
	// 将获取到的列表赋值给 r 结构体
	r.List = list
	// 将获取到的总数赋值给 r 结构体
	r.Total = total
	// 将 search 结构体中的页码赋值给 r 结构体
	r.Page = search.Page
	// 将 search 结构体中的每页数量赋值给 r 结构体
	r.PageSize = search.PageSize

	// 将 r 结构体转换为 v1.ManageListRes 类型
	resp := v1.ManageListRes(r)

	// 返回结果
	return &resp, nil
}
