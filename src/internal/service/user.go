package service

import (
	"context"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
)

type User interface {
	Login(ctx context.Context, username, password string) (string, error)
	Register(ctx context.Context, user *model.User) (string, error)
}
