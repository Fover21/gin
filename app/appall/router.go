package appall

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.POST("/5118/get/text", getTextHandler)
}
