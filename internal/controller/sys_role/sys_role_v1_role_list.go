/*
 * @Description: 获取角色列表
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 21:40:44
 * @LastEditTime: 2025-07-01 20:48:15
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/sys_role/sys_role_v1_role_list.go
 */
package sys_role

import (
	"context"

	v1 "go-dora-api/api/sys_role/v1"
	"go-dora-api/internal/consts"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// RoleList 函数用于获取角色列表
func (c *ControllerV1) RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	// 创建一个 SysRoleListInput 结构体
	search := new(model.SysRoleListInput)
	// 将请求中的页码赋值给 search 结构体
	search.Page = req.Page
	// 如果页码小于等于0，则将默认页码赋值给 search 结构体
	if search.Page <= 0 {
		search.Page = consts.DEFAULT_PAGE
	}
	// 将请求中的每页数量赋值给 search 结构体
	search.PageSize = req.PageSize
	// 如果每页数量小于等于0，则将默认每页数量赋值给 search 结构体
	if search.PageSize <= 0 {
		search.PageSize = consts.DEFAULT_PAGE_SIZE
	}
	// 将请求中的搜索关键字赋值给 search 结构体
	search.Search = req.Search

	// 调用 service.SysRole().List 函数获取角色列表
	result, total, err := service.SysRole().List(ctx, search)
	if err != nil {
		// 如果发生错误，则返回错误
		return nil, err
	}

	// 创建一个 SysRoleListRes 结构体
	r := new(model.SysRoleListRes)
	// 将获取到的角色列表赋值给 r 结构体
	for _, item := range result {
		r.List = append(r.List, &model.SysRoleItemOutput{
			Id:         item.Id,
			RoleName:   item.RoleName,
			RoleCode:   item.RoleCode,
			RoleDesc:   item.RoleDesc,
			Status:     item.Status,
			CreateTime: item.CreateTime,
		})
	}

	r.Total = total
	// 将 search 结构体中的页码赋值给 r 结构体
	r.Page = search.Page
	// 将 search 结构体中的每页数量赋值给 r 结构体
	r.PageSize = search.PageSize

	// 将 r 结构体转换为 v1.RoleListRes 结构体
	resp := v1.RoleListRes(r)

	// 返回结果
	return &resp, nil
}
