// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IPermissionCache interface {
		// LoadUserPermissions 方法用于加载用户权限
		LoadUserPermissions(ctx context.Context, managerID int64) (err error)
		// 检查用户是否有管理员权限
		CheckManagerPermission(ctx context.Context, path string) (hasPermission bool)
		// 从sPermissionCache结构体中移除指定managerID的用户权限
		RemoveManagerPermissions(ctx context.Context, managerID int64)
	}
)

var (
	localPermissionCache IPermissionCache
)

func PermissionCache() IPermissionCache {
	if localPermissionCache == nil {
		panic("implement not found for interface IPermissionCache, forgot register?")
	}
	return localPermissionCache
}

func RegisterPermissionCache(i IPermissionCache) {
	localPermissionCache = i
}
