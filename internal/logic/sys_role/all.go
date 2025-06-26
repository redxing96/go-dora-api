/*
 * @Description: 获取所有角色
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 22:28:25
 * @LastEditTime: 2025-06-25 22:30:33
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/sys_role/all.go
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

// 获取所有角色
func (s *sSysRole) All(ctx context.Context, in *model.SysRoleAllInput) (out []*entity.SysRole, err error) {
	// 创建查询对象
	query := dao.SysRole.Ctx(ctx)
	// 如果传入的角色状态不为0，则添加查询条件
	if in.Status != 0 {
		query = query.Where(dao.SysRole.Columns().Status, in.Status)
	}
	// 执行查询，并将结果扫描到out中
	err = query.Scan(&out)
	// 如果查询失败，则记录错误日志，并返回错误信息
	if err != nil {
		logger.SystemLogger.Errorf("获取所有角色失败: %v", err)
		err = gerror.NewCode(common_return.ErrorCode("", "get_role_all_failed", 500))
		return
	}
	// 返回查询结果
	return
}
