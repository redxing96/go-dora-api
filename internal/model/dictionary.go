/*
 * @Description: 字典模型
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 09:27:58
 * @LastEditTime: 2025-07-02 13:24:04
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/model/dictionary.go
 */
package model

type DictionaryDetailOutput struct {
	Id     int    `json:"id" dc:"字典ID"`
	Name   string `json:"name" dc:"字典名称"`
	Code   string `json:"code" dc:"字典编码"`
	Type   string `json:"type" dc:"字典类型"`
	Desc   string `json:"desc" dc:"字典描述"`
	Status int    `json:"status" dc:"状态 1-正常 2-禁用"`
	IsEdit int    `json:"is_edit" dc:"是否可编辑 1-是 2-否"`
}

type DictionaryListRes struct {
	BaseOutput
	List []*DictionaryDetailOutput `json:"list" dc:"字典列表"`
}

type DictionaryAddInput struct {
	Type   string `json:"type" dc:"字典类型"`
	Name   string `json:"name" dc:"字典名称"`
	Code   string `json:"code" dc:"字典编码"`
	Desc   string `json:"desc" dc:"字典描述"`
	Status int    `json:"status" dc:"状态 1-正常 2-禁用"`
	IsEdit int    `json:"is_edit" dc:"是否可编辑 1-是 2-否"`
}

type DictionaryAddOutput struct {
	Id   int    `json:"id" dc:"字典ID"`
	Name string `json:"name" dc:"字典名称"`
}

type DictionaryDeleteInput struct {
	Ids []int `json:"ids" dc:"字典ID"`
}

type DictionaryDeleteOutput struct {
	Ids []int `json:"ids" dc:"字典ID"`
}

type DictionaryGetAllInput struct {
	Type     string `json:"type" dc:"字典类型"`
	LikeName string `json:"like_name" dc:"模糊查询名称"`
	LikeCode string `json:"like_code" dc:"模糊查询编码"`
	Status   int    `json:"status" dc:"状态 1-正常 2-禁用"`
	IsEdit   int    `json:"is_edit" dc:"是否可编辑 1-是 2-否"`
}

type DictionaryListInput struct {
	BaseInput
	Type     string `json:"type" dc:"字典类型"`
	LikeName string `json:"like_name" dc:"模糊查询名称"`
	LikeCode string `json:"like_code" dc:"模糊查询编码"`
	Status   int    `json:"status" dc:"状态 1-正常 2-禁用"`
	IsEdit   int    `json:"is_edit" dc:"是否可编辑 1-是 2-否"`
}

type DictionaryUpdateInput struct {
	Id     int    `json:"id" dc:"字典ID"`
	Type   string `json:"type" dc:"字典类型"`
	Name   string `json:"name" dc:"字典名称"`
	Code   string `json:"code" dc:"字典编码"`
	Desc   string `json:"desc" dc:"字典描述"`
	Status int    `json:"status" dc:"状态 1-正常 2-禁用"`
	IsEdit int    `json:"is_edit" dc:"是否可编辑 1-是 2-否"`
}

type DictionaryUpdateOutput struct {
	Id   int    `json:"id" dc:"字典ID"`
	Name string `json:"name" dc:"字典名称"`
}
