package mysql

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
	"github.com/alidevjimmy/db-project-go/pkg/rest_err"
	"github.com/georgysavva/scany/sqlscan"
)

const (
	accountsTable          = "accounts"
	accountsRelationsTable = "accounts_relations"
)

var (
	createAccountQuery                     = fmt.Sprintf("INSERT INTO %s (name,password,tag,email) VALUES (?,?,?,?)", accountsTable)
	getAccountByIDQuery                    = fmt.Sprintf("SELECT * FROM %s WHERE id = ?", accountsTable)
	getAccountByEmailAndPasswordQuery      = fmt.Sprintf("SELECT * FROM %s WHERE email = ? AND password = ?", accountsTable)
	getAccountByNameAndTagAndPasswordQuery = fmt.Sprintf("SELECT * FROM %s WHERE name = ? AND tag = ? AND password = ?", accountsTable)
	getAccountFriendsByAccountIDQuery      = fmt.Sprintf("SELECT * FROM %s JOIN %s ON %s.account_id_2 = %s.id AND %s.friend = 1 WHERE %s.account_id_1 = ?", accountsTable, accountsRelationsTable, accountsRelationsTable, accountsRelationsTable, accountsRelationsTable, accountsRelationsTable)
)

func (m *mysql) CreateAccount(ctx context.Context, account *model.Account) (*int, rest_err.RestErr) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	stmt, err := m.db.PrepareContext(ctx, createAccountQuery)

	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return nil, errR
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, account.Name, account.Password, account.Tag, account.Email)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return nil, errR
	}

	lastID, _ := res.LastInsertId()
	var uID int
	uID = int(lastID)
	return &uID, nil
}

func (m *mysql) GetAccountByID(ctx context.Context, id int) (*model.Account, rest_err.RestErr) {
	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()

	acc := new(model.Account)
	sqlscan.Get(ctx, m.db, acc, getAccountByIDQuery, id)
	if acc.ID == 0 {
		errR := rest_err.NewRestErr(http.StatusNotFound, "Account not found")
		return nil, errR
	}
	return acc, nil
}

func (m *mysql) GetAccountByEmailAndPassword(ctx context.Context, email, password string) (*model.Account, rest_err.RestErr) {
	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()

	acc := new(model.Account)
	sqlscan.Get(ctx, m.db, acc, getAccountByEmailAndPasswordQuery, email, password)
	if acc.ID == 0 {
		errR := rest_err.NewRestErr(http.StatusNotFound, "Account not found")
		return nil, errR
	}

	return acc, nil
}

func (m *mysql) GetAccountByNameAndTagPassword(ctx context.Context, name, tag, password string) (*model.Account, rest_err.RestErr) {
	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()

	acc := new(model.Account)
	sqlscan.Get(ctx, m.db, acc, getAccountByNameAndTagAndPasswordQuery, name, tag, password)
	if acc.ID == 0 {
		errR := rest_err.NewRestErr(http.StatusNotFound, "Account not found")
		return nil, errR
	}

	return acc, nil
}

func (m *mysql) GetAccountFriendsByAccountID(ctx context.Context, accountID int) ([]*model.Account, rest_err.RestErr) {
	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()


	var accounts []*model.Account

	err := sqlscan.Select(ctx, m.db, accounts ,getAccountFriendsByAccountIDQuery, accountID)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return nil, errR
	}
	
	return accounts, nil
}

func (m *mysql) AddAccountToFriends(ctx context.Context, AccountID, friendID int) rest_err.RestErr {
	return nil
}

func (m *mysql) BlockAccountFriend(ctx context.Context, AccountID, friendID int) rest_err.RestErr {
	return nil
}

func (m *mysql) DeleteAccountFriend(ctx context.Context, AccountID, friendID int) rest_err.RestErr {
	return nil
}

func (m *mysql) GetAccountGameAccountsByAccountID(ctx context.Context, AccountID int) ([]*model.GameAccount, rest_err.RestErr) {
	return nil, nil
}

func (m *mysql) CreateAccountGameAccount(ctx context.Context, AccountID int, gameAccount *model.GameAccount) rest_err.RestErr {
	return nil
}

func (m *mysql) GetAccountGameAccountChampionsByAccountID(ctx context.Context, AccountID int) ([]*model.Champion, rest_err.RestErr) {
	return nil, nil
}

func (m *mysql) CreateAccountGameAccountChampionByChampionNameAndAccountID(ctx context.Context, AccountID int, championName string) rest_err.RestErr {
	return nil
}

func (m *mysql) GetAccountGameAccountChampionSkinsByChampionNameAndAccountID(ctx context.Context, AccountID int, championName string) ([]*model.ChampionSkins, rest_err.RestErr) {
	return nil, nil
}

func (m *mysql) CreateAccountGameAccountChampionSkinByChampionNameAndSkinNameAndAccountID(ctx context.Context, AccountID int, championName, skinName string) rest_err.RestErr {
	return nil
}
