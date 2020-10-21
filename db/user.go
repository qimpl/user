package db

import (
	"github.com/qimpl/authentication/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// GetUserByID search & return user from a given ID.
func GetUserByID(userID uuid.UUID) (*models.User, error) {
	user := new(models.User)
	if err := Db.Model(user).Where("id = ?", userID).Select(); err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUser hash the password & add an user in database.
func CreateUser(user *models.User) (*models.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if _, err := Db.Model(user).Insert(); err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

// UpdateUserByID search & update an user by his ID.
func UpdateUserByID(user *models.User) (*models.User, error) {
	if _, err := Db.Model(user).Where("id = ?", user.ID).Update(); err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUserByID search & delete an user from a given ID.
func DeleteUserByID(userID uuid.UUID) error {
	user := new(models.User)
	if _, err := Db.Model(user).Where("id = ?", userID).Delete(); err != nil {
		return err
	}
	return nil
}

// ResetUserPassword update user password into database.
func ResetUserPassword(userID uuid.UUID, userResetPassword *models.UserResetPasswordBody) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userResetPassword.Password), bcrypt.DefaultCost)
	user, _ := GetUserByID(userID)

	user.Password = string(hashedPassword)

	if _, err := Db.Model(user).Set("password = ?", user.Password).Where("id = ?", user.ID).Update(); err != nil {
		return err
	}
	return nil
}

// UpdateUserAccountStatus activate or desactivate & update is_enabled field into database
func UpdateUserAccountStatus(userID uuid.UUID, state bool) error {
	user, _ := GetUserByID(userID)

	if _, err := Db.Model(user).Set("is_enabled = ?", state).Where("id = ?", user.ID).Update(); err != nil {
		return err
	}
	return nil
}
