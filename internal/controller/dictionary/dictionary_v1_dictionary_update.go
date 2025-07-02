/*
 * @Description: 更新字典
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 09:30:38
 * @LastEditTime: 2025-07-02 10:00:18
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/dictionary/dictionary_v1_dictionary_update.go
 */
package dictionary

import (
	"context"

	v1 "go-dora-api/api/dictionary/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// DictionaryUpdate 函数用于更新字典信息
func (c *ControllerV1) DictionaryUpdate(ctx context.Context, req *v1.DictionaryUpdateReq) (res *v1.DictionaryUpdateRes, err error) {
	// 调用 service 层的 Dictionary() 方法，并传入 DictionaryUpdateInput 结构体
	result, err := service.Dictionary().Update(ctx, &model.DictionaryUpdateInput{
		Id:     req.Id,     // 字典ID
		Name:   req.Name,   // 字典名称
		Code:   req.Code,   // 字典编码
		Type:   req.Type,   // 字典类型
		Desc:   req.Desc,   // 字典描述
		Status: req.Status, // 字典状态
	})

	// 如果更新失败，则返回错误
	if err != nil {
		return nil, err
	}

	// 创建 DictionaryUpdateRes 结构体
	res = new(v1.DictionaryUpdateRes)
	res.Id = result.Id
	res.Name = result.Name

	// 返回更新后的字典信息
	return res, nil
}
