package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func goodsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "goodsHandler!",
	})
}

func checkoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "checkoutHandler!",
	})
}
