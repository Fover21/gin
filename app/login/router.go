package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers(e *gin.Engine) {
	e.POST("/login", loginHandler)
	e.GET("/login", loginHandler)
	e.GET("/auth", JWTAuthMiddleware(), authHandler)

	e.GET("/login/index", func(c *gin.Context) {
		// name-模板最上方define定义的名字 - 没有起名字 默认文件名
		c.HTML(http.StatusOK, "login/index.tmpl", gin.H{
			"title":   "login/index  -  tmpl",
			"keyword": "keyword",
		})
	})

}
