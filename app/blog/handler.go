package blog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func postHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "postHandler!",
	})
}

func commentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "commentHandler!",
	})
}
