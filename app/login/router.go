package login

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.POST("/login", loginHandler)
	e.GET("/login", loginHandler)
	e.GET("/auth", JWTAuthMiddleware(), authHandler)
}
