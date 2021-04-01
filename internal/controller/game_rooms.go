package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/umarkotak/lul-er_BE/internal/models"
	"github.com/umarkotak/lul-er_BE/internal/service"
	"github.com/umarkotak/lul-er_BE/internal/utils"
)

func GetGameRooms(c *gin.Context) {

	results, err := service.GetGameRooms()
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	utils.RenderSuccess(c, results)
}

func CreateGameRoom(c *gin.Context) {
	ctxUsername, _ := c.Get("LUL-USERNAME")
	username := fmt.Sprintf("%v", ctxUsername)

	var create_game_room models.GameRoom
	c.BindJSON(&create_game_room)

	create_game_room.RoomMasterUsername = username
	result, err := service.CreateGameRoom(create_game_room)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	utils.RenderSuccess(c, result)
}

func JoinGameRoom(c *gin.Context) {
	ctxUsername, _ := c.Get("LUL-USERNAME")
	username := fmt.Sprintf("%v", ctxUsername)
	gameRoomID := c.Param("game_room_id")

	var gameRoom models.GameRoom
	gameRoom.ID = gameRoomID

	results, err := service.JoinGameRoom(gameRoom, username)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	utils.RenderSuccess(c, results)
}
