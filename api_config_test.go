package oasis_test

import (
	"testing"

	"github.com/evgenymarkov/oasis"
	"github.com/stretchr/testify/assert"
)

func TestAPIConfigDefaults(t *testing.T) {
	config := oasis.NewAPIConfig()

	assert.Equal(t, "/api", config.DocsUIPath)
	assert.Equal(t, "/api/openapi.json", config.DocumentPath)
}

func TestAPIConfigOverrides(t *testing.T) {
	config := oasis.NewAPIConfig().
		SetDocsUIPath("/docs").
		SetDocumentPath("/docs/openapi.json")

	assert.Equal(t, "/docs", config.DocsUIPath)
	assert.Equal(t, "/docs/openapi.json", config.DocumentPath)
}
