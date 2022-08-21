package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}

// New initializes a form struct, embeds url.Values object
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if form field is in post and is not empty.
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank.")
		return false
	}
	return true
}

// Valid returns true if there are no errors, otherwise false.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
