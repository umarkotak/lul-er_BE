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

func JoinGameRoom(roomID string, user string) (models.GameRoom, error) {

	var ID models.GameRoom
	if roomID == "" {
		return ID, nil
	}

	fbGameRoomsRef := config.GetConfig().FbGameRoomsRef
	fbGameRoomRef2 := fbGameRoomsRef.Child(roomID).Child("game_players")
	fmt.Println(fbGameRoomRef2)

	// a := "test"
	// fbGameRoomRef2.Set(context.Background(), a)

	fbGameRoomRef := fbGameRoomsRef.Child(roomID)

	fmt.Println(user)
	err := fbGameRoomRef.Get(context.Background(), &ID)
	if err != nil {
		return ID, err
	}

	return ID, nil

}
