package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
	"github.com/alidevjimmy/db-project-go/pkg/rest_err"
	"github.com/georgysavva/scany/sqlscan"
)

func (m *mysql) GetChampions(ctx context.Context) ([]*model.Champion, rest_err.RestErr) {
	query := fmt.Sprintf("SELECT * FROM %s", championsTable)

	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()

	var champs []*model.Champion

	err := sqlscan.Select(ctx, m.db, &champs, query)

	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return nil, errR
	}

	return champs, nil
}

func (m *mysql) GetChampionByName(ctx context.Context, name string) (*model.Champion, rest_err.RestErr) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE name = ?", championsTable)

	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()

	champ := new(model.Champion)
	err := sqlscan.Get(ctx, m.db, champ, query, name)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			errR := rest_err.NewRestErr(http.StatusNotFound, "champion not found")
			return nil, errR
		}
		errR := rest_err.NewRestErr(http.StatusNotFound, err.Error())
		return nil, errR
	}
	return champ, nil
}

func (m *mysql) GetChampionSkins(ctx context.Context, champName string) ([]*model.ChampionSkins, rest_err.RestErr) {
	query := fmt.Sprintf(`
	SELECT * FROM %s WHERE champion_name = ?
	`,
		championSkinsTable,
	)
	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()

	var skins []*model.ChampionSkins

	err := sqlscan.Select(ctx, m.db, &skins, query, champName)

	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return nil, errR
	}

	return skins, nil
}

func (m *mysql) GetChampionSkin(ctx context.Context, skinName string) (*model.ChampionSkins, rest_err.RestErr) {
	query := fmt.Sprintf(`
	SELECT * FROM %s WHERE name = ?
	`,
		championSkinsTable,
	)

	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleFunc()

	skin := new(model.ChampionSkins)

	err := sqlscan.Get(ctx, m.db, skin, query, skinName)

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			errR := rest_err.NewRestErr(http.StatusNotFound, "skin not found")
			return nil, errR
		}
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())
		return nil, errR
	}

	return skin, nil
}
