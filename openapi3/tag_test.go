package openapi3_test

import (
	"encoding/json"
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTag(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		tag := openapi3.NewTag("payments")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "payments", tag.Name)
			assert.Equal(t, "", tag.Description)
			assert.Nil(t, tag.ExternalDocs)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"name": "payments",
			})
			gotBytes, gotErr := json.Marshal(tag)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})

	t.Run("WithDescription", func(t *testing.T) {
		tag := openapi3.NewTag("payments").
			SetDescription("Payments operations")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "payments", tag.Name)
			assert.Equal(t, "Payments operations", tag.Description)
			assert.Nil(t, tag.ExternalDocs)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"name":        "payments",
				"description": "Payments operations",
			})
			gotBytes, gotErr := json.Marshal(tag)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})

	t.Run("WithExternalDocs", func(t *testing.T) {
		tag := openapi3.NewTag("payments").
			SetExternalDocs(
				openapi3.NewExternalDocumentation("https://yandex.com").
					SetDescription("Popular search engine"),
			)

		t.Run("Values", func(t *testing.T) {
			docs := openapi3.NewExternalDocumentation("https://yandex.com").
				SetDescription("Popular search engine")

			assert.Equal(t, "payments", tag.Name)
			assert.Equal(t, "", tag.Description)
			assert.Equal(t, docs, tag.ExternalDocs)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"name": "payments",
				"externalDocs": map[string]any{
					"url":         "https://yandex.com",
					"description": "Popular search engine",
				},
			})
			gotBytes, gotErr := json.Marshal(tag)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})
}
