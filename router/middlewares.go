package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/qimpl/authentication/models"
	"github.com/qimpl/authentication/services"

	"github.com/gorilla/mux"
)

func jwtTokenVerificationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, route := range unProtectedRoutes {
			uriParams := regexp.MustCompile("\\{(.*?)\\}").FindAllStringSubmatch(route, -1)

			if len(uriParams) != 0 {
				for _, uriParam := range uriParams {
					route = strings.Replace(route, uriParam[0], mux.Vars(r)[uriParam[1]], 1)
				}
			}

			if strings.Contains(strings.TrimPrefix(r.RequestURI, "/api/v1"), route) {
				next.ServeHTTP(w, r)

				return
			}
		}

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer")
		if len(authHeader) != 2 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			var badRequest *models.BadRequest
			json.NewEncoder(w).Encode(badRequest.GetError("Malformed Authorization HTTP header"))

			return
		}

		if _, err := services.ValidateJwtToken(authHeader[1]); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			var Unauthorized *models.Unauthorized
			json.NewEncoder(w).Encode(Unauthorized.GetError(fmt.Sprintf("Invalid jwt token: %s", err)))

			return
		}

		next.ServeHTTP(w, r)
	})
}
