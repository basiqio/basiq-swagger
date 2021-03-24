// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TransactionsClass transactions class
//
// swagger:model TransactionsClass
type TransactionsClass struct {

	// Class Code
	Code string `json:"code,omitempty"`

	// Class Details
	Title string `json:"title,omitempty"`
}

// Validate validates this transactions class
func (m *TransactionsClass) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TransactionsClass) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionsClass) UnmarshalBinary(b []byte) error {
	var res TransactionsClass
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
