package rendering

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/evgenymarkov/oasis/openapi3"
)

var errDocumentRender = errors.New("failed to render document")

func NewDocumentHandler(document *openapi3.Document) http.HandlerFunc {
	var (
		documentBytes     []byte
		documentRenderErr error
	)

	return func(response http.ResponseWriter, _ *http.Request) {
		if documentBytes == nil && documentRenderErr == nil {
			documentBytes, documentRenderErr = json.Marshal(document)
		}

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
