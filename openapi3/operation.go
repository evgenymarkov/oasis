package openapi3

// Operation represents a single API operation on a path.
type Operation struct {
	// Unique string used to identify the operation.
	OperationID string `json:"operationId"`

	// A short summary of what the operation does.
	Summary string `json:"summary,omitempty"`

	// A verbose explanation of the operation behavior.
	Description string `json:"description,omitempty"`

	// Declares this operation to be deprecated.
	Deprecated bool `json:"deprecated,omitempty"`

	// Additional external documentation for this operation.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`

	// A list of tags for API documentation control.
	Tags []string `json:"tags,omitempty"`
}

// NewOperation creates a new Operation with a required operation ID.
func NewOperation(operationID string) *Operation {
	return &Operation{
		OperationID:  operationID,
		Summary:      "",
		Description:  "",
		Deprecated:   false,
		ExternalDocs: nil,
		Tags:         make([]string, 0),
	}
}

// SetSummary sets the summary.
func (o *Operation) SetSummary(summary string) *Operation {
	o.Summary = summary

	return o
}

// SetDescription sets the description.
func (o *Operation) SetDescription(description string) *Operation {
	o.Description = description

	return o
}

// MarkAsDeprecated marks the operation as deprecated.
func (o *Operation) MarkAsDeprecated() *Operation {
	o.Deprecated = true

	return o
}

// SetExternalDocs sets the external documentation.
func (o *Operation) SetExternalDocs(externalDocs *ExternalDocumentation) *Operation {
	o.ExternalDocs = externalDocs

	return o
}

// SetTags sets the tags.
func (o *Operation) SetTags(tags ...string) *Operation {
	o.Tags = tags

	return o
}
