package model

import (
	"time"
)

type Account struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	Tag       string    `json:"tag"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ConnectedAccounts struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Type      string    `json:"type"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRelations struct {
	U1ID    int  `json:"u1_id"`
	U2ID    int  `json:"u2_id"`
	Blocked bool `json:"blocked"`
	Friend  bool `json:"friend"`
}

type UserSettings struct {
	UserID             int  `json:"user_id"`
	Is2FaEnabled       bool `json:"is_2fa_enabled"`
	IsSubscribedToNews bool `json:"is_subscribed_to_news"`
}
