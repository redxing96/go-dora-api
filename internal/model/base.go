/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 10:53:30
 * @LastEditTime: 2025-06-25 10:54:37
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/model/base.go
 */
package model

type BaseInput struct {
	Page     int    `d:"1" json:"page" dc:"当前页"`
	PageSize int    `d:"20" json:"page_size" dc:"每页条数"`
	Search   string `d:"" json:"search" dc:"搜索关键字"`
	Type     int    `d:"0" json:"type" dc:"类型"`
}

type BaseOutput struct {
	Page     int `json:"page" dc:"当前页"`
	PageSize int `json:"page_size" dc:"每页条数"`
	Total    int `json:"total" dc:"总条数"`
	List     any `json:"list" dc:"列表"`
}
