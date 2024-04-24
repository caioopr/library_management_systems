package middlewares

import (
	"api/src/auth"
	"api/src/responses"
	"net/http"
)

// Authorize verifies if the user is authenticated
func Authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
