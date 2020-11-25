package services

import (
	"testing"

	"github.com/qimpl/authentication/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateJwtToken(t *testing.T) {
	token := CreateJwtToken(&models.User{
		ID:        uuid.New(),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@test.com",
		IsOwner:   true,
		IsAdmin:   false,
	})

	assert.IsType(t, &models.TokenHash{}, token)
	assert.IsType(t, new(string), &token.Token)
	assert.Regexp(t, "^[\\w-_]{36}.[\\w-_]{234}.[\\w-_]{43}$", token.Token)
}

func TestLogin(t *testing.T) {
	var (
		token *models.TokenHash
		err   error
	)

	token, err = Login("jean.dupond@qimpl.io", "jean")

	assert.NoError(t, err)
	assert.IsType(t, &models.TokenHash{}, token)
	assert.IsType(t, new(string), &token.Token)
	assert.Regexp(t, "^[\\w-_]{36}.[\\w-_]+.[\\w-_]{43}$", token.Token)

	token, err = Login("jean.dupond@qimpl.io", "wrongPassword")

	assert.Error(t, err)
	assert.Nil(t, token)
}

func TestValidateJwtToken(t *testing.T) {
	var (
		isValid bool
		err     error
	)

	isValid, err = ValidateJwtToken("notAToken")

	assert.False(t, isValid)
	assert.EqualError(t, err, "token contains an invalid number of segments")

	isValid, err = ValidateJwtToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjRiODRlMmZhLTZhOTAtNDQ2MC1iMzY1LTYwMDJhODM0NDdiMCIsImZpcnN0X25hbWUiOiJKb2huIiwibGFzdF9uYW1lIjoiRG9lIiwiZW1haWwiOiJqb2huQHRlc3QuY29tIiwiaXNfb3duZXIiOnRydWUsImlzX2FkbWluIjpmYWxzZSwiZXhwIjoyMDIwMTEyNTExNTIwNCwiaXNzIjoiUWltcGwifQ.xM081W2dOhZ9q9hD7ghHqFZNT2CMMc0cBvRNOFiAo6w")

	assert.True(t, isValid)
	assert.NoError(t, err)

	isValid, err = ValidateJwtToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")

	assert.False(t, isValid)
	assert.EqualError(t, err, "signature is invalid")
}
