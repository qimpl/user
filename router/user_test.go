package router

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	responseRecorder, _ = executeRequest(createTestRequest(t, "GET", "/v1/user", nil))
	checkResponseCode(t, http.StatusOK, responseRecorder.Code)
}

func TestGetUserByID(t *testing.T) {
	responseRecorder, _ = executeRequest(
		createTestRequest(t, "GET", "/v1/user/8a76b8db-0490-4ff4-8bdb-dcd18b582c2f", nil),
	)

	checkResponseCode(t, http.StatusOK, responseRecorder.Code)

	responseRecorder, errorResponse = executeRequest(
		createTestRequest(t, "GET", "/v1/user/37d66db1-d834-43f1-8980-ebeec1c66e06", nil),
	)

	checkResponseCode(t, http.StatusBadRequest, responseRecorder.Code)
	assert.EqualValues(t, "An error occurred during user retrieval", errorResponse.Message)
}

func TestGetPartialUserByID(t *testing.T) {
	responseRecorder, _ = executeRequest(
		createTestRequest(t, "GET", "/v1/user/8a76b8db-0490-4ff4-8bdb-dcd18b582c2f/partial", nil),
	)

	checkResponseCode(t, http.StatusOK, responseRecorder.Code)

	responseRecorder, errorResponse = executeRequest(
		createTestRequest(t, "GET", "/v1/user/37d66db1-d834-43f1-8980-ebeec1c66e06/partial", nil),
	)

	checkResponseCode(t, http.StatusBadRequest, responseRecorder.Code)
	assert.EqualValues(t, "An error occurred during partial user retrieval", errorResponse.Message)
}

func TestRegister(t *testing.T) {
	responseRecorder, _ = executeRequest(createTestRequest(
		t,
		"POST",
		"/v1/user/register",
		[]byte(`
			{
				"civility": "Mrs",
				"first_name": "Register",
				"last_name": "Test",
				"email": "register@test.com",
				"mobile_phone_number": "+33100000000",
				"birthdate": "1999-10-05T00:00:00Z",
				"country": "FR",
				"state": "59",
				"street": "1 rue du test",
				"city": "Test sous Lens",
				"zip": "62000"
			}
		`),
	))

	checkResponseCode(t, http.StatusCreated, responseRecorder.Code)

	// Try creating same user
	responseRecorder, errorResponse = executeRequest(createTestRequest(
		t,
		"POST",
		"/v1/user/register",
		[]byte(`
			{
				"civility": "Mrs",
				"first_name": "Register",
				"last_name": "Test",
				"email": "register@test.com",
				"mobile_phone_number": "+33100000000",
				"birthdate": "1999-10-05T00:00:00Z",
				"country": "FR",
				"state": "59",
				"street": "1 rue du test",
				"city": "Test sous Lens",
				"zip": "62000"
			}
		`),
	))

	checkResponseCode(t, http.StatusBadRequest, responseRecorder.Code)
	assert.EqualValues(t, "An error occurred during user creation", errorResponse.Message)

	responseRecorder, errorResponse = executeRequest(
		createTestRequest(t, "POST", "/v1/user/register", []byte(`{"zipp": "62000"}`)),
	)

	checkResponseCode(t, http.StatusBadRequest, responseRecorder.Code)
	assert.EqualValues(t, "An error occurred during user creation", errorResponse.Message)
}

func TestAnonymizeUserByID(t *testing.T) {
	responseRecorder, _ = executeRequest(
		createTestRequest(t, "PUT", "/v1/user/7c1a2a5f-5deb-4cb1-a020-75357827606b/anonymize", nil),
	)

	checkResponseCode(t, http.StatusNoContent, responseRecorder.Code)

	responseRecorder, errorResponse = executeRequest(
		createTestRequest(t, "PUT", "/v1/user/37d66db1-d834-43f1-8980-ebeec1c66e06/anonymize", nil),
	)

	checkResponseCode(t, http.StatusNotFound, responseRecorder.Code)
	assert.EqualValues(t, "User does not exist", errorResponse.Message)
}

func TestResetPassword(t *testing.T) {
	responseRecorder, _ = executeRequest(createTestRequest(
		t,
		"PUT",
		"/v1/user/7c1a2a5f-5deb-4cb1-a020-75357827606b/reset/password",
		[]byte(`{"password": "newPassword"}`),
	))

	checkResponseCode(t, http.StatusNoContent, responseRecorder.Code)

	responseRecorder, errorResponse = executeRequest(createTestRequest(
		t,
		"PUT",
		"/v1/user/7c1a2a5f-5deb-4cb1-a020-75357827606b/reset/password",
		[]byte(`{"password": 1}`),
	))

	checkResponseCode(t, http.StatusUnprocessableEntity, responseRecorder.Code)
	assert.EqualValues(t, "Malformed body", errorResponse.Message)

	responseRecorder, errorResponse = executeRequest(createTestRequest(
		t,
		"PUT",
		"/v1/user/37d66db1-d834-43f1-8980-ebeec1c66e06/reset/password",
		[]byte(`{"password": "newPassword"}`),
	))

	checkResponseCode(t, http.StatusBadRequest, responseRecorder.Code)
	assert.EqualValues(t, "An error occurred during user password reseting", errorResponse.Message)
}

func TestValidateUserAccount(t *testing.T) {
	responseRecorder, _ = executeRequest(
		createTestRequest(t, "PUT", "/v1/user/7c1a2a5f-5deb-4cb1-a020-75357827606b/validate", nil),
	)
	checkResponseCode(t, http.StatusNoContent, responseRecorder.Code)

	responseRecorder, errorResponse = executeRequest(
		createTestRequest(t, "PUT", "/v1/user/37d66db1-d834-43f1-8980-ebeec1c66e06/validate", nil),
	)

	checkResponseCode(t, http.StatusBadRequest, responseRecorder.Code)
	assert.EqualValues(t, "An error occurred during user account activation", errorResponse.Message)
}

func TestDeactivateUserAccount(t *testing.T) {
	responseRecorder, _ = executeRequest(
		createTestRequest(t, "PUT", "/v1/user/7c1a2a5f-5deb-4cb1-a020-75357827606b/deactivate", nil),
	)

	checkResponseCode(t, http.StatusNoContent, responseRecorder.Code)

	responseRecorder, errorResponse = executeRequest(
		createTestRequest(t, "PUT", "/v1/user/37d66db1-d834-43f1-8980-ebeec1c66e06/deactivate", nil),
	)

	checkResponseCode(t, http.StatusBadRequest, responseRecorder.Code)
	assert.EqualValues(t, "An error occurred during user account deactivation", errorResponse.Message)
}
func TestStripeVerificationIntent(t *testing.T) {
	responseRecorder, errorResponse = executeRequest(
		createTestRequest(t, "POST", "/v1/user/7c1a2a5f-5deb-4cb1-a020-75357827606b/identity/verification_intent", nil),
	)

	checkResponseCode(t, http.StatusCreated, responseRecorder.Code)

	responseRecorder, errorResponse = executeRequest(
		createTestRequest(t, "POST", "/v1/user/37d66db1-d834-43f1-8980-ebeec1c66e06/identity/verification_intent", nil),
	)

	checkResponseCode(t, http.StatusNotFound, responseRecorder.Code)
	assert.EqualValues(t, "User does not exist", errorResponse.Message)
}

func TestDeleteUserByID(t *testing.T) {
	responseRecorder, _ = executeRequest(createTestRequest(t, "DELETE", "/v1/user/7c1a2a5f-5deb-4cb1-a020-75357827606b", nil))

	checkResponseCode(t, http.StatusNoContent, responseRecorder.Code)

	responseRecorder, errorResponse = executeRequest(createTestRequest(t, "DELETE", "/v1/user/37d66db1-d834-43f1-8980-ebeec1c66e06", nil))

	checkResponseCode(t, http.StatusBadRequest, responseRecorder.Code)
	assert.EqualValues(t, "An error occurred during user deletion", errorResponse.Message)
}
