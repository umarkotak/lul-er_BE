package models

// type PlayersData struct {
// 	Users GamePlayer
// 	// Users map[GamePlayer]int
// }

type GamePlayers []GamePlayer

type GameRoom struct {
	ID                 string      `json:"id"`
	RoomTitle          string      `json:"title"`
	RoomMasterUsername string      `json:"room_master_username"`
	Mode               string      `json:"mode"`
	GamePlayers        GamePlayers `json:"game_players"`
	// GamePlayers        GamePlayer `json:"game_players"`
	CurrentPlayerCount int `json:"current_player_count"`
	MaxPlayerCount     int `json:"max_player_count"`
}
