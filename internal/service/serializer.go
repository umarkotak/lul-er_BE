package service

import (
	"fmt"
	"sort"

	"github.com/umarkotak/lul-er_BE/internal/models"
)

func SerializeGameRoomDetail(gameRoom models.GameRoom, username string) (models.SerializedGameRoom, error) {
	var myPlayer models.GamePlayer
	serializedGameFields := []models.SerializedGameField{}

	for i := 1; i <= 100; i++ {
		fieldIdx := fmt.Sprintf("idx_%v", i)
		gameField := gameRoom.GameBoard.GameFields[fieldIdx]
		gamePlayers := gameField.GamePlayers
		fieldSerializedGamePlayers := []models.GamePlayer{}
		delete(gamePlayers, "placeholder")
		for _, gamePlayer := range gamePlayers {
			fieldSerializedGamePlayers = append(fieldSerializedGamePlayers, gamePlayer)
			if gamePlayer.Username == username {
				myPlayer = gamePlayer
			}
		}
		sort.Slice(fieldSerializedGamePlayers[:], func(i, j int) bool {
			return fieldSerializedGamePlayers[i].TurnIndex < fieldSerializedGamePlayers[j].TurnIndex
		})
		serializedGameField := models.SerializedGameField{
			Index:       gameField.Index,
			IndexNo:     gameField.IndexNo,
			FieldType:   gameField.FieldType,
			FieldEffect: gameField.FieldEffect,
			GamePlayers: fieldSerializedGamePlayers,
		}
		serializedGameFields = append(serializedGameFields, serializedGameField)
	}

	serializedGamePlayers := []models.GamePlayer{}
	for _, gamePlayer := range gameRoom.GamePlayers {
		serializedGamePlayers = append(serializedGamePlayers, gamePlayer)
	}

	sort.Slice(serializedGamePlayers[:], func(i, j int) bool {
		return serializedGamePlayers[i].TurnIndex < serializedGamePlayers[j].TurnIndex
	})

	serialiedGameBoard := models.SerializedGameBoaard{
		GameFields: serializedGameFields,
	}

	serializedGameRoom := models.SerializedGameRoom{
		ID:                 gameRoom.ID,
		RoomTitle:          gameRoom.RoomTitle,
		RoomMasterUsername: gameRoom.RoomMasterUsername,
		Mode:               gameRoom.Mode,
		CurrentPlayerCount: gameRoom.CurrentPlayerCount,
		MaxPlayerCount:     gameRoom.MaxPlayerCount,
		Status:             gameRoom.Status,
		GamePlayers:        serializedGamePlayers,
		GameBoard:          serialiedGameBoard,
		MyPlayer:           myPlayer,
	}

	return serializedGameRoom, nil
}
