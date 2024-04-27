package api

import (
	"net/http"

	"github.com/evgenymarkov/oasis"
	"github.com/evgenymarkov/oasis/openapi3"
)

func NewAPI(mux *http.ServeMux) *oasis.API {
	api := oasis.NewAPI(
		mux,
		oasis.NewAPIConfig().
			SetDocumentPath("/api/openapi.json").
			SetSwaggerUIPath("/api").
			SetSwaggerUITitle("API Docs"),
		openapi3.NewDocument().
			SetTitle("Greeting API").
			SetVersion("1.0.0").
			SetTags(
				openapi3.NewTag("greetings").
					SetDescription("Greetings operations").
					SetExternalDocs(
						openapi3.NewExternalDocumentation(wikiHelloURL).
							SetDescription(wikiHelloDescription),
					),
			),
	)

	api.Get(
		"/greeting/{name}",
		GetGreetingHandler,
		GetGreetingOperation,
	)

	return api
}
