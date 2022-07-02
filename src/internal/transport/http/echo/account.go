package echo

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/alidevjimmy/db-project-go/internal/entity/model"
	"github.com/alidevjimmy/db-project-go/internal/pkg/logger"
	"github.com/alidevjimmy/db-project-go/internal/service"
	"github.com/alidevjimmy/db-project-go/internal/transport/http/request"
	"github.com/alidevjimmy/db-project-go/internal/transport/http/response"
	"github.com/labstack/echo/v4"
)

type accountController struct {
	logger  logger.Logger
	account service.Account
}

func (h *accountController) register(c echo.Context) error {
	data := new(request.Register)
	if err := (&echo.DefaultBinder{}).Bind(data, c); err != nil {
		errResp := response.ErrorResp{
			Error: response.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		}
		return c.JSON(errResp.Error.Code, errResp)
	}
	if err := validate.Struct(data); err != nil {
		errResp := response.ErrorResp{
			Error: response.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		}
		return c.JSON(errResp.Error.Code, errResp)
	}
	account := model.Account{
		Name:     data.Name,
		Password: data.Password,
		Email:    data.Email,
		Tag:      data.Tag,
	}
	token, err := h.account.Register(context.Background(), &account)
	if err != nil {
		errResp := response.ErrorResp{
			Error: response.Error{
				Code:    err.StatusCode(),
				Message: err.Error(),
			},
		}
		return c.JSON(err.StatusCode(), errResp)
	}
	resp := response.Register{
		Data: response.RegisterData{
			Token: *token,
		},
	}
	return c.JSON(http.StatusCreated, resp)
}

func (h *accountController) login(c echo.Context) error {
	data := new(request.Login)
	if err := (&echo.DefaultBinder{}).Bind(data, c); err != nil {
		errResp := response.ErrorResp{
			Error: response.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		}
		return c.JSON(errResp.Error.Code, errResp)
	}
	if err := validate.Struct(data); err != nil {
		errResp := response.ErrorResp{
			Error: response.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		}
		return c.JSON(errResp.Error.Code, errResp)
	}

	username := data.Username
	password := data.Password

	unameSplit := strings.Split(username, "#")

	if len(unameSplit) != 2 {
		errResp := response.ErrorResp{
			Error: response.Error{
				Code:    http.StatusBadRequest,
				Message: "Invalid username",
			},
		}
		return c.JSON(errResp.Error.Code, errResp)
	}
	token, err := h.account.Login(context.Background(), unameSplit[0], unameSplit[1], password)
	if err != nil {
		errResp := response.ErrorResp{
			Error: response.Error{
				Code:    err.StatusCode(),
				Message: err.Error(),
			},
		}
		return c.JSON(err.StatusCode(), errResp)
	}
	resp := response.Register{
		Data: response.RegisterData{
			Token: *token,
		},
	}
	return c.JSON(http.StatusCreated, resp)
}

func (h *accountController) getAccount(c echo.Context) error {
	accIDParam := c.Param("account_id")
	accID, _ := strconv.Atoi(accIDParam)

	acc, err := h.account.GetAccount(context.Background(), accID)
	if err != nil {
		errR := response.ErrorResp{
			Error: response.Error{
				Code:    err.StatusCode(),
				Message: err.Error(),
			},
		}
		return c.JSON(errR.Error.Code, errR)
	}
	resp := response.GetAccount{
		Data: response.GetAccountData{
			ID:        acc.ID,
			Name:      acc.Name,
			Tag:       acc.Tag,
			Email:     acc.Email,
			CreatedAt: acc.CreatedAt,
			UpdatedAt: acc.CreatedAt,
		},
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *accountController) addAccountToFriends(c echo.Context) error {
	accIDParam := c.Param("account_id")
	accID, _ := strconv.Atoi(accIDParam)

	tIDParam := c.Param("target_account_id")
	tID, _ := strconv.Atoi(tIDParam)

	err := h.account.AddAccountToFriends(context.Background(), accID, tID)
	if err != nil {
		errR := response.ErrorResp{
			Error: response.Error{
				Code:    err.StatusCode(),
				Message: err.Error(),
			},
		}
		return c.JSON(errR.Error.Code, errR)
	}
	resp := response.AddAccountToFriends{
		Data: response.AddAccountToFriendsData{
			Message: "account added to your friend list",
		},
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *accountController) blockAccountFriend(c echo.Context) error {
	accIDParam := c.Param("account_id")
	accID, _ := strconv.Atoi(accIDParam)

	tIDParam := c.Param("target_account_id")
	tID, _ := strconv.Atoi(tIDParam)

	err := h.account.BlockAccountFriend(context.Background(), accID, tID)
	if err != nil {
		errR := response.ErrorResp{
			Error: response.Error{
				Code:    err.StatusCode(),
				Message: err.Error(),
			},
		}
		return c.JSON(errR.Error.Code, errR)
	}
	resp := response.BlockAccountFriend{
		Data: response.BlockAccountFriendData{
			Message: "account blocked",
		},
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *accountController) createAccountGameAccountChampion(c echo.Context) error {
	accIDParam := c.Param("account_id")
	accID, _ := strconv.Atoi(accIDParam)

	champName := c.Param("champion_name")

	err := h.account.CreateAccountGameAccountChampionByChampionNameAndAccountID(context.Background(), accID, champName)
	if err != nil {
		errR := response.ErrorResp{
			Error: response.Error{
				Code:    err.StatusCode(),
				Message: err.Error(),
			},
		}
		return c.JSON(errR.Error.Code, errR)
	}
	resp := response.BlockAccountFriend{
		Data: response.BlockAccountFriendData{
			Message: "champion added to your account",
		},
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *accountController) createAccountGameAccountChampionSkin(c echo.Context) error {
	accIDParam := c.Param("account_id")
	accID, _ := strconv.Atoi(accIDParam)

	champName := c.Param("champion_name")
	skinName := c.Param("skin_name")

	err := h.account.CreateAccountGameAccountChampionSkinByChampionNameAndSkinNameAndAccountID(context.Background(), accID, champName, skinName)
	if err != nil {
		errR := response.ErrorResp{
			Error: response.Error{
				Code:    err.StatusCode(),
				Message: err.Error(),
			},
		}
		return c.JSON(errR.Error.Code, errR)
	}
	resp := response.CreateAccountGameAccountChampionSkin{
		Data: response.CreateAccountGameAccountChampionSkinData{
			Message: "skin added to your account",
		},
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *accountController) deleteAccountFriend(c echo.Context) error {
	accIDParam := c.Param("account_id")
	accID, _ := strconv.Atoi(accIDParam)

	tIDParam := c.Param("target_account_id")
	tID, _ := strconv.Atoi(tIDParam)

	err := h.account.DeleteAccountFriend(context.Background(), accID, tID)
	if err != nil {
		errR := response.ErrorResp{
			Error: response.Error{
				Code:    err.StatusCode(),
				Message: err.Error(),
			},
		}
		return c.JSON(errR.Error.Code, errR)
	}
	resp := response.DeleteAccountFriend{
		Data: response.DeleteAccountFriendData{
			Message: "account removed from friend list",
		},
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *accountController) getAccountFriends(c echo.Context) error {
	accIDParam := c.Param("account_id")
	accID, _ := strconv.Atoi(accIDParam)

	accs, err := h.account.GetAccountFriendsByAccountID(context.Background(), accID)
	if err != nil {
		errR := response.ErrorResp{
			Error: response.Error{
				Code:    err.StatusCode(),
				Message: err.Error(),
			},
		}
		return c.JSON(errR.Error.Code, errR)
	}
	respAccs := []response.GetAccountFriednsData{}
	for _, acc := range accs {
		a := response.GetAccountFriednsData{
			ID:    acc.ID,
			Name:  acc.Name,
			Tag:   acc.Tag,
			Email: acc.Email,
		}
		respAccs = append(respAccs, a)
	}
	resp := response.GetAccountFriends{
		Data: respAccs,
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *accountController) getAccountGameAccountChampionSkins(c echo.Context) error {
	accIDParam := c.Param("account_id")
	accID, _ := strconv.Atoi(accIDParam)

	champName := c.Param("champion_name")

	champs, err := h.account.GetAccountGameAccountChampionSkinsByChampionNameAndAccountID(context.Background(), accID, champName)
	if err != nil {
		errR := response.ErrorResp{
			Error: response.Error{
				Code:    err.StatusCode(),
				Message: err.Error(),
			},
		}
		return c.JSON(errR.Error.Code, errR)
	}
	respSkins := []response.GetAccountGameAccountChampionSkinsData{}
	for _, skn := range champs {
		a := response.GetAccountGameAccountChampionSkinsData{
			Name:         skn.Name,
			ChampionName: skn.ChampionName,
		}
		respSkins = append(respSkins, a)
	}
	resp := response.GetAccountGameAccountChampionSkins{
		Data: respSkins,
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *accountController) getAccountGameAccountChampions(c echo.Context) error {
	accIDParam := c.Param("account_id")
	accID, _ := strconv.Atoi(accIDParam)

	champs, err := h.account.GetAccountGameAccountChampionsByAccountID(context.Background(), accID)
	if err != nil {
		errR := response.ErrorResp{
			Error: response.Error{
				Code:    err.StatusCode(),
				Message: err.Error(),
			},
		}
		return c.JSON(errR.Error.Code, errR)
	}
	respGAcc := []response.GetAccountGameAccountChampionsData{}
	for _, champ := range champs {
		a := response.GetAccountGameAccountChampionsData{
			ChampionName: champ.Name,
		}
		respGAcc = append(respGAcc, a)
	}
	resp := response.GetAccountGameAccountChampions{
		Data: respGAcc,
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *accountController) getAccountGameAccounts(c echo.Context) error {
	accIDParam := c.Param("account_id")
	accID, _ := strconv.Atoi(accIDParam)

	gAccs, err := h.account.GetAccountGameAccountsByAccountID(context.Background(), accID)
	if err != nil {
		errR := response.ErrorResp{
			Error: response.Error{
				Code:    err.StatusCode(),
				Message: err.Error(),
			},
		}
		return c.JSON(errR.Error.Code, errR)
	}
	respGAcc := []response.GetAccountGameAccountsData{}
	for _, gAcc := range gAccs {
		a := response.GetAccountGameAccountsData{
			AccountID:      gAcc.AccountID,
			Name:           gAcc.Name,
			Level:          gAcc.Level,
			Avatar:         gAcc.Avatar,
			AvatarBorderID: gAcc.AvatarBorderID,
			GameCredit:     gAcc.GameCredit,
			BlueEssence:    gAcc.BlueEssence,
			OrangeEssence:  gAcc.OrangeEssence,
			MythicEssence:  gAcc.MythicEssence,
			CreatedAt:      gAcc.CreatedAt,
			UpdatedAt:      gAcc.UpdatedAt,
		}
		respGAcc = append(respGAcc, a)
	}
	resp := response.GetAccountGameAccounts{
		Data: respGAcc,
	}
	return c.JSON(http.StatusOK, resp)
}
