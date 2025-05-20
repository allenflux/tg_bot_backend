package cmd

import (
	"context"
	"tg_bot_backend/internal/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"tg_bot_backend/internal/controller/hello"
	"tg_bot_backend/internal/controller/managment"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, middleware.JWTAuth)
				group.Bind(
					hello.NewV1(),
					managment.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
