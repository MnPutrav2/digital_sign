package middleware

import (
	"net/http"
	"strings"

	"github.com/MnPutrav2/go_architecture/pkg/response"
)

func CTJson(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ct := r.Header.Get("Content-Type")
		if !strings.HasPrefix(ct, "application/json") {
			response.Message("invalid content type, must be application/json", "invalid content type, must be application/json", "WARN", 400, w, r)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func CTFormData(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ct := r.Header.Get("Content-Type")
		if !strings.HasPrefix(ct, "multipart/form-data") {
			response.Message("invalid content type, must be multipart/form-data", "invalid content type, must be multipart/form-data", "WARN", 400, w, r)
			return
		}

		next.ServeHTTP(w, r)
	}
}
