package cmd

import (
	"context"
	"seckill/internal/controller/seckill"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"seckill/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
				)
				group.Group("/api", func(group *ghttp.RouterGroup) {
					group.Middleware(ghttp.MiddlewareHandlerResponse) // 统一返回格式中间件
					group.Bind(
						seckill.NewV1(), // 注册秒杀控制器
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
