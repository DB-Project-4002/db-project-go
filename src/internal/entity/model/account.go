package model

import (
	"time"
)

type Account struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Password  string    `json:"-" db:"password"`
	Tag       string    `json:"tag" db:"tag"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Blocked   bool      `json:"is_blocked" db:"blocked"`
}

type ConnectedAccounts struct {
	ID        int       `json:"id" db:"id"`
	AccountID int       `json:"account_id" db:"account_id"`
	Type      string    `json:"type" db:"type"`
	Address   string    `json:"address" db:"address"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type AccountRelations struct {
	A1ID      int       `json:"u1_id" db:"account_id_1"`
	A2ID      int       `json:"u2_id" db:"account_id_2"`
	Blocked   bool      `json:"blocked" db:"blocked"`
	Friend    bool      `json:"friend" db:"friend"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type AccountSettings struct {
	AccountID          int       `json:"account_id" db:"account_id"`
	Is2FaEnabled       bool      `json:"is_2fa_enabled" db:"is_2fa_enabled"`
	IsSubscribedToNews bool      `json:"is_subscribed_to_news" db:"is_subscribed_to_news"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}
