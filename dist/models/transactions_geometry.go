// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TransactionsGeometry transactions geometry
//
// swagger:model TransactionsGeometry
type TransactionsGeometry struct {

	// Latitude
	Lat string `json:"lat,omitempty"`

	// Longitude
	Lng string `json:"lng,omitempty"`
}

// Validate validates this transactions geometry
func (m *TransactionsGeometry) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TransactionsGeometry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionsGeometry) UnmarshalBinary(b []byte) error {
	var res TransactionsGeometry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
