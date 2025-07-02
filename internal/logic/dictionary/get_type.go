/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 13:28:22
 * @LastEditTime: 2025-07-02 13:44:49
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/dictionary/get_type.go
 */
package dictionary

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"

	"github.com/gogf/gf/v2/errors/gerror"
)

// 获取字典类型
func (s *sDictionary) GetType(ctx context.Context) (out []string, err error) {
	var types = []struct{ Type string }{}
	// 从数据库中获取字典类型
	err = dao.Dictionary.Ctx(ctx).Group(dao.Dictionary.Columns().Type).Fields(dao.Dictionary.Columns().Type).Scan(&types)
	// 如果获取失败，则记录错误日志，并返回错误信息
	if err != nil {
		logger.SystemLogger.Errorf("获取字典类型失败: %s", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_dictionary_type_failed", 500))
		return
	}
	for _, v := range types {
		out = append(out, v.Type)
	}
	return
}
