package account

import (
	"context"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
	"github.com/alidevjimmy/db-project-go/internal/pkg/logger"
	"github.com/alidevjimmy/db-project-go/internal/repository"
	"github.com/alidevjimmy/db-project-go/internal/service"
	"github.com/alidevjimmy/db-project-go/pkg/hash"
	"github.com/alidevjimmy/db-project-go/pkg/rest_err"
)

type account struct {
	mysql  repository.Mysql
	logger logger.Logger
}

func New(mysql repository.Mysql, logger logger.Logger) service.Account {
	return &account{
		mysql:  mysql,
		logger: logger,
	}
}

func (u *account) Register(ctx context.Context, account *model.Account) (string, rest_err.RestErr) {
	account.Password = hash.GenerateSha256(account.Password)
	if err := u.mysql.CreateAccount(ctx, account); err != nil {
		return "", err
	}
	// generate jwt token
	return "", nil
}

func (u *account) Login(ctx context.Context, username, password string) (string, rest_err.RestErr) {
	return "", nil
}
