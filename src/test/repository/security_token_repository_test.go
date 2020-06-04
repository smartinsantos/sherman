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

func TestCreateOrUpdateTokenShouldInsert(t *testing.T) {
	st := &auth.SecurityToken{
		ID: uuid.New().String(),
		UserID: "some-user-id",
		Token: "some-user-token",
		Type: "some-token-type",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var securityTokenRepo auth.SecurityTokenRepository = &mysqlds.SecurityTokenRepository{ DB: db }

	mock.
		ExpectQuery("SELECT id FROM security_tokens").
		WithArgs(st.UserID, st.Type).
		WillReturnRows(sqlmock.NewRows([]string{ "id" }))

	mock.
		ExpectExec("INSERT security_tokens SET").
		WithArgs(st.ID, st.UserID, st.Token, st.Type, st.CreatedAt, st.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1,1))

	err = securityTokenRepo.CreateOrUpdateToken(st)

	assert.NoError(t, err)
}

func TestCreateOrUpdateTokenShouldUpdate(t *testing.T) {
	st := &auth.SecurityToken{
		ID: uuid.New().String(),
		UserID: "some-user-id",
		Token: "some-user-token",
		Type: "some-token-type",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var securityTokenRepo auth.SecurityTokenRepository = &mysqlds.SecurityTokenRepository{ DB: db }

	rows := sqlmock.NewRows([]string{"id"}).AddRow(st.ID)

	mock.
		ExpectQuery("SELECT id FROM security_tokens").
		WithArgs(st.UserID, st.Type).
		WillReturnRows(rows)

	mock.
		ExpectExec("UPDATE security_tokens SET").
		WithArgs(st.Token, st.UpdatedAt, st.ID).
		WillReturnResult(sqlmock.NewResult(1,1))

	err = securityTokenRepo.CreateOrUpdateToken(st)

	assert.NoError(t, err)
}

func TestCreateOrUpdateTokenShouldThrowError(t *testing.T) {
	st := &auth.SecurityToken{
		ID: uuid.New().String(),
		UserID: "some-user-id",
		Token: "some-user-token",
		Type: "some-token-type",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db.Close()

	var securityTokenRepo auth.SecurityTokenRepository = &mysqlds.SecurityTokenRepository{ DB: db }

	err = securityTokenRepo.CreateOrUpdateToken(st)

	assert.Error(t, err)
}

func TestGetTokenByMetadataShouldReturnToken(t *testing.T) {
	st := &auth.SecurityToken{
		ID: uuid.New().String(),
		UserID: "some-user-id",
		Token: "some-user-token",
		Type: "some-token-type",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tmd := &auth.TokenMetadata{
		UserID: st.UserID,
		Type: st.Type,
		Token: st.Token,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var securityTokenRepo auth.SecurityTokenRepository = &mysqlds.SecurityTokenRepository{ DB: db }

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
}

func TestGetTokenByMetadataShouldThrowError(t *testing.T) {
	st := &auth.SecurityToken{
		ID: uuid.New().String(),
		UserID: "some-user-id",
		Token: "some-user-token",
		Type: "some-token-type",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tmd := &auth.TokenMetadata{
		UserID: "some-other-user-id",
		Type: st.Type,
		Token: st.Token,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var securityTokenRepo auth.SecurityTokenRepository = &mysqlds.SecurityTokenRepository{ DB: db }

	mock.
		ExpectQuery("SELECT id, user_id, token, type, created_at, updated_at FROM security_tokens").
		WithArgs(st.UserID, st.Type).
		WillReturnError(exception.NewNotFoundError("token not found"))

	_, err = securityTokenRepo.GetTokenByMetadata(tmd)

	assert.Error(t, err)
}

func TestRemoveTokenMetadata(t *testing.T) {
	tmd := &auth.TokenMetadata{
		UserID: "some-user-id",
		Type: "some-token-type",
		Token: "some-user-token",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var securityTokenRepo auth.SecurityTokenRepository = &mysqlds.SecurityTokenRepository{ DB: db }

	mock.
		ExpectExec("DELETE FROM security_tokens WHERE").
		WithArgs(tmd.UserID, tmd.Type).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = securityTokenRepo.RemoveTokenByMetadata(tmd)

	assert.NoError(t, err)
}