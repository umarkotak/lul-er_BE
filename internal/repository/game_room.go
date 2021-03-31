package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/umarkotak/lul-er_BE/internal/config"
	"github.com/umarkotak/lul-er_BE/internal/models"
)

func GetGameRooms() ([]models.GameRoom, error) {
	gameRooms := []models.GameRoom{}
	var tempGameRooms map[string]models.GameRoom

	fbGameRoomsRef := config.GetConfig().FbGameRoomsRef
	fbGameRoomsRef.Get(context.Background(), &tempGameRooms)

	for _, tempGameRoom := range tempGameRooms {
		gameRooms = append(gameRooms, tempGameRoom)
	}

	return gameRooms, nil
}

func CreateGameRoom(gameRoom models.GameRoom) (models.GameRoom, error) {

	fbGameRoomsRef := config.GetConfig().FbGameRoomsRef
	gameRoom.ID = fmt.Sprintf("%v", time.Now().Unix())

	fbGameRoomRef := fbGameRoomsRef.Child(gameRoom.ID)
	fbGameRoomRef.Set(context.Background(), gameRoom)

	return gameRoom, nil
}

func JoinGameRoom(gameRoom models.GameRoom, username string) (models.GameRoom, error) {

	var ID models.GameRoom

	fbGameRoomsRef := config.GetConfig().FbGameRoomsRef
	fbGameRoomRef := fbGameRoomsRef.Child(gameRoom.ID).Child("game_players").Child(username)

	gamePlayer := models.GamePlayer{
		Username: username,
		Status:   "joined",
	}

	fbGameRoomRef.Set(context.Background(), gamePlayer)

	fbGameRoomRef2 := fbGameRoomsRef.Child(gameRoom.ID)
	err := fbGameRoomRef2.Get(context.Background(), &ID)
	if err != nil {
		return ID, err
	}

	return ID, nil

}
