package openapi3_test

import (
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
)

func TestParameter(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		t.Run("Path", func(t *testing.T) {
			parameter := openapi3.NewParameter("petId", "path")

			t.Run("Values", func(t *testing.T) {
				assert.Equal(t, "petId", parameter.Name)
				assert.Equal(t, "path", parameter.In)
				assert.True(t, parameter.Required)
			})

			t.Run("Serialization", func(t *testing.T) {
				assertObjectSerialization(
					t,
					map[string]any{
						"name":     "petId",
						"in":       "path",
						"required": true,
					},
					parameter,
				)
			})
		})

		t.Run("Query", func(t *testing.T) {
			parameter := openapi3.NewParameter("showExtra", "query")

			t.Run("Values", func(t *testing.T) {
				assert.Equal(t, "showExtra", parameter.Name)
				assert.Equal(t, "query", parameter.In)
				assert.False(t, parameter.Required)
			})

			t.Run("Serialization", func(t *testing.T) {
				assertObjectSerialization(
					t,
					map[string]any{
						"name": "showExtra",
						"in":   "query",
					},
					parameter,
				)
			})
		})
	})

	t.Run("WithDescription", func(t *testing.T) {
		parameter := openapi3.NewParameter("petId", "path").
			SetDescription("ID of pet to return")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "petId", parameter.Name)
			assert.Equal(t, "path", parameter.In)
			assert.Equal(t, "ID of pet to return", parameter.Description)
			assert.True(t, parameter.Required)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"name":        "petId",
					"in":          "path",
					"description": "ID of pet to return",
					"required":    true,
				},
				parameter,
			)
		})
	})

	t.Run("WithRequiredFlag", func(t *testing.T) {
		parameter := openapi3.NewParameter("showExtra", "query").
			MarkAsRequired()

		t.Run("Values", func(t *testing.T) {
			assert.True(t, parameter.Required)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"name":     "showExtra",
					"in":       "query",
					"required": true,
				},
				parameter,
			)
		})
	})

	t.Run("WithDeprecatedFlag", func(t *testing.T) {
		parameter := openapi3.NewParameter("showExtra", "query").
			MarkAsDeprecated()

		t.Run("Values", func(t *testing.T) {
			assert.True(t, parameter.Deprecated)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"name":       "showExtra",
					"in":         "query",
					"deprecated": true,
				},
				parameter,
			)
		})
	})
}
