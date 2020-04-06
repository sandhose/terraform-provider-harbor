// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RetentionSelector retention selector
// swagger:model RetentionSelector
type RetentionSelector struct {

	// decoration
	Decoration string `json:"decoration,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`

	// pattern
	Pattern string `json:"pattern,omitempty"`
}

// Validate validates this retention selector
func (m *RetentionSelector) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RetentionSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetentionSelector) UnmarshalBinary(b []byte) error {
	var res RetentionSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
