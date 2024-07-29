// Code generated by hertz generator.

package main

import (
	"context"
	"fmt"

	"hertz-examples/hz_demo/hertz-server/infra/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	// Use the configuration entry to generate a hertz server
	// Additional configuration items for use can be found at: https://www.cloudwego.io/zh/docs/hertz/reference/config/
	h := server.Default(
		server.WithHostPorts("127.0.0.1:8080"),
		server.WithKeepAlive(true), // 默认就开启长连接
	)
	rpc.InitClient()

	// This in can also do some business initialization work and the user decides what this part is

	// We can define some global middleware here, such as link tracing, telemetry, cross-domain requests, etc.
	// Currently, hertz provides a series of middleware extensions that you can refer to if you need them: https://github.com/hertz-contrib
	h.Use(func(ctx context.Context, c *app.RequestContext) {
		fmt.Println("pre-handler")
		c.Next(c)
		fmt.Println("post-handler")
	})

	register(h)
	h.Spin()
}
