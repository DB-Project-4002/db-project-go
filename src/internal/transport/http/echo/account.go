package echo

import (
	"context"
	"net/http"

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
			Token: token,
		},
	}
	return c.JSON(http.StatusCreated, resp)
}

func (h *accountController) login(c echo.Context) error {
	return nil
}
