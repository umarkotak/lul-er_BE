package models

type SerializedGameField struct {
	Index       string       `json:"index"`
	IndexNo     int          `json:"index_no"`
	FieldType   string       `json:"field_type"`
	FieldEffect string       `json:"field_effect"`
	GamePlayers []GamePlayer `json:"game_players"`
}

type SerializedGameBoaard struct {
	GameFields []SerializedGameField `json:"game_fields"`
}

type SerializedGameRoom struct {
	ID                 string               `json:"id"`
	RoomTitle          string               `json:"room_title"`
	RoomMasterUsername string               `json:"room_master_username"`
	Mode               string               `json:"mode"`
	CurrentPlayerCount int                  `json:"current_player_count"`
	MaxPlayerCount     int                  `json:"max_player_count"`
	Status             string               `json:"status"`
	GamePlayers        []GamePlayer         `json:"game_players"`
	GameBoard          SerializedGameBoaard `json:"game_board"`
	MyPlayer           GamePlayer           `json:"my_player"`
}
