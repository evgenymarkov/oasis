package openapi3_test

import (
	"encoding/json"
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServerVariable(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		serverVariable := openapi3.NewServerVariable("default-token")

		t.Run("Values", func(t *testing.T) {
			assert.Nil(t, serverVariable.Enum)
			assert.Equal(t, "default-token", serverVariable.Default)
			assert.Equal(t, "", serverVariable.Description)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"default": "default-token",
			})
			gotBytes, gotErr := json.Marshal(serverVariable)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})

	t.Run("WithEnum", func(t *testing.T) {
		serverVariable := openapi3.NewServerVariable("default-token").
			SetEnum([]string{"default-token", "custom-token"})

		t.Run("Values", func(t *testing.T) {
			assert.ElementsMatch(t, []string{"default-token", "custom-token"}, serverVariable.Enum)
			assert.Equal(t, "default-token", serverVariable.Default)
			assert.Equal(t, "", serverVariable.Description)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"enum":    []string{"default-token", "custom-token"},
				"default": "default-token",
			})
			gotBytes, gotErr := json.Marshal(serverVariable)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})

	t.Run("WithDescription", func(t *testing.T) {
		serverVariable := openapi3.NewServerVariable("default-token").
			SetDescription("Authentication token")

		t.Run("Values", func(t *testing.T) {
			assert.Nil(t, serverVariable.Enum)
			assert.Equal(t, "default-token", serverVariable.Default)
			assert.Equal(t, "Authentication token", serverVariable.Description)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"default":     "default-token",
				"description": "Authentication token",
			})
			gotBytes, gotErr := json.Marshal(serverVariable)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})
}
