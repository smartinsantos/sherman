package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"sherman/src/domain/auth"
	"sherman/src/repository/mysqlds"
	"sherman/src/utils/exception"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	u := &auth.User{
		ID: uuid.New().String(),
		FirstName: "first",
		LastName: "last",
		EmailAddress: "some@email.com",
		Password: "some-password",
		Active: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var userRepo auth.UserRepository = &mysqlds.UserRepository{ DB: db }

	mock.
		ExpectExec("INSERT users SET").
		WithArgs(u.ID, u.FirstName, u.LastName, u.EmailAddress, u.Password, u.Active, u.CreatedAt, u.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1,1))

	err = userRepo.CreateUser(u)

	assert.NoError(t, err)
}

func TestCreateUserShouldTrowError(t *testing.T) {
	u := &auth.User{
		ID: uuid.New().String(),
		FirstName: "first",
		LastName: "last",
		EmailAddress: "some@email.com",
		Password: "some-password",
		Active: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var userRepo auth.UserRepository = &mysqlds.UserRepository{ DB: db }

	sqlmock.
		NewRows([]string{"id", "first_name", "last_name", "email_address", "password", "active", "created_at", "updated_at"}).
		AddRow(u.ID, u.FirstName, u.LastName, u.EmailAddress, u.Password, u.Active, u.CreatedAt, u.UpdatedAt)

	mock.
		ExpectExec("INSERT users SET").
		WithArgs(u.ID, u.FirstName, u.LastName, u.EmailAddress, u.Password, u.Active, u.CreatedAt, u.UpdatedAt).
		WillReturnError(exception.NewDuplicateEntryError("user already exist"))

	err = userRepo.CreateUser(u)

	assert.Error(t, err)
}