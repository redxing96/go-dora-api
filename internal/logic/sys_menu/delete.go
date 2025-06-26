/*
 * @Description: 删除路由权限
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 20:30:20
 * @LastEditTime: 2025-06-25 20:37:08
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/sys_menu/delete.go
 */
package sys_menu

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"

	"github.com/gogf/gf/v2/errors/gerror"
)

// 删除路由权限
func (s *sSysMenu) Delete(ctx context.Context, in *model.MenuDeleteInput) (res *model.MenuDeleteOutput, err error) {
	// 根据传入的Ids删除路由权限
	result, err := dao.SysMenu.Ctx(ctx).WhereIn(dao.SysMenu.Columns().Id, in.Ids).Delete()
	if err != nil {
		// 如果删除失败，记录错误日志
		logger.SystemLogger.Errorf("删除路由权限失败: %v", err)
		// 返回错误信息
		err = gerror.NewCode(common_return.ErrorCode("", "delete_menu_error", 500))
		return nil, err
	}

	// 获取受影响的行数
	rowsAffected, err := result.RowsAffected()

	// 如果受影响的行数为0或者出现错误，记录错误日志并返回错误信息
	if err != nil || rowsAffected == 0 {
		logger.SystemLogger.Errorf("删除路由权限失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "delete_menu_error", 500))
		return nil, err
	}

	// 返回删除成功的Ids
	return &model.MenuDeleteOutput{
		Ids: in.Ids,
	}, nil
}
