package openapi3_test

import (
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
)

func TestOperationDefaults(t *testing.T) {
	operation := openapi3.NewOperation()

	assert.Equal(t, "", operation.OperationID)
	assert.Equal(t, "", operation.Summary)
}

func TestOperationOverrides(t *testing.T) {
	operation := openapi3.NewOperation().
		SetOperationID("get-greeting").
		SetSummary("Get greeting")

	assert.Equal(t, "get-greeting", operation.OperationID)
	assert.Equal(t, "Get greeting", operation.Summary)
}

func TestOperationMarshaling(t *testing.T) {
	operation := openapi3.NewOperation().
		SetOperationID("get-greeting").
		SetSummary("Get greeting")

	assertObjectSerialization(
		t,
		map[string]any{
			"operationId": "get-greeting",
			"summary":     "Get greeting",
		},
		operation,
	)
}
