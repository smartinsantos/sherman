package mysqlds

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	_ "sherman/src/app/testing"
	"sherman/src/app/utils/terr"
	"sherman/src/domain/auth"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	u := &auth.User{
		ID:           uuid.New().String(),
		FirstName:    "first",
		LastName:     "last",
		EmailAddress: "some@email.com",
		Password:     "some-password",
		Active:       true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	t.Run("should insert", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}
		defer db.Close()

		userRepo := NewUserRepository(db)

		mock.
			ExpectExec("INSERT users SET").
			WithArgs(u.ID, u.FirstName, u.LastName, u.EmailAddress, u.Password, u.Active, u.CreatedAt, u.UpdatedAt).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err = userRepo.CreateUser(u)

		assert.NoError(t, err)
	})

	t.Run("should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}
		defer db.Close()

		userRepo := NewUserRepository(db)

		returnError := terr.NewDuplicateEntryError("duplicate")
		mock.
			ExpectExec("INSERT users SET").
			WithArgs(u.ID, u.FirstName, u.LastName, u.EmailAddress, u.Password, u.Active, u.CreatedAt, u.UpdatedAt).
			WillReturnError(returnError)

		err = userRepo.CreateUser(u)

		expectedError := terr.NewDuplicateEntryError("user already exist")
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
	})
}

func TestGetUserByID(t *testing.T) {
	u := &auth.User{
		ID:           uuid.New().String(),
		FirstName:    "first",
		LastName:     "last",
		EmailAddress: "some@email.com",
		Password:     "some-password",
		Active:       true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	t.Run("should return a user", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}
		defer db.Close()

		userRepo := NewUserRepository(db)

		rows := sqlmock.
			NewRows([]string{"id", "first_name", "last_name", "email_address", "password", "active", "created_at", "updated_at"}).
			AddRow(u.ID, u.FirstName, u.LastName, u.EmailAddress, u.Password, u.Active, u.CreatedAt, u.UpdatedAt)

		mock.
			ExpectQuery("SELECT id, first_name, last_name, email_address, password, active, created_at, updated_at FROM users").
			WithArgs(u.ID).
			WillReturnRows(rows)

		user, err := userRepo.GetUserByID(u.ID)

		assert.EqualValues(t, u, &user)
		assert.NoError(t, err)
	})

	t.Run("should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}
		defer db.Close()

		userRepo := NewUserRepository(db)
		wrongID := "some-wrong-user-id"

		returnError := errors.New("no rows")
		mock.
			ExpectQuery("SELECT id, first_name, last_name, email_address, password, active, created_at, updated_at FROM users").
			WithArgs(wrongID).
			WillReturnError(returnError)

		_, err = userRepo.GetUserByID(wrongID)

		expectedError := terr.NewNotFoundError("user not found")
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
	})
}

func TestGetUserByEmail(t *testing.T) {
	u := &auth.User{
		ID:           uuid.New().String(),
		FirstName:    "first",
		LastName:     "last",
		EmailAddress: "some@email.com",
		Password:     "some-password",
		Active:       true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	t.Run("should return a user", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}
		defer db.Close()

		userRepo := NewUserRepository(db)

		rows := sqlmock.
			NewRows([]string{"id", "first_name", "last_name", "email_address", "password", "active", "created_at", "updated_at"}).
			AddRow(u.ID, u.FirstName, u.LastName, u.EmailAddress, u.Password, u.Active, u.CreatedAt, u.UpdatedAt)

		mock.
			ExpectQuery("SELECT id, first_name, last_name, email_address, password, active, created_at, updated_at FROM users").
			WithArgs(u.EmailAddress).
			WillReturnRows(rows)

		user, err := userRepo.GetUserByEmail(u.EmailAddress)

		assert.EqualValues(t, u, &user)
		assert.NoError(t, err)
	})

	t.Run("should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}
		defer db.Close()

		userRepo := NewUserRepository(db)

		wrongEmail := "wrongg@email.com"

		returnError := errors.New("no rows")
		mock.
			ExpectQuery("SELECT id, first_name, last_name, email_address, password, active, created_at, updated_at FROM users").
			WithArgs(wrongEmail).
			WillReturnError(returnError)

		_, err = userRepo.GetUserByEmail(wrongEmail)

		expectedError := terr.NewNotFoundError("user not found")
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
	})
}
