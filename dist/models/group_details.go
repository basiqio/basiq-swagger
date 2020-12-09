// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GroupDetails group details
//
// swagger:model GroupDetails
type GroupDetails struct {

	// Group Code
	// Example: 451
	Code string `json:"code,omitempty"`

	// Group Details
	// Example: Cafes, Restaurants and Takeaway Food Services
	Title string `json:"title,omitempty"`
}

// Validate validates this group details
func (m *GroupDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GroupDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupDetails) UnmarshalBinary(b []byte) error {
	var res GroupDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
