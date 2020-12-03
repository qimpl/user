package db

import (
	"testing"
	"time"

	"github.com/qimpl/authentication/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	users, err := GetAllUsers()

	assert.NoError(t, err)
	assert.IsType(t, []models.User{}, users)
	assert.EqualValues(t, uuid.MustParse("8a76b8db-0490-4ff4-8bdb-dcd18b582c2f"), users[0].ID)
	assert.EqualValues(t, uuid.MustParse("2e202ea5-d635-4c2f-9bde-2f358ec98e94"), users[1].ID)
}

func TestGetUserByID(t *testing.T) {
	var (
		user *models.User
		err  error
	)

	user, err = GetUserByID(uuid.MustParse("8a76b8db-0490-4ff4-8bdb-dcd18b582c2f"))

	assert.NoError(t, err)
	assert.IsType(t, &models.User{}, user)
	assert.EqualValues(t, uuid.MustParse("8a76b8db-0490-4ff4-8bdb-dcd18b582c2f"), user.ID)
	assert.EqualValues(t, "Jean", user.FirstName)

	// Test get inexistant ID
	_, err = GetUserByID(uuid.MustParse("93fb088d-cdac-4cb9-8bf3-a5652ff25ffc"))
	assert.EqualError(t, err, "pg: no rows in result set")
}

func TestGetPartialUserByID(t *testing.T) {
	var (
		partialUser *models.PartialUser
		err         error
	)

	partialUser, err = GetPartialUserByID(uuid.MustParse("8a76b8db-0490-4ff4-8bdb-dcd18b582c2f"))

	assert.NoError(t, err)
	assert.IsType(t, &models.PartialUser{}, partialUser)
	assert.EqualValues(t, "Jean", partialUser.FirstName)
	assert.EqualValues(t, "Dupond", partialUser.LastName)
	assert.True(t, partialUser.IsVerified)

	// Test user with undefined UserVerifications data
	partialUser, err = GetPartialUserByID(uuid.MustParse("2e202ea5-d635-4c2f-9bde-2f358ec98e94"))

	assert.NoError(t, err)
	assert.IsType(t, &models.PartialUser{}, partialUser)
	assert.False(t, partialUser.IsVerified)

	partialUser, err = GetPartialUserByID(uuid.MustParse("27fd09df-794d-4f33-8a8e-358f42172a78"))

	assert.NoError(t, err)
	assert.IsType(t, &models.PartialUser{}, partialUser)
	assert.EqualValues(t, "Nicolas", partialUser.FirstName)
	assert.EqualValues(t, "Henry", partialUser.LastName)
	assert.False(t, partialUser.IsVerified)

	// Test get inexistant ID
	_, err = GetPartialUserByID(uuid.MustParse("93fb088d-cdac-4cb9-8bf3-a5652ff25ffc"))
	assert.EqualError(t, err, "pg: no rows in result set")
}

func TestCreateUser(t *testing.T) {
	var (
		user *models.User
		err  error
	)

	user, err = CreateUser(&models.User{
		Civility:          "Mrs",
		FirstName:         "Jane",
		LastName:          "Doe",
		Email:             "jane.doe@test.com",
		MobilePhoneNumber: "+33100000000",
		Birthdate:         time.Now(),
		Country:           "FR",
		State:             "59",
		Street:            "1 rue du test",
		City:              "Test sous Lens",
		Zip:               "62000",
	})

	assert.NoError(t, err)
	assert.EqualValues(t, "Jane", user.FirstName)
	assert.EqualValues(t, "Test sous Lens", user.City)
	assert.IsType(t, &models.NotificationPreferences{}, user.NotificationPreferences)
	assert.True(t, user.NotificationPreferences.OnEmail)
	assert.Empty(t, user.Password)

	// Test with missing fields
	user, err = CreateUser(&models.User{
		FirstName: "John",
	})

	assert.EqualError(t, err, "ERROR #23502 null value in column \"last_name\" violates not-null constraint")
	assert.Nil(t, user)
}

func TestUpdateUserByID(t *testing.T) {
	var err error

	err = UpdateUserByID(&models.User{
		ID:                uuid.MustParse("2e202ea5-d635-4c2f-9bde-2f358ec98e94"),
		Civility:          "Mr",
		FirstName:         "John",
		LastName:          "Doe",
		Email:             "john.doe@test.com",
		MobilePhoneNumber: "+33123456789",
		Birthdate:         time.Now(),
		Country:           "FR",
		State:             "59",
		Street:            "1 rue du test",
		City:              "Lille",
		Zip:               "59000",
		Password:          "test",
		NotificationPreferences: &models.NotificationPreferences{
			UserID: uuid.MustParse("2e202ea5-d635-4c2f-9bde-2f358ec98e94"),
		},
	})

	assert.NoError(t, err)

	// Test inexistant user
	err = UpdateUserByID(&models.User{
		ID:                uuid.MustParse("93fb088d-cdac-4cb9-8bf3-a5652ff25ffc"),
		Civility:          "Mr",
		FirstName:         "John",
		LastName:          "Doe",
		Email:             "john.doe@test.com",
		MobilePhoneNumber: "+33123456789",
		Birthdate:         time.Now(),
		Country:           "FR",
		State:             "59",
		Street:            "1 rue du test",
		City:              "Lille",
		Zip:               "59000",
	})

	assert.EqualError(t, err, "pg: no rows in result set")

	// Test missing fields
	err = UpdateUserByID(&models.User{
		ID:       uuid.MustParse("27fd09df-794d-4f33-8a8e-358f42172a78"),
		Civility: "Mr",
	})

	assert.EqualError(t, err, "ERROR #23502 null value in column \"first_name\" violates not-null constraint")
}

func TestDeleteUserByID(t *testing.T) {
	assert.NoError(
		t,
		DeleteUserByID(uuid.MustParse("7c1a2a5f-5deb-4cb1-a020-75357827606b")),
	)

	assert.EqualError(
		t,
		DeleteUserByID(uuid.MustParse("7c1a2a5f-5deb-4cb1-a020-75357827606b")),
		"pg: no rows in result set",
	)
}

func TestResetUserPassword(t *testing.T) {
	assert.NoError(
		t,
		ResetUserPassword(
			uuid.MustParse("27fd09df-794d-4f33-8a8e-358f42172a78"),
			&models.UserResetPasswordBody{
				Password: "newPassword",
			},
		),
	)

	assert.EqualError(
		t,
		ResetUserPassword(
			uuid.MustParse("93fb088d-cdac-4cb9-8bf3-a5652ff25ffc"),
			&models.UserResetPasswordBody{
				Password: "newPassword",
			},
		),
		"pg: no rows in result set",
	)
}

func TestUpdateUserAccountStatus(t *testing.T) {
	assert.NoError(
		t,
		UpdateUserAccountStatus(uuid.MustParse("27fd09df-794d-4f33-8a8e-358f42172a78"), true),
	)

	assert.EqualError(
		t,
		UpdateUserAccountStatus(uuid.MustParse("93fb088d-cdac-4cb9-8bf3-a5652ff25ffc"), true),
		"pg: no rows in result set",
	)
}

func TestCreateOrUpdateUserVerification(t *testing.T) {
	assert.NoError(
		t,
		CreateOrUpdateUserVerification(&models.UserVerifications{
			UserID:           uuid.MustParse("8a76b8db-0490-4ff4-8bdb-dcd18b582c2f"),
			VerificationType: "identity_card",
		}),
	)

	assert.NoError(
		t,
		CreateOrUpdateUserVerification(&models.UserVerifications{
			UserID: uuid.MustParse("8a76b8db-0490-4ff4-8bdb-dcd18b582c2f"),
			Status: "succeeded",
		}),
	)

	assert.EqualError(
		t,
		CreateOrUpdateUserVerification(&models.UserVerifications{
			VerificationType: "identity_card",
		}),
		"ERROR #23502 null value in column \"user_id\" violates not-null constraint",
	)
}
