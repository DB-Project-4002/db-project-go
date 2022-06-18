package model

type GameAccount struct {
	UserID         string `json:"user_id"`
	Name           string `json:"name"`
	Level          int    `json:"level"`
	Avatar         string `json:"avatar"`
	AvatarBorderID int    `json:"avatar_border_id"`
	GameCredit     int    `json:"game_credit"`
	BlueEssence    int    `json:"blue_essence"`
	OrangeEssence  int    `json:"orange_essence"`
	MythicEssence  int    `json:"mythic_essence"`
}
