package db

import (
	"github.com/qimpl/authentication/models"

	"golang.org/x/crypto/bcrypt"
)

// Login select user from the database and verify if input password is equal to the database password hash.
func Login(email string, password string) (*models.User, error) {
	var user models.User
	if err := Db.Model(&user).Where("email = ?", email).Select(); err != nil {
		return nil, err
	}

	errLogin := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errLogin != nil {
		return nil, errLogin
	}

	user.Password = ""
	return &user, nil
}
