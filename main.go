package main

import (
	"fmt"
	"net/http"

	"github.com/MnPutrav2/digital_sign/route"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("POST /sign", route.Sign())

	fmt.Println("server running")
	http.ListenAndServe(":8080", mux)
}
