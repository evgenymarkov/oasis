package oasis

import (
	"net/http"
	"slices"
	"strings"
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

var pingOperation = openapi3.NewOperation().
	SetOperationID("ping").
	SetSummary("Ping server")

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func TestAPIEmpty(t *testing.T) {
	router := chi.NewRouter()

	NewAPI(
		router,
		NewAPIConfig().
			SetDocsUIPath("/api").
			SetJSONDocumentPath("/api/openapi.json").
			SetYAMLDocumentPath("/api/openapi.yaml"),
		openapi3.NewDocument().
			SetTitle("Greeting API").
			SetVersion("1.0.0"),
	)

	assert.Empty(t, router.Routes())
}

func TestAPIWithOperations(t *testing.T) {
	router := chi.NewRouter()

	api := NewAPI(
		router,
		NewAPIConfig().
			SetDocsUIPath("/api").
			SetJSONDocumentPath("/api/openapi.json").
			SetYAMLDocumentPath("/api/openapi.yaml"),
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
	assert.Len(t, routes, 9)

	methods := []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	}

	for _, method := range methods {
		route := routes[slices.IndexFunc(routes, func(r chi.Route) bool {
			return r.Pattern == "/ping-"+strings.ToLower(method)
		})]
		assert.Nil(t, route.SubRoutes)
		assert.Len(t, route.Handlers, 1)
		assert.Contains(t, route.Handlers, method)
	}
}
