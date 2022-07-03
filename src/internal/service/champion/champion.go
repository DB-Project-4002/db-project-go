package champion

import (
	"context"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
	"github.com/alidevjimmy/db-project-go/internal/pkg/logger"
	"github.com/alidevjimmy/db-project-go/internal/repository"
	"github.com/alidevjimmy/db-project-go/internal/service"
	"github.com/alidevjimmy/db-project-go/pkg/rest_err"
)

type champion struct {
	mysql  repository.Mysql
	logger logger.Logger
}

func New(mysql repository.Mysql, logger logger.Logger) service.Champion {
	return &champion{
		mysql:  mysql,
		logger: logger,
	}
}

func (c *champion) GetChampion(ctx context.Context, name string) (*model.Champion, rest_err.RestErr) {
	champ, err := c.mysql.GetChampionByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return champ, nil
}

func (c *champion) GetChampionSkin(ctx context.Context, skinName string) (*model.ChampionSkins, rest_err.RestErr) {
	champ, err := c.mysql.GetChampionSkin(ctx, skinName)
	if err != nil {
		return nil, err
	}

	return champ, nil
}

func (c *champion) GetChampionSkins(ctx context.Context, champName string) ([]*model.ChampionSkins, rest_err.RestErr) {
	champ, err := c.mysql.GetChampionSkins(ctx, champName)
	if err != nil {
		return nil, err
	}

	return champ, nil
}

func (c *champion) GetChampions(ctx context.Context) ([]*model.Champion, rest_err.RestErr) {
	champ, err := c.mysql.GetChampions(ctx)
	if err != nil {
		return nil, err
	}

	return champ, nil
}
