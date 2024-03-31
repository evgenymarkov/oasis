package openapi3

// ExternalDocumentation allows referencing an external resource for extended documentation.
//
// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#external-documentation-object
type ExternalDocumentation struct {
	// The URL for the target documentation.
	// This MUST be in the form of a URL.
	URL string `json:"url"`

	// A description of the target documentation.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
}

// NewExternalDocumentation creates new ExternalDocumentation.
func NewExternalDocumentation(url string) *ExternalDocumentation {
	return &ExternalDocumentation{
		URL:         url,
		Description: "",
	}
}

// SetDescription sets description for ExternalDocumentation.
func (t *ExternalDocumentation) SetDescription(description string) *ExternalDocumentation {
	t.Description = description

	return t
}
