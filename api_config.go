package oasis

const (
	apiConfigDocsPath    = "/docs"
	apiConfigOpenAPIPath = "/openapi"
	apiConfigSchemasPath = "/schemas"
)

// Config represents a configuration for a new API.
// See `oasis.NewDefaultAPIConfig()` as a starting point.
type APIConfig struct {
	// DocsPath is the path to the API documentation.
	// If set to `/docs` it will allow clients to get `/docs`
	// to view the Swagger UI in a browser.
	DocsPath string

	// OpenAPIPath is the path to the OpenAPI spec without extension.
	// If set to `/openapi` it will allow clients to get `/openapi.json` or `/openapi.yaml`.
	OpenAPIPath string

	// SchemasPath is the path to the API schemas.
	// If set to `/schemas` it will allow clients to get `/schemas/{schema}`
	// to view the schema in a browser or for use in editors like VSCode
	// to provide autocomplete & validation.
	SchemasPath string
}

// NewDefaultAPIConfig returns a default configuration for a new API.
// It is a good starting point for creating your own configuration.
//
// The `/openapi.[json|yaml]`, `/docs`, and `/schemas` paths are
// set up to serve the OpenAPI specification, Swagger UI, and schemas.
//
//	// Create the API config.
//	apiConfig := oasis.NewDefaultAPIConfig()
//
//	// Create the API using the config.
//	api := oasis.NewAPI(apiConfig)
func NewDefaultAPIConfig() *APIConfig {
	return &APIConfig{
		DocsPath:    apiConfigDocsPath,
		OpenAPIPath: apiConfigOpenAPIPath,
		SchemasPath: apiConfigSchemasPath,
	}
}
