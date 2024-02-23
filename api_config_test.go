package oasis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIConfigDefaults(t *testing.T) {
	config := NewAPIConfig()

	assert.Equal(t, "Service API", config.OpenAPI.Title)
	assert.Equal(t, "0.0.1", config.OpenAPI.Version)

	assert.Equal(t, "/docs", config.Routing.DocsPath)
	assert.Equal(t, "/openapi", config.Routing.SchemaPath)
}

func TestAPIConfigOverrides(t *testing.T) {
	config := NewAPIConfig().
		SetAPITitle("Greeting API").
		SetAPIVersion("1.0.0").
		SetDocsPath("/private/docs").
		SetSchemaPath("/private/openapi")

	assert.Equal(t, "Greeting API", config.OpenAPI.Title)
	assert.Equal(t, "1.0.0", config.OpenAPI.Version)

	assert.Equal(t, "/private/docs", config.Routing.DocsPath)
	assert.Equal(t, "/private/openapi", config.Routing.SchemaPath)
}
