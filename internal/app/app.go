package app

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/umarkotak/lul-er_BE/internal/config"
	"github.com/umarkotak/lul-er_BE/internal/controller"
	"github.com/umarkotak/lul-er_BE/internal/utils"
)

func Start() {
	config.InitConfig()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())

	// check server
	router.GET("/ping", controller.Ping)

	// user api

	userRouter := router.Group("/users")
	userRouter.POST("/login", controller.Login)
	userRouter.POST("/register", controller.Register)

	// game service

	gameRouter := router.Group("/game_rooms")
	gameRouter.Use(AuthMiddleware())
	gameRouter.GET("/", controller.GetGameRooms)
	gameRouter.GET("/:game_room_id", controller.GetGameRoom)
	gameRouter.POST("/", controller.CreateGameRoom)
	gameRouter.POST("/:game_room_id/join", controller.JoinGameRoom)
	gameRouter.POST("/:game_room_id/move", controller.GamePlayerMove)

	router.Run(":" + getPort())
	// router.Run(":" + "3000")
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

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")

		if bearerToken == "" {
			utils.RenderError(c, 401, "Missing auth token")
			return
		}

		splitBearerToken := strings.Split(bearerToken, " ")

		if len(splitBearerToken) != 2 {
			utils.RenderError(c, 401, "Invalid auth token format")
			return
		}

		jwtToken := splitBearerToken[1]

		result, err := utils.DecodeToken(jwtToken)
		if err != nil {
			utils.RenderError(c, 401, "Invalid auth token")
			return
		}

		c.Set("LUL-USERNAME", result.Issuer)
		c.Next()
	}
}
