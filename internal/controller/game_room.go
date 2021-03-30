package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/umarkotak/lul-er_BE/internal/models"
	"github.com/umarkotak/lul-er_BE/internal/service"
	"github.com/umarkotak/lul-er_BE/internal/utils"
)

func GetGameRooms(c *gin.Context) {

	a := c.Request.Header.Get("Authorization")

	var roomData models.GameRoom
	c.BindJSON(&roomData)

	username, err := service.DecodeToken(a)
	if err != nil {
		utils.RenderError(c, 401, err.Error())
		return
	}
	fmt.Println(username)
	result, err := service.JoinGameRoom(roomData)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	utils.RenderSuccess(c, result)
}

func CreateGameRoom(c *gin.Context) {

	var create_game_room models.GameRoom
	c.BindJSON(&create_game_room)

	result, err := service.CreateGameRoom(create_game_room)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	utils.RenderSuccess(c, result)
}
