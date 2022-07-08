package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
	"github.com/alidevjimmy/db-project-go/pkg/rest_err"
	"github.com/georgysavva/scany/sqlscan"
)

const (
	accountsTable                = "accounts"
	accountsRelationsTable       = "accounts_relations"
	gameAccountsTable            = "game_accounts"
	championsTable               = "champions"
	championSkiknsOwnershipTable = "champion_skin_ownerships"
	championSkinsTable           = "champion_skins"
)

func (m *mysql) CreateAccount(ctx context.Context, account *model.Account) (*int, rest_err.RestErr) {
	query := fmt.Sprintf("INSERT INTO %s (name,password,tag,email) VALUES (?,?,?,?)", accountsTable)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	res, err := m.db.ExecContext(ctx, query, account.Name, account.Password, account.Tag, account.Email)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return nil, errR
	}

	lastID, _ := res.LastInsertId()
	var uID int
	uID = int(lastID)

	gameAccQuery := fmt.Sprintf("INSERT INTO %s (account_id, name, level, avatar, avatar_border_id) VALUES (?,?,?,?,?)", gameAccountsTable)

	res, err = m.db.ExecContext(ctx, gameAccQuery, uID, account.Name, 1, "avatar", 1)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return nil, errR
	}

	return &uID, nil
}

func (m *mysql) GetAccountByID(ctx context.Context, id int) (*model.Account, rest_err.RestErr) {
	querry := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", accountsTable)

	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()

	acc := new(model.Account)
	sqlscan.Get(ctx, m.db, acc, querry, id)
	if acc.ID == 0 {
		errR := rest_err.NewRestErr(http.StatusNotFound, "Account not found")
		return nil, errR
	}
	return acc, nil
}

func (m *mysql) GetAccountByEmailAndPassword(ctx context.Context, email, password string) (*model.Account, rest_err.RestErr) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE email = ? AND password = ?", accountsTable)

	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()

	acc := new(model.Account)
	sqlscan.Get(ctx, m.db, acc, query, email, password)
	if acc.ID == 0 {
		errR := rest_err.NewRestErr(http.StatusNotFound, "Account not found")
		return nil, errR
	}

	return acc, nil
}

func (m *mysql) GetAccountByNameAndTagPassword(ctx context.Context, name, tag, password string) (*model.Account, rest_err.RestErr) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE name = ? AND tag = ? AND password = ?", accountsTable)

	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()

	acc := new(model.Account)
	sqlscan.Get(ctx, m.db, acc, query, name, tag, password)
	if acc.ID == 0 {
		errR := rest_err.NewRestErr(http.StatusNotFound, "Account not found")
		return nil, errR
	}

	return acc, nil
}

func (m *mysql) GetAccountFriendsByAccountID(ctx context.Context, accountID int) ([]*model.Account, rest_err.RestErr) {
	query := fmt.Sprintf("SELECT %s.* FROM %s JOIN %s ON %s.account_id_2 = %s.id AND %s.friend = 1 WHERE %s.account_id_1 = ?", accountsTable, accountsTable, accountsRelationsTable, accountsRelationsTable, accountsTable, accountsRelationsTable, accountsRelationsTable)

	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()

	var accounts []*model.Account

	err := sqlscan.Select(ctx, m.db, &accounts, query, accountID)

	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return nil, errR
	}

	return accounts, nil
}

func (m *mysql) AddAccountToFriends(ctx context.Context, accountID, friendID int) rest_err.RestErr {
	query := fmt.Sprintf("INSERT INTO %s (account_id_1,account_id_2,friend) VALUES (?,?,1)", accountsRelationsTable)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	res, err := m.db.ExecContext(ctx, query, accountID, friendID)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return errR
	}
	if r, _ := res.RowsAffected(); r == 0 {
		errR := rest_err.NewRestErr(http.StatusNotFound, "account not found")
		return errR
	}
	return nil
}
func (m *mysql) AddAccountToFriendsByUsername(ctx context.Context, accountID int, friendUsername string) rest_err.RestErr {
	username := strings.Split(friendUsername, "#")
	if len(username) != 2 {
		errR := rest_err.NewRestErr(http.StatusBadRequest, "username is not valid")
		return errR
	}
	user := new(model.Account)
	userQuery := fmt.Sprintf("SELECT * FROM %s WHERE name = ? AND tag = ?", accountsTable)
	if err := sqlscan.Get(ctx, m.db, user, userQuery, username[0], username[1]); err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			errR := rest_err.NewRestErr(http.StatusNotFound, "account not found")
			return errR
		}
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return errR
	}
	query := fmt.Sprintf("INSERT INTO %s (account_id_1,account_id_2,friend) VALUES (?,?,1)", accountsRelationsTable)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	res, err := m.db.ExecContext(ctx, query, accountID, user.ID)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return errR
	}
	if r, _ := res.RowsAffected(); r == 0 {
		errR := rest_err.NewRestErr(http.StatusNotFound, "account not found")
		return errR
	}
	return nil
}

func (m *mysql) BlockAccountFriend(ctx context.Context, accountID, friendID int) rest_err.RestErr {
	query := fmt.Sprintf("UPDATE %s SET blocked = 1, friend = 0 WHERE account_id_1 = ? AND account_id_2 = ?", accountsRelationsTable)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	res, err := m.db.ExecContext(ctx, query, accountID, friendID)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return errR
	}
	if r, _ := res.RowsAffected(); r == 0 {
		errR := rest_err.NewRestErr(http.StatusNotFound, "account not found")
		return errR
	}

	return nil
}

func (m *mysql) DeleteAccountFriend(ctx context.Context, accountID, friendID int) rest_err.RestErr {
	query := fmt.Sprintf("DELETE FROM %s WHERE account_id_1 = ? AND account_id_2 = ?", accountsRelationsTable)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	res, err := m.db.ExecContext(ctx, query, accountID, friendID)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return errR
	}
	if r, _ := res.RowsAffected(); r == 0 {
		errR := rest_err.NewRestErr(http.StatusNotFound, "account not found")
		return errR
	}
	return nil
}

func (m *mysql) GetAccountGameAccountsByAccountID(ctx context.Context, accountID int) ([]*model.GameAccount, rest_err.RestErr) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE account_id = ?", gameAccountsTable)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	var gameAccounts []*model.GameAccount
	err := sqlscan.Select(ctx, m.db, &gameAccounts, query, accountID)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return nil, errR
	}
	return gameAccounts, nil
}

func (m *mysql) CreateAccountGameAccount(ctx context.Context, accountID int, gameAccount *model.GameAccount) rest_err.RestErr {
	query := fmt.Sprintf("INSERT INTO %s (account_id, name, level, avatar, avatar_border_id, game_credit, blue_essence, orange_essence, mythic_essence) VALUES (?,?,?,?,?,?,?,?,?)", gameAccountsTable)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	res, err := m.db.ExecContext(ctx, query, accountID, gameAccount.Name, gameAccount.Level, gameAccount.Avatar, gameAccount.AvatarBorderID, gameAccount.GameCredit, gameAccount.BlueEssence, gameAccount.OrangeEssence, gameAccount.MythicEssence)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return errR
	}
	if r, _ := res.RowsAffected(); r == 0 {
		errR := rest_err.NewRestErr(http.StatusNotFound, "account not found")
		return errR
	}
	return nil
}

func (m *mysql) GetAccountGameAccountChampionsByAccountID(ctx context.Context, accountID int) ([]*model.Champion, rest_err.RestErr) {
	query := fmt.Sprintf(`SELECT name FROM %s JOIN %s ON %s.name = %s.champion_name WHERE %s.account_id = ?`, championsTable, championSkiknsOwnershipTable, championsTable, championSkiknsOwnershipTable, championSkiknsOwnershipTable)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	var champions []*model.Champion
	err := sqlscan.Select(ctx, m.db, &champions, query, accountID)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return nil, errR
	}
	return champions, nil
}

func (m *mysql) CreateAccountGameAccountChampionByChampionNameAndAccountID(ctx context.Context, accountID int, championName string) rest_err.RestErr {
	query := fmt.Sprintf("INSERT INTO %s (account_id, champion_name, champion_skin_name) VALUES (?,?,?)", championSkiknsOwnershipTable)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	res, err := m.db.ExecContext(ctx, query, accountID, championName, championName)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return errR
	}
	if r, _ := res.RowsAffected(); r == 0 {
		errR := rest_err.NewRestErr(http.StatusNotFound, "not found")
		return errR
	}
	return nil
}

func (m *mysql) GetAccountGameAccountChampionSkinsByChampionNameAndAccountID(ctx context.Context, accountID int, championName string) ([]*model.ChampionSkins, rest_err.RestErr) {
	query := fmt.Sprintf(`SELECT %s.name, %s.champion_name FROM %s JOIN %s ON %s.champion_name = %s.champion_name WHERE %s.account_id = ?`, championSkinsTable, championSkinsTable, championSkinsTable, championSkiknsOwnershipTable, championSkiknsOwnershipTable, championSkinsTable, championSkiknsOwnershipTable)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	var championSkins []*model.ChampionSkins
	err := sqlscan.Select(ctx, m.db, &championSkins, query, accountID)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return nil, errR
	}

	return championSkins, nil
}

func (m *mysql) CreateAccountGameAccountChampionSkinByChampionNameAndSkinNameAndAccountID(ctx context.Context, accountID int, championName, skinName string) rest_err.RestErr {
	query := fmt.Sprintf("INSERT INTO %s (account_id, champion_name, champion_skin_name) VALUES (?,?,?)", championSkiknsOwnershipTable)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	res, err := m.db.ExecContext(ctx, query, accountID, championName, skinName)
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return errR
	}
	if r, _ := res.RowsAffected(); r == 0 {
		errR := rest_err.NewRestErr(http.StatusNotFound, "account not found")
		return errR
	}
	return nil
}
