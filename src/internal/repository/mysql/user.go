package mysql

import (
	"context"
	"fmt"
	"time"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
)

const (
	usersTable = "users"
)

func (m *mysql) CreateUser(ctx context.Context, user *model.User) error {
	query := fmt.Sprintf("INSERT INTO %s (name,password,tag,email) VALUES (?,?,?,?)", usersTable)
	ctx , cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	
	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		m.logger.Error(err.Error())
		return err
	}
	defer stmt.Close()
	res , err := stmt.ExecContext(ctx, user.Name, user.Password, user.Tag, user.Email)
	if err != nil {
		m.logger.Error(err.Error())
		return err
	}
	rows , err := res.RowsAffected()
	if err != nil {
		m.logger.Error(err.Error())
		return err
	}
	m.logger.Info(fmt.Sprintf("%d row affected", rows))
	return nil
}

func (m *mysql) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	return nil, nil
}

func (m *mysql) GetUserByEmailAndPassword(ctx context.Context, email, password string) (*model.User, error) {
	return nil, nil
}

func (m *mysql) GetUserFriendsByUserID(ctx context.Context, userID int) ([]*model.User, error) {
	return nil, nil
}

func (m *mysql) AddUserToFriends(ctx context.Context, userID, friendID int) error {
	return nil
}

func (m *mysql) BlockUserFriend(ctx context.Context, userID, friendID int) error {
	return nil
}

func (m *mysql) DeleteUserFriend(ctx context.Context, userID, friendID int) error {
	return nil
}

func (m *mysql) GetUserGameAccountsByUserID(ctx context.Context, userID int) ([]*model.GameAccount, error) {
	return nil, nil
}

func (m *mysql) CreateUserGameAccount(ctx context.Context, userID int, gameAccount *model.GameAccount) error {
	return nil
}

func (m *mysql) GetUserGameAccountChampionsByUserID(ctx context.Context, userID int) ([]*model.Champion, error) {
	return nil, nil
}

func (m *mysql) CreateUserGameAccountChampionByChampionNameAndUserID(ctx context.Context, userID int, championName string) error {
	return nil
}

func (m *mysql) GetUserGameAccountChampionSkinsByChampionNameAndUserID(ctx context.Context, userID int, championName string) ([]*model.ChampionSkins, error) {
	return nil, nil
}

func (m *mysql) CreateUserGameAccountChampionSkinByChampionNameAndSkinNameAndUserID(ctx context.Context, userID int, championName, skinName string) error {
	return nil
}
