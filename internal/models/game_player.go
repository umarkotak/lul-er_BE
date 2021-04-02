package models

type GamePlayer struct {
	Username string `json:"username"`
	Status   string `json:"status"`
	Position int    `json:"position"`
}
