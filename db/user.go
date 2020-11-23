package db

import (
	"github.com/qimpl/authentication/models"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// GetAllUsers get all users from the database
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := Db.Model(&users).Order("created_at").Select(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByID search & return user from a given ID.
func GetUserByID(userID uuid.UUID) (*models.User, error) {
	user := new(models.User)
	if err := Db.Model(user).Relation("NotificationPreferences").Where("? = ?", pg.Ident("user.id"), userID).First(); err != nil {
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

	notificationPreferences := new(models.NotificationPreferences)
	notificationPreferences.UserID = user.ID
	notificationPreferences.OnEmail = true
	if _, err := Db.Model(notificationPreferences).Returning("*").Insert(); err != nil {
		return nil, err
	}

	user.NotificationPreferences = notificationPreferences
	user.Password = ""
	return user, nil
}

// UpdateUserByID search & update an user by his ID.
func UpdateUserByID(user *models.User) error {
	_, err := GetUserByID(user.ID)
	if err != nil {
		return err
	}

	_, err = Db.Model(user).Where("id = ?", user.ID).Update()

	if err != nil {
		return err
	}

	notificationPreferences := new(models.NotificationPreferences)
	notificationPreferences = user.NotificationPreferences
	_, err = Db.Model(notificationPreferences).Where("user_id = ?", user.ID).Update()

	return err
}

// DeleteUserByID search & delete an user from a given ID.
func DeleteUserByID(userID uuid.UUID) error {
	user, err := GetUserByID(userID)
	if err != nil {
		return err
	}

	_, err = Db.Model(user).Where("id = ?", userID).Delete()
	return err
}

// ResetUserPassword update user password into database.
func ResetUserPassword(userID uuid.UUID, userResetPassword *models.UserResetPasswordBody) error {
	user, err := GetUserByID(userID)
	if err != nil {
		return err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userResetPassword.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	_, err = Db.Model(user).Set("password = ?", user.Password).Where("id = ?", user.ID).Update()

	return err
}

// UpdateUserAccountStatus activate or desactivate & update is_enabled field into database
func UpdateUserAccountStatus(userID uuid.UUID, state bool) error {
	user, err := GetUserByID(userID)
	if err != nil {
		return err
	}

	_, err = Db.Model(user).Set("is_enabled = ?", state).Where("id = ?", user.ID).Update()

	return err
}

// CreateUserVerification create a new entry into user_verification table for a given user.
func CreateUserVerification(userVerification *models.UserVerification) error {
	if _, err := Db.Model(userVerification).Insert(); err != nil {
		return err
	}
	return nil
}
