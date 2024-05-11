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

	// Additional external documentation for the OpenAPI document.
	//
	// Default: nil
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`

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

	// A short summary of the API.
	//
	// Default: ""
	Summary string `json:"summary,omitempty"`

	// A description of the API.
	// CommonMark syntax MAY be used for rich text representation.
	//
	// Default: ""
	Description string `json:"description,omitempty"`

	// A URL to the Terms of Service for the API.
	// This MUST be in the form of a URL.
	//
	// Default: ""
	TermsOfService string `json:"termsOfService,omitempty"`

	// The contact information for the exposed API.
	//
	// Default: ""
	Contact *Contact `json:"contact,omitempty"`
}

// NewDocument creates new OpenAPI document with default values.
func NewDocument() *Document {
	return &Document{
		OpenAPI: "3.1.0",
		Info: &DocumentInfo{
			Title:          "API",
			Version:        "0.0.1",
			Summary:        "",
			Description:    "",
			TermsOfService: "",
			Contact:        nil,
		},
		ExternalDocs: nil,
		Tags:         make([]*Tag, 0),
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

// SetSummary method sets OpenAPI document summary.
func (c *Document) SetSummary(summary string) *Document {
	c.Info.Summary = summary

	return c
}

// SetDescription method sets OpenAPI document description.
func (c *Document) SetDescription(description string) *Document {
	c.Info.Description = description

	return c
}

// SetTermsOfService method sets Terms of Service for the API.
func (c *Document) SetTermsOfService(termsURL string) *Document {
	c.Info.TermsOfService = termsURL

	return c
}

// SetContact method sets contact information for the exposed API.
func (c *Document) SetContact(contact *Contact) *Document {
	c.Info.Contact = contact

	return c
}

// SetExternalDocs sets external documentation link to OpenAPI document.
func (c *Document) SetExternalDocs(externalDocs *ExternalDocumentation) *Document {
	c.ExternalDocs = externalDocs

	return c
}

// SetTags method sets additional meta information to OpenAPI document.
func (c *Document) SetTags(tags ...*Tag) *Document {
	c.Tags = append(c.Tags, tags...)

	return c
}
