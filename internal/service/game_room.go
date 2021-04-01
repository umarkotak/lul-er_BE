package service

import (
	"context"
	"errors"

	"github.com/umarkotak/lul-er_BE/internal/models"
	"github.com/umarkotak/lul-er_BE/internal/repository"
	"google.golang.org/appengine/log"
)

func GetGameRooms() ([]models.GameRoom, error) {
	gameRooms, err := repository.GetGameRooms()

	return gameRooms, err
}

func CreateGameRoom(gameRoom models.GameRoom) (models.GameRoom, error) {

	gamePlayer := models.GamePlayer{
		Username: gameRoom.RoomMasterUsername,
		Status:   "joined",
	}
	gameRoom.GamePlayers = map[string]models.GamePlayer{}
	gameRoom.GamePlayers[gameRoom.RoomMasterUsername] = gamePlayer

	gameRoom, err := repository.CreateGameRoom(gameRoom)
	return gameRoom, err
}

func JoinGameRoom(gameRoom models.GameRoom, username string) (models.GameRoom, error) {
	gameRoom, err := repository.GetGameRoom(gameRoom.ID)

	if err != nil {
		log.Errorf(context.Background(), "Error GetGameRoom %v", err)
		return gameRoom, err
	}

	if gameRoom.ID == "" {
		return gameRoom, errors.New("game room not found")
	}

	gamePlayer := models.GamePlayer{
		Username: username,
		Status:   "joined",
	}
	gameRoom.GamePlayers[username] = gamePlayer

	repository.UpdateGameRoom(gameRoom)

	return gameRoom, nil
}
