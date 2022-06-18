package mysql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
)

const (
	usersTable          = "accounts"
	usersRelationsTable = "accounts_relations"
)


var (
	createUserQuery                = fmt.Sprintf("INSERT INTO %s (name,password,tag,email) VALUES (?,?,?,?)", usersTable)
	getUserByIDQuery               = fmt.Sprintf("SELECT * FROM %s WHERE id = ?", usersTable)
	getUserByEmailAndPasswordQuery = fmt.Sprintf("SELECT * FROM %s WHERE email = ? AND password = ?", usersTable)
	getUserFriendsByUserIDQuery    = fmt.Sprintf("SELECT * FROM %s JOIN %s ON %s.account_id_2 = %s.id AND %s.friend = 1 WHERE %s.account_id_1 = ?", usersTable, usersRelationsTable, usersRelationsTable, usersRelationsTable, usersRelationsTable,usersRelationsTable)
)

func (m *mysql) CreateUser(ctx context.Context, user *model.User) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	stmt, err := m.db.PrepareContext(ctx, createUserQuery)

	if err != nil {
		m.logger.Error(err.Error())
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, user.Name, user.Password, user.Tag, user.Email)
	if err != nil {
		m.logger.Error(err.Error())
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		m.logger.Error(err.Error())
		return err
	}

	m.logger.Info(fmt.Sprintf("%d row affected", rows))
	return nil
}

func (m *mysql) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()

	stmt, err := m.db.PrepareContext(ctx, getUserByIDQuery)
	if err != nil {
		m.logger.Error(err.Error())
		return nil, err
	}
	defer stmt.Close()

	u := new(model.User)

	row := stmt.QueryRowContext(ctx, id)
	if row.Err() != nil {
		errMsg := errors.New("user not found")
		m.logger.Error(errMsg.Error())
		return nil, errors.New(errMsg.Error())
	}

	if err = row.Scan(&u.ID, &u.Name, &u.Password, &u.Tag, &u.Email, &u.CreatedAt, &u.UpdatedAt); err != nil {
		m.logger.Error(err.Error())
		return nil, err
	}

	return u, nil
}

func (m *mysql) GetUserByEmailAndPassword(ctx context.Context, email, password string) (*model.User, error) {
	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()

	stmt, err := m.db.PrepareContext(ctx, getUserByEmailAndPasswordQuery)
	if err != nil {
		m.logger.Error(err.Error())
		return nil, err
	}
	defer stmt.Close()

	u := new(model.User)

	row := stmt.QueryRowContext(ctx, email, password)
	if row.Err() != nil {
		errMsg := errors.New("user not found")
		m.logger.Error(errMsg.Error())
		return nil, errors.New(errMsg.Error())
	}

	if err = row.Scan(&u.ID, &u.Name, &u.Password, &u.Tag, &u.Email, &u.CreatedAt, &u.UpdatedAt); err != nil {
		m.logger.Error(err.Error())
		return nil, err
	}

	return u, nil
}

func (m *mysql) GetUserFriendsByUserID(ctx context.Context, userID int) ([]*model.User, error) {
	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()

	stmt, err := m.db.PrepareContext(ctx, getUserFriendsByUserIDQuery)
	if err != nil {
		m.logger.Error(err.Error())
		return nil, err
	}
	defer stmt.Close()

	var users []*model.User

	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		m.logger.Error(err.Error())
		return nil, errors.New(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var u model.User
		err = rows.Scan(&u.ID, &u.Name, &u.Password, &u.Tag, &u.Email, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			m.logger.Error(err.Error())
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
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
