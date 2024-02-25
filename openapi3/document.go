package openapi3

// Document is the root object of the OpenAPI document.
//
// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#openapi-object
type Document struct {
	// OpenAPI is the version number of the OpenAPI Specification.
	//
	// Default: "3.1.0"
	OpenAPI string

	// Info provides metadata about the API.
	//
	// Default: DocumentInfo{
	// 	Title:   "API",
	//	Version: "0.0.1",
	// }
	Info DocumentInfo
}

// DocumentInfo provides metadata about the API.
//
// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#info-object
type DocumentInfo struct {
	// The title of the API.
	//
	// Default: "API"
	Title string

	// The version of the OpenAPI document.
	//
	// Default: "0.0.1"
	Version string
}

// NewDocument creates new OpenAPI document with default values.
func NewDocument() *Document {
	return &Document{
		OpenAPI: "3.1.0",
		Info: DocumentInfo{
			Title:   "API",
			Version: "0.0.1",
		},
	}
}

// SetTitle method sets OpenAPI document title.
func (c *Document) SetTitle(title string) *Document {
	c.Info.Title = title

	return c
}

// SetVersion method sets OpenAPI document version.
func (c *Document) SetVersion(version string) *Document {
	c.Info.Version = version

	return c
}
