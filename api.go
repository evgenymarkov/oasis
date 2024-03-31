package oasis

import (
	"net/http"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/evgenymarkov/oasis/rendering"
)

// API struct is a wrapper on top of multiplexer for registering operations.
type API struct {
	mux      *http.ServeMux
	config   *APIConfig
	document *openapi3.Document
}

// NewAPI method creates new API struct and fills it up.
func NewAPI(
	mux *http.ServeMux,
	config *APIConfig,
	document *openapi3.Document,
) *API {
	mux.HandleFunc(
		http.MethodGet+" "+config.DocumentPath,
		rendering.NewDocumentHandler(document),
	)

	mux.HandleFunc(
		http.MethodGet+" "+config.SwaggerUIPath,
		rendering.NewSwaggerUIHandler(rendering.SwaggerUIConfig{
			BaseURL:   config.SwaggerUIPath,
			PageTitle: config.SwaggerUITitle,
			Document:  document,
		}),
	)

	for _, staticFile := range rendering.GetSwaggerUIStaticFiles() {
		mux.HandleFunc(
			http.MethodGet+" "+config.SwaggerUIPath+"/"+staticFile,
			rendering.NewSwaggerUIStaticHandler(staticFile),
		)
	}

	return &API{
		mux:      mux,
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
	a.mux.HandleFunc(http.MethodGet+" "+pattern, handler)
}

// Head method registers an handler for HTTP HEAD requests matching the pattern
// and updates operations table in OpenAPI document.
func (a *API) Head(
	pattern string,
	handler http.HandlerFunc,
	operation *openapi3.Operation,
) {
	a.mux.HandleFunc(http.MethodHead+" "+pattern, handler)
}

// Post method registers an handler for HTTP POST requests matching the pattern
// and updates operations table in OpenAPI document.
func (a *API) Post(
	pattern string,
	handler http.HandlerFunc,
	operation *openapi3.Operation,
) {
	a.mux.HandleFunc(http.MethodPost+" "+pattern, handler)
}

// Put method registers an handler for HTTP PUT requests matching the pattern
// and updates operations table in OpenAPI document.
func (a *API) Put(
	pattern string,
	handler http.HandlerFunc,
	operation *openapi3.Operation,
) {
	a.mux.HandleFunc(http.MethodPut+" "+pattern, handler)
}

// Patch method registers an handler for HTTP PATCH requests matching the pattern
// and updates operations table in OpenAPI document.
func (a *API) Patch(
	pattern string,
	handler http.HandlerFunc,
	operation *openapi3.Operation,
) {
	a.mux.HandleFunc(http.MethodPatch+" "+pattern, handler)
}

// Delete method registers an handler for HTTP DELETE requests matching the pattern
// and updates operations table in OpenAPI document.
func (a *API) Delete(
	pattern string,
	handler http.HandlerFunc,
	operation *openapi3.Operation,
) {
	a.mux.HandleFunc(http.MethodDelete+" "+pattern, handler)
}

// Connect method registers an handler for HTTP CONNECT requests matching the pattern
// and updates operations table in OpenAPI document.
func (a *API) Connect(
	pattern string,
	handler http.HandlerFunc,
	operation *openapi3.Operation,
) {
	a.mux.HandleFunc(http.MethodConnect+" "+pattern, handler)
}

// Options method registers an handler for HTTP OPTIONS requests matching the pattern
// and updates operations table in OpenAPI document.
func (a *API) Options(
	pattern string,
	handler http.HandlerFunc,
	operation *openapi3.Operation,
) {
	a.mux.HandleFunc(http.MethodOptions+" "+pattern, handler)
}

// Trace method registers an handler for HTTP TRACE requests matching the pattern
// and updates operations table in OpenAPI document.
func (a *API) Trace(
	pattern string,
	handler http.HandlerFunc,
	operation *openapi3.Operation,
) {
	a.mux.HandleFunc(http.MethodTrace+" "+pattern, handler)
}
