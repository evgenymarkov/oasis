package api

import (
	"net/http"

	"github.com/evgenymarkov/oasis/openapi3"
)

var GetGreetingOperation = openapi3.NewOperation("GetGreeting").
	SetSummary("Get a greeting").
	AddParameter(
		openapi3.NewParameter("name", "path").
			SetDescription("Name of the person to greet"),
	).
	AddParameter(
		openapi3.NewParameter("verbose", "query").
			SetDescription("Show verbose greeting"),
	).
	SetTags(greetingsTag)

func GetGreetingHandler(response http.ResponseWriter, request *http.Request) {
	name := request.PathValue("name")
	verbose := request.URL.Query().Get("verbose")

	greeting := "Hello, " + name + "!"

	if verbose == "true" {
		greeting += " Have a great day!"
	}

	_, _ = response.Write([]byte(greeting))
}
