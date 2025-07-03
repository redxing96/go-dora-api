/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-29 11:51:10
 * @LastEditTime: 2025-06-29 12:01:10
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/manage/update.go
 */
package manage

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"
	"go-dora-api/internal/model/do"
	"go-dora-api/utility/common"

	"github.com/gogf/gf/v2/errors/gerror"
)

// 更新管理员信息
func (s *sManage) Update(ctx context.Context, in *model.UpdateInput) (res *model.UpdateOutput, err error) {
	// 使用dao.Manage.Ctx(ctx)获取数据库连接，并使用Where方法指定更新条件，使用Data方法指定更新数据，使用Update方法执行更新操作

	data := do.Manage{}

	if in.Account != "" {
		data.Account = in.Account
	}
	if in.Email != "" {
		data.Email = in.Email
	}
	if in.Phone != "" {
		data.Phone = in.Phone
	}
	if in.Avatar != "" {
		data.Avatar = in.Avatar
	}

	if in.Password != "" {
		data.Password = common.GenerateFromPassword(in.Password)
	}

	if data.Account == "" && data.Email == "" && data.Phone == "" && data.Avatar == "" && in.Password == "" {
		logger.SystemLogger.Errorf("更新管理员信息失败: %v", in.Id)
		err = gerror.NewCode(common_return.ErrorCode("", "update_manage_failed", 500))
		return
	}

	result, err := dao.Manage.Ctx(ctx).Where(dao.Manage.Columns().Id, in.Id).Data(data).Update()
	if err != nil {
		// 如果更新失败，记录错误日志，并返回错误信息
		logger.SystemLogger.Errorf("更新管理员信息失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "update_manage_failed", 500))
		return
	}

	// 获取受影响的行数
	rowsAffected, err := result.RowsAffected()

	// 如果受影响的行数为0，说明管理员不存在，记录错误日志，并返回错误信息
	if err != nil || rowsAffected == 0 {
		logger.SystemLogger.Errorf("更新管理员信息失败: %v", in.Id)
		err = gerror.NewCode(common_return.ErrorCode("", "manage_update_not_found", 500))
		return
	}

	// 构造更新后的管理员信息
	res = &model.UpdateOutput{
		Id:      in.Id,
		Account: in.Account,
		Email:   in.Email,
		Phone:   in.Phone,
		Avatar:  in.Avatar,
	}
	return
}
