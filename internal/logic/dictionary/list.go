package dictionary

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/errors/gerror"
)

// List函数用于查询字典列表
func (s *sDictionary) List(ctx context.Context, in *model.DictionaryListInput) (out []*model.DictionaryDetailOutput, total int, err error) {
	// 创建查询对象
	query := dao.Dictionary.Ctx(ctx)

	if in.Search != "" {
		query = query.WhereLike(dao.Dictionary.Columns().Type, "%"+in.Search+"%").
			WhereOrLike(dao.Dictionary.Columns().Name, "%"+in.Search+"%").
			WhereOrLike(dao.Dictionary.Columns().Code, "%"+in.Search+"%")
	}

	// 如果输入参数中的状态不为0，则添加查询条件
	if in.Status != 0 {
		query = query.Where(dao.Dictionary.Columns().Status, in.Status)
	}

	// 按照id降序排列
	query.Order(dao.Dictionary.Columns().Id, "desc")

	// 执行查询，并将结果扫描到out中，同时返回总记录数
	err = query.ScanAndCount(&out, &total, true)
	if err != nil {
		logger.SystemLogger.Errorf("查询字典列表失败: %s", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_dictionary_list_failed", 500))
		return
	}
	return
}
