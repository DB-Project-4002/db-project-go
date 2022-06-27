package mysql

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
	"github.com/alidevjimmy/db-project-go/pkg/rest_err"
)

const (
	usersTable          = "accounts"
	usersRelationsTable = "accounts_relations"
)

var (
	createUserQuery                = fmt.Sprintf("INSERT INTO %s (name,password,tag,email) VALUES (?,?,?,?)", usersTable)
	getUserByIDQuery               = fmt.Sprintf("SELECT * FROM %s WHERE id = ?", usersTable)
	getUserByEmailAndPasswordQuery = fmt.Sprintf("SELECT * FROM %s WHERE email = ? AND password = ?", usersTable)
	getUserFriendsByUserIDQuery    = fmt.Sprintf("SELECT * FROM %s JOIN %s ON %s.account_id_2 = %s.id AND %s.friend = 1 WHERE %s.account_id_1 = ?", usersTable, usersRelationsTable, usersRelationsTable, usersRelationsTable, usersRelationsTable, usersRelationsTable)
)

func (m *mysql) CreateAccount(ctx context.Context, account *model.Account) rest_err.RestErr {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	stmt, err := m.db.PrepareContext(ctx, createUserQuery)

	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return errR
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, account.Name, account.Password, account.Tag, account.Email)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return errR
	}

	_, err = res.RowsAffected()
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return errR
	}

	return nil
}

func (m *mysql) GetUserByID(ctx context.Context, id int) (*model.Account, rest_err.RestErr) {
	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()

	stmt, err := m.db.PrepareContext(ctx, getUserByIDQuery)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return nil, errR
	}
	defer stmt.Close()

	u := new(model.Account)

	row := stmt.QueryRowContext(ctx, id)
	if row.Err() != nil {
		errMsg := errors.New("user not found")
		errR := rest_err.NewRestErr(http.StatusNotFound, errMsg.Error())
		return nil, errR
	}

	if err = row.Scan(&u.ID, &u.Name, &u.Password, &u.Tag, &u.Email, &u.CreatedAt, &u.UpdatedAt); err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return nil, errR
	}

	return u, nil
}

func (m *mysql) GetUserByEmailAndPassword(ctx context.Context, email, password string) (*model.Account, rest_err.RestErr) {
	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()

	stmt, err := m.db.PrepareContext(ctx, getUserByEmailAndPasswordQuery)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return nil, errR
	}
	defer stmt.Close()

	u := new(model.Account)

	row := stmt.QueryRowContext(ctx, email, password)
	if row.Err() != nil {
		errMsg := errors.New("user not found")
		errR := rest_err.NewRestErr(http.StatusInternalServerError, errMsg.Error())

		return nil, errR
	}

	if err = row.Scan(&u.ID, &u.Name, &u.Password, &u.Tag, &u.Email, &u.CreatedAt, &u.UpdatedAt); err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())

		return nil, errR
	}

	return u, nil
}

func (m *mysql) GetUserFriendsByUserID(ctx context.Context, userID int) ([]*model.Account, rest_err.RestErr) {
	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()

	stmt, err := m.db.PrepareContext(ctx, getUserFriendsByUserIDQuery)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())

		return nil, errR
	}
	defer stmt.Close()

	var users []*model.Account

	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())

		return nil, errR
	}
	defer rows.Close()

	for rows.Next() {
		var u model.Account
		err = rows.Scan(&u.ID, &u.Name, &u.Password, &u.Tag, &u.Email, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())

			return nil, errR
		}
		users = append(users, &u)
	}

	return users, nil
}

func (m *mysql) AddUserToFriends(ctx context.Context, userID, friendID int) rest_err.RestErr {
	return nil
}

func (m *mysql) BlockUserFriend(ctx context.Context, userID, friendID int) rest_err.RestErr {
	return nil
}

func (m *mysql) DeleteUserFriend(ctx context.Context, userID, friendID int) rest_err.RestErr {
	return nil
}

func (m *mysql) GetUserGameAccountsByUserID(ctx context.Context, userID int) ([]*model.GameAccount, rest_err.RestErr) {
	return nil, nil
}

func (m *mysql) CreateUserGameAccount(ctx context.Context, userID int, gameAccount *model.GameAccount) rest_err.RestErr {
	return nil
}

func (m *mysql) GetUserGameAccountChampionsByUserID(ctx context.Context, userID int) ([]*model.Champion, rest_err.RestErr) {
	return nil, nil
}

func (m *mysql) CreateUserGameAccountChampionByChampionNameAndUserID(ctx context.Context, userID int, championName string) rest_err.RestErr {
	return nil
}

func (m *mysql) GetUserGameAccountChampionSkinsByChampionNameAndUserID(ctx context.Context, userID int, championName string) ([]*model.ChampionSkins, rest_err.RestErr) {
	return nil, nil
}

func (m *mysql) CreateUserGameAccountChampionSkinByChampionNameAndSkinNameAndUserID(ctx context.Context, userID int, championName, skinName string) rest_err.RestErr {
	return nil
}
