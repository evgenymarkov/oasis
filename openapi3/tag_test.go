package openapi3_test

import (
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
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
			assertObjectSerialization(
				t,
				map[string]any{"name": "payments"},
				tag,
			)
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
			assertObjectSerialization(
				t,
				map[string]any{
					"name":        "payments",
					"description": "Payments operations",
				},
				tag,
			)
		})
	})

	t.Run("WithExternalDocs", func(t *testing.T) {
		tag := openapi3.NewTag("payments").
			SetExternalDocs(
				openapi3.NewExternalDocumentation("https://example.com").
					SetDescription("Popular search engine"),
			)

		t.Run("Values", func(t *testing.T) {
			docs := openapi3.NewExternalDocumentation("https://example.com").
				SetDescription("Popular search engine")

			assert.Equal(t, "payments", tag.Name)
			assert.Equal(t, "", tag.Description)
			assert.Equal(t, docs, tag.ExternalDocs)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"name": "payments",
					"externalDocs": map[string]any{
						"url":         "https://example.com",
						"description": "Popular search engine",
					},
				},
				tag,
			)
		})
	})
}
