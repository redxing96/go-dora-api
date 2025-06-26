/*
 * @Description: 日志字段
 * @Author: redxing96@163.com
 * @Date: 2025-06-23 16:42:18
 * @LastEditTime: 2025-06-23 16:42:31
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/utility/logger/field.go
 */
package logger

// 字段结构体
type LogField struct {
	Key   string
	Value interface{}
}

// 获取key值
func (l LogField) GetKey() string {
	return l.Key
}

// 获取value值
func (l LogField) GetValue() interface{} {
	return l.Value
}

// 统一封装zap.Any结构体
func Field(key string, value interface{}) IField {
	return LogField{
		Key:   key,
		Value: value,
	}
}
