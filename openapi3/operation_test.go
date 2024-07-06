package openapi3_test

import (
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
)

func TestOperation(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		operation := openapi3.NewOperation("GetPet")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "GetPet", operation.OperationID)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{"operationId": "GetPet"},
				operation,
			)
		})
	})

	t.Run("WithSummary", func(t *testing.T) {
		operation := openapi3.NewOperation("GetPet").
			SetSummary("Get a pet by ID")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "GetPet", operation.OperationID)
			assert.Equal(t, "Get a pet by ID", operation.Summary)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"operationId": "GetPet",
					"summary":     "Get a pet by ID",
				},
				operation,
			)
		})
	})

	t.Run("WithDescription", func(t *testing.T) {
		operation := openapi3.NewOperation("GetPet").
			SetDescription("Returns a single pet")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "GetPet", operation.OperationID)
			assert.Equal(t, "Returns a single pet", operation.Description)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"operationId": "GetPet",
					"description": "Returns a single pet",
				},
				operation,
			)
		})
	})

	t.Run("WithDeprecatedFlag", func(t *testing.T) {
		operation := openapi3.NewOperation("GetPet").
			MarkAsDeprecated()

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "GetPet", operation.OperationID)
			assert.True(t, operation.Deprecated)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"operationId": "GetPet",
					"deprecated":  true,
				},
				operation,
			)
		})
	})

	t.Run("WithExternalDocs", func(t *testing.T) {
		externalDocs := openapi3.NewExternalDocumentation("https://example.com/docs").
			SetDescription("Additional documentation")
		operation := openapi3.NewOperation("GetPet").
			SetExternalDocs(externalDocs)

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "GetPet", operation.OperationID)
			assert.Equal(t, externalDocs, operation.ExternalDocs)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"operationId": "GetPet",
					"externalDocs": map[string]any{
						"url":         "https://example.com/docs",
						"description": "Additional documentation",
					},
				},
				operation,
			)
		})
	})

	t.Run("WithTags", func(t *testing.T) {
		operation := openapi3.NewOperation("GetPet").
			SetTags("pets", "read")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "GetPet", operation.OperationID)
			assert.ElementsMatch(t, []string{"pets", "read"}, operation.Tags)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"operationId": "GetPet",
					"tags":        []string{"pets", "read"},
				},
				operation,
			)
		})
	})
}
