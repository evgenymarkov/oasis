package openapi3_test

import (
	"encoding/json"
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExternalDocumentation(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		externalDocs := openapi3.NewExternalDocumentation("https://example.com")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "https://example.com", externalDocs.URL)
			assert.Equal(t, "", externalDocs.Description)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"url": "https://example.com",
			})
			gotBytes, gotErr := json.Marshal(externalDocs)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})

	t.Run("WithDescription", func(t *testing.T) {
		externalDocs := openapi3.NewExternalDocumentation("https://example.com").
			SetDescription("Popular search engine")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "https://example.com", externalDocs.URL)
			assert.Equal(t, "Popular search engine", externalDocs.Description)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"url":         "https://example.com",
				"description": "Popular search engine",
			})
			gotBytes, gotErr := json.Marshal(externalDocs)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})
}
