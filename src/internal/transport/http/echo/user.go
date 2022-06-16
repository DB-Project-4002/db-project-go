package echo

import (
	"github.com/alidevjimmy/db-project-go/internal/pkg/logger"
	"github.com/alidevjimmy/db-project-go/internal/service"
	"github.com/labstack/echo/v4"
)

type userController struct {
	logger  logger.Logger
	account service.User
}

func (h *userController) login(c echo.Context) error {
	return nil
}

func (h *userController) register(c echo.Context) error {
	return nil
}
