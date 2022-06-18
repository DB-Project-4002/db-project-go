package repository

import (
	"context"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
)

type Mysql interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	GetUserByEmailAndPassword(ctx context.Context, email, password string) (*model.User, error)

	GetUserFriendsByUserID(ctx context.Context, userID int) ([]*model.User, error)
	AddUserToFriends(ctx context.Context, userID, friendID int) error
	BlockUserFriend(ctx context.Context, userID, friendID int) error
	DeleteUserFriend(ctx context.Context, userID, friendID int) error

	GetUserGameAccountsByUserID(ctx context.Context, userID int) ([]*model.GameAccount, error)
	CreateUserGameAccount(ctx context.Context, userID int, gameAccount *model.GameAccount) error

	GetUserGameAccountChampionsByUserID(ctx context.Context, userID int) ([]*model.Champion, error)
	CreateUserGameAccountChampionByChampionNameAndUserID(ctx context.Context, userID int, championName string) error

	GetUserGameAccountChampionSkinsByChampionNameAndUserID(ctx context.Context, userID int, championName string) ([]*model.ChampionSkins, error)
	CreateUserGameAccountChampionSkinByChampionNameAndSkinNameAndUserID(ctx context.Context, userID int, championName, skinName string) error
}
