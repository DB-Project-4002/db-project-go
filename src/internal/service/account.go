package service

import (
	"context"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
	"github.com/alidevjimmy/db-project-go/pkg/rest_err"
)

type Account interface {
	Login(ctx context.Context, username, password string) (string, rest_err.RestErr)
	Register(ctx context.Context, account *model.Account) (string, rest_err.RestErr)
}
