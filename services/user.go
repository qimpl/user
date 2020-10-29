package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/qimpl/authentication/models"
)

// AnonymizeUser is used to anonymize user account data.
func AnonymizeUser(user *models.User) *models.User {
	user.FirstName = anonymizeField(user.FirstName)
	user.LastName = anonymizeField(user.LastName)
	user.Email = anonymizeField(user.Email)
	user.Street = anonymizeField(user.Street)
	user.Zip = anonymizeField(user.Zip)
	user.City = anonymizeField(user.City)
	user.Birthdate = anonymizeDatetime(user.Birthdate)

	return user
}

func anonymizeField(field string) string {
	var res = ""
	for range field {
		v := rand.Intn(126-33) + 33
		newChar := string(rune(v))
		res = fmt.Sprintf("%s%s", res, newChar)
	}
	return res
}

func anonymizeDatetime(date time.Time) time.Time {
	year, _, _ := date.Date()
	return time.Date(year, time.January, 01, 0, 0, 0, 0, time.UTC)
}
