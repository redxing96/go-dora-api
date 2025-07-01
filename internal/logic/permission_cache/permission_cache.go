/*
 * @Description: 权限缓存服务
 * @Author: redxing96@163.com
 * @Date: 2025-06-26 13:03:18
 * @LastEditTime: 2025-07-01 12:04:37
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/permission_cache/permission_cache.go
 */
package permissioncache

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/service"
	"slices"
	"sync"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/patrickmn/go-cache"
	"github.com/spf13/cast"
)

// 白名单
var whiteList = []string{
	"/manage/v1/auth/login",
	"/manage/v1/auth/logout",
	"/manage/v1/menu/get_auth_menu",
	"/manage/v1/manager/self",
}

// sPermissionCache 权限缓存服务实现
type sPermissionCache struct {
	cache        *cache.Cache
	permissionMu sync.RWMutex
}

func New() *sPermissionCache {
	return &sPermissionCache{
		cache: cache.New(24*time.Hour, 48*time.Hour),
	}
}

func init() {
	service.RegisterPermissionCache(New())
}

// LoadUserPermissions 方法用于加载用户权限
func (s *sPermissionCache) LoadUserPermissions(ctx context.Context, managerID int64) (err error) {
	// 调用 service.SysMenu().GetManageMenu 方法获取管理员菜单
	menuList, err := service.SysMenu().GetManageMenu(ctx, managerID, []int{})
	// 如果获取菜单失败，则记录错误日志，并返回错误
	if err != nil {
		logger.SystemLogger.Errorf("获取管理员菜单失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_manager_menu_failed", 500))
		return
	}

	// 创建一个空的权限集合
	permissions := make(map[string]any)
	// 遍历菜单列表，将菜单的路径添加到权限集合中
	for _, menu := range menuList {
		permissions[menu.Path] = struct{}{}
	}

	// 将权限集合存储到缓存中
	s.cache.Set(cast.ToString(managerID), permissions, cache.DefaultExpiration)
	return
}

// 检查用户是否有管理员权限
func (s *sPermissionCache) CheckManagerPermission(ctx context.Context, path string) (hasPermission bool) {

	// 如果路径在白名单中，则直接返回true
	if slices.Contains(whiteList, path) {
		return true
	}

	// 从上下文中获取用户是否为超级管理员
	isSuper := cast.ToInt(ctx.Value("is_super"))
	// 如果用户为超级管理员，则直接返回true
	if isSuper == 1 {
		return true
	}

	// 从上下文中获取用户ID
	managerID := cast.ToInt64(ctx.Value("manager_id"))

	// 读取权限缓存
	s.permissionMu.RLock()
	defer s.permissionMu.RUnlock()

	// 从缓存中获取用户权限
	permissions, ok := s.cache.Get(cast.ToString(managerID))
	if !ok {
		// 如果缓存中不存在用户权限，则记录错误日志
		logger.SystemLogger.Errorf("用户权限缓存不存在: %v", managerID)
		return
	}

	// 将权限转换为map类型
	data := cast.ToStringMap(permissions)
	// 判断用户权限是否存在
	_, exists := data[path]
	return exists
}

// 从sPermissionCache结构体中移除指定managerID的用户权限
func (s *sPermissionCache) RemoveManagerPermissions(ctx context.Context, managerID int64) {
	// 加锁，防止并发操作
	s.permissionMu.Lock()
	// 在函数结束时解锁
	defer s.permissionMu.Unlock()

	// 从缓存中删除指定managerID的用户权限
	s.cache.Delete(cast.ToString(managerID))
}
