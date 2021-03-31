package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/umarkotak/lul-er_BE/internal/models"
	"github.com/umarkotak/lul-er_BE/internal/service"
	"github.com/umarkotak/lul-er_BE/internal/utils"
)

func GetGameRooms(c *gin.Context) {

	a := c.Request.Header.Get("Authorization")
	b := c.Param("game-room-id")

	var roomData models.GameRoom
	// c.BindJSON(&roomData)
	roomData.ID = b

	username, err := service.DecodeToken(a)
	if err != nil {
		utils.RenderError(c, 401, err.Error())
		return
	}

	// fmt.Println(username)
	// fmt.Println(b)
	result, err := service.JoinGameRoom(roomData, username)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	utils.RenderSuccess(c, result)
}

func CreateGameRoom(c *gin.Context) {

	// a := c.Request.Header.Get("Authorization")

	var create_game_room models.GameRoom
	c.BindJSON(&create_game_room)

	username := "umarkotak"
	// if err != nil {
	// 	utils.RenderError(c, 401, err.Error())
	// 	return
	// }

	create_game_room.RoomMasterUsername = username
	result, err := service.CreateGameRoom(create_game_room)
	if err != nil {
		utils.RenderError(c, 400, err.Error())
		return
	}

	utils.RenderSuccess(c, result)
}
