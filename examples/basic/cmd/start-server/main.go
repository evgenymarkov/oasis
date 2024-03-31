package main

import (
	"net/http"

	"github.com/evgenymarkov/oasis/examples/basic/internal/api"
)

const (
	serverHost = "localhost"
	serverPort = "3000"
	serverAddr = serverHost + ":" + serverPort
)

func main() {
	mux := http.NewServeMux()
	api.NewAPI(mux)

	if startErr := http.ListenAndServe(serverAddr, mux); startErr != nil {
		panic(startErr)
	}
}
