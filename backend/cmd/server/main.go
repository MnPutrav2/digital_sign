package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"net/http"
	"os"

	"github.com/MnPutrav2/go_architecture/config"
	"github.com/MnPutrav2/go_architecture/internal/http/handler"
	passphraseModel "github.com/MnPutrav2/go_architecture/internal/model/passphrase"
	signModel "github.com/MnPutrav2/go_architecture/internal/model/sign"
	signRepository "github.com/MnPutrav2/go_architecture/internal/repository/sign"
	verifRepository "github.com/MnPutrav2/go_architecture/internal/repository/verif"
	signService "github.com/MnPutrav2/go_architecture/internal/service/sign"
	verifService "github.com/MnPutrav2/go_architecture/internal/service/verif"
	"github.com/MnPutrav2/go_architecture/pkg/middleware"
	"github.com/MnPutrav2/go_architecture/pkg/query"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	mux := http.NewServeMux()

	// <----- Entry ----->

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey

	signRepo := signRepository.InitSignRepository(db)
	signServ := signService.InitSignService(signRepo)
	mux.HandleFunc("POST /sign", middleware.Chain(handler.Sign(signServ, publicKey, privateKey), middleware.CORS))

	verifRepo := verifRepository.InitVerifRepository(db)
	verifServ := verifService.InitVerifService(verifRepo)
	mux.HandleFunc("POST /verif", middleware.Chain(handler.Verif(verifServ, publicKey)))
	// <----- Last ----->

	query.InitDB(db).Migrate(
		signModel.Sign{},
		passphraseModel.Passphrase{},
	)

	listen := os.Getenv("LISTEN_PROD")
	srv := &http.Server{
		Addr:    listen,
		Handler: mux,
	}

	fmt.Println("Server listen in port " + listen)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
