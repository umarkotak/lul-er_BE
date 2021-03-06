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

func GetGameRoom(c *gin.Context) {
	ctxUsername, _ := c.Get("LUL-USERNAME")
	username := fmt.Sprintf("%v", ctxUsername)

	gameRoomID := c.Param("game_room_id")
	gameRoom, err := service.GetGameRoom(gameRoomID)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	serializedGameRoom, err := service.SerializeGameRoomDetail(gameRoom, username)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	utils.RenderSuccess(c, serializedGameRoom)
}

func CreateGameRoom(c *gin.Context) {
	ctxUsername, _ := c.Get("LUL-USERNAME")
	username := fmt.Sprintf("%v", ctxUsername)

	var create_game_room models.GameRoom
	c.BindJSON(&create_game_room)

	create_game_room.RoomMasterUsername = username
	gameRoom, err := service.CreateGameRoom(create_game_room)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	serializedGameRoom, err := service.SerializeGameRoomDetail(gameRoom, username)

	utils.RenderSuccess(c, serializedGameRoom)
}

func JoinGameRoom(c *gin.Context) {
	ctxUsername, _ := c.Get("LUL-USERNAME")
	username := fmt.Sprintf("%v", ctxUsername)
	gameRoomID := c.Param("game_room_id")

	var gameRoom models.GameRoom
	gameRoom.ID = gameRoomID

	gameRoom, err := service.JoinGameRoom(gameRoom, username)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	serializedGameRoom, err := service.SerializeGameRoomDetail(gameRoom, username)

	utils.RenderSuccess(c, serializedGameRoom)
}

func StartGameRoom(c *gin.Context) {
	ctxUsername, _ := c.Get("LUL-USERNAME")
	username := fmt.Sprintf("%v", ctxUsername)
	gameRoomID := c.Param("game_room_id")

	var gameRoom models.GameRoom
	gameRoom.ID = gameRoomID

	gameRoom, err := service.StartGameRoom(gameRoom, username)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	serializedGameRoom, err := service.SerializeGameRoomDetail(gameRoom, username)

	utils.RenderSuccess(c, serializedGameRoom)
}

func LeaveGameRoom(c *gin.Context) {
	ctxUsername, _ := c.Get("LUL-USERNAME")
	username := fmt.Sprintf("%v", ctxUsername)
	gameRoomID := c.Param("game_room_id")

	var gameRoom models.GameRoom
	gameRoom.ID = gameRoomID

	result, err := service.LeaveGameRoom(gameRoom, username)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	utils.RenderSuccess(c, result)
}

func GamePlayerGenerateMove(c *gin.Context) {
	ctxUsername, _ := c.Get("LUL-USERNAME")
	username := fmt.Sprintf("%v", ctxUsername)
	gameRoomID := c.Param("game_room_id")

	var gameRoom models.GameRoom
	gameRoom.ID = gameRoomID

	gameRoom, err := service.GenerateMove(gameRoom, username)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	serializedGameRoom, err := service.SerializeGameRoomDetail(gameRoom, username)

	utils.RenderSuccess(c, serializedGameRoom)
}

func GamePlayerExecuteItem(c *gin.Context) {
	utils.RenderSuccess(c, map[string]interface{}{})
}

func GamePlayerExecuteMove(c *gin.Context) {
	ctxUsername, _ := c.Get("LUL-USERNAME")
	username := fmt.Sprintf("%v", ctxUsername)
	gameRoomID := c.Param("game_room_id")

	var gameRoom models.GameRoom
	gameRoom.ID = gameRoomID

	gameRoom, err := service.ExecuteMove(gameRoom, username)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	serializedGameRoom, err := service.SerializeGameRoomDetail(gameRoom, username)

	utils.RenderSuccess(c, serializedGameRoom)
}
