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
