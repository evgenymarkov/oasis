package openapi3

// Contact provides contact information for the exposed API.
//
// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#contact-object
type Contact struct {
	// The identifying name of the contact person/organization.
	Name string `json:"name"`

	// The URL pointing to the contact information.
	// This MUST be in the form of a URL.
	URL string `json:"url,omitempty"`

	// The email address of the contact person/organization.
	// This MUST be in the form of an email address.
	Email string `json:"email,omitempty"`
}

// NewContact creates new Contact.
func NewContact(name string) *Contact {
	return &Contact{
		Name:  name,
		URL:   "",
		Email: "",
	}
}

// SetURL sets the URL pointing to the contact information.
func (c *Contact) SetURL(url string) *Contact {
	c.URL = url

	return c
}

// SetEmail sets the email address of the contact person/organization.
func (c *Contact) SetEmail(email string) *Contact {
	c.Email = email

	return c
}
