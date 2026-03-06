package middleware

import (
	"net/http"

	"github.com/MnPutrav2/go_architecture/pkg/response"
)

func RoleAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx, ok := UserClaims(r.Context())
		if !ok {
			response.Message("unauthorization", "unauthorization", "WARN", 401, w, r)
			return
		}

		if ctx.Role != "admin" {
			response.Message("Access requires administrator privileges.", "Access requires administrator privileges.", "WARN", 403, w, r)
			return
		}

		next.ServeHTTP(w, r)
	}
}
