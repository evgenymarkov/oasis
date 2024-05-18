package rendering_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/evgenymarkov/oasis/rendering"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDocumentHandler(t *testing.T) {
	document := openapi3.NewDocument().
		SetTitle("Greeting API").
		SetVersion("1.0.0").
		SetTags(
			openapi3.NewTag("orders").
				SetDescription("Orders operations").
				SetExternalDocs(
					openapi3.NewExternalDocumentation("https://market.example.com").
						SetDescription("E-commerce platform"),
				),
			openapi3.NewTag("payments").
				SetDescription("Payments operations").
				SetExternalDocs(
					openapi3.NewExternalDocumentation("https://bank.example.com").
						SetDescription("New fancy digital bank"),
				),
		)

	documentHandler := rendering.NewDocumentHandler(document)

	testServer := httptest.NewServer(documentHandler)
	defer testServer.Close()

	request, requestErr := http.NewRequestWithContext(
		context.TODO(),
		http.MethodGet,
		testServer.URL,
		nil,
	)
	require.NoError(t, requestErr)

	response, responseErr := http.DefaultClient.Do(request)
	require.NoError(t, responseErr)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))

	gotBytes, readErr := io.ReadAll(response.Body)
	response.Body.Close()
	require.NoError(t, readErr)

	wantBytes, wantErr := json.Marshal(map[string]any{
		"openapi":           "3.1.0",
		"jsonSchemaDialect": "https://spec.openapis.org/oas/3.1/dialect/base",
		"info": map[string]any{
			"title":   "Greeting API",
			"version": "1.0.0",
		},
		"servers": []any{},
		"tags": []any{
			map[string]any{
				"name":        "orders",
				"description": "Orders operations",
				"externalDocs": map[string]any{
					"url":         "https://market.example.com",
					"description": "E-commerce platform",
				},
			},
			map[string]any{
				"name":        "payments",
				"description": "Payments operations",
				"externalDocs": map[string]any{
					"url":         "https://bank.example.com",
					"description": "New fancy digital bank",
				},
			},
		},
	})
	require.NoError(t, wantErr)

	assert.JSONEq(t, string(wantBytes), string(gotBytes))
}
