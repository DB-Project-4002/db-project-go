package service

import (
	"context"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
	"github.com/alidevjimmy/db-project-go/pkg/rest_err"
)

type Champion interface {
	GetChampions(ctx context.Context) ([]*model.Champion, rest_err.RestErr)
	GetChampion(ctx context.Context, name string) (*model.Champion, rest_err.RestErr)
	GetChampionSkins(ctx context.Context, champName string) ([]*model.ChampionSkins, rest_err.RestErr)
	GetChampionSkin(ctx context.Context, skinName string) (*model.ChampionSkins, rest_err.RestErr)
}
