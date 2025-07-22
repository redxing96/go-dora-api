/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-23 16:26:20
 * @LastEditTime: 2025-07-02 10:05:44
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/cmd/cmd.go
 */
package cmd

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"

	"go-dora-api/internal/consts"
	"go-dora-api/internal/controller/auth"
	"go-dora-api/internal/controller/dictionary"
	"go-dora-api/internal/controller/file"
	"go-dora-api/internal/controller/manager"
	"go-dora-api/internal/controller/menu"
	"go-dora-api/internal/controller/sys_role"
	"go-dora-api/internal/service"
	"go-dora-api/utility/common"
)

const (
	swaggerUIPageContent = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <meta name="description" content="SwaggerUI" />
  <title>SwaggerUI</title>
  <link rel="stylesheet" href="../resource/public/swagger-ui/css/swagger-ui.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="../resource/public/swagger-ui/js/swagger-ui-bundle.js" crossorigin></script>
<script>
  window.onload = () => {
    window.ui = SwaggerUIBundle({
      url: '/api.json',
      dom_id: '#swagger-ui',
    });
  };
</script>
</body>
</html>
`
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 设置语言
			languageSetting()

			// 初始化定时任务
			service.Crontab().InitData(ctx)

			s := g.Server()
			s.SetServerRoot(".") // 设置静态文件目录
			s.AddStaticPath("/files", common.GetAbsPath(consts.UPLOADS_PATH))
			s.SetClientMaxBodySize(consts.Max_UPLOAD_SIZE) // 设置最大上传大小
			openApiSetting(s.GetOpenApi())                 // 设置OpenAPI
			// 设置开发模式
			develop, _ := g.Cfg().Get(ctx, "server.develop", "false")
			if develop.Bool() {
				s.EnablePProf()
			}
			s.Group("/", func(group *ghttp.RouterGroup) {
				// 设置语言
				group.Middleware(func(r *ghttp.Request) {
					lang := consts.LANG_EN
					if l := r.Header.Get("Lang"); l != "" {
						lang = l
					}

					if strings.ToUpper(lang) == "ZH-CN" {
						lang = consts.LANG_CN
					}

					// 设置语言
					g.I18n().SetLanguage(lang)
					r.SetCtxVar("lang", lang)
					r.Middleware.Next()
				})

				// 添加跨域中间件
				group.Middleware(service.Middleware().MiddlewareCORS)
				// 添加全局响应中间件
				group.Middleware(service.Middleware().MiddlewareHandlerResponse)
				// 添加访问日志中间件
				group.Middleware(service.Middleware().MiddlewareAccessLog)
				// 添加全局捕获异常中间件
				group.Middleware(service.Middleware().MiddlewareErrorHandler)
				// 添加限流中间件
				group.Middleware(service.Middleware().RateLimitMiddleware)
				// 添加数据库日志中间件
				group.Middleware(service.Middleware().MiddlewareDbLog)
				// 添加SwaggerUI路由
				if develop.Bool() {
					group.GET("/swagger", func(r *ghttp.Request) {
						r.Response.Write(swaggerUIPageContent)
					})
				}

				group.GET("/ws", func(r *ghttp.Request) {
					service.WebSocket().HandleWsConnection(r.Response.Writer, r.Request)
				})

				// 客户端接口
				apiGroup := group.Group("/api")

				// 绑定控制器-不验证token
				apiGroup.Bind()

				// 绑定控制器-验证token
				apiGroup.Middleware(service.Middleware().MiddlewareClientJWT)
				apiGroup.Bind()

				// 管理端接口
				manageGroup := group.Group("/manage")

				// 绑定控制器-不验证token
				manageGroup.Bind(
					auth.NewV1().Login,
				)

				// 绑定控制器-验证token
				manageGroup.Middleware(service.Middleware().MiddlewareManageJWT)
				manageGroup.Bind()

				// 绑定控制器-验证权限
				manageGroup.Middleware(service.Middleware().MiddlewarePermission)
				manageGroup.Bind(
					menu.NewV1(),
					sys_role.NewV1(),
					auth.NewV1().Logout,
					manager.NewV1(),
					dictionary.NewV1(),
					file.NewV1(),
				)

			})
			s.Run()
			return nil
		},
	}
)

// 设置OpenAPI
func openApiSetting(api *goai.OpenApiV3) {
	api.Info.Version = consts.APP_VERSION
	api.Info.Title = "go-dora-api"
	api.Info.Description = "This is a set of go-dora-api programming interfaces "
	api.Components.SecuritySchemes = map[string]goai.SecuritySchemeRef{
		"api_key": {
			Value: &goai.SecurityScheme{
				Type:        "apiKey",
				Description: "登录token",
				Name:        "Authorization",
				In:          "header",
			},
		},
	}
}

// 设置语言
func languageSetting() {
	g.I18n().SetLanguage(service.Translate().GetDefaultLanguage())
}
