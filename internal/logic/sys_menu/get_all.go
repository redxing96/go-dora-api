/*
 * @Description: 获取所有菜单
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 10:25:04
 * @LastEditTime: 2025-07-01 17:42:43
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/sys_menu/get_all.go
 */
package sys_menu

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"
	"go-dora-api/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

// 获取所有系统菜单
func (s *sSysMenu) GetAll(ctx context.Context, in *model.GetAllSysMenuInput) (res []*entity.SysMenu, total int, err error) {
	query := dao.SysMenu.Ctx(ctx)

	if in.Search != "" {
		query = query.WhereLike(dao.SysMenu.Columns().Title, "%"+in.Search+"%")
	}

	if in.Page > 0 && in.PageSize > 0 {
		query = query.Page(in.Page, in.PageSize)
	}

	if len(in.Type) > 0 {
		query = query.WhereIn(dao.SysMenu.Columns().Type, in.Type)
	}

	if in.Status > 0 {
		query = query.WhereIn(dao.SysMenu.Columns().Status, in.Status)
	}

	query.Order(dao.SysMenu.Columns().Sort + " desc")

	// 从数据库中查询所有系统菜单
	err = query.ScanAndCount(&res, &total, true)
	// 如果查询失败，记录错误日志，并返回错误信息
	if err != nil {
		logger.SystemLogger.Errorf("get_all_sys_menu_error: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_all_sys_menu_error", 500))
		return
	}

	return
}
