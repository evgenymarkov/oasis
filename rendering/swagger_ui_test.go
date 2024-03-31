package rendering_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/evgenymarkov/oasis/rendering"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var staticFiles = []string{
	"favicon-32x32.png",
	"index.css",
	"oauth2-redirect.html",
	"swagger-ui-bundle.js",
	"swagger-ui-bundle.js.map",
	"swagger-ui.css",
	"swagger-ui.css.map",
}

func TestSwaggerUIHandler(t *testing.T) {
	document := openapi3.NewDocument().
		SetTitle("Greeting API").
		SetVersion("1.0.0")

	request, requestErr := http.NewRequestWithContext(
		context.TODO(),
		http.MethodGet,
		"/api",
		nil,
	)
	require.NoError(t, requestErr)

	response := httptest.NewRecorder()
	handler := rendering.NewSwaggerUIHandler(rendering.SwaggerUIConfig{
		BaseURL:   "/api",
		PageTitle: "API Docs",
		Document:  document,
	})
	handler.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "text/html", response.Header().Get("Content-Type"))
}

func TestSwaggerUIStaticFiles(t *testing.T) {
	assert.ElementsMatch(t, staticFiles, rendering.GetSwaggerUIStaticFiles())
}

func TestSwaggerUIStaticHandler(t *testing.T) {
	for _, file := range staticFiles {
		t.Run(file, func(t *testing.T) {
			request, requestErr := http.NewRequestWithContext(
				context.TODO(),
				http.MethodGet,
				"/"+file,
				nil,
			)
			require.NoError(t, requestErr)

			response := httptest.NewRecorder()
			handler := rendering.NewSwaggerUIStaticHandler(file)
			handler.ServeHTTP(response, request)

			assert.Equal(t, http.StatusOK, response.Code)
		})
	}
}
