package manage

import "go-dora-api/internal/service"

type sManage struct{}

func New() *sManage {
	return &sManage{}
}

func init() {
	service.RegisterManage(New())
}
