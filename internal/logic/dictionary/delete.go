package dictionary

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/errors/gerror"
)

// 删除字典
func (s *sDictionary) Delete(ctx context.Context, in *model.DictionaryDeleteInput) (out *model.DictionaryDeleteOutput, err error) {
	// 从数据库中删除指定id的字典
	result, err := dao.Dictionary.Ctx(ctx).WhereIn(dao.Dictionary.Columns().Id, in.Ids).Delete()
	if err != nil {
		// 如果删除失败，记录错误日志，并返回错误信息
		logger.SystemLogger.Errorf("删除字典失败: %s", err)
		err = gerror.NewCode(common_return.ErrorCode("", "delete_dictionary_failed", 500))
		return
	}
	// 获取删除的行数
	rows, err := result.RowsAffected()
	if err != nil {
		// 如果获取行数失败，记录错误日志，并返回错误信息
		logger.SystemLogger.Errorf("删除字典失败: %s", err)
		err = gerror.NewCode(common_return.ErrorCode("", "delete_dictionary_failed", 500))
		return
	}
	// 如果删除的行数为0，记录错误日志，并返回错误信息
	if rows <= 0 {
		logger.SystemLogger.Errorf("删除字典失败: %s", err)
		err = gerror.NewCode(common_return.ErrorCode("", "delete_dictionary_failed", 500))
		return
	}

	// 创建返回结果
	res := new(model.DictionaryDeleteOutput)

	// 设置返回结果中的id
	res.Ids = in.Ids

	// 返回结果和错误信息
	return res, nil
}
