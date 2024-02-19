# Oasis

A simple library for building HTTP REST/RPC APIs in Go backed by OpenAPI 3 and JSON Schema.

## Features

- Declarative Operation interface on top of net/http router:
  - Operation & models documentation
  - Request params (path, query, or header)
  - Request body
  - Responses (including errors)
  - Response headers
- Annotated Go types for input and output models
  - Generates JSON Schema from Go types
  - Automatic input model validation & error handling
  - Static typing for path/query/header params, bodies, response headers, etc
- Swagger UI rendering using generated OpenAPI schema
- Support for middlewares compatible with net/http
- Content negotiation between server and client
  - Support for [JSON](https://tools.ietf.org/html/rfc8259)
  - Support for [CBOR](https://datatracker.ietf.org/doc/html/rfc7049)
  - Support for [MsgPack](https://msgpack.org)

## Example

Here is a basic hello world example with Oasis:

```go
package main

import (
	"context"
	"fmt"
  "net/http"

	"github.com/evgenymarkov/oasis"
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
  // Create a new API
  api := oasis.NewAPI(oasis.NewAPIConfig("My API", "1.0.0"))

  // Register GET /greeting/{name}
  api.Register(oasis.Operation{
    OperationID: "get-greeting",
    Summary:     "Get a greeting",
    Method:      http.MethodGet,
    Path:        "/greeting/{name}",
  }, func(ctx context.Context, input *GreetingInput) (*GreetingOutput, error) {
    response := &GreetingOutput{}
    response.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)

    return response, nil
  })

  // Create a new server
  server := oasis.NewServer(oasis.NewServerConfig("localhost", 3000))

  // Start handling incoming requests
  listenErr := server.ListenAndServe(api, func() {
    fmt.Printf("server listening on %s:%s\n", host, port)
  })
  if listenErr != nil {
    fmt.Printf("failed to start server due to error: %s\n", listenErr)
  }
}
```
