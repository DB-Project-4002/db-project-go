package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/alidevjimmy/db-project-go/internal/config"
	"github.com/alidevjimmy/db-project-go/internal/pkg/logger/zap"
	"github.com/alidevjimmy/db-project-go/internal/repository/mysql"
	"github.com/alidevjimmy/db-project-go/internal/service/account"
	"github.com/alidevjimmy/db-project-go/internal/service/champion"
	"github.com/alidevjimmy/db-project-go/internal/transport/http/echo"
	"github.com/alidevjimmy/db-project-go/internal/transport/http/middlewares"
	"github.com/alidevjimmy/db-project-go/pkg/jwt"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap/zapcore"
)

var serveCMD = &cli.Command{
	Name:    "serve",
	Aliases: []string{"s"},
	Usage:   "serve http",
	Action:  serve,
}

func serve(c *cli.Context) error {
	cfg := new(config.Config)
	config.ReadYAML("config.yml", cfg)
	f, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	logger := zap.New(f, zapcore.ErrorLevel)

	mysqlRepo, err := mysql.New(cfg.Mysql, logger)
	if err != nil {
		return err
	}

	jwtpkg := jwt.New(cfg.App.Secret)
	accountSrv := account.New(mysqlRepo, logger, jwtpkg)
	champSrv := champion.New(mysqlRepo, logger)

	accMdwr := middlewares.NewAccountMiddleware(jwtpkg)

	restServer := echo.New(logger, accountSrv, accMdwr, champSrv)
	go func() {
		if err := restServer.Start(cfg.App.Address); err != nil {
			logger.Error(fmt.Sprintf("error happen while serving: %v", err))
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("\nReceived an interrupt, closing connections...")

	if err := restServer.Shutdown(); err != nil {
		fmt.Println("\nRest server doesn't shutdown in 10 seconds")
	}

	return nil
}
