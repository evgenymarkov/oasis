package openapi3_test

import (
	"testing"

	"github.com/evgenymarkov/oasis/openapi3"
	"github.com/stretchr/testify/assert"
)

const (
	jsonSchemaDialect = "https://spec.openapis.org/oas/3.1/dialect/base"
)

func TestDocument(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		document := openapi3.NewDocument()

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "3.1.0", document.OpenAPI)
			assert.Equal(t, jsonSchemaDialect, document.JSONSchemaDialect)
			assert.Equal(t, "API", document.Info.Title)
			assert.Equal(t, "0.0.1", document.Info.Version)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"openapi":           "3.1.0",
					"jsonSchemaDialect": jsonSchemaDialect,
					"info": map[string]any{
						"title":   "API",
						"version": "0.0.1",
					},
					"servers": []any{},
					"paths":   map[string]any{},
					"tags":    []any{},
				},
				document,
			)
		})
	})

	t.Run("WithCustomTitle", func(t *testing.T) {
		document := openapi3.NewDocument().
			SetTitle("Greeting API")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "Greeting API", document.Info.Title)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"openapi":           "3.1.0",
					"jsonSchemaDialect": jsonSchemaDialect,
					"info": map[string]any{
						"title":   "Greeting API",
						"version": "0.0.1",
					},
					"servers": []any{},
					"paths":   map[string]any{},
					"tags":    []any{},
				},
				document,
			)
		})
	})

	t.Run("WithCustomVersion", func(t *testing.T) {
		document := openapi3.NewDocument().
			SetVersion("1.0.0")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "1.0.0", document.Info.Version)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"openapi":           "3.1.0",
					"jsonSchemaDialect": jsonSchemaDialect,
					"info": map[string]any{
						"title":   "API",
						"version": "1.0.0",
					},
					"servers": []any{},
					"paths":   map[string]any{},
					"tags":    []any{},
				},
				document,
			)
		})
	})

	t.Run("WithSummary", func(t *testing.T) {
		document := openapi3.NewDocument().
			SetSummary("API for greetings")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "API for greetings", document.Info.Summary)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"openapi":           "3.1.0",
					"jsonSchemaDialect": jsonSchemaDialect,
					"info": map[string]any{
						"title":   "API",
						"summary": "API for greetings",
						"version": "0.0.1",
					},
					"servers": []any{},
					"paths":   map[string]any{},
					"tags":    []any{},
				},
				document,
			)
		})
	})

	t.Run("WithDescription", func(t *testing.T) {
		document := openapi3.NewDocument().
			SetDescription("_Oasis_ is a library for Go web apps")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "_Oasis_ is a library for Go web apps", document.Info.Description)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"openapi":           "3.1.0",
					"jsonSchemaDialect": jsonSchemaDialect,
					"info": map[string]any{
						"title":       "API",
						"description": "_Oasis_ is a library for Go web apps",
						"version":     "0.0.1",
					},
					"servers": []any{},
					"paths":   map[string]any{},
					"tags":    []any{},
				},
				document,
			)
		})
	})

	t.Run("WithTermsOfService", func(t *testing.T) {
		document := openapi3.NewDocument().
			SetTermsOfService("https://example.com/legal/rules/")

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, "https://example.com/legal/rules/", document.Info.TermsOfService)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"openapi":           "3.1.0",
					"jsonSchemaDialect": jsonSchemaDialect,
					"info": map[string]any{
						"title":          "API",
						"termsOfService": "https://example.com/legal/rules/",
						"version":        "0.0.1",
					},
					"servers": []any{},
					"paths":   map[string]any{},
					"tags":    []any{},
				},
				document,
			)
		})
	})

	t.Run("WithContact", func(t *testing.T) {
		document := openapi3.NewDocument().
			SetContact(
				openapi3.NewContact("API Support").
					SetURL("https://example.com/support").
					SetEmail("greeting-api@support.example.com"),
			)

		t.Run("Values", func(t *testing.T) {
			assert.Equal(
				t,
				&openapi3.Contact{
					Name:  "API Support",
					URL:   "https://example.com/support",
					Email: "greeting-api@support.example.com",
				},
				document.Info.Contact,
			)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"openapi":           "3.1.0",
					"jsonSchemaDialect": jsonSchemaDialect,
					"info": map[string]any{
						"title":   "API",
						"version": "0.0.1",
						"contact": map[string]any{
							"name":  "API Support",
							"url":   "https://example.com/support",
							"email": "greeting-api@support.example.com",
						},
					},
					"servers": []any{},
					"paths":   map[string]any{},
					"tags":    []any{},
				},
				document,
			)
		})
	})

	t.Run("WithLicense", func(t *testing.T) {
		document := openapi3.NewDocument().
			SetLicense(
				openapi3.NewLicense("MIT").
					SetIdentifier("MIT"),
			)

		t.Run("Values", func(t *testing.T) {
			assert.Equal(
				t,
				&openapi3.License{
					Name:       "MIT",
					Identifier: "MIT",
					URL:        "",
				},
				document.Info.License,
			)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"openapi":           "3.1.0",
					"jsonSchemaDialect": jsonSchemaDialect,
					"info": map[string]any{
						"title":   "API",
						"version": "0.0.1",
						"license": map[string]any{
							"name":       "MIT",
							"identifier": "MIT",
						},
					},
					"servers": []any{},
					"paths":   map[string]any{},
					"tags":    []any{},
				},
				document,
			)
		})
	})

	t.Run("WithExternalDocs", func(t *testing.T) {
		wikiHelloURL := "https://wikipedia.org/wiki/Hello,_world!"
		wikiHelloDescription := "Free online encyclopedia, created and edited by volunteers"

		document := openapi3.NewDocument().
			SetExternalDocs(
				openapi3.NewExternalDocumentation(wikiHelloURL).
					SetDescription(wikiHelloDescription),
			)

		t.Run("Values", func(t *testing.T) {
			assert.Equal(
				t,
				&openapi3.ExternalDocumentation{
					URL:         wikiHelloURL,
					Description: wikiHelloDescription,
				},
				document.ExternalDocs,
			)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"openapi":           "3.1.0",
					"jsonSchemaDialect": jsonSchemaDialect,
					"info": map[string]any{
						"title":   "API",
						"version": "0.0.1",
					},
					"externalDocs": map[string]any{
						"url":         wikiHelloURL,
						"description": wikiHelloDescription,
					},
					"servers": []any{},
					"paths":   map[string]any{},
					"tags":    []any{},
				},
				document,
			)
		})
	})

	t.Run("WithServers", func(t *testing.T) {
		server1 := openapi3.NewServer("https://example.com").
			SetDescription("Production")
		server2 := openapi3.NewServer("https://test.example.com").
			SetDescription("Testing")
		document := openapi3.NewDocument().
			SetServers(server1, server2)

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, []*openapi3.Server{server1, server2}, document.Servers)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"openapi":           "3.1.0",
					"jsonSchemaDialect": jsonSchemaDialect,
					"info": map[string]any{
						"title":   "API",
						"version": "0.0.1",
					},
					"servers": []any{
						map[string]any{
							"url":         "https://example.com",
							"description": "Production",
						},
						map[string]any{
							"url":         "https://test.example.com",
							"description": "Testing",
						},
					},
					"paths": map[string]any{},
					"tags":  []any{},
				},
				document,
			)
		})
	})

	t.Run("WithPaths", func(t *testing.T) {
		getOperation := openapi3.NewOperation("GetOperation").
			SetSummary("Get operation")
		postOperation := openapi3.NewOperation("SaveOperation").
			SetSummary("Post operation")
		document := openapi3.NewDocument().
			AddOperation("/example", "GET", getOperation).
			AddOperation("/example", "POST", postOperation)

		t.Run("Values", func(t *testing.T) {
			assert.Equal(t, getOperation, document.Paths["/example"].Get)
			assert.Equal(t, postOperation, document.Paths["/example"].Post)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"openapi":           "3.1.0",
					"jsonSchemaDialect": jsonSchemaDialect,
					"info": map[string]any{
						"title":   "API",
						"version": "0.0.1",
					},
					"servers": []any{},
					"paths": map[string]any{
						"/example": map[string]any{
							"get": map[string]any{
								"operationId": "GetOperation",
								"summary":     "Get operation",
							},
							"post": map[string]any{
								"operationId": "SaveOperation",
								"summary":     "Post operation",
							},
						},
					},
					"tags": []any{},
				},
				document,
			)
		})
	})

	t.Run("WithTags", func(t *testing.T) {
		document := openapi3.NewDocument().
			SetTags(
				openapi3.NewTag("orders").
					SetDescription("Orders operations").
					SetExternalDocs(
						openapi3.NewExternalDocumentation("https://market.example.com").
							SetDescription("E-commerce platform"),
					),
				openapi3.NewTag("payments").
					SetDescription("Payments operations").
					SetExternalDocs(
						openapi3.NewExternalDocumentation("https://bank.example.com").
							SetDescription("New fancy digital bank"),
					),
			)

		t.Run("Values", func(t *testing.T) {
			assert.Equal(
				t,
				[]*openapi3.Tag{
					openapi3.NewTag("orders").
						SetDescription("Orders operations").
						SetExternalDocs(
							openapi3.NewExternalDocumentation("https://market.example.com").
								SetDescription("E-commerce platform"),
						),
					openapi3.NewTag("payments").
						SetDescription("Payments operations").
						SetExternalDocs(
							openapi3.NewExternalDocumentation("https://bank.example.com").
								SetDescription("New fancy digital bank"),
						),
				},
				document.Tags,
			)
		})

		t.Run("Serialization", func(t *testing.T) {
			assertObjectSerialization(
				t,
				map[string]any{
					"openapi":           "3.1.0",
					"jsonSchemaDialect": jsonSchemaDialect,
					"info": map[string]any{
						"title":   "API",
						"version": "0.0.1",
					},
					"servers": []any{},
					"paths":   map[string]any{},
					"tags": []any{
						map[string]any{
							"name":        "orders",
							"description": "Orders operations",
							"externalDocs": map[string]any{
								"url":         "https://market.example.com",
								"description": "E-commerce platform",
							},
						},
						map[string]any{
							"name":        "payments",
							"description": "Payments operations",
							"externalDocs": map[string]any{
								"url":         "https://bank.example.com",
								"description": "New fancy digital bank",
							},
						},
					},
				},
				document,
			)
		})
	})
}
