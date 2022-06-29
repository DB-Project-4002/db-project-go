package model

import "time"

type Champion struct {
	Name             string    `json:"name" db:"name"`
	BlueEssencePrice int       `json:"blue_essence_price" db:"blue_essence_price"`
	GameCreditPrice  int       `json:"game_credit_price" db:"game_credit_price"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

type ChampionSkins struct {
	Name               string    `json:"name" db:"name"`
	ChampionName       string    `json:"champion_name" db:"champion_name"`
	OrangeEssencePrice int       `json:"orange_essence_price" db:"orange_essence_price"`
	GameCreditPrice    int       `json:"game_credit_price" db:"game_credit_price"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}

type ChampionSkiknsOwnership struct {
	AccountID        int       `json:"account_id" db:"account_id"`
	ChampionName     string    `json:"champion_name" db:"champion_name"`
	ChampionSkinName string    `json:"champion_skin_name" db:"champion_skin_name"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}
