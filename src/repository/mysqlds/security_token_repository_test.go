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

func TestCreateOrUpdateToken(t *testing.T) {
	st := &auth.SecurityToken{
		ID:        uuid.New().String(),
		UserID:    "some-user-id",
		Token:     "some-user-token",
		Type:      "some-token-type",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("should insert", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var securityTokenRepo auth.SecurityTokenRepository = &SecurityTokenRepository{DB: db}

		mock.
			ExpectQuery("SELECT id FROM security_tokens").
			WithArgs(st.UserID, st.Type).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))

		mock.
			ExpectExec("INSERT security_tokens SET").
			WithArgs(st.ID, st.UserID, st.Token, st.Type, st.CreatedAt, st.UpdatedAt).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err = securityTokenRepo.CreateOrUpdateToken(st)

		assert.NoError(t, err)
	})

	t.Run("should update", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var securityTokenRepo auth.SecurityTokenRepository = &SecurityTokenRepository{DB: db}

		rows := sqlmock.NewRows([]string{"id"}).AddRow(st.ID)

		mock.
			ExpectQuery("SELECT id FROM security_tokens").
			WithArgs(st.UserID, st.Type).
			WillReturnRows(rows)

		mock.
			ExpectExec("UPDATE security_tokens SET").
			WithArgs(st.Token, st.UpdatedAt, st.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err = securityTokenRepo.CreateOrUpdateToken(st)

		assert.NoError(t, err)
	})

	t.Run("should return an error", func(t *testing.T) {
		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		db.Close()

		var securityTokenRepo auth.SecurityTokenRepository = &SecurityTokenRepository{DB: db}

		err = securityTokenRepo.CreateOrUpdateToken(st)

		assert.Error(t, err)
	})

}

func TestGetTokenByMetadata(t *testing.T) {
	st := &auth.SecurityToken{
		ID:        uuid.New().String(),
		UserID:    "some-user-id",
		Token:     "some-user-token",
		Type:      "some-token-type",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tmd := &auth.TokenMetadata{
		UserID: st.UserID,
		Type:   st.Type,
		Token:  st.Token,
	}

	t.Run("should return a token", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var securityTokenRepo auth.SecurityTokenRepository = &SecurityTokenRepository{DB: db}

		rows := sqlmock.
			NewRows([]string{"id", "user_id", "token", "type", "created_at", "updated_at"}).
			AddRow(st.ID, st.UserID, st.Token, st.Type, st.CreatedAt, st.UpdatedAt)
		mock.
			ExpectQuery("SELECT id, user_id, token, type, created_at, updated_at FROM security_tokens").
			WithArgs(st.UserID, st.Type).
			WillReturnRows(rows)

		token, err := securityTokenRepo.GetTokenByMetadata(tmd)

		assert.EqualValues(t, st, &token)
		assert.NoError(t, err)
	})

	t.Run("should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var securityTokenRepo auth.SecurityTokenRepository = &SecurityTokenRepository{DB: db}

		expectedError := exception.NewNotFoundError("token not found")

		mock.
			ExpectQuery("SELECT id, user_id, token, type, created_at, updated_at FROM security_tokens").
			WithArgs(st.UserID, st.Type).
			WillReturnError(expectedError)

		_, err = securityTokenRepo.GetTokenByMetadata(tmd)

		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
	})
}

func TestRemoveTokenMetadata(t *testing.T) {
	tmd := &auth.TokenMetadata{
		UserID: "some-user-id",
		Type:   "some-token-type",
		Token:  "some-user-token",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var securityTokenRepo auth.SecurityTokenRepository = &SecurityTokenRepository{DB: db}

	mock.
		ExpectExec("DELETE FROM security_tokens WHERE").
		WithArgs(tmd.UserID, tmd.Type).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = securityTokenRepo.RemoveTokenByMetadata(tmd)

	assert.NoError(t, err)
}
