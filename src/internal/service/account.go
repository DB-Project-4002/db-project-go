package service

import (
	"context"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
	"github.com/alidevjimmy/db-project-go/pkg/rest_err"
)

type Account interface {
	Login(ctx context.Context, name, tag, password string) (*string, *int, rest_err.RestErr)
	Register(ctx context.Context, account *model.Account) (*string, *int, rest_err.RestErr)
	GetAccount(ctx context.Context, accountID int) (*model.Account, rest_err.RestErr)

	GetAccountFriendsByAccountID(ctx context.Context, accountID int) ([]*model.Account, rest_err.RestErr)
	AddAccountToFriends(ctx context.Context, accountID, friendID int) rest_err.RestErr
	AddAccountToFriendsByUsername(ctx context.Context, accountID int, friendUsername string) rest_err.RestErr
	BlockAccountFriend(ctx context.Context, accountID, friendID int) rest_err.RestErr
	DeleteAccountFriend(ctx context.Context, accountID, friendID int) rest_err.RestErr

	GetAccountGameAccountsByAccountID(ctx context.Context, accountID int) ([]*model.GameAccount, rest_err.RestErr)
	CreateAccountGameAccount(ctx context.Context, AccountID int, gameAccount *model.GameAccount) rest_err.RestErr

	GetAccountGameAccountChampionsByAccountID(ctx context.Context, accountID int) ([]*model.Champion, rest_err.RestErr)
	CreateAccountGameAccountChampionByChampionNameAndAccountID(ctx context.Context, accountID int, championName string) rest_err.RestErr

	GetAccountGameAccountChampionSkinsByChampionNameAndAccountID(ctx context.Context, accountID int, championName string) ([]*model.ChampionSkiknsOwnership, rest_err.RestErr)
	CreateAccountGameAccountChampionSkinByChampionNameAndSkinNameAndAccountID(ctx context.Context, accountID int, championName, skinName string) rest_err.RestErr
}
