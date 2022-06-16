package model

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  int       `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
