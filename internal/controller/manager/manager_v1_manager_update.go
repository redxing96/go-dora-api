/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-29 11:50:54
 * @LastEditTime: 2025-06-29 11:58:24
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/manager/manager_v1_manager_update.go
 */
package manager

import (
	"context"

	v1 "go-dora-api/api/manager/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// ManagerUpdate函数用于更新管理员的个人信息
func (c *ControllerV1) ManagerUpdate(ctx context.Context, req *v1.ManagerUpdateReq) (res *v1.ManagerUpdateRes, err error) {
	// 调用service.Manage().Update函数更新管理员的个人信息
	result, err := service.Manage().Update(ctx, &model.UpdateInput{
		Id:      req.Id,      // 管理员ID
		Account: req.Account, // 管理员账号
		Email:   req.Email,   // 管理员邮箱
		Phone:   req.Phone,   // 管理员电话
		Avatar:  req.Avatar,  // 管理员头像
	})

	// 如果更新失败，则返回错误
	if err != nil {
		return
	}

	// 将更新后的管理员信息封装成ManagerDetailRes结构体
	r := &model.ManagerDetailRes{
		Id:      result.Id,
		Account: result.Account,
		Email:   result.Email,
		Phone:   result.Phone,
		Avatar:  result.Avatar,
	}
	// 将ManagerDetailRes结构体封装成ManagerUpdateRes结构体
	resp := v1.ManagerUpdateRes(r)

	// 返回更新后的管理员信息
	return &resp, nil
}
