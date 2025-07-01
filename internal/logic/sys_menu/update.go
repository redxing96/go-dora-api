/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 21:45:26
 * @LastEditTime: 2025-07-01 12:14:34
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/sys_menu/update.go
 */
package sys_menu

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"
	"go-dora-api/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// 更新路由权限
func (s *sSysMenu) Update(ctx context.Context, in *model.MenuUpdateInput) (out *model.MenuUpdateOutput, err error) {
	// 更新路由权限
	result, err := dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Id, in.Id).Data(do.SysMenu{
		Pid:           in.Pid,           // 父级ID
		Type:          in.Type,          // 类型
		Path:          in.Path,          // 路径
		Sort:          in.Sort,          // 排序
		Component:     in.Component,     // 组件
		Title:         in.Title,         // 标题
		Icon:          in.Icon,          // 图标
		IsHidden:      in.IsHidden,      // 是否隐藏
		IsKeepAlive:   in.IsKeepAlive,   // 是否缓存
		ActiveMenu:    in.ActiveMenu,    // 激活菜单
		IsLargeScreen: in.IsLargeScreen, // 是否大屏幕
		Link:          in.Link,          // 是否链接
		Remark:        in.Remark,        // 备注
		Status:        in.Status,        // 状态
		UpdateTime:    gtime.Now(),      // 更新时间
	}).Update()
	if err != nil {
		// 更新失败，记录日志
		logger.SystemLogger.Errorf("更新路由权限失败: %v", err)
		// 返回错误信息
		err = gerror.NewCode(common_return.ErrorCode("", "update_menu_failed", 500))
		return
	}
	// 获取受影响的行数
	res, err := result.RowsAffected()
	if err != nil || res == 0 {
		// 更新失败，记录日志
		logger.SystemLogger.Errorf("更新路由权限失败: %v", err)
		// 返回错误信息
		err = gerror.NewCode(common_return.ErrorCode("", "update_menu_failed", 500))
		return
	}

	// 返回更新结果
	out = &model.MenuUpdateOutput{
		Id: in.Id,
	}

	return
}
