package openapi3_test

import (
	"encoding/json"
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDocumentDefaults(t *testing.T) {
	document := openapi3.NewDocument()

	assert.Equal(t, "3.1.0", document.OpenAPI)
	assert.Equal(t, "API", document.Info.Title)
	assert.Equal(t, "0.0.1", document.Info.Version)
	assert.Empty(t, document.Tags)
}

func TestDocumentOverrides(t *testing.T) {
	document := openapi3.NewDocument().
		SetTitle("Greeting API").
		SetVersion("1.0.0").
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

	assert.Equal(t, "3.1.0", document.OpenAPI)
	assert.Equal(t, "Greeting API", document.Info.Title)
	assert.Equal(t, "1.0.0", document.Info.Version)
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
}

func TestDocumentMarshaling(t *testing.T) {
	document := openapi3.NewDocument().
		SetTitle("Greeting API").
		SetVersion("1.0.0").
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

	wantBytes, wantErr := json.Marshal(map[string]any{
		"openapi": "3.1.0",
		"info": map[string]any{
			"title":   "Greeting API",
			"version": "1.0.0",
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
	wantStr := string(wantBytes)

	gotBytes, gotErr := json.Marshal(document)
	gotStr := string(gotBytes)

	require.NoError(t, wantErr)
	require.NoError(t, gotErr)
	assert.JSONEq(t, wantStr, gotStr)
}
