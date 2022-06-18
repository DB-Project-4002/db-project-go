package mysql

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alidevjimmy/db-project-go/internal/entity/model"
	"github.com/stretchr/testify/assert"
)

var (
	user = model.User{
		Name:     "ali",
		Tag:      "ali:007",
		Password: "password",
		Email:    "ali@email.com",
	}
)

func TestCreateUser(t *testing.T) {
	prep := testMock.ExpectPrepare(createUserQuery)
	prep.ExpectExec().WithArgs(user.Name, user.Password, user.Tag, user.Email).WillReturnResult(sqlmock.NewResult(1, 1))

	err := testMysql.CreateUser(context.Background(), &user)
	assert.NoError(t, err)

	err = testMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetUserByID(t *testing.T) {
	prep := testMock.ExpectPrepare(getUserByIDQuery)
	user.ID = 1
	prep.ExpectQuery().WithArgs(user.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "tag", "password", "email", "created_at", "updated_at"}).AddRow(user.ID, user.Name, user.Tag, user.Password, user.Email, user.CreatedAt, user.UpdatedAt))

	rowUser, err := testMysql.GetUserByID(context.Background(), user.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, rowUser)

	err = testMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetUserByEmailAndPassword(t *testing.T) {
	prep := testMock.ExpectPrepare(getUserByEmailAndPasswordQuery)
	user.ID = 1
	prep.ExpectQuery().WithArgs(user.Email, user.Password).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "tag", "password", "email", "created_at", "updated_at"}).AddRow(user.ID, user.Name, user.Tag, user.Password, user.Email, user.CreatedAt, user.UpdatedAt))

	rowUser, err := testMysql.GetUserByEmailAndPassword(context.Background(), user.Email, user.Password)
	assert.NoError(t, err)
	assert.NotEmpty(t, rowUser)

	err = testMock.ExpectationsWereMet()
	assert.NoError(t, err)
}
