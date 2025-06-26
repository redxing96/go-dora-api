/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 20:28:19
 * @LastEditTime: 2025-06-25 20:39:17
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/controller/menu/menu_v1_menu_delete.go
 */
package menu

import (
	"context"

	v1 "go-dora-api/api/menu/v1"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"
)

// MenuDelete 删除菜单
func (c *ControllerV1) MenuDelete(ctx context.Context, req *v1.MenuDeleteReq) (res *v1.MenuDeleteRes, err error) {
	// 调用service.SysMenu().Delete方法删除菜单
	result, err := service.SysMenu().Delete(ctx, &model.MenuDeleteInput{
		Ids: req.Ids,
	})
	if err != nil {
		// 如果删除失败，则返回错误
		return
	}

	// 构造返回结果
	resp := v1.MenuDeleteRes{
		Ids: result.Ids,
	}

	// 返回结果和错误
	return &resp, nil
}
