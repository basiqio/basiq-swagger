// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubdivisionDetails subdivision details
//
// swagger:model SubdivisionDetails
type SubdivisionDetails struct {

	// Subdivision Code
	// Example: 45
	Code string `json:"code,omitempty"`

	// Subdivision Details
	// Example: Food and Beverage Services
	Title string `json:"title,omitempty"`
}

// Validate validates this subdivision details
func (m *SubdivisionDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this subdivision details based on context it is used
func (m *SubdivisionDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SubdivisionDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubdivisionDetails) UnmarshalBinary(b []byte) error {
	var res SubdivisionDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
