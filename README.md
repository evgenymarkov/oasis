# Oasis

A simple library for building HTTP REST/RPC APIs in Go backed by OpenAPI 3 and JSON Schema.

## Features

-   Step-by-step migration from raw handlers to documented ones
-   Swagger UI rendering using generated OpenAPI schema
-   Declarative interface for API operations:
    -   Operation & models documentation
    -   Request params (path, query, or header)
    -   Request body
    -   Responses (including errors)
    -   Response headers
-   Annotated Go types for input and output models
    -   Generates JSON Schema from Go types
    -   Automatic input model validation & error handling
    -   Static typing for path/query/header params, bodies, response headers, etc

## Example

Here is a basic hello world example with Oasis:

```go
package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/evgenymarkov/oasis"
	"github.com/go-chi/chi/v5"
)

const (
	host = "localhost"
	port = 3000
)

// GreetingInput represents the greeting operation request.
type GreetingInput struct {
	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
}

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func main() {
	// Create a new mux
	mux := chi.NewRouter()

	// Create a new API
	api := oasis.NewAPI(mux, oasis.APIConfig{
		Title:   "My API",
		Version: "1.0.0",
	})

	// Register GET /greeting/{name}
	api.RegisterOperation(
		oasis.Operation{
			ID: 	 "get-greeting",
			Method:  http.MethodGet,
			Path:    "/greeting/{name}",
			Summary: "Get a greeting",
		},
		func(ctx oasis.Context, input *GreetingInput) (*GreetingOutput, error) {
			response := &GreetingOutput{}
			response.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)

			return response, nil
		},
	)

	// Start handling incoming requests
	http.ListenAndServe(host+":"+port, mux)
}
```
