package mysql

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alidevjimmy/db-project-go/internal/entity/model"
	"github.com/stretchr/testify/assert"
)

var (
	user = &model.User{
		Name:     "ali",
		Tag:      "ali:007",
		Password: "password",
		Email:    "ali@email.com",
	}
)

func TestCreateUser(t *testing.T) {
	query := fmt.Sprintf("INSERT INTO %s (name,password,tag,email) VALUES (?,?,?,?)", usersTable)
	testMock.ExpectBegin()
	testMock.ExpectExec(query).WithArgs(user.Name, user.Password, user.Tag, user.Email).WillReturnResult(sqlmock.NewResult(1, 1))
	testMock.ExpectCommit()

	err := testMysql.CreateUser(context.Background(), user)
	assert.NoError(t, err)

	err = testMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

// func TestCreateDuplicateUserEmail(t *testing.T) {

// }

// func TestCreateUserError(t *testing.T) {
// 	mysql, mock := NewMock()

// 	query := "SELECT * FROM users WHERE email = ?"
// }
