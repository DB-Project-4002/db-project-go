package user

import (
	"context"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
	"github.com/alidevjimmy/db-project-go/internal/pkg/logger"
	"github.com/alidevjimmy/db-project-go/internal/repository"
	"github.com/alidevjimmy/db-project-go/internal/service"
)

type user struct {
	mysql  repository.Mysql
	logger logger.Logger
}

func New(mysql repository.Mysql, logger logger.Logger) service.User {
	return &user{
		mysql:  mysql,
		logger: logger,
	}
}

func (u *user) Login(ctx context.Context, username, password string) (string, error) {
	return "", nil
}

func (u *user) Register(ctx context.Context, user *model.User) (string, error) {
	return "", nil
}
