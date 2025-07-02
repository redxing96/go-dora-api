/*
 * @Description: 添加字典
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 09:30:38
 * @LastEditTime: 2025-07-02 09:54:46
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/dictionary/dictionary_v1_dictionary_add.go
 */
package dictionary

import (
	"context"

	v1 "go-dora-api/api/dictionary/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

func (c *ControllerV1) DictionaryAdd(ctx context.Context, req *v1.DictionaryAddReq) (res *v1.DictionaryAddRes, err error) {

	result, err := service.Dictionary().Add(ctx, &model.DictionaryAddInput{
		Name:   req.Name,
		Code:   req.Code,
		Type:   req.Type,
		Desc:   req.Desc,
		Status: req.Status,
	})

	if err != nil {
		return nil, err
	}

	res = new(v1.DictionaryAddRes)
	res.Id = result.Id
	res.Name = result.Name

	return res, nil
}
