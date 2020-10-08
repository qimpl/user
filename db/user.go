package db

import (
	"github.com/qimpl/authentication/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// GetUserByID search & return user from a given ID.
func GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user = &models.User{ID: userID}
	if err := Db.Select(user); err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUser hash the password & add an user in database.
func CreateUser(user *models.User) (*models.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := Db.Insert(user); err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

// UpdateUser search & update an user.
func UpdateUser(user *models.User) (*models.User, error) {
	if err := Db.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUserByID search & delete an user from a given ID.
func DeleteUserByID(userID uuid.UUID) error {
	return Db.Delete(&models.User{ID: userID})
}
