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

func TestEmptyAPIHandlers(t *testing.T) {
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
		{method: http.MethodGet, pattern: "/api/swagger-ui.css"},
		{method: http.MethodGet, pattern: "/api/swagger-ui.css.map"},
	})
}

func TestAPIWithOperationsHandlers(t *testing.T) {
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
	api.Options("/ping-options", pingHandler, pingOperation)
	api.Trace("/ping-trace", pingHandler, pingOperation)

	checkHandlersRegistered(t, mux, []endpoint{
		{method: http.MethodGet, pattern: "/ping-get"},
		{method: http.MethodHead, pattern: "/ping-head"},
		{method: http.MethodPost, pattern: "/ping-post"},
		{method: http.MethodPut, pattern: "/ping-put"},
		{method: http.MethodPatch, pattern: "/ping-patch"},
		{method: http.MethodDelete, pattern: "/ping-delete"},
		{method: http.MethodOptions, pattern: "/ping-options"},
		{method: http.MethodTrace, pattern: "/ping-trace"},
	})

	checkOpenAPIOperations(t, api.Document(), []endpoint{
		{method: http.MethodGet, pattern: "/ping-get"},
		{method: http.MethodHead, pattern: "/ping-head"},
		{method: http.MethodPost, pattern: "/ping-post"},
		{method: http.MethodPut, pattern: "/ping-put"},
		{method: http.MethodPatch, pattern: "/ping-patch"},
		{method: http.MethodDelete, pattern: "/ping-delete"},
		{method: http.MethodOptions, pattern: "/ping-options"},
		{method: http.MethodTrace, pattern: "/ping-trace"},
	})
}

func TestAPIWithOverlappingPaths(t *testing.T) {
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

	api.Get(
		"/resource",
		pingHandler,
		openapi3.NewOperation().
			SetOperationID("get-resource").
			SetSummary("Get Resource"),
	)
	api.Post(
		"/resource",
		pingHandler,
		openapi3.NewOperation().
			SetOperationID("post-resource").
			SetSummary("Create Resource"),
	)
	api.Get(
		"/resource/{id}",
		pingHandler,
		openapi3.NewOperation().
			SetOperationID("get-resource-by-id").
			SetSummary("Get Resource by ID"),
	)
	api.Put(
		"/resource/{id}",
		pingHandler,
		openapi3.NewOperation().
			SetOperationID("put-resource-by-id").
			SetSummary("Update Resource by ID"),
	)

	checkHandlersRegistered(t, mux, []endpoint{
		{method: http.MethodGet, pattern: "/resource"},
		{method: http.MethodPost, pattern: "/resource"},
		{method: http.MethodGet, pattern: "/resource/{id}"},
		{method: http.MethodPut, pattern: "/resource/{id}"},
	})

	checkOpenAPIOperations(t, api.Document(), []endpoint{
		{method: http.MethodGet, pattern: "/resource"},
		{method: http.MethodPost, pattern: "/resource"},
		{method: http.MethodGet, pattern: "/resource/{id}"},
		{method: http.MethodPut, pattern: "/resource/{id}"},
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

func checkOpenAPIOperations(t *testing.T, document *openapi3.Document, endpoints []endpoint) {
	t.Helper()

	for _, endpoint := range endpoints {
		t.Run(endpoint.method+endpoint.pattern, func(t *testing.T) {
			pathItem, exists := document.Paths[endpoint.pattern]
			require.True(t, exists)

			var operation *openapi3.Operation

			switch endpoint.method {
			case http.MethodGet:
				operation = pathItem.Get
			case http.MethodHead:
				operation = pathItem.Head
			case http.MethodPost:
				operation = pathItem.Post
			case http.MethodPut:
				operation = pathItem.Put
			case http.MethodPatch:
				operation = pathItem.Patch
			case http.MethodDelete:
				operation = pathItem.Delete
			case http.MethodOptions:
				operation = pathItem.Options
			case http.MethodTrace:
				operation = pathItem.Trace
			}

			require.NotNil(t, operation)
			assert.NotEmpty(t, operation.OperationID)
			assert.NotEmpty(t, operation.Summary)
		})
	}
}
