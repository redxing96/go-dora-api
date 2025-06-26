/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-26 19:46:28
 * @LastEditTime: 2025-06-26 20:14:10
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/utility/redis_client/redis_client.go
 */
package redis_client

import (
	"context"
	"go-dora-api/internal/logger"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	config        *Config
	singleClient  *redis.Client
	clusterClient *redis.ClusterClient
	isConnected   int32 // 使用原子变量管理连接状态
	cancelFunc    context.CancelFunc
}

var redisMetricsEnabled bool

// NewRedisClient 根据配置初始化 Redis 单机或集群客户端
func NewRedisClient(cfg *Config) *RedisClient {
	client := &RedisClient{config: cfg}
	redisMetricsEnabled = cfg.EnableMetrics

	if cfg.Mode == "cluster" {
		client.clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        cfg.Addrs,
			Username:     cfg.Username,
			Password:     cfg.Password,
			DialTimeout:  cfg.DialTimeout,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
		})
	} else {
		client.singleClient = redis.NewClient(&redis.Options{
			Addr:         cfg.Addrs[0],
			Username:     cfg.Username,
			Password:     cfg.Password,
			DB:           cfg.DB,
			DialTimeout:  cfg.DialTimeout,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
		})
	}

	ctx, cancel := context.WithCancel(context.Background())
	client.cancelFunc = cancel
	go client.pingLoop(ctx)
	return client
}

// pingLoop 是后台任务，定期 PING Redis，维护连接状态并输出日志/metrics
func (c *RedisClient) pingLoop(ctx context.Context) {
	ticker := time.NewTicker(c.config.PingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			var err error
			if c.config.Mode == "cluster" {
				err = c.clusterClient.Ping(ctx).Err()
			} else {
				err = c.singleClient.Ping(ctx).Err()
			}

			if err != nil {
				if atomic.LoadInt32(&c.isConnected) == 1 {
					if c.config.EnableLogging {
						logger.RedisLogger.Errorf("Redis disconnected: %v", err)
					}
					atomic.StoreInt32(&c.isConnected, 0)
				}
				// 可在此增加 metrics fail count
			} else {
				if atomic.LoadInt32(&c.isConnected) == 0 {
					if c.config.EnableLogging {
						logger.RedisLogger.Infof("Redis connected successfully")
					}
					atomic.StoreInt32(&c.isConnected, 1)
				}
				// 可在此增加 metrics success count
			}
		}
	}
}

// GetClient 返回可用的 Redis 客户端（自动判断集群或单机）
func (c *RedisClient) GetClient() redis.UniversalClient {
	if c.config.Mode == "cluster" {
		return c.clusterClient
	}
	return c.singleClient
}

// Close 优雅关闭 Redis 客户端及健康检测 goroutine
func (c *RedisClient) Close() error {
	if c.cancelFunc != nil {
		c.cancelFunc()
	}
	if c.config.Mode == "cluster" {
		return c.clusterClient.Close()
	}
	return c.singleClient.Close()
}
