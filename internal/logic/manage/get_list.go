/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 17:16:10
 * @LastEditTime: 2025-07-02 17:21:09
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/manage/get_list.go
 */
package manage

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/errors/gerror"
)

// 获取管理员列表
func (s *sManage) GetList(ctx context.Context, in *model.ManageListInput) (out []*model.GetManageDetailOutput, total int, err error) {
	// 创建查询对象
	query := dao.Manage.Ctx(ctx)
	// 如果有搜索条件，则添加搜索条件
	if in.Search != "" {
		query = query.WhereLike(dao.Manage.Columns().Account, "%"+in.Search+"%").
			WhereOrLike(dao.Manage.Columns().Email, "%"+in.Search+"%").
			WhereOrLike(dao.Manage.Columns().Phone, "%"+in.Search+"%")
	}
	// 设置查询条件-禁用不查询
	query = query.Where(dao.Manage.Columns().IsDelete, 2)
	// 设置分页
	query = query.Page(in.Page, in.PageSize)
	// 查询并统计结果
	err = query.Order("id desc").ScanAndCount(&out, &total, true)
	// 如果查询失败，则记录错误日志并返回错误
	if err != nil {
		logger.SystemLogger.Errorf("获取管理员列表失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_manage_list_error", 500))
		return
	}
	// 返回结果
	return
}
