package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/umarkotak/lul-er_BE/internal/config"
	"github.com/umarkotak/lul-er_BE/internal/controller"
	"github.com/umarkotak/lul-er_BE/internal/utils"
)

func Start() {
	config.InitConfig()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(CORSMiddleware())

	router.GET("/ping", controller.Ping)
	router.POST("/users/register", controller.Register)
	router.POST("/users/login", controller.Login)

	router.Run(":" + getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	return port
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			utils.RenderSuccess(c, nil)
			return
		}

		c.Next()
	}
}
