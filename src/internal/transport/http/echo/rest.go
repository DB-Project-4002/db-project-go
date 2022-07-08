package echo

import (
	"context"
	"time"

	"github.com/alidevjimmy/db-project-go/internal/pkg/logger"
	"github.com/alidevjimmy/db-project-go/internal/service"
	"github.com/alidevjimmy/db-project-go/internal/transport/http"
	"github.com/alidevjimmy/db-project-go/internal/transport/http/middlewares"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

type rest struct {
	echo               *echo.Echo
	accountController  *accountController
	accountMiddleware  *middlewares.AccountMiddleware
	championController *championController
}

func New(logger logger.Logger, accSrv service.Account, accMdwr middlewares.AccountMiddleware, champSrv service.Champion) http.Rest {
	return &rest{
		echo: echo.New(),
		accountController: &accountController{
			logger:  logger,
			account: accSrv,
		},
		championController: &championController{
			logger:   logger,
			champion: champSrv,
		},
		accountMiddleware: &middlewares.AccountMiddleware{
			JwtPkg: accMdwr.JwtPkg,
		},
	}
}

func (r *rest) Start(address string) error {
	r.echo.Use(middleware.Recover())

	r.routing()

	return r.echo.Start(address)
}

func (r *rest) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return r.echo.Shutdown(ctx)
}
