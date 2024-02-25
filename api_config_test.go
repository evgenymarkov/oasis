package oasis_test

import (
	"testing"

	"github.com/evgenymarkov/oasis"
	"github.com/stretchr/testify/assert"
)

func TestAPIConfigDefaults(t *testing.T) {
	config := oasis.NewAPIConfig()

	assert.Equal(t, "/api", config.DocsUIPath)
	assert.Equal(t, "/api/openapi.json", config.JSONDocumentPath)
	assert.Equal(t, "/api/openapi.yaml", config.YAMLDocumentPath)
}

func TestAPIConfigOverrides(t *testing.T) {
	config := oasis.NewAPIConfig().
		SetDocsUIPath("/docs").
		SetJSONDocumentPath("/docs/openapi.json").
		SetYAMLDocumentPath("/docs/openapi.yaml")

	assert.Equal(t, "/docs", config.DocsUIPath)
	assert.Equal(t, "/docs/openapi.json", config.JSONDocumentPath)
	assert.Equal(t, "/docs/openapi.yaml", config.YAMLDocumentPath)
}
