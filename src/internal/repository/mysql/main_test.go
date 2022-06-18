package mysql

import (
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alidevjimmy/db-project-go/internal/pkg/logger/zap"
	"github.com/alidevjimmy/db-project-go/internal/repository"
	"go.uber.org/zap/zapcore"
)

var (
	testMysql repository.Mysql
	testMock  sqlmock.Sqlmock
)

func NewMock() (repository.Mysql, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	f, err := os.OpenFile("../../../logs/test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	logger := zap.New(f, zapcore.ErrorLevel)

	return &mysql{
		db:     db,
		logger: logger,
	}, mock
}

func TestMain(m *testing.M) {
	testMysql, testMock = NewMock()

	os.Exit(m.Run())
}
