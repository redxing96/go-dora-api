// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_role

import (
	"context"

	"go-dora-api/api/sys_role/v1"
)

type ISysRoleV1 interface {
	RoleAdd(ctx context.Context, req *v1.RoleAddReq) (res *v1.RoleAddRes, err error)
	RoleAll(ctx context.Context, req *v1.RoleAllReq) (res *v1.RoleAllRes, err error)
	RoleDelete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error)
	RoleDetail(ctx context.Context, req *v1.RoleDetailReq) (res *v1.RoleDetailRes, err error)
	RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error)
	RoleUpdate(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.RoleUpdateRes, err error)
}
