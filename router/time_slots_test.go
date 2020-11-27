package router

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTimeSlot(t *testing.T) {
	responseRecorder, _ = executeRequest(createTestRequest(
		t,
		"POST",
		"/v1/time-slots",
		[]byte(`{"weekday": "0", "start_time": "14:00:00", "end_time": "18:00:00", "user_id": "27fd09df-794d-4f33-8a8e-358f42172a78"}`),
	))

	checkResponseCode(t, http.StatusCreated, responseRecorder.Code)

	responseRecorder, errorResponse = executeRequest(createTestRequest(
		t,
		"POST",
		"/v1/time-slots",
		[]byte(`{"weekday": "7", "start_time": "14:00:00", "end_time": "18:00:00", "user_id": "27fd09df-794d-4f33-8a8e-358f42172a78"}`),
	))

	checkResponseCode(t, http.StatusBadRequest, responseRecorder.Code)
	assert.EqualValues(t, "An error occurred during time slot creation", errorResponse.Message)

	responseRecorder, errorResponse = executeRequest(createTestRequest(
		t,
		"POST",
		"/v1/time-slots",
		[]byte(`{"weekday": "1", "start_time": "14:00:00", "end_time": "18:00:00", "user_id": ""}`),
	))

	checkResponseCode(t, http.StatusUnprocessableEntity, responseRecorder.Code)
	assert.EqualValues(t, "Malformed body", errorResponse.Message)
}

func TestUpdateTimeSlotByID(t *testing.T) {
	responseRecorder, _ = executeRequest(createTestRequest(
		t,
		"PUT",
		"/v1/time-slots/08760256-f183-4b07-9f28-27b954d6cbcb",
		[]byte(`{"start_time": "14:00:00", "end_time": "18:00:00"}`),
	))

	checkResponseCode(t, http.StatusOK, responseRecorder.Code)

	responseRecorder, errorResponse = executeRequest(createTestRequest(
		t,
		"PUT",
		"/v1/time-slots/08760256-f183-4b07-9f28-27b954d6cbcb",
		[]byte(`{"start_time": "14:00:00"}`),
	))

	checkResponseCode(t, http.StatusBadRequest, responseRecorder.Code)
	assert.EqualValues(t, "An error occurred during time slot update", errorResponse.Message)

	responseRecorder, errorResponse = executeRequest(createTestRequest(
		t,
		"PUT",
		"/v1/time-slots/37d66db1-d834-43f1-8980-ebeec1c66e06",
		[]byte(`{"start_time": "14:00:00", "end_time": "18:00:00"}`),
	))

	checkResponseCode(t, http.StatusNotFound, responseRecorder.Code)
	assert.EqualValues(t, "The given time slot ID doesn't exist", errorResponse.Message)

	responseRecorder, errorResponse = executeRequest(
		createTestRequest(t, "PUT", "/v1/time-slots/08760256-f183-4b07-9f28-27b954d6cbcb", []byte(`{"start_time": 14}`)),
	)

	checkResponseCode(t, http.StatusUnprocessableEntity, responseRecorder.Code)
	assert.EqualValues(t, "Malformed body", errorResponse.Message)
}

func TestGetTimeSlotByID(t *testing.T) {
	responseRecorder, _ = executeRequest(
		createTestRequest(t, "GET", "/v1/time-slots/user/8a76b8db-0490-4ff4-8bdb-dcd18b582c2f", nil),
	)

	checkResponseCode(t, http.StatusOK, responseRecorder.Code)

	responseRecorder, errorResponse = executeRequest(
		createTestRequest(t, "GET", "/v1/time-slots/user/37d66db1-d834-43f1-8980-ebeec1c66e06", nil),
	)

	checkResponseCode(t, http.StatusNotFound, responseRecorder.Code)
	assert.EqualValues(t, "The given user ID doesn't exist", errorResponse.Message)
}
