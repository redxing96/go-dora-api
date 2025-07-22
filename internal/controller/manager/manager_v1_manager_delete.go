package manager

import (
	"context"

	v1 "go-dora-api/api/manager/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// ManagerDelete函数用于删除Manager
func (c *ControllerV1) ManagerDelete(ctx context.Context, req *v1.ManagerDeleteReq) (res *v1.ManagerDeleteRes, err error) {
	// 调用service.Manage().Update函数，将req.Ids和IsDelete设置为1，更新数据库
	_, err = service.Manage().Update(ctx, &model.UpdateInput{
		Ids:      req.Ids,
		IsDelete: 1,
	})

	// 如果更新失败，则返回错误
	if err != nil {
		return
	}

	// 创建ManagerDeleteRes结构体，将req.Ids赋值给Ids
	resp := v1.ManagerDeleteRes{
		Ids: req.Ids,
	}

	// 返回ManagerDeleteRes结构体和nil
	return &resp, nil
}
