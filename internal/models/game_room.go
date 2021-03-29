package models

type GameRoom struct {
	ID 					string `json:"id"`
	RoomMasterUsername 	string `json:"room_master_username"`
	Mode 				string `json:"mode"`
	CurrentPlayerCount 	string `json:"current_player_count"`
	MaxPlayerCount 		string `json:"max_player_count"`
}