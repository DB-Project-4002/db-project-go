package response

import "time"

type (
	LoginData struct {
		Token     string `json:"token"`
		AccountID int    `json:"account_id"`
	}
	Login struct {
		Data LoginData `json:"data"`
	}

	RegisterData struct {
		Token     string `json:"token"`
		AccountID int    `json:"account_id"`
	}
	Register struct {
		Data RegisterData `json:"data"`
	}

	GetAccountData struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Tag       string    `json:"tag"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	GetAccount struct {
		Data GetAccountData `json:"data"`
	}

	GetAccountFriednsData struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Tag     string `json:"tag"`
		Email   string `json:"email"`
		Blocked bool   `json:"is_blocked"`
	}

	GetAccountFriends struct {
		Data []GetAccountFriednsData `json:"data"`
	}

	AddAccountToFriendsData struct {
		Message string `json:"message"`
	}

	AddAccountToFriends struct {
		Data AddAccountToFriendsData `json:"data"`
	}

	BlockAccountFriendData struct {
		Message string `json:"message"`
	}

	BlockAccountFriend struct {
		Data BlockAccountFriendData `json:"data"`
	}

	DeleteAccountFriendData struct {
		Message string `json:"message"`
	}

	DeleteAccountFriend struct {
		Data DeleteAccountFriendData `json:"data"`
	}

	GetAccountGameAccountsData struct {
		AccountID      int       `json:"account_id"`
		Name           string    `json:"name"`
		Level          int       `json:"level"`
		Avatar         string    `json:"avatar"`
		AvatarBorderID int       `json:"avatar_border_id"`
		GameCredit     int       `json:"game_credit"`
		BlueEssence    int       `json:"blue_essence"`
		OrangeEssence  int       `json:"orange_essence"`
		MythicEssence  int       `json:"mythic_essence"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	GetAccountGameAccounts struct {
		Data []GetAccountGameAccountsData `json:"data"`
	}

	CreateAccountGameAccountData struct {
		Message string `json:"message"`
	}

	CreateAccountGameAccount struct {
		Data CreateAccountGameAccountData `json:"data"`
	}

	GetAccountGameAccountChampionsData struct {
		ChampionName string `json:"champion_name"`
	}

	GetAccountGameAccountChampions struct {
		Data []GetAccountGameAccountChampionsData `json:"data"`
	}

	CreateAccountGameAccountChampionData struct {
		Message string `json:"message"`
	}

	CreateAccountGameAccountChampion struct {
		Data CreateAccountGameAccountChampionData `json:"data"`
	}

	CreateAccountGameAccountChampionSkinData struct {
		Message string `json:"message"`
	}

	CreateAccountGameAccountChampionSkin struct {
		Data CreateAccountGameAccountChampionSkinData `json:"data"`
	}

	GetAccountGameAccountChampionSkinsData struct {
		Name string `json:"name"`
	}

	GetAccountGameAccountChampionSkins struct {
		Data []GetAccountGameAccountChampionSkinsData `json:"data"`
	}
)
