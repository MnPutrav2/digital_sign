package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	jwtEnc "github.com/MnPutrav2/go_architecture/pkg/auth/jwt"
	"github.com/MnPutrav2/go_architecture/pkg/response"
	"github.com/MnPutrav2/go_architecture/pkg/util"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load()
var jwtKey = []byte(os.Getenv("JWT_SECURE_KEY"))

func Token(r *http.Request) string {

	auth := r.Header.Get("Authorization")
	// Check Header
	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		return ""
	}

	if split[0] != "Bearer" {
		return ""
	}

	return split[1]
}

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		split := strings.SplitN(auth, " ", 2)

		if split[0] != "Bearer" {
			response.Message("token need bearer", "token need bearer", "WARN", 401, w, r)
			return
		}

		claim, err := jwtEnc.ValidateJWT(split[1])
		if err != nil {
			response.Message("unauthorization", err.Error(), "WARN", 401, w, r)
			return
		}

		ctx := context.WithValue(r.Context(), util.UserClaimsKey, claim)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
