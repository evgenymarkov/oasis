package oasis_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/evgenymarkov/oasis"
	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type endpoint struct {
	method  string
	pattern string
}

var pingOperation = openapi3.NewOperation().
	SetOperationID("ping").
	SetSummary("Ping server")

func pingHandler(response http.ResponseWriter, _ *http.Request) {
	response.Write([]byte("pong"))
}

func TestAPIEmpty(t *testing.T) {
	mux := http.NewServeMux()

	oasis.NewAPI(
		mux,
		oasis.NewAPIConfig().
			SetDocumentPath("/api/openapi.json").
			SetSwaggerUIPath("/api").
			SetSwaggerUITitle("API Docs"),
		openapi3.NewDocument().
			SetTitle("Greeting API").
			SetVersion("1.0.0"),
	)

	checkHandlersRegistered(t, mux, []endpoint{
		{method: http.MethodGet, pattern: "/api"},
		{method: http.MethodGet, pattern: "/api/favicon-32x32.png"},
		{method: http.MethodGet, pattern: "/api/index.css"},
		{method: http.MethodGet, pattern: "/api/oauth2-redirect.html"},
		{method: http.MethodGet, pattern: "/api/openapi.json"},
		{method: http.MethodGet, pattern: "/api/swagger-ui-bundle.js"},
		{method: http.MethodGet, pattern: "/api/swagger-ui-bundle.js.map"},
		{method: http.MethodGet, pattern: "/api/swagger-ui-standalone-preset.js"},
		{method: http.MethodGet, pattern: "/api/swagger-ui-standalone-preset.js.map"},
		{method: http.MethodGet, pattern: "/api/swagger-ui.css"},
		{method: http.MethodGet, pattern: "/api/swagger-ui.css.map"},
	})
}

func TestAPIWithOperations(t *testing.T) {
	mux := http.NewServeMux()

	api := oasis.NewAPI(
		mux,
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

	checkHandlersRegistered(t, mux, []endpoint{
		{method: http.MethodGet, pattern: "/api"},
		{method: http.MethodGet, pattern: "/api/favicon-32x32.png"},
		{method: http.MethodGet, pattern: "/api/index.css"},
		{method: http.MethodGet, pattern: "/api/oauth2-redirect.html"},
		{method: http.MethodGet, pattern: "/api/openapi.json"},
		{method: http.MethodGet, pattern: "/api/swagger-ui-bundle.js"},
		{method: http.MethodGet, pattern: "/api/swagger-ui-bundle.js.map"},
		{method: http.MethodGet, pattern: "/api/swagger-ui-standalone-preset.js"},
		{method: http.MethodGet, pattern: "/api/swagger-ui-standalone-preset.js.map"},
		{method: http.MethodGet, pattern: "/api/swagger-ui.css"},
		{method: http.MethodGet, pattern: "/api/swagger-ui.css.map"},
		{method: http.MethodGet, pattern: "/ping-get"},
		{method: http.MethodHead, pattern: "/ping-head"},
		{method: http.MethodPost, pattern: "/ping-post"},
		{method: http.MethodPut, pattern: "/ping-put"},
		{method: http.MethodPatch, pattern: "/ping-patch"},
		{method: http.MethodDelete, pattern: "/ping-delete"},
		{method: http.MethodConnect, pattern: "/ping-connect"},
		{method: http.MethodOptions, pattern: "/ping-options"},
		{method: http.MethodTrace, pattern: "/ping-trace"},
	})
}

func checkHandlersRegistered(t *testing.T, mux *http.ServeMux, endpoints []endpoint) {
	t.Helper()

	for _, endpoint := range endpoints {
		t.Run(endpoint.method+endpoint.pattern, func(t *testing.T) {
			request, requestErr := http.NewRequestWithContext(
				context.TODO(),
				endpoint.method,
				endpoint.pattern,
				nil,
			)
			require.NoError(t, requestErr)

			_, internalPattern := mux.Handler(request)
			assert.Equal(t, endpoint.method+" "+endpoint.pattern, internalPattern)
		})
	}
}
