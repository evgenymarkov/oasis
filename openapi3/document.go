package openapi3

// Document is the root object of the OpenAPI document.
//
// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#openapi-object
type Document struct {
	// OpenAPI is the version number of the OpenAPI Specification.
	//
	// Default: "3.1.0"
	OpenAPI string `json:"openapi"`

	// Info provides metadata about the API.
	//
	// Default: DocumentInfo{
	// 	Title:   "API",
	//	Version: "0.0.1",
	// }
	Info *DocumentInfo `json:"info"`

	// Additional meta information of the OpenAPI document.
	//
	// Default: []
	Tags []*Tag `json:"tags"`
}

// DocumentInfo provides metadata about the API.
//
// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#info-object
type DocumentInfo struct {
	// The title of the API.
	//
	// Default: "API"
	Title string `json:"title"`

	// The version of the OpenAPI document.
	//
	// Default: "0.0.1"
	Version string `json:"version"`
}

// NewDocument creates new OpenAPI document with default values.
func NewDocument() *Document {
	return &Document{
		OpenAPI: "3.1.0",
		Info: &DocumentInfo{
			Title:   "API",
			Version: "0.0.1",
		},
		Tags: make([]*Tag, 0),
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

// SetTags method sets additional meta information to OpenAPI document.
func (c *Document) SetTags(tags ...*Tag) *Document {
	c.Tags = append(c.Tags, tags...)

	return c
}
