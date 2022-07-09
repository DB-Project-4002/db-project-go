package echo

import (
	"context"
	"net/http"

	"github.com/alidevjimmy/db-project-go/internal/pkg/logger"
	"github.com/alidevjimmy/db-project-go/internal/service"
	"github.com/alidevjimmy/db-project-go/internal/transport/http/response"
	"github.com/labstack/echo/v4"
)

type championController struct {
	logger   logger.Logger
	champion service.Champion
}

func (h *championController) getChampions(c echo.Context) error {
	champs, err := h.champion.GetChampions(context.Background())
	if err != nil {
		errR := response.ErrorResp{
			Error: response.Error{
				Code:    err.StatusCode(),
				Message: err.Error(),
			},
		}
		return c.JSON(errR.Error.Code, errR)
	}
	respChamps := []response.GetChampionsData{}
	for _, champ := range champs {
		a := response.GetChampionsData{
			Name:             champ.Name,
			BlueEssencePrice: champ.BlueEssencePrice,
			GameCreditPrice:  champ.GameCreditPrice,
			CreatedAt:        champ.CreatedAt,
			UpdatedAt:        champ.UpdatedAt,
		}
		respChamps = append(respChamps, a)
	}
	resp := response.GetChampions{
		Data: respChamps,
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *championController) getChampionSkins(c echo.Context) error {
	champName := c.Param("champion_name")

	skins, err := h.champion.GetChampionSkins(context.Background(), champName)
	if err != nil {
		errR := response.ErrorResp{
			Error: response.Error{
				Code:    err.StatusCode(),
				Message: err.Error(),
			},
		}
		return c.JSON(errR.Error.Code, errR)
	}
	respSkins := []response.GetChampionSkinsData{}
	for _, skin := range skins {
		a := response.GetChampionSkinsData{
			Name:               skin.Name,
			ChampionName:       skin.ChampionName,
			GameCreditPrice:    skin.GameCreditPrice,
			OrangeEssencePrice: skin.OrangeEssencePrice,
			CreatedAt:          skin.CreatedAt,
			UpdatedAt:          skin.UpdatedAt,
		}
		respSkins = append(respSkins, a)
	}
	resp := response.GetChampionSkins{
		Data: respSkins,
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *championController) getChampion(c echo.Context) error {
	champName := c.Param("champion_name")

	champ, err := h.champion.GetChampion(context.Background(), champName)
	if err != nil {
		errR := response.ErrorResp{
			Error: response.Error{
				Code:    err.StatusCode(),
				Message: err.Error(),
			},
		}
		return c.JSON(errR.Error.Code, errR)
	}
	resp := response.GetChampion{
		Data: response.GetChampionData{
			Name:             champ.Name,
			BlueEssencePrice: champ.BlueEssencePrice,
			GameCreditPrice:  champ.GameCreditPrice,
			CreatedAt:        champ.CreatedAt,
			UpdatedAt:        champ.UpdatedAt,
		},
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *championController) getChampionSkin(c echo.Context) error {
	skinName := c.Param("skin_name")

	skin, err := h.champion.GetChampionSkin(context.Background(), skinName)
	if err != nil {
		errR := response.ErrorResp{
			Error: response.Error{
				Code:    err.StatusCode(),
				Message: err.Error(),
			},
		}
		return c.JSON(errR.Error.Code, errR)
	}
	resp := response.GetChampionSkin{
		Data: response.GetChampionSkinData{
			Name:               skin.Name,
			GameCreditPrice:    skin.GameCreditPrice,
			ChampionName:       skin.ChampionName,
			OrangeEssencePrice: skin.OrangeEssencePrice,
			CreatedAt:          skin.CreatedAt,
			UpdatedAt:          skin.UpdatedAt,
		},
	}
	return c.JSON(http.StatusOK, resp)
}
