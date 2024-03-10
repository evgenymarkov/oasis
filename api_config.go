package oasis

// APIConfig struct describes configuration options for API.
type APIConfig struct {
	// DocsUIPath is the path to the API documentation UI.
	//
	// Default: "/api"
	DocsUIPath string

	// DocumentPath is the path to the API specification in JSON format.
	//
	// Default: "/api/openapi.json"
	DocumentPath string
}

// NewAPIConfig method creates new APIConfig with default settings.
func NewAPIConfig() *APIConfig {
	return &APIConfig{
		DocsUIPath:   "/api",
		DocumentPath: "/api/openapi.json",
	}
}

// SetDocsUIPath method sets custom path for the API documentation UI.
func (c *APIConfig) SetDocsUIPath(path string) *APIConfig {
	c.DocsUIPath = path

	return c
}

// SetDocumentPath method sets custom path for the API specification in JSON format.
func (c *APIConfig) SetDocumentPath(path string) *APIConfig {
	c.DocumentPath = path

	return c
}
