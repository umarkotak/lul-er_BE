package service

import (
	"fmt"

	"github.com/umarkotak/lul-er_BE/internal/models"
)

func CreateGameRoom(reqGameRoom models.GameRoom) (models.GameRoom, error) {
	fmt.Println("hello world", reqGameRoom)
	return reqGameRoom, nil
}
