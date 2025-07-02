/*
 * @Description: 添加字典
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 09:33:52
 * @LastEditTime: 2025-07-02 09:44:26
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/dictionary/add.go
 */
package dictionary

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"
	"go-dora-api/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// 向字典中添加一个新条目
func (s *sDictionary) Add(ctx context.Context, in *model.DictionaryAddInput) (out *model.DictionaryAddOutput, err error) {
	// 向数据库中插入新条目，并获取插入后的ID
	result, err := dao.Dictionary.Ctx(ctx).Data(do.Dictionary{
		Name:       in.Name,
		Code:       in.Code,
		Type:       in.Type,
		Desc:       in.Desc,
		Status:     in.Status,
		CreateTime: gtime.Now(),
	}).InsertAndGetId()
	// 如果插入失败或者插入的ID小于等于0，则返回错误
	if err != nil || result <= 0 {
		logger.SystemLogger.Errorf("添加字典失败: %s", err)
		err = gerror.NewCode(common_return.ErrorCode("", "add_dictionary_failed", 500))
		return
	}

	data := new(model.DictionaryAddOutput)
	// 将插入后的ID赋值给输出参数
	data.Id = int(result)
	// 将输入参数中的名称赋值给输出参数
	data.Name = in.Name

	return data, nil
}
