package main

import (
	"net/http"

	"github.com/evgenymarkov/oasis"
	"github.com/evgenymarkov/oasis/openapi3"
)

const (
	serverHost = "localhost"
	serverPort = "3000"
	serverAddr = serverHost + ":" + serverPort
)

func main() {
	// Create multiplexer
	mux := http.NewServeMux()

	// Create API wrapper
	api := oasis.NewAPI(
		mux,
		oasis.NewAPIConfig().
			SetDocumentPath("/api/openapi.json").
			SetSwaggerUIPath("/api").
			SetSwaggerUITitle("API Docs"),
		openapi3.NewDocument().
			SetTitle("Greeting API").
			SetVersion("1.0.0"),
	)

	// Register operations
	api.Get(
		"/greeting/{name}",
		GetGreetingHandler,
		GetGreetingOperation,
	)

	// Start handling incoming requests
	if startErr := http.ListenAndServe(serverAddr, mux); startErr != nil {
		panic(startErr)
	}
}
