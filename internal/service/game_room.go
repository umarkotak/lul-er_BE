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
	gameRoom, err := repository.CreateGameRoom(gameRoom)
	return gameRoom, err
}

func JoinGameRoom(reqUser models.GameRoom, username string) (models.GameRoom, error) {
	gameRoom, err := repository.JoinGameRoom(reqUser.ID, username)

	if err != nil {
		log.Errorf(context.Background(), "Error GetUserByUsername %v", err)
		return gameRoom, err
	}

	if gameRoom.ID == "" {
		return gameRoom, errors.New("Room not found")
	}

	return gameRoom, nil
}
