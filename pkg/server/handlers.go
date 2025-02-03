package server

import (
	"github.com/gin-gonic/gin"
)

func PingPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"Status": "true",
	})
}
