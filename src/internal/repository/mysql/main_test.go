package mysql

import (
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alidevjimmy/db-project-go/internal/repository"
)

var (
	testMysql repository.Mysql
	testMock  sqlmock.Sqlmock
)

func NewMock() (repository.Mysql, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	return &mysql{
		db:     db,
		logger: nil,
	}, mock
}

func TestMain(m *testing.M) {
	testMysql, testMock = NewMock()

	os.Exit(m.Run())
}
