package routers

import (
	"gin_one/middleware/cors"
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	r := gin.Default()

	// 全局中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 使用跨域中间件
	r.Use(cors.Cors())

	for _, opt := range options {
		opt(r)
	}
	return r
}
