package oasis

import (
	"github.com/evgenymarkov/oasis/openapi3"
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
	return &API{
		router:   router,
		config:   config,
		document: document,
	}
}
