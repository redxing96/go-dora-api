/*
 * @Description: 获取管理员详情
 * @Author: redxing96@163.com
 * @Date: 2025-06-24 20:04:34
 * @LastEditTime: 2025-06-29 11:50:21
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/manage/get_manage_detail.go
 */
package manage

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"
	"go-dora-api/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

// 根据账户获取管理员详情
func (s *sManage) GetManageDetailToAccount(ctx context.Context, in *model.GetManageDetailInput) (res *model.GetManageDetailOutput, err error) {
	// 定义一个管理员列表
	var (
		list []*entity.Manage
	)
	// 根据账户查询管理员
	err = dao.Manage.Ctx(ctx).Where(dao.Manage.Columns().Account, in.Account).Scan(&list)
	// 如果查询失败，记录错误日志，并返回错误
	if err != nil {
		logger.SystemLogger.Errorf("获取管理员详情失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "manager_not_found", 500))
		return
	}

	// 如果查询结果为空，记录错误日志，并返回错误
	if len(list) == 0 {
		logger.SystemLogger.Errorf("管理员不存在: %v", in.Account)
		err = gerror.NewCode(common_return.ErrorCode("", "manager_not_found", 500))
		return
	}

	// 构造返回结果
	res = &model.GetManageDetailOutput{
		Id:       list[0].Id,
		Account:  list[0].Account,
		Email:    list[0].Email,
		Phone:    list[0].Phone,
		Avatar:   list[0].Avatar,
		Status:   list[0].Status,
		IsSuper:  list[0].IsSuper,
		Password: list[0].Password,
	}

	return
}

// 根据管理员ID获取管理员详情
func (s *sManage) GetManageDetailToId(ctx context.Context, in *model.GetManageDetailInput) (res *model.GetManageDetailOutput, err error) {
	// 定义一个管理员列表
	var (
		list []*entity.Manage
	)
	// 根据管理员ID查询管理员详情
	err = dao.Manage.Ctx(ctx).Where(dao.Manage.Columns().Id, in.Id).Scan(&list)
	// 如果查询失败，记录错误日志，并返回错误信息
	if err != nil {
		logger.SystemLogger.Errorf("获取管理员详情失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "manager_not_found", 500))
		return
	}
	// 如果查询结果为空，记录错误日志，并返回错误信息
	if len(list) == 0 {
		logger.SystemLogger.Errorf("管理员不存在: %v", in.Id)
		err = gerror.NewCode(common_return.ErrorCode("", "manager_not_found", 500))
		return
	}

	// 构造管理员详情输出
	res = &model.GetManageDetailOutput{
		Id:         list[0].Id,
		Account:    list[0].Account,
		Status:     list[0].Status,
		IsSuper:    list[0].IsSuper,
		Password:   list[0].Password,
		Email:      list[0].Email,
		Phone:      list[0].Phone,
		Avatar:     list[0].Avatar,
		CreateTime: list[0].CreateTime,
		UpdateTime: list[0].UpdateTime,
	}

	return
}
