package repository

import (
	"context"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
	"github.com/alidevjimmy/db-project-go/pkg/rest_err"
)

type Mysql interface {
	CreateAccount(ctx context.Context, account *model.Account) (*int, rest_err.RestErr)
	GetUserByID(ctx context.Context, id int) (*model.Account, rest_err.RestErr)
	GetUserByEmailAndPassword(ctx context.Context, email, password string) (*model.Account, rest_err.RestErr)
	GetUserByNameAndTagPassword(ctx context.Context, name, tag, password string) (*model.Account, rest_err.RestErr)

	GetUserFriendsByUserID(ctx context.Context, userID int) ([]*model.Account, rest_err.RestErr)
	AddUserToFriends(ctx context.Context, userID, friendID int) rest_err.RestErr
	BlockUserFriend(ctx context.Context, userID, friendID int) rest_err.RestErr
	DeleteUserFriend(ctx context.Context, userID, friendID int) rest_err.RestErr

	GetUserGameAccountsByUserID(ctx context.Context, userID int) ([]*model.GameAccount, rest_err.RestErr)
	CreateUserGameAccount(ctx context.Context, userID int, gameAccount *model.GameAccount) rest_err.RestErr

	GetUserGameAccountChampionsByUserID(ctx context.Context, userID int) ([]*model.Champion, rest_err.RestErr)
	CreateUserGameAccountChampionByChampionNameAndUserID(ctx context.Context, userID int, championName string) rest_err.RestErr

	GetUserGameAccountChampionSkinsByChampionNameAndUserID(ctx context.Context, userID int, championName string) ([]*model.ChampionSkins, rest_err.RestErr)
	CreateUserGameAccountChampionSkinByChampionNameAndSkinNameAndUserID(ctx context.Context, userID int, championName, skinName string) rest_err.RestErr
}
