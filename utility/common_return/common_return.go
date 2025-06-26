/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-23 17:07:25
 * @LastEditTime: 2025-06-23 17:16:46
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/utility/common_return/common_return.go
 */
package common_return

import (
	"fmt"
	"go-dora-api/internal/consts"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

// 业务码
type bussCode struct {
	version string
	code    int
	message string
	detail  BussDetail
}

// 业务码详情
type BussDetail struct {
	Detail interface{}
	Args   []interface{}
}

// 创建业务码
func New(code int, detail any, message string, args ...any) gcode.Code {
	return bussCode{
		version: consts.APP_VERSION,
		code:    code,
		message: message,
		detail: BussDetail{
			Detail: detail,
			Args:   args,
		},
	}
}

// 获取业务码
func (c bussCode) Code() int {
	return c.code
}

// 获取业务码消息
func (c bussCode) Message() string {
	//多语言实现
	msg := g.I18n().GetContent(gctx.New(), c.message)
	if msg == "" {
		return fmt.Sprintf(c.message, c.detail.Args...)
	}
	return fmt.Sprintf(msg, c.detail.Args...)
}

// 获取业务码详情
func (c bussCode) Detail() interface{} {
	return c.detail
}

// 获取业务码字符串
func (c bussCode) String() string {
	if c.detail.Detail != nil {
		return fmt.Sprintf(`%d:%s %v`, c.code, c.Message(), c.detail.Detail)
	}
	if c.message != "" {
		return fmt.Sprintf(`%d:%s`, c.code, c.Message())
	}
	return fmt.Sprintf(`%d`, c.code)
}
