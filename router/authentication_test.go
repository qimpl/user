package router

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticate(t *testing.T) {
	responseRecorder, _ = executeRequest(
		createTestRequest(t, "POST", "/v1/authenticate", []byte(`{"email": "jean.dupond@qimpl.io", "password": "jean"}`)),
	)

	checkResponseCode(t, http.StatusOK, responseRecorder.Code)

	responseRecorder, errorResponse = executeRequest(createTestRequest(
		t,
		"POST",
		"/v1/authenticate",
		[]byte(`{"email": "jean.dupond@qimpl.io", "password": "wrongPassword"}`),
	))

	checkResponseCode(t, http.StatusBadRequest, responseRecorder.Code)
	assert.EqualValues(t, "An error occurred during user authentication", errorResponse.Message)

	responseRecorder, errorResponse = executeRequest(
		createTestRequest(t, "POST", "/v1/authenticate", []byte(`{"email": "jean.dupond@qimpl.io", "password": 1}`)),
	)

	checkResponseCode(t, http.StatusUnprocessableEntity, responseRecorder.Code)
	assert.EqualValues(t, "Malformed body", errorResponse.Message)
}
