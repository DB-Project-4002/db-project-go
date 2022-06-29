package account

import (
	"context"
	"net/http"
	"time"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
	"github.com/alidevjimmy/db-project-go/internal/pkg/logger"
	"github.com/alidevjimmy/db-project-go/internal/repository"
	"github.com/alidevjimmy/db-project-go/internal/service"
	"github.com/alidevjimmy/db-project-go/pkg/hash"
	"github.com/alidevjimmy/db-project-go/pkg/jwt"
	"github.com/alidevjimmy/db-project-go/pkg/rest_err"
)

type account struct {
	mysql  repository.Mysql
	logger logger.Logger
	jwtpkg jwt.Jwt
}

func New(mysql repository.Mysql, logger logger.Logger, jwtpkg jwt.Jwt) service.Account {
	return &account{
		mysql:  mysql,
		logger: logger,
		jwtpkg: jwtpkg,
	}
}

func (u *account) Register(ctx context.Context, account *model.Account) (*string, rest_err.RestErr) {
	account.Password = hash.GenerateSha256(account.Password)
	uID, err := u.mysql.CreateAccount(ctx, account)
	if err != nil {
		return nil, err
	}
	claims := map[string]interface{}{
		"exp": time.Now().Unix() + int64(time.Hour*24*30*2),
		"iat": time.Now().Unix(),
		"sub": uID,
	}
	token, errR := u.jwtpkg.GenerateToken(claims)
	if err != nil {
		err := rest_err.NewRestErr(http.StatusInternalServerError, errR.Error())
		return nil, err
	}
	return &token, nil
}

func (u *account) Login(ctx context.Context, name, tag, password string) (*string, rest_err.RestErr) {
	password = hash.GenerateSha256(password)
	account, err := u.mysql.GetAccountByNameAndTagPassword(ctx, name, tag, password)
	if err != nil {
		return nil, err
	}
	claims := map[string]interface{}{
		"exp": time.Now().Unix() + int64(time.Hour*24*30*2),
		"iat": time.Now().Unix(),
		"sub": account.ID,
	}
	token, errR := u.jwtpkg.GenerateToken(claims)
	if err != nil {
		err := rest_err.NewRestErr(http.StatusInternalServerError, errR.Error())
		return nil, err
	}
	return &token, nil
}
