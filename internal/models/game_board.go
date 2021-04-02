package models

type GameBoard struct {
	GameFields map[string]GameField `json:"game_fields"`
}
