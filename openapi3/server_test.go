package openapi3_test

import (
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		server := openapi3.NewServer("https://example.com")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "https://example.com", server.URL)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{"url": "https://example.com"},
				server,
			)
		})
	})

	t.Run("WithDescription", func(t *testing.T) {
		server := openapi3.NewServer("https://example.com").
			SetDescription("Server with description")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "https://example.com", server.URL)
			assert.Equal(t, "Server with description", server.Description)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"url":         "https://example.com",
					"description": "Server with description",
				},
				server,
			)
		})
	})

	t.Run("WithVariables", func(t *testing.T) {
		server := openapi3.NewServer("https://{username}.example.com").
			AddVariable("username", openapi3.NewServerVariable("user1").
				SetDescription("Username for the server").
				SetEnum([]string{"user1", "user2"})).
			AddVariable("port", openapi3.NewServerVariable("8080").
				SetDescription("Port for the server").
				SetEnum([]string{"8080", "9090"}))

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "https://{username}.example.com", server.URL)
			assert.Contains(t, server.Variables, "username")
			assert.Contains(t, server.Variables, "port")
			assert.Equal(t, "user1", server.Variables["username"].Default)
			assert.Equal(t, "Username for the server", server.Variables["username"].Description)
			assert.ElementsMatch(t, []string{"user1", "user2"}, server.Variables["username"].Enum)
			assert.Equal(t, "8080", server.Variables["port"].Default)
			assert.Equal(t, "Port for the server", server.Variables["port"].Description)
			assert.ElementsMatch(t, []string{"8080", "9090"}, server.Variables["port"].Enum)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"url": "https://{username}.example.com",
					"variables": map[string]any{
						"username": map[string]any{
							"default":     "user1",
							"description": "Username for the server",
							"enum":        []string{"user1", "user2"},
						},
						"port": map[string]any{
							"default":     "8080",
							"description": "Port for the server",
							"enum":        []string{"8080", "9090"},
						},
					},
				},
				server,
			)
		})
	})
}
