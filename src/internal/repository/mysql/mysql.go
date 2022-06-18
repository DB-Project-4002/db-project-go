package mysql

import (
	"database/sql"
	"fmt"

	"github.com/alidevjimmy/db-project-go/internal/config"
	"github.com/alidevjimmy/db-project-go/internal/pkg/logger"
	"github.com/alidevjimmy/db-project-go/internal/repository"
)

type mysql struct {
	db     *sql.DB
	logger logger.Logger
}

func New(cfg config.Mysql, logger logger.Logger) (repository.Mysql, error) {
	db, err := sql.Open("mysql", dsn(cfg))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &mysql{db: db, logger: logger}, nil
}

func dsn(cfg config.Mysql) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.DBName)
}
