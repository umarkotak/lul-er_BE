package repository

import (
	"context"

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
