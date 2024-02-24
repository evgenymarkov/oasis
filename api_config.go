package oasis

// APIConfig struct describes configuration options for API.
type APIConfig struct {
	// DocsUIPath is the path to the API documentation UI.
	//
	// Default: "/api".
	DocsUIPath string

	// JSONDocumentPath is the path to the API specification in JSON format.
	//
	// Default: "/api/openapi.json".
	JSONDocumentPath string

	// YAMLDocumentPath is the path to the API specification in YAML format.
	//
	// Default: "/api/openapi.yaml"
	YAMLDocumentPath string
}

// NewAPIConfig method creates new APIConfig with default settings.
func NewAPIConfig() *APIConfig {
	return &APIConfig{
		DocsUIPath:       "/api",
		JSONDocumentPath: "/api/openapi.json",
		YAMLDocumentPath: "/api/openapi.yaml",
	}
}

// SetDocsUIPath method sets custom path for the API documentation UI.
func (c *APIConfig) SetDocsUIPath(path string) *APIConfig {
	c.DocsUIPath = path
	return c
}

// SetJSONDocumentPath method sets custom path for the API specification in JSON format.
func (c *APIConfig) SetJSONDocumentPath(path string) *APIConfig {
	c.JSONDocumentPath = path
	return c
}

// SetYAMLDocumentPath method sets custom path for the API specification in YAML format.
func (c *APIConfig) SetYAMLDocumentPath(path string) *APIConfig {
	c.YAMLDocumentPath = path
	return c
}
