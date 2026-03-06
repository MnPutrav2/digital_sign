package handler

import (
	"context"
	"crypto/rsa"
	"fmt"
	"net/http"
	"time"

	verifService "github.com/MnPutrav2/go_architecture/internal/service/verif"
	"github.com/MnPutrav2/go_architecture/pkg/response"
)

// Entry
func Verif(service verifService.VerifService, public *rsa.PublicKey) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, close := context.WithTimeout(r.Context(), time.Second*5)
		defer close()

		sign := r.FormValue("sign")
		file, header, err := r.FormFile("file")
		fmt.Println("size: ", header.Size)
		if err != nil {
			response.Message("failed decode body", err.Error(), "WARN", http.StatusBadRequest, w, r)
			return
		}

		if err := service.Verif(ctx, file, public, sign); err != nil {
			response.Message("failed", err.Error(), "WARN", http.StatusBadRequest, w, r)
			return
		}

		response.Message("valid", "valid", "INFO", http.StatusOK, w, r)
	}
}
