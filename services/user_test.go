package services

import (
	"testing"
	"time"

	"github.com/qimpl/authentication/models"

	"github.com/stretchr/testify/assert"
)

func TestAnonymizeUser(t *testing.T) {
	anonimyzedUser := AnonymizeUser(&models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@test.com",
		Street:    "45 rue du test",
		Zip:       "59000",
		City:      "Lille",
		Birthdate: time.Date(2020, 10, 06, 0, 0, 0, 0, time.UTC),
	})

	assert.NotEqualValues(t, "John", anonimyzedUser.FirstName)
	assert.NotEqualValues(t, "Doe", anonimyzedUser.LastName)
	assert.NotEqualValues(t, "john@test.com", anonimyzedUser.Email)
	assert.NotEqualValues(t, "45 rue du test", anonimyzedUser.Street)
	assert.NotEqualValues(t, "59000", anonimyzedUser.Zip)
	assert.NotEqualValues(t, "Lille", anonimyzedUser.City)
	assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", anonimyzedUser.Birthdate.String())
}

func TestAnonymizeField(t *testing.T) {
	assert.NotEqualValues(t, "test", anonymizeField("test"))
}

func TestAnonymizeDatetime(t *testing.T) {
	assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", anonymizeDatetime(time.Now()).String())
}
