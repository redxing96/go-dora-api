/*
 * @Description: 删除字典
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 09:30:38
 * @LastEditTime: 2025-07-02 09:55:26
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/dictionary/dictionary_v1_dictionary_delete.go
 */
package dictionary

import (
	"context"

	v1 "go-dora-api/api/dictionary/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

func (c *ControllerV1) DictionaryDelete(ctx context.Context, req *v1.DictionaryDeleteReq) (res *v1.DictionaryDeleteRes, err error) {
	result, err := service.Dictionary().Delete(ctx, &model.DictionaryDeleteInput{
		Ids: req.Ids,
	})

	if err != nil {
		return nil, err
	}

	res = new(v1.DictionaryDeleteRes)
	res.Ids = result.Ids

	return res, nil
}
