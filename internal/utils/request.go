package utils

import (
	"github.com/gin-gonic/gin"
)

func RenderSuccess(c *gin.Context, data interface{}) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
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
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	c.JSON(
		statusCode,
		gin.H{
			"success": false,
			"data":    map[string]interface{}{},
			"error":   message,
		},
	)
	c.AbortWithStatus(statusCode)
}
