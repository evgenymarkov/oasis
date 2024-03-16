package oasis_test

import (
	"net/http"
	"slices"
	"testing"

	"github.com/evgenymarkov/oasis"
	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

var pingOperation = openapi3.NewOperation().
	SetOperationID("ping").
	SetSummary("Ping server")

func pingHandler(response http.ResponseWriter, _ *http.Request) {
	response.Write([]byte("pong"))
}

func TestAPIEmpty(t *testing.T) {
	router := chi.NewRouter()

	oasis.NewAPI(
		router,
		oasis.NewAPIConfig().
			SetDocumentPath("/api/openapi.json").
			SetSwaggerUIPath("/api").
			SetSwaggerUITitle("API Docs"),
		openapi3.NewDocument().
			SetTitle("Greeting API").
			SetVersion("1.0.0"),
	)

	routes := router.Routes()
	assert.Len(t, routes, 11)
	checkHandlersRegistered(t, routes, []string{
		"/api",
		"/api/favicon-32x32.png",
		"/api/index.css",
		"/api/oauth2-redirect.html",
		"/api/openapi.json",
		"/api/swagger-ui-bundle.js",
		"/api/swagger-ui-bundle.js.map",
		"/api/swagger-ui-standalone-preset.js",
		"/api/swagger-ui-standalone-preset.js.map",
		"/api/swagger-ui.css",
		"/api/swagger-ui.css.map",
	})
}

func TestAPIWithOperations(t *testing.T) {
	router := chi.NewRouter()

	api := oasis.NewAPI(
		router,
		oasis.NewAPIConfig().
			SetDocumentPath("/api/openapi.json").
			SetSwaggerUIPath("/api").
			SetSwaggerUITitle("API Docs"),
		openapi3.NewDocument().
			SetTitle("Greeting API").
			SetVersion("1.0.0"),
	)

	api.Get("/ping-get", pingHandler, pingOperation)
	api.Head("/ping-head", pingHandler, pingOperation)
	api.Post("/ping-post", pingHandler, pingOperation)
	api.Put("/ping-put", pingHandler, pingOperation)
	api.Patch("/ping-patch", pingHandler, pingOperation)
	api.Delete("/ping-delete", pingHandler, pingOperation)
	api.Connect("/ping-connect", pingHandler, pingOperation)
	api.Options("/ping-options", pingHandler, pingOperation)
	api.Trace("/ping-trace", pingHandler, pingOperation)

	routes := router.Routes()
	assert.Len(t, routes, 20)
	checkHandlersRegistered(t, routes, []string{
		"/api",
		"/api/favicon-32x32.png",
		"/api/index.css",
		"/api/oauth2-redirect.html",
		"/api/openapi.json",
		"/api/swagger-ui-bundle.js",
		"/api/swagger-ui-bundle.js.map",
		"/api/swagger-ui-standalone-preset.js",
		"/api/swagger-ui-standalone-preset.js.map",
		"/api/swagger-ui.css",
		"/api/swagger-ui.css.map",
		"/ping-get",
		"/ping-head",
		"/ping-post",
		"/ping-put",
		"/ping-patch",
		"/ping-delete",
		"/ping-connect",
		"/ping-options",
		"/ping-trace",
	})
}

func checkHandlersRegistered(t *testing.T, routes []chi.Route, patterns []string) {
	t.Helper()

	for _, pattern := range patterns {
		assert.NotEqual(t, -1, slices.IndexFunc(routes, func(r chi.Route) bool {
			return r.Pattern == pattern
		}))
	}
}
