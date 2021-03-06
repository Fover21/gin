package routers

import (
	"gin_one/middleware/cors"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

//gin中间件中使用goroutine
//当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）。

// 初始化
func Init() *gin.Engine {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// gin.Default()默认使用了Logger和Recovery中间件
	r := gin.Default()
	// 不使用任何中间件
	// r:=gin.New()

	// 全局中间件
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())

	// 使用跨域中间件
	r.Use(cors.Cors())

	// 在配置模板之前的时候加载   -  以static开头的去"./statics" 这里找
	r.Static("/static", "./statics")

	// 配置模板
	// Gin框架默认都是使用单模板，如果需要使用block template功能，可以通过"github.com/gin-contrib/multitemplate"
	r.LoadHTMLGlob("templates/**/*")

	for _, opt := range options {
		opt(r)
	}
	return r
}
