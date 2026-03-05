package errs

import (
	"errors"
	"net/http"
)

func Map(err error) (string, string, int) {
	switch {
	case errors.Is(err, IsErrorDuplicate):
		return "duplicate", "WARN", http.StatusBadRequest

	default:
		return "failed", "ERROR", http.StatusBadRequest
	}
}
