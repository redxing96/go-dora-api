/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 12:07:00
 * @LastEditTime: 2025-06-25 21:52:43
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/sys_menu/add.go
 */
package sys_menu

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"

	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"
	"go-dora-api/internal/model/do"
)

// 添加菜单
func (s *sSysMenu) Add(ctx context.Context, in *model.MenuAddInput) (res *model.MenuAddOutput, err error) {
	// 调用dao.SysMenu的Ctx方法，传入上下文参数，并使用Data方法传入菜单数据
	id, err := dao.SysMenu.Ctx(ctx).Data(do.SysMenu{
		Pid:           in.Pid,           // 父级菜单ID
		Type:          in.Type,          // 菜单类型
		Path:          in.Path,          // 菜单路径
		Sort:          in.Sort,          // 菜单排序
		Component:     in.Component,     // 菜单组件
		Title:         in.Title,         // 菜单标题
		Icon:          in.Icon,          // 菜单图标
		Hidden:        in.Hidden,        // 是否隐藏
		KeepAlive:     in.KeepAlive,     // 是否缓存
		ActiveMenu:    in.ActiveMenu,    // 激活菜单
		AlwaysShow:    in.AlwaysShow,    // 是否总是显示
		IsLargeScreen: in.IsLargeScreen, // 是否大屏
		IsFirstLevel:  in.IsFirstLevel,  // 是否一级菜单
		IsSecondLevel: in.IsSecondLevel, // 是否二级菜单
		NoRedirect:    in.NoRedirect,    // 是否重定向
		IsLink:        in.IsLink,        // 是否链接
		Remark:        in.Remark,        // 备注
		Status:        in.Status,        // 状态
		CreateTime:    gtime.Now(),      // 创建时间
	}).InsertAndGetId() // 插入数据并获取ID
	if err != nil {
		// 如果插入数据失败，记录错误日志，并返回错误信息
		logger.SystemLogger.Errorf("添加菜单失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "add_menu_error", 500))
		return nil, err
	}

	// 如果插入数据成功，返回菜单ID
	res = &model.MenuAddOutput{
		Id: id,
	}
	return
}
