/*
 * @Description: 添加管理员
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 10:08:11
 * @LastEditTime: 2025-07-02 10:14:59
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/manage/add.go
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
	"github.com/gogf/gf/v2/os/gtime"
)

// 添加管理员
func (s *sManage) Add(ctx context.Context, in *model.ManagerAddInput) (res *model.ManagerAddOutput, err error) {
	// 向数据库中插入管理员信息
	id, err := dao.Manage.Ctx(ctx).Data(do.Manage{
		Account:    in.Account,                               // 账号
		Password:   common.GenerateFromPassword(in.Password), // 密码
		Email:      in.Email,                                 // 邮箱
		Phone:      in.Phone,                                 // 手机号
		Avatar:     in.Avatar,                                // 头像
		Status:     in.Status,                                // 状态
		IsSuper:    in.IsSuper,                               // 是否为超级管理员
		CreateTime: gtime.Now(),                              // 创建时间
	}).OmitEmpty().InsertAndGetId()
	if err != nil {
		// 插入失败，记录日志并返回错误
		logger.SystemLogger.Errorf("添加管理员失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "add_manage_error", 500))
		return nil, err
	}

	// 向数据库中插入管理员角色信息
	dao.SysManagerRole.Ctx(ctx).Data(do.SysManagerRole{
		ManagerId: int(id),    // 管理员ID
		RoleId:    in.RoleIds, // 角色ID
	}).OmitEmpty().Insert()

	// 返回管理员信息
	return &model.ManagerAddOutput{
		Id:      int(id),    // 管理员ID
		Account: in.Account, // 账号
	}, nil
}
