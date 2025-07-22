/*
 * @Description: 获取字典列表
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 09:30:38
 * @LastEditTime: 2025-07-02 09:59:34
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/dictionary/dictionary_v1_dictionary_list.go
 */
package dictionary

import (
	"context"

	v1 "go-dora-api/api/dictionary/v1"
	"go-dora-api/internal/consts"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// DictionaryList 函数用于获取字典列表
func (c *ControllerV1) DictionaryList(ctx context.Context, req *v1.DictionaryListReq) (res *v1.DictionaryListRes, err error) {
	// 创建一个新的字典列表输入对象
	search := new(model.DictionaryListInput)
	// 设置字典状态
	search.Status = req.Status
	// 设置页码
	search.Page = req.Page
	// 搜索
	search.Search = req.Search
	// 如果页码小于等于0，则设置为默认页码
	if search.Page <= 0 {
		search.Page = consts.DEFAULT_PAGE
	}
	// 设置每页显示数量
	search.PageSize = req.PageSize
	// 如果每页显示数量小于等于0，则设置为默认每页显示数量
	if search.PageSize <= 0 {
		search.PageSize = consts.DEFAULT_PAGE_SIZE
	}

	// 调用服务层获取字典列表
	result, total, err := service.Dictionary().List(ctx, search)

	// 如果出现错误，则返回错误
	if err != nil {
		return nil, err
	}

	// 创建一个新的字典列表响应对象
	r := new(model.DictionaryListRes)
	// 设置字典列表
	r.List = result
	// 设置总数量
	r.Total = total
	// 设置当前页码
	r.Page = search.Page
	// 设置每页显示数量
	r.PageSize = search.PageSize

	// 将响应对象转换为v1类型
	resp := v1.DictionaryListRes(r)

	// 返回响应对象和错误
	return &resp, nil
}
