package models

type GameRoom struct {
	ID                 string `json:"id"`
	RoomTitle          string `json:"title"`
	RoomMasterUsername string `json:"room_master_username"`
	Mode               string `json:"mode"`
	CurrentPlayerCount int    `json:"current_player_count"`
	MaxPlayerCount     int    `json:"max_player_count"`
}
