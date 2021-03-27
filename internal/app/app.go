package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/umarkotak/lul-er_BE/internal/config"
	"github.com/umarkotak/lul-er_BE/internal/controller"
)

func Start() {
	config.InitConfig()

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/ping", controller.Ping)

	router.POST("/users/register", controller.Register)

	router.Run(":" + getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	return port
}
