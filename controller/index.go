package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(webStart),
	fx.Invoke(genImg),
)

func webStart(lc fx.Lifecycle) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// 在这里可以设置服务器的停止逻辑
			return nil
		},
	})
	return r
}
