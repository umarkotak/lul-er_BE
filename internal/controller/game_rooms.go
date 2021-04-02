package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/umarkotak/lul-er_BE/internal/models"
	"github.com/umarkotak/lul-er_BE/internal/repository"
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

func GetGameRoom(c *gin.Context) {
	gameRoomID := c.Param("game_room_id")
	result, err := service.GetGameRoom(gameRoomID)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	utils.RenderSuccess(c, result)
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

	result, err := service.JoinGameRoom(gameRoom, username)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	utils.RenderSuccess(c, result)
}

func LeaveGameRoom(c *gin.Context) {

}

func GamePlayerMove(c *gin.Context) {
	ctxUsername, _ := c.Get("LUL-USERNAME")
	username := fmt.Sprintf("%v", ctxUsername)
	gameRoomID := c.Param("game_room_id")

	gameRoom, _ := repository.GetGameRoom(gameRoomID)
	gamePlayer := gameRoom.GamePlayers[username]

	moveCount := 5

	oldFieldIdx := fmt.Sprintf("idx_%v", gamePlayer.Position)
	gamePlayer.Position = gamePlayer.Position + moveCount

	gameRoom.GamePlayers[username] = gamePlayer

	fieldIdx := fmt.Sprintf("idx_%v", gamePlayer.Position)
	movedGamePlayer := models.GamePlayer{
		Username: username,
		Status:   "active",
	}
	gameRoom.GameBoard.GameFields[fieldIdx].GamePlayers[username] = movedGamePlayer
	gameRoom.GameBoard.GameFields[oldFieldIdx].GamePlayers[username] = models.GamePlayer{}

	repository.UpdateGameRoom(gameRoom)

	utils.RenderSuccess(c, gameRoom)
}
