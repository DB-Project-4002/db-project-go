package mysql

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/alidevjimmy/db-project-go/internal/config"
	"github.com/alidevjimmy/db-project-go/internal/pkg/logger"
	"github.com/alidevjimmy/db-project-go/internal/repository"
	"github.com/alidevjimmy/db-project-go/pkg/rest_err"
)

type mysql struct {
	db     *sql.DB
	logger logger.Logger
}

func New(cfg config.Mysql, logger logger.Logger) (repository.Mysql, rest_err.RestErr) {
	db, err := sql.Open("mysql", dsn(cfg))
	if err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())

		return nil, errR
	}

	if err := db.Ping(); err != nil {
		errR := rest_err.NewRestErr(http.StatusInternalServerError, err.Error())

		return nil, errR
	}
	return &mysql{db: db, logger: logger}, nil
}

func dsn(cfg config.Mysql) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.DBName)
}
