package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/qimpl/authentication/models"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var (
	router           *mux.Router
	responseRecorder *httptest.ResponseRecorder
	errorResponse    *models.ErrorResponse
)

func init() {
	router = mux.NewRouter()
	apiRouter := router.PathPrefix("/v1").Subrouter()

	if os.Getenv("ENV") != "dev" {
		router.Use(jwtTokenVerificationMiddleware)
	}

	createAuthenticationRouter(apiRouter)
	createUserRouter(apiRouter)
	createTimeSlotsRouter(apiRouter)
}

func createTestRequest(t *testing.T, method, url string, payload []byte) *http.Request {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		t.Errorf("Creating '%s %s' request failed!", method, url)
	}

	return req
}

func executeRequest(req *http.Request) (*httptest.ResponseRecorder, *models.ErrorResponse) {
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	var errorResponse *models.ErrorResponse

	if err := json.Unmarshal(recorder.Body.Bytes(), &errorResponse); err != nil {
		return recorder, nil
	}

	return recorder, errorResponse
}

func checkResponseCode(t *testing.T, expected int, actual int) {
	assert.Equal(
		t,
		expected,
		actual,
		fmt.Sprintf("Expected response code %d. Got %d", expected, actual),
	)
}
