package repository

import (
	"context"
	"errors"
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

func GetGameRoom(gameRoomID string) (models.GameRoom, error) {
	var gameRoom models.GameRoom

	if gameRoomID == "" {
		return gameRoom, errors.New("game room id can't be blank")
	}

	fbGameRoomsRef := config.GetConfig().FbGameRoomsRef
	fbGameRoomRef := fbGameRoomsRef.Child(gameRoomID)

	err := fbGameRoomRef.Get(context.Background(), &gameRoom)
	if err != nil {
		return gameRoom, err
	}

	return gameRoom, nil
}

func CreateGameRoom(gameRoom models.GameRoom) (models.GameRoom, error) {
	fbGameRoomsRef := config.GetConfig().FbGameRoomsRef
	gameRoom.ID = fmt.Sprintf("GAMEROOM-%v", time.Now().Unix())

	fbGameRoomRef := fbGameRoomsRef.Child(gameRoom.ID)
	fbGameRoomRef.Set(context.Background(), gameRoom)

	return gameRoom, nil
}

func UpdateGameRoom(gameRoom models.GameRoom) (models.GameRoom, error) {
	if gameRoom.ID == "" {
		return gameRoom, errors.New("game room id can't be blank")
	}

	fbGameRoomsRef := config.GetConfig().FbGameRoomsRef
	fbGameRoomRef := fbGameRoomsRef.Child(gameRoom.ID)

	err := fbGameRoomRef.Set(context.Background(), gameRoom)
	if err != nil {
		return gameRoom, err
	}

	return gameRoom, nil
}
