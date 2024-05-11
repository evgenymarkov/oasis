package openapi3_test

import (
	"encoding/json"
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDocument(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		document := openapi3.NewDocument()

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "3.1.0", document.OpenAPI)
			assert.Equal(t, "API", document.Info.Title)
			assert.Equal(t, "0.0.1", document.Info.Version)
			assert.Empty(t, document.Tags)
			assert.Nil(t, document.ExternalDocs)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"openapi": "3.1.0",
				"info": map[string]any{
					"title":   "API",
					"version": "0.0.1",
				},
				"tags": []any{},
			})
			gotBytes, gotErr := json.Marshal(document)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})

	t.Run("WithCustomTitle", func(t *testing.T) {
		document := openapi3.NewDocument().
			SetTitle("Greeting API")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "3.1.0", document.OpenAPI)
			assert.Equal(t, "Greeting API", document.Info.Title)
			assert.Equal(t, "0.0.1", document.Info.Version)
			assert.Empty(t, document.Tags)
			assert.Nil(t, document.ExternalDocs)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"openapi": "3.1.0",
				"info": map[string]any{
					"title":   "Greeting API",
					"version": "0.0.1",
				},
				"tags": []any{},
			})
			gotBytes, gotErr := json.Marshal(document)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})

	t.Run("WithCustomVersion", func(t *testing.T) {
		document := openapi3.NewDocument().
			SetVersion("1.0.0")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "3.1.0", document.OpenAPI)
			assert.Equal(t, "API", document.Info.Title)
			assert.Equal(t, "1.0.0", document.Info.Version)
			assert.Empty(t, document.Tags)
			assert.Nil(t, document.ExternalDocs)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"openapi": "3.1.0",
				"info": map[string]any{
					"title":   "API",
					"version": "1.0.0",
				},
				"tags": []any{},
			})
			gotBytes, gotErr := json.Marshal(document)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})

	t.Run("WithTags", func(t *testing.T) {
		document := openapi3.NewDocument().
			SetTags(
				openapi3.NewTag("orders").
					SetDescription("Orders operations").
					SetExternalDocs(
						openapi3.NewExternalDocumentation("https://market.yandex.ru").
							SetDescription("E-commerce platform"),
					),
				openapi3.NewTag("payments").
					SetDescription("Payments operations").
					SetExternalDocs(
						openapi3.NewExternalDocumentation("https://bank.yandex.ru").
							SetDescription("New fancy digital bank"),
					),
			)

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "3.1.0", document.OpenAPI)
			assert.Equal(t, "API", document.Info.Title)
			assert.Equal(t, "0.0.1", document.Info.Version)
			assert.Equal(
				t,
				[]*openapi3.Tag{
					openapi3.NewTag("orders").
						SetDescription("Orders operations").
						SetExternalDocs(
							openapi3.NewExternalDocumentation("https://market.yandex.ru").
								SetDescription("E-commerce platform"),
						),
					openapi3.NewTag("payments").
						SetDescription("Payments operations").
						SetExternalDocs(
							openapi3.NewExternalDocumentation("https://bank.yandex.ru").
								SetDescription("New fancy digital bank"),
						),
				},
				document.Tags,
			)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"openapi": "3.1.0",
				"info": map[string]any{
					"title":   "API",
					"version": "0.0.1",
				},
				"tags": []any{
					map[string]any{
						"name":        "orders",
						"description": "Orders operations",
						"externalDocs": map[string]any{
							"url":         "https://market.yandex.ru",
							"description": "E-commerce platform",
						},
					},
					map[string]any{
						"name":        "payments",
						"description": "Payments operations",
						"externalDocs": map[string]any{
							"url":         "https://bank.yandex.ru",
							"description": "New fancy digital bank",
						},
					},
				},
			})
			gotBytes, gotErr := json.Marshal(document)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})

	t.Run("WithExternalDocs", func(t *testing.T) {
		wikiHelloURL := "https://wikipedia.org/wiki/Hello,_world!"
		wikiHelloDescription := "Free online encyclopedia, created and edited by volunteers"

		document := openapi3.NewDocument().
			SetExternalDocs(
				openapi3.NewExternalDocumentation(wikiHelloURL).
					SetDescription(wikiHelloDescription),
			)

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "3.1.0", document.OpenAPI)
			assert.Equal(t, "API", document.Info.Title)
			assert.Equal(t, "0.0.1", document.Info.Version)
			assert.Empty(t, document.Tags)
			assert.Equal(
				t,
				&openapi3.ExternalDocumentation{
					URL:         wikiHelloURL,
					Description: wikiHelloDescription,
				},
				document.ExternalDocs,
			)
		})

		t.Run("Serialization", func(t *testing.T) {
			wantBytes, wantErr := json.Marshal(map[string]any{
				"openapi": "3.1.0",
				"info": map[string]any{
					"title":   "API",
					"version": "0.0.1",
				},
				"tags": []any{},
				"externalDocs": map[string]any{
					"url":         wikiHelloURL,
					"description": wikiHelloDescription,
				},
			})
			gotBytes, gotErr := json.Marshal(document)

			require.NoError(t, wantErr)
			require.NoError(t, gotErr)
			assert.JSONEq(t, string(wantBytes), string(gotBytes))
		})
	})
}
