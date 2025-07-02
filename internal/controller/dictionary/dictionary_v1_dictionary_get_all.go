/*
 * @Description: 获取所有字典
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 10:02:31
 * @LastEditTime: 2025-07-02 10:02:42
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/dictionary/dictionary_v1_dictionary_get_all.go
 */
package dictionary

import (
	"context"

	v1 "go-dora-api/api/dictionary/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// DictionaryGetAll 获取所有字典
func (c *ControllerV1) DictionaryGetAll(ctx context.Context, req *v1.DictionaryGetAllReq) (res *v1.DictionaryGetAllRes, err error) {
	// 调用服务层获取所有字典
	result, err := service.Dictionary().GetAll(ctx, &model.DictionaryGetAllInput{
		Type:     req.Type,
		LikeName: req.LikeName,
		LikeCode: req.LikeCode,
		Status:   req.Status,
	})

	// 如果出现错误，则返回错误
	if err != nil {
		return nil, err
	}

	// 将结果转换为响应
	resp := v1.DictionaryGetAllRes(result)

	// 返回响应和错误
	return &resp, nil
}
