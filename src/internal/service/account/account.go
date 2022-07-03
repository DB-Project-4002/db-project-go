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

func (a *account) Register(ctx context.Context, account *model.Account) (*string, rest_err.RestErr) {
	account.Password = hash.GenerateSha256(account.Password)
	uID, err := a.mysql.CreateAccount(ctx, account)
	if err != nil {
		return nil, err
	}
	claims := map[string]interface{}{
		"exp": time.Now().Unix() + int64(time.Hour*24*30*2),
		"iat": time.Now().Unix(),
		"sub": uID,
	}
	token, errR := a.jwtpkg.GenerateToken(claims)
	if err != nil {
		err := rest_err.NewRestErr(http.StatusInternalServerError, errR.Error())
		return nil, err
	}
	return &token, nil
}

func (a *account) Login(ctx context.Context, name, tag, password string) (*string, rest_err.RestErr) {
	password = hash.GenerateSha256(password)
	account, err := a.mysql.GetAccountByNameAndTagPassword(ctx, name, tag, password)
	if err != nil {
		return nil, err
	}
	claims := map[string]interface{}{
		"exp": time.Now().Unix() + int64(time.Hour*24*30*2),
		"iat": time.Now().Unix(),
		"sub": account.ID,
	}
	token, errR := a.jwtpkg.GenerateToken(claims)
	if err != nil {
		err := rest_err.NewRestErr(http.StatusInternalServerError, errR.Error())
		return nil, err
	}
	return &token, nil
}

func (a *account) GetAccount(ctx context.Context, accountID int) (*model.Account, rest_err.RestErr) {
	acc, err := a.mysql.GetAccountByID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (a *account) AddAccountToFriends(ctx context.Context, accountID int, friendID int) rest_err.RestErr {
	err := a.mysql.AddAccountToFriends(ctx, accountID, friendID)
	if err != nil {
		return err
	}

	return nil
}
func (a *account) AddAccountToFriendsByUsername(ctx context.Context, accountID int, friendUsername string) rest_err.RestErr {
	err := a.mysql.AddAccountToFriendsByUsername(ctx, accountID, friendUsername)
	if err != nil {
		return err
	}

	return nil
}

func (a *account) BlockAccountFriend(ctx context.Context, accountID int, friendID int) rest_err.RestErr {
	err := a.mysql.BlockAccountFriend(context.Background(), accountID, friendID)
	if err != nil {
		return err
	}

	return nil
}

func (a *account) CreateAccountGameAccount(ctx context.Context, AccountID int, gameAccount *model.GameAccount) rest_err.RestErr {
	err := a.mysql.CreateAccountGameAccount(context.Background(), AccountID, gameAccount)
	if err != nil {
		return err
	}

	return nil
}

func (a *account) CreateAccountGameAccountChampionByChampionNameAndAccountID(ctx context.Context, accountID int, championName string) rest_err.RestErr {
	err := a.mysql.CreateAccountGameAccountChampionByChampionNameAndAccountID(context.Background(), accountID, championName)
	if err != nil {
		return err
	}

	return nil
}

func (a *account) CreateAccountGameAccountChampionSkinByChampionNameAndSkinNameAndAccountID(ctx context.Context, accountID int, championName string, skinName string) rest_err.RestErr {
	err := a.mysql.CreateAccountGameAccountChampionSkinByChampionNameAndSkinNameAndAccountID(context.Background(), accountID, championName, skinName)
	if err != nil {
		return err
	}

	return nil
}

func (a *account) DeleteAccountFriend(ctx context.Context, accountID int, friendID int) rest_err.RestErr {
	err := a.mysql.DeleteAccountFriend(context.Background(), accountID, friendID)
	if err != nil {
		return err
	}

	return nil
}

func (a *account) GetAccountFriendsByAccountID(ctx context.Context, accountID int) ([]*model.Account, rest_err.RestErr) {
	accs, err := a.mysql.GetAccountFriendsByAccountID(context.Background(), accountID)
	if err != nil {
		return nil, err
	}

	return accs, nil
}

func (a *account) GetAccountGameAccountChampionSkinsByChampionNameAndAccountID(ctx context.Context, accountID int, championName string) ([]*model.ChampionSkins, rest_err.RestErr) {
	skns, err := a.mysql.GetAccountGameAccountChampionSkinsByChampionNameAndAccountID(context.Background(), accountID, championName)
	if err != nil {
		return nil, err
	}

	return skns, nil
}

func (a *account) GetAccountGameAccountChampionsByAccountID(ctx context.Context, accountID int) ([]*model.Champion, rest_err.RestErr) {
	chmps, err := a.mysql.GetAccountGameAccountChampionsByAccountID(context.Background(), accountID)
	if err != nil {
		return nil, err
	}

	return chmps, nil
}

func (a *account) GetAccountGameAccountsByAccountID(ctx context.Context, accountID int) ([]*model.GameAccount, rest_err.RestErr) {
	gas, err := a.mysql.GetAccountGameAccountsByAccountID(context.Background(), accountID)
	if err != nil {
		return nil, err
	}

	return gas, nil
}
