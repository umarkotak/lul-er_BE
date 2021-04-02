package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/umarkotak/lul-er_BE/internal/models"
	"github.com/umarkotak/lul-er_BE/internal/repository"
	"google.golang.org/appengine/log"
)

func GetGameRooms() ([]models.GameRoom, error) {
	gameRooms, err := repository.GetGameRooms()

	return gameRooms, err
}

func GetGameRoom(gameRoomID string) (models.GameRoom, error) {
	gameRoom, err := repository.GetGameRoom(gameRoomID)

	return gameRoom, err
}

func CreateGameRoom(gameRoom models.GameRoom) (models.GameRoom, error) {
	gameMode, err := repository.GetGameMode(gameRoom.Mode)
	if err != nil {
		return gameRoom, err
	}

	if gameMode.ID == "" {
		return gameRoom, errors.New(fmt.Sprintf("game mode %v not found", gameRoom.Mode))
	}

	gamePlayer := models.GamePlayer{
		Username:      gameRoom.RoomMasterUsername,
		Status:        "active",
		Position:      1,
		TurnIndex:     1,
		TurnStatus:    "active",
		TurnSubStatus: "move_phase",
	}
	gameRoom.GamePlayers = map[string]models.GamePlayer{}
	gameRoom.GamePlayers[gameRoom.RoomMasterUsername] = gamePlayer
	gameRoom.CurrentPlayerCount = 1
	gameRoom.Status = "initiated"
	gameRoom.MaxPlayerCount = gameMode.MaxGamePlayers
	tempGameBoard := gameMode.GameBoard
	tempGameBoard.GameFields["idx_1"].GamePlayers[gameRoom.RoomMasterUsername] = gamePlayer
	gameRoom.GameBoard = tempGameBoard

	gameRoom, err = repository.CreateGameRoom(gameRoom)
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

	if gameRoom.GamePlayers[username].Status != "" {
		return gameRoom, errors.New("already joined room")
	}

	if gameRoom.Status != "initiated" {
		return gameRoom, errors.New("game already started")
	}

	currentLen := len(gameRoom.GamePlayers)
	turnIndex := currentLen + 1

	gamePlayer := models.GamePlayer{
		Username:      username,
		Status:        "active",
		Position:      1,
		TurnIndex:     turnIndex,
		TurnStatus:    "waiting",
		TurnSubStatus: "waiting",
	}
	gameRoom.GamePlayers[username] = gamePlayer
	tempGameBoard := gameRoom.GameBoard
	tempGameBoard.GameFields["idx_1"].GamePlayers[username] = gamePlayer
	gameRoom.GameBoard = tempGameBoard
	gameRoom.CurrentPlayerCount += 1

	gameRoom, err = repository.UpdateGameRoom(gameRoom)
	if err != nil {
		return gameRoom, errors.New(fmt.Sprintf("Error UpdateGameRoom: %v", err))
	}

	return gameRoom, nil
}

func LeaveGameRoom(gameRoom models.GameRoom, username string) (models.GameRoom, error) {
	return gameRoom, nil
}

func StartGameRoom(gameRoom models.GameRoom, username string) (models.GameRoom, error) {
	gameRoom, err := repository.GetGameRoom(gameRoom.ID)

	if err != nil {
		log.Errorf(context.Background(), "Error GetGameRoom %v", err)
		return gameRoom, err
	}

	if gameRoom.ID == "" {
		return gameRoom, errors.New("game room not found")
	}

	if gameRoom.RoomMasterUsername != username {
		return gameRoom, errors.New("only room master can start the game")
	}

	gameRoom.Status = "started"

	gameRoom, err = repository.UpdateGameRoom(gameRoom)
	if err != nil {
		return gameRoom, errors.New(fmt.Sprintf("Error UpdateGameRoom: %v", err))
	}

	return gameRoom, nil
}
