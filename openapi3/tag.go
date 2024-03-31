package openapi3

// Tag is an additional metadata for document and operations.
// Tags may be used by some tools that works with OpenAPI documents.
//
// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#tag-object
type Tag struct {
	// The name of the tag.
	Name string `json:"name"`

	// A description for the tag.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`

	// Additional external documentation for this tag.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
}

// NewTag creates new Tag.
func NewTag(name string) *Tag {
	return &Tag{
		Name:         name,
		Description:  "",
		ExternalDocs: nil,
	}
}

// SetDescription sets description for Tag.
func (t *Tag) SetDescription(description string) *Tag {
	t.Description = description

	return t
}

// SetExternalDocs sets external documentation link for Tag.
func (t *Tag) SetExternalDocs(externalDocs *ExternalDocumentation) *Tag {
	t.ExternalDocs = externalDocs

	return t
}
