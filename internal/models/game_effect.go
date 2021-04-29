package models

type GameEffect struct {
	Id           string `json:"id"`
	EffectType   string `json:"effect_type"`
	EffectValue  int    `json:"effect_value"`
	EffectValueS string `json:"effect_value_s"`
}
