package openapi3

// License information for the exposed API.
//
// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#license-object
type License struct {
	// The license name used for the API.
	Name string `json:"name"`

	// An SPDX license expression for the API.
	// The identifier field is mutually exclusive of the url field.
	Identifier string `json:"identifier,omitempty"`

	// A URL to the license used for the API.
	// This MUST be in the form of a URL.
	// The url field is mutually exclusive of the identifier field.
	URL string `json:"url,omitempty"`
}

// NewLicense creates new License.
func NewLicense(name string) *License {
	return &License{
		Name:       name,
		Identifier: "",
		URL:        "",
	}
}

// SetIdentifier sets an SPDX license identifier.
// The identifier field is mutually exclusive of the url field.
func (l *License) SetIdentifier(identifier string) *License {
	l.Identifier = identifier

	return l
}

// SetIdentifier sets a URL to the license text.
// The url field is mutually exclusive of the identifier field.
func (l *License) SetURL(url string) *License {
	l.URL = url

	return l
}
