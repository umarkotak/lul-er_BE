package models

type GameRoom struct {
	GamePlayers        map[string]GamePlayer
	MaxPlayers         int
	Mode               string
	RoomMasterUsername string
}
