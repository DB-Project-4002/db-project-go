package model

type Champion struct {
	Name             string `json:"name"`
	BlueEssencePrice int    `json:"blue_essence_price"`
	GameCreditPrice  int    `json:"game_credit_price"`
	DefaultSkinName  string `json:"default_skin_name"` // foreign key to ChampionSkins
}

type ChampionSkins struct {
	Name               string `json:"name"`
	ChampionName       string `json:"champion_name"`
	OrangeEssencePrice int    `json:"orange_essence_price"`
	GameCreditPrice    int    `json:"game_credit_price"`
}

type ChampionSkiknsOwnership struct {
	UserID           int    `json:"user_id"`
	ChampionName     string `json:"champion_name"`
	ChampionSkinName string `json:"champion_skin_name"`
}
