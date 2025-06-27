// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/go-redis/redis/v8"
)

type (
	IRedis interface {
		// 获取redis客户端
		GetClient() redis.UniversalClient
	}
)

var (
	localRedis IRedis
)

func Redis() IRedis {
	if localRedis == nil {
		panic("implement not found for interface IRedis, forgot register?")
	}
	return localRedis
}

func RegisterRedis(i IRedis) {
	localRedis = i
}
