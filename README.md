# Oasis

A simple library for building HTTP REST/RPC APIs in Go backed by OpenAPI.

## Features

-	OpenAPI definitions for your endpoints in Go code
-   Swagger UI rendering using compiled OpenAPI document
-   Step-by-step migration from raw handlers to documented ones

## Example

Here is a basic hello world example with Oasis:

```go
package main

import (
	"net/http"

	"github.com/evgenymarkov/oasis"
	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/go-chi/chi/v5"
)

const (
	host = "localhost"
	port = 3000
)

func main() {
	// Create router
	router := chi.NewRouter()

	// Create API wrapper
	api := oasis.NewAPI(
		router,
		oasis.NewAPIConfig().
			SetDocsUIPath("/api").
			SetJSONDocumentPath("/api/openapi.json").
			SetYAMLDocumentPath("/api/openapi.yaml"),
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
	http.ListenAndServe(host+":"+port, router)
}
```

```go
package main

import (
	"net/http"

	"github.com/evgenymarkov/oasis"
	"github.com/evgenymarkov/oasis/openapi3"
)

var GetGreetingOperation = openapi3.NewOperation().
	SetOperationID("get-greeting").
	SetSummary("Get a greeting")

func GetGreetingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
```
