package models

type GamePlayer struct {
	Username      string `json:"username"`
	Status        string `json:"status"` // [active, leave]
	Position      int    `json:"position"`
	TurnIndex     int    `json:"turn_index"`
	TurnStatus    string `json:"turn_status"`     // [active, waiting]
	TurnSubStatus string `json:"turn_sub_status"` // [waiting, move_phase, item_phase, execution_phase]
	MoveSize      int    `json:"move_size"`
}
