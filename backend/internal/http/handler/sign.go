package handler

import (
	"context"
	"crypto/rsa"
	"fmt"
	"net/http"
	"time"

	signService "github.com/MnPutrav2/go_architecture/internal/service/sign"
	"github.com/MnPutrav2/go_architecture/pkg/response"
)

// Entry
func Sign(serv signService.SignService, publickey *rsa.PublicKey, privateKey *rsa.PrivateKey) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, close := context.WithTimeout(r.Context(), time.Second*5)
		defer close()

		x := r.FormValue("x")
		y := r.FormValue("y")
		file, _, err := r.FormFile("file")
		if err != nil {
			response.Message("failed decode body", err.Error(), "WARN", http.StatusBadRequest, w, r)
			return
		}

		buf, err := serv.Signer(ctx, file, x, y, privateKey, publickey)
		if err != nil {
			response.Message("error", err.Error(), "WARN", http.StatusBadRequest, w, r)
			return
		}

		fmt.Println("size: ", buf.Len())
		response.File(buf, "success", "INFO", w, r)
	}
}
