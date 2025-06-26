/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-23 16:26:20
 * @LastEditTime: 2025-06-24 19:46:55
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/main.go
 */
package main

import (
	_ "go-dora-api/internal/packed"

	_ "go-dora-api/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"go-dora-api/internal/cmd"

	_ "go-dora-api/internal/migrations"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
