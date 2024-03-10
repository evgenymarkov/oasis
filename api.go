package oasis

import (
	"net/http"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/evgenymarkov/oasis/rendering"
	"github.com/go-chi/chi/v5"
)

// API struct is a wrapper on top of router for registering operations.
type API struct {
	router   chi.Router
	config   *APIConfig
	document *openapi3.Document
}

// NewAPI method creates new API struct and fills it up.
func NewAPI(
	router chi.Router,
	config *APIConfig,
	document *openapi3.Document,
) *API {
	router.Get(
		config.DocumentPath,
		rendering.NewDocumentHandler(document),
	)

	return &API{
		router:   router,
		config:   config,
		document: document,
	}
}

// Get method registers an handler for HTTP GET requests matching the pattern
// and updates operations table in OpenAPI document.
func (a *API) Get(
	pattern string,
	handler http.HandlerFunc,
	operation *openapi3.Operation,
) {
	a.router.Get(pattern, handler)
}

// Head method registers an handler for HTTP HEAD requests matching the pattern
// and updates operations table in OpenAPI document.
func (a *API) Head(
	pattern string,
	handler http.HandlerFunc,
	operation *openapi3.Operation,
) {
	a.router.Head(pattern, handler)
}

// Post method registers an handler for HTTP POST requests matching the pattern
// and updates operations table in OpenAPI document.
func (a *API) Post(
	pattern string,
	handler http.HandlerFunc,
	operation *openapi3.Operation,
) {
	a.router.Post(pattern, handler)
}

// Put method registers an handler for HTTP PUT requests matching the pattern
// and updates operations table in OpenAPI document.
func (a *API) Put(
	pattern string,
	handler http.HandlerFunc,
	operation *openapi3.Operation,
) {
	a.router.Put(pattern, handler)
}

// Patch method registers an handler for HTTP PATCH requests matching the pattern
// and updates operations table in OpenAPI document.
func (a *API) Patch(
	pattern string,
	handler http.HandlerFunc,
	operation *openapi3.Operation,
) {
	a.router.Patch(pattern, handler)
}

// Delete method registers an handler for HTTP DELETE requests matching the pattern
// and updates operations table in OpenAPI document.
func (a *API) Delete(
	pattern string,
	handler http.HandlerFunc,
	operation *openapi3.Operation,
) {
	a.router.Delete(pattern, handler)
}

// Connect method registers an handler for HTTP CONNECT requests matching the pattern
// and updates operations table in OpenAPI document.
func (a *API) Connect(
	pattern string,
	handler http.HandlerFunc,
	operation *openapi3.Operation,
) {
	a.router.Connect(pattern, handler)
}

// Options method registers an handler for HTTP OPTIONS requests matching the pattern
// and updates operations table in OpenAPI document.
func (a *API) Options(
	pattern string,
	handler http.HandlerFunc,
	operation *openapi3.Operation,
) {
	a.router.Options(pattern, handler)
}

// Trace method registers an handler for HTTP TRACE requests matching the pattern
// and updates operations table in OpenAPI document.
func (a *API) Trace(
	pattern string,
	handler http.HandlerFunc,
	operation *openapi3.Operation,
) {
	a.router.Trace(pattern, handler)
}
