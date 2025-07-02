/*
 * @Description: 根据类型获取所有字典
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 09:34:16
 * @LastEditTime: 2025-07-02 09:52:11
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/dictionary/get_all.go
 */
package dictionary

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/errors/gerror"
)

// GetAll函数用于获取所有字典数据
func (s *sDictionary) GetAll(ctx context.Context, in *model.DictionaryGetAllInput) (out []*model.DictionaryDetailOutput, err error) {
	// 创建查询对象
	query := dao.Dictionary.Ctx(ctx)

	// 如果输入参数中的类型不为空，则添加查询条件
	if in.Type != "" {
		query = query.Where(dao.Dictionary.Columns().Type, in.Type)
	}
	// 如果输入参数中的名称不为空，则添加查询条件
	if in.LikeName != "" {
		query = query.Where(dao.Dictionary.Columns().Name, "like", "%"+in.LikeName+"%")
	}
	// 如果输入参数中的编码不为空，则添加查询条件
	if in.LikeCode != "" {
		query = query.Where(dao.Dictionary.Columns().Code, "like", "%"+in.LikeCode+"%")
	}
	// 如果输入参数中的状态不为0，则添加查询条件
	if in.Status != 0 {
		query = query.Where(dao.Dictionary.Columns().Status, in.Status)
	}

	// 添加排序条件，按照id降序排列
	query.Order(dao.Dictionary.Columns().Id, "desc")

	// 执行查询，并将结果扫描到out中
	err = query.Scan(&out)

	if err != nil {
		logger.SystemLogger.Errorf("获取字典数据失败: %s", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_dictionary_failed", 500))
		return
	}

	// 返回查询结果和错误信息
	return out, err
}
