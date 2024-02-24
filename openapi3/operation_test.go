package openapi3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperationDefaults(t *testing.T) {
	operation := NewOperation()

	assert.Equal(t, "", operation.OperationID)
	assert.Equal(t, "", operation.Summary)
}

func TestOperationOverrides(t *testing.T) {
	operation := NewOperation().
		SetOperationID("get-greeting").
		SetSummary("Get greeting")

	assert.Equal(t, "get-greeting", operation.OperationID)
	assert.Equal(t, "Get greeting", operation.Summary)
}
