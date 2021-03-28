package utils

import (
	"github.com/gin-gonic/gin"
)

func RenderSuccess(c *gin.Context, data interface{}) {
	c.JSON(
		200,
		gin.H{
			"success": true,
			"data":    data,
			"error":   "",
		},
	)
}

func RenderError(c *gin.Context, statusCode int, message string) {
	c.JSON(
		statusCode,
		gin.H{
			"success": false,
			"data":    map[string]interface{}{},
			"error":   message,
		},
	)
}
