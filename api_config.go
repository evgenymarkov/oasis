package oasis

// APIConfig struct describes configuration options for API.
type APIConfig struct {
	// DocumentPath is the path to the API specification in JSON format.
	//
	// Default: "/api/openapi.json"
	DocumentPath string

	// SwaggerUIPath is the path to the API documentation UI.
	//
	// Default: "/api"
	SwaggerUIPath string

	// SwaggerUITitle is the title of API documentation UI page.
	//
	// Default: "Swagger UI"
	SwaggerUITitle string
}

// NewAPIConfig method creates new APIConfig with default settings.
func NewAPIConfig() *APIConfig {
	return &APIConfig{
		DocumentPath:   "/api/openapi.json",
		SwaggerUIPath:  "/api",
		SwaggerUITitle: "Swagger UI",
	}
}

// SetDocumentPath method sets custom path for the API specification in JSON format.
func (c *APIConfig) SetDocumentPath(path string) *APIConfig {
	c.DocumentPath = path

	return c
}

// SetSwaggerUIPath method sets custom path for the API documentation UI.
func (c *APIConfig) SetSwaggerUIPath(path string) *APIConfig {
	c.SwaggerUIPath = path

	return c
}

// SetSwaggerUITitle method sets custom title of the API documentation UI page.
func (c *APIConfig) SetSwaggerUITitle(title string) *APIConfig {
	c.SwaggerUITitle = title

	return c
}
