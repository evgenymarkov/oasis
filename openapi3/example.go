package openapi3

// Example represents an example object in OpenAPI.
//
// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#example-object
type Example struct {
	// Value is the value of the example.
	Value any `json:"value"`

	// Summary is a short description of the example.
	Summary string `json:"summary,omitempty"`

	// Description is a long description of the example.
	Description string `json:"description,omitempty"`
}

// NewExample creates new Example.
func NewExample(value any) *Example {
	return &Example{
		Value:       value,
		Summary:     "",
		Description: "",
	}
}

// SetSummary sets the short description of the example.
func (e *Example) SetSummary(summary string) *Example {
	e.Summary = summary

	return e
}

// SetDescription sets the long description of the example.
func (e *Example) SetDescription(description string) *Example {
	e.Description = description

	return e
}
