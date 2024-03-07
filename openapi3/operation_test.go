package openapi3_test

import (
	"encoding/json"
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	wantBytes, wantErr := json.Marshal(map[string]any{
		"operationId": "get-greeting",
		"summary":     "Get greeting",
	})
	wantStr := string(wantBytes)

	gotBytes, gotErr := json.Marshal(operation)
	gotStr := string(gotBytes)

	require.NoError(t, wantErr)
	require.NoError(t, gotErr)
	assert.JSONEq(t, wantStr, gotStr)
}
