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
			SetSummary("API for greetings").
			SetDescription("_Oasis_ is a library for Go web apps").
			SetContact(
				openapi3.NewContact("API Support").
					SetURL("https://example.com/support").
					SetEmail("greeting-api@support.example.com"),
			).
			SetLicense(
				openapi3.NewLicense("MIT").
					SetIdentifier("MIT"),
			).
			SetExternalDocs(
				openapi3.NewExternalDocumentation("https://spec.openapis.org/oas/latest.html").
					SetDescription("OpenAPI specification document"),
			),
	)

	api.Get(
		"/greeting/{name}",
		GetGreetingHandler,
		GetGreetingOperation,
	)

	return api
}
