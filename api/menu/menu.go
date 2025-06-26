// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package menu

import (
	"context"

	"go-dora-api/api/menu/v1"
)

type IMenuV1 interface {
	MenuAdd(ctx context.Context, req *v1.MenuAddReq) (res *v1.MenuAddRes, err error)
	MenuDelete(ctx context.Context, req *v1.MenuDeleteReq) (res *v1.MenuDeleteRes, err error)
	MenuAllList(ctx context.Context, req *v1.MenuAllListReq) (res *v1.MenuAllListRes, err error)
	GetAuthMenu(ctx context.Context, req *v1.GetAuthMenuReq) (res *v1.GetAuthMenuRes, err error)
	MenuUpdate(ctx context.Context, req *v1.MenuUpdateReq) (res *v1.MenuUpdateRes, err error)
}
