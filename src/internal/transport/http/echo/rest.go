package echo

import (
	"context"
	"time"

	"github.com/alidevjimmy/db-project-go/internal/pkg/logger"
	"github.com/alidevjimmy/db-project-go/internal/service"
	"github.com/alidevjimmy/db-project-go/internal/transport/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type rest struct {
	echo           *echo.Echo
	userController *userController
}

func New(logger logger.Logger, accSrv service.User) http.Rest {
	return &rest{
		echo: echo.New(),
		userController: &userController{
			logger:  logger,
			account: accSrv,
		}}
}

func (r *rest) Start(address string) error {
	r.echo.Use(middleware.Recover())

	r.routing()

	return r.echo.Start(address)
}

func (r *rest) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // use config for time
	defer cancel()

	return r.echo.Shutdown(ctx)
}
