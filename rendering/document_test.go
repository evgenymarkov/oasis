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
		SetVersion("1.0.0")
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
		"openapi": "3.1.0",
		"info": map[string]any{
			"title":   "Greeting API",
			"version": "1.0.0",
		},
	})
	require.NoError(t, wantErr)

	assert.JSONEq(t, string(wantBytes), string(gotBytes))
}
