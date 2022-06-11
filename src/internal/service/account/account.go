package account

import (
	"context"

	"github.com/alidevjimmy/db-project-go/internal/config"
	"github.com/alidevjimmy/db-project-go/internal/entity/model"
	"github.com/alidevjimmy/db-project-go/internal/pkg/logger"
	"github.com/alidevjimmy/db-project-go/internal/repository"
	"github.com/alidevjimmy/db-project-go/internal/service"
)

type acc struct {
	cfg    config.Account
	mysql  repository.Mysql
	logger logger.Logger
}

func New(cfg config.Account, mysql repository.Mysql, logger logger.Logger) service.Account {
	return &acc{
		cfg:    cfg,
		mysql:  mysql,
		logger: logger,
	}
}

func (a *acc) Login(ctx context.Context, username, password string) (string, error) {
	// start span with context

	user, err := a.mysql.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	_ = user
	// check password

	// create token
	token := ""

	return token, nil
}

func (a *acc) Register(ctx context.Context, user *model.User) (string, error) {
	return "", nil
}
