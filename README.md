# Oasis

A simple library for building HTTP REST/RPC APIs in Go backed by OpenAPI.

## Features

-   OpenAPI definitions for your endpoints in Go code
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
```

```go
package main

import (
	"net/http"

	"github.com/evgenymarkov/oasis/openapi3"
)

var GetGreetingOperation = openapi3.NewOperation().
	SetOperationID("GetGreeting").
	SetSummary("Get a greeting")

func GetGreetingHandler(response http.ResponseWriter, _ *http.Request) {
	_, _ = response.Write([]byte("Hello, world!"))
}
```

## Development

1. [Install Go 1.22](https://github.com/go-nv/goenv) to run tests.
2. [Install Taskfile](https://taskfile.dev/installation) to run tasks.
3. [Install golangci-lint](https://golangci-lint.run/usage/install) to lint code.

Now you can use any commands from the `Taskfile.yaml` and edit code on your machine.
