package service

import (
	"errors"
	"fmt"

	"github.com/umarkotak/lul-er_BE/internal/models"
	"github.com/umarkotak/lul-er_BE/internal/repository"
	"github.com/umarkotak/lul-er_BE/internal/utils"
)

func GenerateMove(gameRoom models.GameRoom, username string) (models.GameRoom, error) {
	gameRoom, err := repository.GetGameRoom(gameRoom.ID)

	if err != nil {
		return gameRoom, errors.New(fmt.Sprintf("Error GetGameRoom: %v", err))
	}

	if gameRoom.ID == "" {
		return gameRoom, errors.New("game room not found")
	}

	moveCount := utils.GenerateRandomNumber(1, 6)
	gamePlayer := gameRoom.GamePlayers[username]
	gamePlayer.MoveSize = moveCount
	gamePlayer.TurnSubStatus = "item_phase"
	gameRoom.GamePlayers[username] = gamePlayer

	repository.UpdateGameRoom(gameRoom)

	return gameRoom, nil
}

func ExecuteItem(gameRoom models.GameRoom, username string) (models.GameRoom, error) {
	gameRoom, err := repository.GetGameRoom(gameRoom.ID)

	if err != nil {
		return gameRoom, errors.New(fmt.Sprintf("Error GetGameRoom: %v", err))
	}

	if gameRoom.ID == "" {
		return gameRoom, errors.New("game room not found")
	}

	return gameRoom, nil
}

func ExecuteMove(gameRoom models.GameRoom, username string) (models.GameRoom, error) {
	gameRoom, err := repository.GetGameRoom(gameRoom.ID)

	if err != nil {
		return gameRoom, errors.New(fmt.Sprintf("Error GetGameRoom: %v", err))
	}

	if gameRoom.ID == "" {
		return gameRoom, errors.New("game room not found")
	}

	gamePlayer := gameRoom.GamePlayers[username]

	oldFieldIdx := fmt.Sprintf("idx_%v", gamePlayer.Position)
	gamePlayer.Position = gamePlayer.Position + gamePlayer.MoveSize
	gamePlayer.MoveSize = 0
	gamePlayer.TurnStatus = "waiting"
	gamePlayer.TurnSubStatus = "waiting"
	turnIndex := gamePlayer.TurnIndex

	gameRoom.GamePlayers[username] = gamePlayer

	fieldIdx := fmt.Sprintf("idx_%v", gamePlayer.Position)
	gameRoom.GameBoard.GameFields[fieldIdx].GamePlayers[username] = gamePlayer
	delete(gameRoom.GameBoard.GameFields[oldFieldIdx].GamePlayers, username)

	// changing turn to next player
	var nextIndex int
	if turnIndex == len(gameRoom.GamePlayers) {
		nextIndex = 1
	} else {
		nextIndex = turnIndex + 1
	}

	var nextGamePlayer models.GamePlayer

	for _, tempGamePlayer := range gameRoom.GamePlayers {
		if tempGamePlayer.TurnIndex == nextIndex {
			nextGamePlayer = tempGamePlayer
			break
		}
	}

	if nextGamePlayer.Username == "" {
		return gameRoom, errors.New("next game player not found")
	}
	nextGamePlayer.TurnStatus = "active"
	nextGamePlayer.TurnSubStatus = "move_phase"
	gameRoom.GamePlayers[nextGamePlayer.Username] = nextGamePlayer

	repository.UpdateGameRoom(gameRoom)

	return gameRoom, nil
}
