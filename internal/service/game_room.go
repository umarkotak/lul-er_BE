package service

import (
	"github.com/umarkotak/lul-er_BE/internal/models"
	"github.com/umarkotak/lul-er_BE/internal/repository"
)

func GetGameRooms() ([]models.GameRoom, error) {
	gameRooms, err := repository.GetGameRooms()

	return gameRooms, err
}
