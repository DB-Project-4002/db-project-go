package repository

import (
	"context"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
	"github.com/alidevjimmy/db-project-go/pkg/rest_err"
)

type Mysql interface {
	CreateAccount(ctx context.Context, account *model.Account) (*int, rest_err.RestErr)
	GetAccountByID(ctx context.Context, id int) (*model.Account, rest_err.RestErr)
	GetAccountByEmailAndPassword(ctx context.Context, email, password string) (*model.Account, rest_err.RestErr)
	GetAccountByNameAndTagPassword(ctx context.Context, name, tag, password string) (*model.Account, rest_err.RestErr)

	GetAccountFriendsByAccountID(ctx context.Context, AccountID int) ([]*model.Account, rest_err.RestErr)
	AddAccountToFriends(ctx context.Context, AccountID, friendID int) rest_err.RestErr
	BlockAccountFriend(ctx context.Context, AccountID, friendID int) rest_err.RestErr
	DeleteAccountFriend(ctx context.Context, AccountID, friendID int) rest_err.RestErr

	GetAccountGameAccountsByAccountID(ctx context.Context, AccountID int) ([]*model.GameAccount, rest_err.RestErr)
	CreateAccountGameAccount(ctx context.Context, AccountID int, gameAccount *model.GameAccount) rest_err.RestErr

	GetAccountGameAccountChampionsByAccountID(ctx context.Context, AccountID int) ([]*model.Champion, rest_err.RestErr)
	CreateAccountGameAccountChampionByChampionNameAndAccountID(ctx context.Context, AccountID int, championName string) rest_err.RestErr

	GetAccountGameAccountChampionSkinsByChampionNameAndAccountID(ctx context.Context, AccountID int, championName string) ([]*model.ChampionSkins, rest_err.RestErr)
	CreateAccountGameAccountChampionSkinByChampionNameAndSkinNameAndAccountID(ctx context.Context, AccountID int, championName, skinName string) rest_err.RestErr
}
