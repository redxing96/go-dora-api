/*
 * @Description: redis 逻辑
 * @Author: redxing96@163.com
 * @Date: 2025-06-26 21:03:52
 * @LastEditTime: 2025-06-26 21:09:56
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/redis/redis.go
 */
package redis

import (
	"go-dora-api/internal/service"
	"go-dora-api/utility/redis_client"

	"github.com/go-redis/redis/v8"
)

type sRedis struct {
	client *redis_client.RedisClient
}

// 创建一个新的sRedis实例
func New() *sRedis {
	// 加载redis客户端配置
	cfg := redis_client.LoadConfig()
	// 创建redis客户端
	redis := redis_client.NewRedisClient(cfg)
	// 返回sRedis实例
	return &sRedis{
		client: redis,
	}
}

func init() {
	service.RegisterRedis(New())
}

// 获取redis客户端
func (s *sRedis) GetClient() redis.UniversalClient {
	// 返回redis客户端
	return s.client.GetClient()
}
