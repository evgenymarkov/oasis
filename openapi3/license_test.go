package openapi3_test

import (
	"encoding/json"
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLicense(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		license := openapi3.NewLicense("Apache 2.0")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "Apache 2.0", license.Name)
			assert.Equal(t, "", license.Identifier)
			assert.Equal(t, "", license.URL)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"name": "Apache 2.0",
			})
			gotBytes, gotErr := json.Marshal(license)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})

	t.Run("WithIdentifier", func(t *testing.T) {
		license := openapi3.NewLicense("Apache 2.0").
			SetIdentifier("Apache-2.0")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "Apache 2.0", license.Name)
			assert.Equal(t, "Apache-2.0", license.Identifier)
			assert.Equal(t, "", license.URL)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"name":       "Apache 2.0",
				"identifier": "Apache-2.0",
			})
			gotBytes, gotErr := json.Marshal(license)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})

	t.Run("WithURL", func(t *testing.T) {
		license := openapi3.NewLicense("Apache 2.0").
			SetURL("https://www.apache.org/licenses/LICENSE-2.0")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "Apache 2.0", license.Name)
			assert.Equal(t, "", license.Identifier)
			assert.Equal(t, "https://www.apache.org/licenses/LICENSE-2.0", license.URL)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"name": "Apache 2.0",
				"url":  "https://www.apache.org/licenses/LICENSE-2.0",
			})
			gotBytes, gotErr := json.Marshal(license)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})
}
