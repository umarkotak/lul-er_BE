package models

type GameMode struct {
	ID             string    `json:"id"`
	MaxGamePlayers int       `json:"max_game_players"`
	GameBoard      GameBoard `json:"game_board"`
}
