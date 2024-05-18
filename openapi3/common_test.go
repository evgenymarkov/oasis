package openapi3_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertObjectSerialization(t *testing.T, wantMap map[string]any, gotObject any) {
	t.Helper()

	wantBytes, wantErr := json.Marshal(wantMap)
	gotBytes, gotErr := json.Marshal(gotObject)

	require.NoError(t, wantErr)
	require.NoError(t, gotErr)
	assert.JSONEq(t, string(wantBytes), string(gotBytes))
}
