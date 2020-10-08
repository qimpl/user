package handlers

import (
	"net/http"
)

// HealthyCheck is a public method that allow users to check if the current API is alive or not.
// @Summary Healthy Check
// @Description Check if the current API is alive
// @Success 200 {string} string
// @Failure 404 {string} string
// @Router /healthy [get]
func HealthyCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
