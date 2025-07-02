// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dictionary

import (
	"context"

	"go-dora-api/api/dictionary/v1"
)

type IDictionaryV1 interface {
	DictionaryAdd(ctx context.Context, req *v1.DictionaryAddReq) (res *v1.DictionaryAddRes, err error)
	DictionaryDelete(ctx context.Context, req *v1.DictionaryDeleteReq) (res *v1.DictionaryDeleteRes, err error)
	DictionaryGetAll(ctx context.Context, req *v1.DictionaryGetAllReq) (res *v1.DictionaryGetAllRes, err error)
	DictionaryGetType(ctx context.Context, req *v1.DictionaryGetTypeReq) (res *v1.DictionaryGetTypeRes, err error)
	DictionaryList(ctx context.Context, req *v1.DictionaryListReq) (res *v1.DictionaryListRes, err error)
	DictionaryUpdate(ctx context.Context, req *v1.DictionaryUpdateReq) (res *v1.DictionaryUpdateRes, err error)
}
