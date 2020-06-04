package mysqlds

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"sherman/src/domain/auth"
	"sherman/src/utils/exception"
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
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var userRepo auth.UserRepository = &UserRepository{DB: db}

		mock.
			ExpectExec("INSERT users SET").
			WithArgs(u.ID, u.FirstName, u.LastName, u.EmailAddress, u.Password, u.Active, u.CreatedAt, u.UpdatedAt).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err = userRepo.CreateUser(u)

		assert.NoError(t, err)
	})

	t.Run("should throw an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var userRepo auth.UserRepository = &UserRepository{DB: db}

		expectedError := exception.NewDuplicateEntryError("user already exist")

		sqlmock.
			NewRows([]string{"id", "first_name", "last_name", "email_address", "password", "active", "created_at", "updated_at"}).
			AddRow(u.ID, u.FirstName, u.LastName, u.EmailAddress, u.Password, u.Active, u.CreatedAt, u.UpdatedAt)

		err = userRepo.CreateUser(u)

		mock.
			ExpectExec("INSERT users SET").
			WithArgs(u.ID, u.FirstName, u.LastName, u.EmailAddress, u.Password, u.Active, u.CreatedAt, u.UpdatedAt).
			WillReturnError(expectedError)

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
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var userRepo auth.UserRepository = &UserRepository{DB: db}

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
		wrongID := "some-wrong-user-id"

		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var userRepo auth.UserRepository = &UserRepository{DB: db}

		expectedError := exception.NewNotFoundError("user not found")

		sqlmock.
			NewRows([]string{"id", "first_name", "last_name", "email_address", "password", "active", "created_at", "updated_at"}).
			AddRow(u.ID, u.FirstName, u.LastName, u.EmailAddress, u.Password, u.Active, u.CreatedAt, u.UpdatedAt)

		_, err = userRepo.GetUserByID(wrongID)

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
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var userRepo auth.UserRepository = &UserRepository{DB: db}

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
		wrongEmail := "wrongg@email.com"

		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var userRepo auth.UserRepository = &UserRepository{DB: db}

		expectedError := exception.NewNotFoundError("user not found")

		sqlmock.
			NewRows([]string{"id", "first_name", "last_name", "email_address", "password", "active", "created_at", "updated_at"}).
			AddRow(u.ID, u.FirstName, u.LastName, u.EmailAddress, u.Password, u.Active, u.CreatedAt, u.UpdatedAt)

		_, err = userRepo.GetUserByEmail(wrongEmail)

		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
	})
}
