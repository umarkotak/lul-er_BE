package models

type User struct {
	Username         string `json:"username"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	AuthToken        string `json:"auth_token"`
	ActiveGameRoomID string `json:"active_game_room_id"`
}
