package oasis_test

import (
	"testing"

	"github.com/evgenymarkov/oasis"
	"github.com/stretchr/testify/assert"
)

func TestAPIConfigDefaults(t *testing.T) {
	config := oasis.NewAPIConfig()

	assert.Equal(t, "/api/openapi.json", config.DocumentPath)
	assert.Equal(t, "/api", config.SwaggerUIPath)
	assert.Equal(t, "Swagger UI", config.SwaggerUITitle)
}

func TestAPIConfigOverrides(t *testing.T) {
	config := oasis.NewAPIConfig().
		SetDocumentPath("/docs/openapi.json").
		SetSwaggerUIPath("/docs").
		SetSwaggerUITitle("API Docs")

	assert.Equal(t, "/docs/openapi.json", config.DocumentPath)
	assert.Equal(t, "/docs", config.SwaggerUIPath)
	assert.Equal(t, "API Docs", config.SwaggerUITitle)
}
