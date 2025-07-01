// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package manager

import (
	"context"

	"go-dora-api/api/manager/v1"
)

type IManagerV1 interface {
	ManagerDetail(ctx context.Context, req *v1.ManagerDetailReq) (res *v1.ManagerDetailRes, err error)
	ManagerSelf(ctx context.Context, req *v1.ManagerSelfReq) (res *v1.ManagerSelfRes, err error)
	ManagerRefreshPass(ctx context.Context, req *v1.ManagerRefreshPassReq) (res *v1.ManagerRefreshPassRes, err error)
	ManagerUpdate(ctx context.Context, req *v1.ManagerUpdateReq) (res *v1.ManagerUpdateRes, err error)
	ManagerUpdateSelf(ctx context.Context, req *v1.ManagerUpdateSelfReq) (res *v1.ManagerUpdateSelfRes, err error)
}
