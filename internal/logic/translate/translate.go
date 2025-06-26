/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-23 16:59:29
 * @LastEditTime: 2025-06-23 17:57:45
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/logic/translate/translate.go
 */
package translate

import (
	"context"
	"fmt"
	"go-dora-api/internal/consts"
	"go-dora-api/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type (
	sTranslate struct{}
)

func New() *sTranslate {
	return &sTranslate{}
}

func init() {
	service.RegisterTranslate(New())
}

// 翻译
func (s *sTranslate) Translate(ctx context.Context, content string, args ...any) string {
	msg := g.I18n().GetContent(ctx, content)
	if msg != "" {
		msg = fmt.Sprintf(msg, args...)
	} else {
		msg = fmt.Sprintf(content, args...)
	}
	return msg
}

// 获取默认语言
func (s *sTranslate) GetDefaultLanguage() string {
	language, _ := g.Cfg().Get(gctx.New(), "server.language", consts.LANG_EN)
	return language.String()
}
