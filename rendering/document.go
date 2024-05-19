package rendering

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"

	"github.com/evgenymarkov/oasis/openapi3"
)

var errDocumentRender = errors.New("failed to render document")

func NewDocumentHandler(document *openapi3.Document) http.HandlerFunc {
	var (
		documentBytes     []byte
		documentRenderErr error
		renderOnce        sync.Once
	)

	return func(response http.ResponseWriter, _ *http.Request) {
		renderOnce.Do(func() {
			documentBytes, documentRenderErr = json.Marshal(document)
		})

		if documentRenderErr != nil {
			message := errors.Join(errDocumentRender, documentRenderErr).Error()
			http.Error(response, message, http.StatusInternalServerError)

			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		_, _ = response.Write(documentBytes)
	}
}
