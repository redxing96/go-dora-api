/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 09:32:59
 * @LastEditTime: 2025-07-02 13:21:23
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/dictionary/update.go
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

// 更新字典
func (s *sDictionary) Update(ctx context.Context, in *model.DictionaryUpdateInput) (out *model.DictionaryUpdateOutput, err error) {
	// 创建查询对象
	query := dao.Dictionary.Ctx(ctx).Where(dao.Dictionary.Columns().Id, in.Id)
	// 更新字典数据
	query.Data(do.Dictionary{
		Name:       in.Name,
		Code:       in.Code,
		Type:       in.Type,
		Desc:       in.Desc,
		Status:     in.Status,
		UpdateTime: gtime.Now(),
	})
	// 执行更新操作
	result, err := query.Update()
	if err != nil {
		// 记录错误日志
		logger.SystemLogger.Errorf("更新字典失败: %s", err)
		// 返回错误信息
		err = gerror.NewCode(common_return.ErrorCode("", "update_dictionary_failed", 500))
		return
	}
	// 获取受影响的行数
	rows, err := result.RowsAffected()
	if err != nil {
		// 记录错误日志
		logger.SystemLogger.Errorf("更新字典失败: %s", err)
		// 返回错误信息
		err = gerror.NewCode(common_return.ErrorCode("", "update_dictionary_failed", 500))
		return
	}
	// 如果受影响的行数为0，则更新失败
	if rows <= 0 {
		// 记录错误日志
		logger.SystemLogger.Errorf("更新字典失败: %s", err)
		// 返回错误信息
		err = gerror.NewCode(common_return.ErrorCode("", "update_dictionary_failed", 500))
		return
	}
	// 创建返回对象
	out = new(model.DictionaryUpdateOutput)
	// 设置返回对象属性
	out.Id = in.Id
	out.Name = in.Name
	// 返回更新结果
	return out, nil
}
