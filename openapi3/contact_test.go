package openapi3_test

import (
	"encoding/json"
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContact(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		contact := openapi3.NewContact("API Support")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "API Support", contact.Name)
			assert.Equal(t, "", contact.URL)
			assert.Equal(t, "", contact.Email)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"name": "API Support",
			})
			gotBytes, gotErr := json.Marshal(contact)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})

	t.Run("WithURL", func(t *testing.T) {
		contact := openapi3.NewContact("API Support").
			SetURL("https://example.com/support")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "API Support", contact.Name)
			assert.Equal(t, "https://example.com/support", contact.URL)
			assert.Equal(t, "", contact.Email)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"name": "API Support",
				"url":  "https://example.com/support",
			})
			gotBytes, gotErr := json.Marshal(contact)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})

	t.Run("WithEmail", func(t *testing.T) {
		contact := openapi3.NewContact("API Support").
			SetEmail("greeting-api@support.example.com")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "API Support", contact.Name)
			assert.Equal(t, "", contact.URL)
			assert.Equal(t, "greeting-api@support.example.com", contact.Email)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"name":  "API Support",
				"email": "greeting-api@support.example.com",
			})
			gotBytes, gotErr := json.Marshal(contact)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})
}
