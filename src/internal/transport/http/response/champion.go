package response

import "time"

type (
	GetChampionsData struct {
		Name             string    `json:"name"`
		BlueEssencePrice int       `json:"blue_essence_price"`
		GameCreditPrice  int       `json:"game_credit_price"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
	}

	GetChampions struct {
		Data []GetChampionsData `json:"data"`
	}
	GetChampionData struct {
		Name             string    `json:"name"`
		BlueEssencePrice int       `json:"blue_essence_price"`
		GameCreditPrice  int       `json:"game_credit_price"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
	}

	GetChampion struct {
		Data GetChampionData `json:"data"`
	}

	GetChampionSkinsData struct {
		Name               string    `json:"name" `
		ChampionName       string    `json:"champion_name" `
		OrangeEssencePrice int       `json:"orange_essence_price" `
		GameCreditPrice    int       `json:"game_credit_price" `
		CreatedAt          time.Time `json:"created_at" `
		UpdatedAt          time.Time `json:"updated_at" `
	}

	GetChampionSkins struct {
		Data []GetChampionSkinsData `json:"data"`
	}
	GetChampionSkinData struct {
		Name               string    `json:"name" `
		ChampionName       string    `json:"champion_name" `
		OrangeEssencePrice int       `json:"orange_essence_price" `
		GameCreditPrice    int       `json:"game_credit_price" `
		CreatedAt          time.Time `json:"created_at" `
		UpdatedAt          time.Time `json:"updated_at" `
	}

	GetChampionSkin struct {
		Data GetChampionSkinData `json:"data"`
	}
)
