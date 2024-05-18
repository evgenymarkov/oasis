package openapi3_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertObjectSerialization(t *testing.T, expected map[string]any, document any) {
	t.Helper()

	wantBytes, wantErr := json.Marshal(expected)
	gotBytes, gotErr := json.Marshal(document)

	require.NoError(t, wantErr)
	require.NoError(t, gotErr)
	assert.JSONEq(t, string(wantBytes), string(gotBytes))
}
