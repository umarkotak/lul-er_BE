package models

type GameField struct {
	Index       string                `json:"index"`
	IndexNo     int                   `json:"index_no"`
	FieldType   string                `json:"field_type"`
	FieldEffect string                `json:"field_effect"`
	GamePlayers map[string]GamePlayer `json:"game_players"`
}
