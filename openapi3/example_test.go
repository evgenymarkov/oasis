package openapi3_test

import (
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		t.Run("String", func(t *testing.T) {
			example := openapi3.NewExample("Sample Value")

			t.Run("Values", func(t *testing.T) {
				assert.Equal(t, "Sample Value", example.Value)
			})

			t.Run("Serialization", func(t *testing.T) {
				assertObjectSerialization(
					t,
					map[string]any{"value": "Sample Value"},
					example,
				)
			})
		})

		t.Run("Map", func(t *testing.T) {
			example := openapi3.NewExample(map[string]any{"key": "value"})

			t.Run("Values", func(t *testing.T) {
				assert.Equal(t, map[string]any{"key": "value"}, example.Value)
			})

			t.Run("Serialization", func(t *testing.T) {
				assertObjectSerialization(
					t,
					map[string]any{"value": map[string]any{"key": "value"}},
					example,
				)
			})
		})
	})

	t.Run("WithSummary", func(t *testing.T) {
		example := openapi3.NewExample("Sample Value").
			SetSummary("Short description")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "Sample Value", example.Value)
			assert.Equal(t, "Short description", example.Summary)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"value":   "Sample Value",
					"summary": "Short description",
				},
				example,
			)
		})
	})

	t.Run("WithDescription", func(t *testing.T) {
		example := openapi3.NewExample("Sample Value").
			SetDescription("Long description")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "Sample Value", example.Value)
			assert.Equal(t, "Long description", example.Description)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"value":       "Sample Value",
					"description": "Long description",
				},
				example,
			)
		})
	})
}
