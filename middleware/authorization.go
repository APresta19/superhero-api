package middleware

import(
	"errors"
	"superheroapi/api_info"
	"net/http"
)

var UnauthorizedError = errors.New("Invalid ID or token")

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		expectedToken := "ABC123"

		if token != expectedToken {
			api_info.SpecificError(w, UnauthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}