/*
 * @Description: 日志配置
 * @Author: redxing96@163.com
 * @Date: 2025-06-23 16:41:46
 * @LastEditTime: 2025-06-23 16:42:04
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/utility/logger/config.go
 */
package logger

// 获取默认日志配置
func DefaultConfig() Config {
	return Config{
		Level:      "info",
		Filename:   "logs/log.log",
		MaxSize:    20,
		MaxAge:     180,
		MaxBackups: 50,
		Compress:   true,
	}
}

// 日志配置
type Config struct {
	// 日志等级(error/warn/info/debug, 默认为info)
	Level string
	// 日志文件路径
	Filename string
	// 日志切片大小, 单位:mb, 默认100mb
	MaxSize int
	// 日志最大保存天数, 默认180天
	MaxAge int
	// 日志文件最大保存个数, 默认100个
	MaxBackups int
	// 是否开启压缩,默认开启
	Compress bool
	// 是否标准输出打印(默认为false, 如果为true，日志会在os.Stdout打印)
	IsOutput bool
	// 是否开启堆栈信息打印
	IsStack bool
	// 堆栈记录等级, IsStack为true时生效(error/warn/info/debug, 默认为error)
	StackLevel string
	// module key值, 为空则不记录
	ModuleKey string
	// module value值, 为空则不记录
	ModuleValue string
	// 堆栈忽略层级
	StackSkip int
}
