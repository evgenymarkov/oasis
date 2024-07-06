package openapi3

// Parameter represents a parameter of OpenAPI operation.
type Parameter struct {
	// Name of the parameter.
	Name string `json:"name"`

	// Location of the parameter (e.g., query, header, path, cookie).
	In string `json:"in"`

	// A brief description of the parameter.
	Description string `json:"description,omitempty"`

	// Determines whether this parameter is mandatory.
	Required bool `json:"required,omitempty"`

	// Specifies that a parameter is deprecated and should be transitioned out of usage.
	Deprecated bool `json:"deprecated,omitempty"`
}

// NewParameter creates a new Parameter.
func NewParameter(name, in string) *Parameter {
	return &Parameter{
		Name:        name,
		In:          in,
		Description: "",
		Required:    in == "path",
		Deprecated:  false,
	}
}

// SetDescription sets the description of the parameter.
func (p *Parameter) SetDescription(description string) *Parameter {
	p.Description = description

	return p
}

// MarkAsRequired sets the parameter as required.
func (p *Parameter) MarkAsRequired() *Parameter {
	p.Required = true

	return p
}

// MarkAsDeprecated marks the parameter as deprecated.
func (p *Parameter) MarkAsDeprecated() *Parameter {
	p.Deprecated = true

	return p
}
