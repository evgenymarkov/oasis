package api

import (
	"net/http"

	"github.com/evgenymarkov/oasis/openapi3"
)

var GetGreetingOperation = openapi3.NewOperation("GetGreeting").
	SetSummary("Get a greeting").
	AddParameter(
		openapi3.NewParameter("name", "path").
			SetDescription("Name of the person to greet").
			AddExample(
				"Russian revolutionist",
				openapi3.NewExample("Anatoly Lunacharsky").
					SetSummary("Russian Marxist revolutionary").
					SetDescription("Russian Marxist revolutionary and the first Bolshevik Soviet People's Commissar"),
			).
			AddExample(
				"Russian poet",
				openapi3.NewExample("Alexander Pushkin").
					SetSummary("Russian poet, playwright, and novelist").
					SetDescription("Alexander Pushkin was a Russian poet, playwright, and novelist in the early 19th century."),
			),
	).
	AddParameter(
		openapi3.NewParameter("verbose", "query").
			SetDescription("Show verbose greeting").
			AddExample("Show extra text", openapi3.NewExample(true)).
			AddExample("Hide extra text", openapi3.NewExample(false)),
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
