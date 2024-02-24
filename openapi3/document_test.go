package openapi3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentDefaults(t *testing.T) {
	document := NewDocument()

	assert.Equal(t, "3.1.0", document.OpenAPI)
	assert.Equal(t, "API", document.Info.Title)
	assert.Equal(t, "0.0.1", document.Info.Version)
}

func TestDocumentOverrides(t *testing.T) {
	document := NewDocument().
		SetTitle("Greeting API").
		SetVersion("1.0.0")

	assert.Equal(t, "3.1.0", document.OpenAPI)
	assert.Equal(t, "Greeting API", document.Info.Title)
	assert.Equal(t, "1.0.0", document.Info.Version)
}
