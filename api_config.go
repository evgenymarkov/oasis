package oasis

// Config represents a configuration for a new API.
type APIConfig struct {
	// OpenAPI describes the configurable fields of the root OpenAPI document.
	OpenAPI apiConfigOpenAPI

	// Routing describes the options that affect the behavior of the router.
	Routing apiConfigRouting
}

type apiConfigOpenAPI struct {
	// Title of the API.
	Title string

	// Version of the OpenAPI document.
	Version string
}

type apiConfigRouting struct {
	// DocsPath is the path to the API documentation UI.
	// If set to `/docs` it will allow clients to get `/docs`
	// to view the Swagger UI in a browser.
	DocsPath string

	// SchemaPath is the path to the API specification without extension.
	// If set to `/openapi` it will allow clients to get `/openapi.json` or `/openapi.yaml`.
	SchemaPath string
}

// NewAPIConfig method creates new APIConfig with default settings.
func NewAPIConfig() *APIConfig {
	return &APIConfig{
		OpenAPI: apiConfigOpenAPI{
			Title:   "API",
			Version: "0.0.1",
		},
		Routing: apiConfigRouting{
			DocsPath:   "/docs",
			SchemaPath: "/openapi",
		},
	}
}

// SetAPITitle method sets OpenAPI document title.
func (c *APIConfig) SetAPITitle(title string) *APIConfig {
	c.OpenAPI.Title = title
	return c
}

// SetAPIVersion method sets OpenAPI document version.
func (c *APIConfig) SetAPIVersion(version string) *APIConfig {
	c.OpenAPI.Version = version
	return c
}

// SetDocsPath method sets custom path for Swagger UI handlers
func (c *APIConfig) SetDocsPath(path string) *APIConfig {
	c.Routing.DocsPath = path
	return c
}

// SetSchemaPath method sets custom path for OpenAPI schema handlers
func (c *APIConfig) SetSchemaPath(path string) *APIConfig {
	c.Routing.SchemaPath = path
	return c
}
