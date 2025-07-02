// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"go-dora-api/internal/model"
	"go-dora-api/internal/model/entity"
)

type (
	ISysRole interface {
		// 添加角色
		Add(ctx context.Context, in *model.SysRoleAddInput) (out *model.SysRoleAddOutput, err error)
		// 获取所有角色
		All(ctx context.Context, in *model.SysRoleAllInput) (out []*entity.SysRole, err error)
		// 删除角色
		Delete(ctx context.Context, in *model.SysRoleDeleteInput) (out *model.SysRoleDeleteOutput, err error)
		Detail(ctx context.Context, in *model.SysRoleDetailInput) (out *model.SysRoleDetailRes, err error)
		// 获取管理员角色ID
		GetManagerRoleIDs(ctx context.Context, managerID int64) (roleIDs []int64, err error)
		GetMenu(ctx context.Context, roleIDs []int64, menuType []int) (menuList []*entity.SysMenu, err error)
		// 根据角色ID获取菜单ID
		GetRoleMenuIDs(ctx context.Context, roleIDs []int64) (menuIds []int64, err error)
		// 根据传入的上下文和参数，获取角色列表
		List(ctx context.Context, in *model.SysRoleListInput) (out []*entity.SysRole, total int, err error)
		// 更新角色信息
		Update(ctx context.Context, in *model.SysRoleUpdateInput) (out *model.SysRoleUpdateOutput, err error)
	}
)

var (
	localSysRole ISysRole
)

func SysRole() ISysRole {
	if localSysRole == nil {
		panic("implement not found for interface ISysRole, forgot register?")
	}
	return localSysRole
}

func RegisterSysRole(i ISysRole) {
	localSysRole = i
}
