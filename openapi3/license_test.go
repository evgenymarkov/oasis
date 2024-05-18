package openapi3_test

import (
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
)

func TestLicense(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		license := openapi3.NewLicense("Apache 2.0")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "Apache 2.0", license.Name)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{"name": "Apache 2.0"},
				license,
			)
		})
	})

	t.Run("WithIdentifier", func(t *testing.T) {
		license := openapi3.NewLicense("Apache 2.0").
			SetIdentifier("Apache-2.0")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "Apache 2.0", license.Name)
			assert.Equal(t, "Apache-2.0", license.Identifier)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"name":       "Apache 2.0",
					"identifier": "Apache-2.0",
				},
				license,
			)
		})
	})

	t.Run("WithURL", func(t *testing.T) {
		license := openapi3.NewLicense("Apache 2.0").
			SetURL("https://www.apache.org/licenses/LICENSE-2.0")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "Apache 2.0", license.Name)
			assert.Equal(t, "https://www.apache.org/licenses/LICENSE-2.0", license.URL)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"name": "Apache 2.0",
					"url":  "https://www.apache.org/licenses/LICENSE-2.0",
				},
				license,
			)
		})
	})
}
