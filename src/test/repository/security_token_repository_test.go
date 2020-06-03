package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"sherman/src/domain/auth"
	"sherman/src/repository/mysqlds"
	"testing"
	"time"
)

func TestCreateOrUpdateToken(t *testing.T) {
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

	mock.
		ExpectQuery("SELECT id FROM security_tokens").
		WithArgs(st.UserID, st.Type).
		WillReturnRows(sqlmock.NewRows([]string{ "id" }))

	mock.
		ExpectExec("INSERT security_tokens SET").
		WithArgs(st.ID, st.UserID, st.Token, st.Type, st.CreatedAt, st.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1,1))

	var securityTokenRepository auth.SecurityTokenRepository = &mysqlds.SecurityTokenRepository{ DB: db }
	err = securityTokenRepository.CreateOrUpdateToken(st)

	

	assert.NoError(t, err)
}