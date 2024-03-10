package main

import (
	"net/http"

	"github.com/evgenymarkov/oasis/openapi3"
)

var GetGreetingOperation = openapi3.NewOperation().
	SetOperationID("get-greeting").
	SetSummary("Get a greeting")

func GetGreetingHandler(response http.ResponseWriter, _ *http.Request) {
	_, _ = response.Write([]byte("Hello, world!"))
}
