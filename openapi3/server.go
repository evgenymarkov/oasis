package openapi3

// Server represents a server object in the OpenAPI 3.1 specification.
//
// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#server-object
type Server struct {
	// URL to the target host. This URL supports Server Variables and MAY be relative,
	// to indicate that the host location is relative to the location where the OpenAPI
	// document is being served. Variable substitutions will be made when a variable is
	// named in {brackets}.
	URL string `json:"url"`

	// An optional string describing the host designated by the URL.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`

	// A map between a variable name and its value. The value is used for substitution
	// in the server's URL template.
	Variables map[string]*ServerVariable `json:"variables,omitempty"`
}

// NewServer creates a new Server object.
func NewServer(url string) *Server {
	return &Server{
		URL:         url,
		Description: "",
		Variables:   make(map[string]*ServerVariable),
	}
}

// SetDescription sets the description for the Server.
func (s *Server) SetDescription(description string) *Server {
	s.Description = description

	return s
}

// AddVariable adds a variable to the Server.
func (s *Server) AddVariable(name string, variable *ServerVariable) *Server {
	s.Variables[name] = variable

	return s
}
