// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClassDetails class details
//
// swagger:model ClassDetails
type ClassDetails struct {

	// Class Code
	// Example: 4511
	Code string `json:"code,omitempty"`

	// Class Details
	// Example: Cafes and Restaurants
	Title string `json:"title,omitempty"`
}

// Validate validates this class details
func (m *ClassDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClassDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClassDetails) UnmarshalBinary(b []byte) error {
	var res ClassDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
