package middleware

import (
	"net/http"

	"github.com/herlianali/goCommerce/pkg/response"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			response.Error(w, 401, "unauthorized")
			return
		}
		next.ServeHTTP(w, r)
	})
}
