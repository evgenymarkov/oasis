package api

import (
	"net/http"

	"github.com/evgenymarkov/oasis/openapi3"
)

var GetGreetingOperation = openapi3.NewOperation("GetGreeting").
	SetSummary("Get a greeting").
	SetTags(greetingsTag)

func GetGreetingHandler(response http.ResponseWriter, _ *http.Request) {
	_, _ = response.Write([]byte("Hello, world!"))
}
