package controller

import (
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(
		200,
		gin.H{
			"success": true,
			"data": map[string]interface{}{
				"ping": "ok",
			},
			"error": "",
		},
	)
}
