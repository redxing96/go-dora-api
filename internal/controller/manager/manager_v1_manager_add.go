/*
 * @Description: 添加管理员
 * @Author: redxing96@163.com
 * @Date: 2025-07-02 10:07:59
 * @LastEditTime: 2025-07-02 10:17:23
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/manager/manager_v1_manager_add.go
 */
package manager

import (
	"context"

	v1 "go-dora-api/api/manager/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

func (c *ControllerV1) ManagerAdd(ctx context.Context, req *v1.ManagerAddReq) (res *v1.ManagerAddRes, err error) {
	result, err := service.Manage().Add(ctx, &model.ManagerAddInput{
		Account:  req.Account,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
		Avatar:   req.Avatar,
		Status:   req.Status,
		IsSuper:  req.IsSuper,
		RoleIds:  req.RoleIds,
	})

	if err != nil {
		return nil, err
	}

	res = new(v1.ManagerAddRes)
	res.Id = result.Id
	res.Account = result.Account

	return res, nil
}
