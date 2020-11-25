package db

import (
	"testing"

	"github.com/qimpl/authentication/models"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	var (
		user *models.User
		err  error
	)

	user, err = Login("jean.dupond@qimpl.io", "jean")

	assert.NoError(t, err)
	assert.EqualValues(t, "Jean", user.FirstName)
	assert.Empty(t, user.Password)

	user, err = Login("jean.dupond@qimpl.io", "wrongpassword")

	assert.Error(t, err)
	assert.EqualError(t, err, "crypto/bcrypt: hashedPassword is not the hash of the given password")

	// Test inexistant user
	user, err = Login("wrong@qimpl.io", "stillWrong")

	assert.Error(t, err)
	assert.EqualError(t, err, "pg: no rows in result set")
}
