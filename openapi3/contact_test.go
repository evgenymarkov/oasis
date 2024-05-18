package openapi3_test

import (
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
)

func TestContact(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		contact := openapi3.NewContact("API Support")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "API Support", contact.Name)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{"name": "API Support"},
				contact,
			)
		})
	})

	t.Run("WithURL", func(t *testing.T) {
		contact := openapi3.NewContact("API Support").
			SetURL("https://example.com/support")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "API Support", contact.Name)
			assert.Equal(t, "https://example.com/support", contact.URL)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"name": "API Support",
					"url":  "https://example.com/support",
				},
				contact,
			)
		})
	})

	t.Run("WithEmail", func(t *testing.T) {
		contact := openapi3.NewContact("API Support").
			SetEmail("greeting-api@support.example.com")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "API Support", contact.Name)
			assert.Equal(t, "greeting-api@support.example.com", contact.Email)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"name":  "API Support",
					"email": "greeting-api@support.example.com",
				},
				contact,
			)
		})
	})
}
