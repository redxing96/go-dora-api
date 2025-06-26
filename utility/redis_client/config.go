/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-26 20:07:13
 * @LastEditTime: 2025-06-26 20:09:47
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/utility/redis_client/config.go
 */
package redis_client

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type Config struct {
	Mode          string        // Redis 模式: "single" 或 "cluster"
	Addrs         []string      // Redis 地址列表：single 模式下仅取第一个地址
	Username      string        // 用户名（可选）
	Password      string        // 密码（可选）
	DB            int           // 数据库 index，仅在 single 模式生效
	DialTimeout   time.Duration // 连接超时时间
	ReadTimeout   time.Duration // 读超时
	WriteTimeout  time.Duration // 写超时
	PingInterval  time.Duration // PING 间隔，用于健康检测
	EnableMetrics bool          // 是否开启 metrics 统计
	EnableLogging bool          // 是否开启日志输出（通过 logger.RedisLogger）
}

// LoadConfig 从 goframe 的配置中读取 Redis 配置
func LoadConfig() *Config {
	cfg := &Config{}
	err := g.Cfg().MustGet(gctx(), "redis").Struct(cfg)
	if err != nil {
		panic("redis config load failed: " + err.Error())
	}
	return cfg
}

// gctx 返回默认上下文
func gctx() context.Context {
	return context.TODO()
}
