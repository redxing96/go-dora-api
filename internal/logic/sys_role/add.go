/*
 * @Description: 添加角色
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 22:19:01
 * @LastEditTime: 2025-06-25 22:27:27
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/sys_role/add.go
 */
package sys_role

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

// 添加角色
func (s *sSysRole) Add(ctx context.Context, in *model.SysRoleAddInput) (out *model.SysRoleAddOutput, err error) {
	// 调用dao.SysRole的Ctx方法，传入上下文和SysRole结构体，将in中的数据赋值给SysRole结构体
	id, err := dao.SysRole.Ctx(ctx).Data(do.SysRole{
		RoleName:   in.RoleName,
		RoleDesc:   in.RoleDesc,
		Status:     in.Status,
		CreateTime: gtime.Now(),
	}).InsertAndGetId()
	// 如果插入失败，记录错误日志，并返回错误信息
	if err != nil {
		logger.SystemLogger.Errorf("添加角色失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "add_role_failed", 500))
		return
	}
	// 如果插入成功，将id赋值给out的Id字段
	out = &model.SysRoleAddOutput{
		Id: int(id),
	}
	return
}
