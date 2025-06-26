/*
 * @Description: 获取角色列表
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 21:57:33
 * @LastEditTime: 2025-06-25 22:13:20
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/sys_role/list.go
 */
package sys_role

import (
	"context"
	"go-dora-api/internal/common_return"
	"go-dora-api/internal/dao"
	"go-dora-api/internal/logger"
	"go-dora-api/internal/model"
	"go-dora-api/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

// 根据传入的上下文和参数，获取角色列表
func (s *sSysRole) List(ctx context.Context, in *model.SysRoleListInput) (out []*entity.SysRole, total int, err error) {
	// 创建查询对象
	query := dao.SysRole.Ctx(ctx)
	// 设置分页参数
	query = query.Page(in.Page, in.PageSize)
	// 如果有搜索条件，则添加搜索条件
	if in.Search != "" {
		query = query.WhereLike(dao.SysRole.Columns().RoleName, "%"+in.Search+"%")
	}
	// 设置排序条件
	query = query.Order(dao.SysRole.Columns().Id + " desc")
	// 执行查询，并将结果扫描到out中，同时获取总条数
	err = query.ScanAndCount(&out, &total, true)
	// 如果查询失败，则记录错误日志，并返回错误信息
	if err != nil {
		logger.SystemLogger.Errorf("获取角色列表失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_role_list_failed", 500))
		return
	}
	// 返回查询结果
	return
}
