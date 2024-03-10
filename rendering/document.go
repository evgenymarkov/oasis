package rendering

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/evgenymarkov/oasis/openapi3"
)

var errDocumentMarshal = errors.New("failed to marshal document")

func NewDocumentHandler(document *openapi3.Document) http.HandlerFunc {
	var (
		marshalErr    error
		documentBytes []byte
	)

	return func(response http.ResponseWriter, _ *http.Request) {
		if marshalErr == nil && documentBytes == nil {
			documentBytes, marshalErr = json.Marshal(document)
		}

		if marshalErr != nil {
			message := errors.Join(errDocumentMarshal, marshalErr).Error()
			http.Error(response, message, http.StatusInternalServerError)

			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		_, _ = response.Write(documentBytes)
	}
}
