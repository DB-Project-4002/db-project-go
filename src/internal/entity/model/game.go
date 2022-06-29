package model

import "time"

type GameAccount struct {
	AccountID      string    `json:"account_id" db:"account_id"`
	Name           string    `json:"name" db:"name"`
	Level          int       `json:"level" db:"level"`
	Avatar         string    `json:"avatar" db:"avatar"`
	AvatarBorderID int       `json:"avatar_border_id" db:"avatar_border_id"`
	GameCredit     int       `json:"game_credit" db:"game_credit"`
	BlueEssence    int       `json:"blue_essence" db:"blue_essence"`
	OrangeEssence  int       `json:"orange_essence" db:"orange_essence"`
	MythicEssence  int       `json:"mythic_essence" db:"mythic_essence"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
