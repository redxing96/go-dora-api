/*
 * @Description: 日志助手
 * @Author: redxing96@163.com
 * @Date: 2025-06-23 16:45:56
 * @LastEditTime: 2025-06-26 20:13:55
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logger/loggerHelper.go
 */
package logger

import (
	"go-dora-api/utility/common"
	"go-dora-api/utility/logger"
)

var (
	SystemLogger logger.ILogger = nil
	RedisLogger  logger.ILogger = nil
)

func init() {
	LoadLogger()
}

func LoadLogger() {
	SystemLogger = InitSystemLogger()
	RedisLogger = InitRedisLogger()
}

func InitSystemLogger() logger.ILogger {
	if SystemLogger != nil {
		return SystemLogger
	}

	cfg := logger.DefaultConfig()
	cfg.Filename = common.GetAbsPath("logs/system.log")
	cfg.Level = "info"         // 日志记录等级
	cfg.ModuleKey = "module"   // 模块key值
	cfg.ModuleValue = "system" // 模块value值
	cfg.Compress = false       // 是否开启压缩
	cfg.IsOutput = true        // 是否标准输出打印
	cfg.IsStack = true         // 是否开启堆栈信息打印
	cfg.StackLevel = "error"   // 堆栈记录等级
	// cfg.StackSkip = 1          // 堆栈忽略层级
	SystemLogger = logger.NewLogger(cfg)
	return SystemLogger
}

func InitRedisLogger() logger.ILogger {
	if RedisLogger != nil {
		return RedisLogger
	}

	cfg := logger.DefaultConfig()
	cfg.Filename = common.GetAbsPath("logs/redis.log")
	cfg.Level = "info"         // 日志记录等级
	cfg.ModuleKey = "module"   // 模块key值
	cfg.ModuleValue = "system" // 模块value值
	cfg.Compress = false       // 是否开启压缩
	cfg.IsOutput = true        // 是否标准输出打印
	cfg.IsStack = true         // 是否开启堆栈信息打印
	cfg.StackLevel = "error"   // 堆栈记录等级
	// cfg.StackSkip = 1          // 堆栈忽略层级
	RedisLogger = logger.NewLogger(cfg)
	return RedisLogger
}

// 日志刷盘
func LoggerFlush() {
	if SystemLogger != nil {
		SystemLogger.Sync()
	}
	if RedisLogger != nil {
		RedisLogger.Sync()
	}
}

// 设置日志记录等级
func LoggerSetLevel(level string) error {
	if SystemLogger != nil {
		err := SystemLogger.SetLevel(level)
		if err != nil {
			return err
		}
	}
	if RedisLogger != nil {
		err := RedisLogger.SetLevel(level)
		if err != nil {
			return err
		}
	}
	return nil
}
