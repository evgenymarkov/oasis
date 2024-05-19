package openapi3_test

import (
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
)

func TestPathItem(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		pathItem := openapi3.NewPathItem()

		t.Run("Values", func(t *testing.T) {
			assert.Nil(t, pathItem.Get)
			assert.Nil(t, pathItem.Head)
			assert.Nil(t, pathItem.Post)
			assert.Nil(t, pathItem.Put)
			assert.Nil(t, pathItem.Patch)
			assert.Nil(t, pathItem.Delete)
			assert.Nil(t, pathItem.Options)
			assert.Nil(t, pathItem.Trace)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(t, map[string]any{}, pathItem)
		})
	})

	t.Run("WithOperations", func(t *testing.T) {
		getOperation := openapi3.NewOperation("GetSomething").
			SetSummary("Get operation")
		postOperation := openapi3.NewOperation("SaveSomething").
			SetSummary("Post operation")
		pathItem := openapi3.NewPathItem().
			SetOperation("GET", getOperation).
			SetOperation("POST", postOperation)

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, getOperation, pathItem.Get)
			assert.Equal(t, postOperation, pathItem.Post)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"get": map[string]any{
						"operationId": "GetSomething",
						"summary":     "Get operation",
					},
					"post": map[string]any{
						"operationId": "SaveSomething",
						"summary":     "Post operation",
					},
				},
				pathItem,
			)
		})
	})
}
