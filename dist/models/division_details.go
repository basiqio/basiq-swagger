// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DivisionDetails division details
//
// swagger:model DivisionDetails
type DivisionDetails struct {

	// Division Code
	// Example: H
	Code string `json:"code,omitempty"`

	// Division Details
	// Example: Accommodation and  Food Services
	Title string `json:"title,omitempty"`
}

// Validate validates this division details
func (m *DivisionDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this division details based on context it is used
func (m *DivisionDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DivisionDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DivisionDetails) UnmarshalBinary(b []byte) error {
	var res DivisionDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
