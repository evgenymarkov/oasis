package openapi3

// ServerVariable represents a variable for server URL template substitution.
//
// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#server-variable-object
type ServerVariable struct {
	// An enumeration of string values to be used if the substitution options are from a limited set.
	Enum []string `json:"enum,omitempty"`

	// The default value to use for substitution,
	// which SHALL be sent if an alternate value is not supplied.
	Default string `json:"default"`

	// An optional description for the server variable.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
}

// NewServerVariable creates a new ServerVariable object.
func NewServerVariable(defaultValue string) *ServerVariable {
	return &ServerVariable{
		Enum:        nil,
		Default:     defaultValue,
		Description: "",
	}
}

// SetEnum sets the enumeration values for the ServerVariable.
func (s *ServerVariable) SetEnum(enumValues []string) *ServerVariable {
	s.Enum = enumValues

	return s
}

// SetDescription sets the description for the ServerVariable.
func (s *ServerVariable) SetDescription(description string) *ServerVariable {
	s.Description = description

	return s
}
