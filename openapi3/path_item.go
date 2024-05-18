package openapi3

// PathItem represents a single path item in OpenAPI 3.1 specification.
//
// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#path-item-object
type PathItem struct {
	// A definition of a GET operation on this path.
	Get *Operation `json:"get,omitempty"`

	// A definition of a HEAD operation on this path.
	Head *Operation `json:"head,omitempty"`

	// A definition of a POST operation on this path.
	Post *Operation `json:"post,omitempty"`

	// A definition of a PUT operation on this path.
	Put *Operation `json:"put,omitempty"`

	// A definition of a PATCH operation on this path.
	Patch *Operation `json:"patch,omitempty"`

	// A definition of a DELETE operation on this path.
	Delete *Operation `json:"delete,omitempty"`

	// A definition of an OPTIONS operation on this path.
	Options *Operation `json:"options,omitempty"`

	// A definition of a TRACE operation on this path.
	Trace *Operation `json:"trace,omitempty"`
}

// NewPathItem creates a new PathItem object.
func NewPathItem() *PathItem {
	return &PathItem{
		Get:     nil,
		Head:    nil,
		Post:    nil,
		Put:     nil,
		Patch:   nil,
		Delete:  nil,
		Options: nil,
		Trace:   nil,
	}
}

// SetOperation sets the operation for the specified HTTP method.
func (p *PathItem) SetOperation(method string, operation *Operation) *PathItem {
	switch method {
	case "GET":
		p.Get = operation
	case "HEAD":
		p.Head = operation
	case "POST":
		p.Post = operation
	case "PUT":
		p.Put = operation
	case "PATCH":
		p.Patch = operation
	case "DELETE":
		p.Delete = operation
	case "OPTIONS":
		p.Options = operation
	case "TRACE":
		p.Trace = operation
	}

	return p
}
