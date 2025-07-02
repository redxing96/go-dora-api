/*
 * @Description: 获取字典类型
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 13:27:45
 * @LastEditTime: 2025-07-02 13:34:35
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/dictionary/dictionary_v1_dictionary_get_type.go
 */
package dictionary

import (
	"context"

	v1 "go-dora-api/api/dictionary/v1"
	"go-dora-api/internal/service"
)

func (c *ControllerV1) DictionaryGetType(ctx context.Context, req *v1.DictionaryGetTypeReq) (res *v1.DictionaryGetTypeRes, err error) {
	result, err := service.Dictionary().GetType(ctx)
	if err != nil {
		return nil, err
	}

	resp := v1.DictionaryGetTypeRes(result)

	return &resp, nil
}
