// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EnrichLocation enrich location
//
// swagger:model EnrichLocation
type EnrichLocation struct {

	// Country
	// Example: Australia
	Country string `json:"country,omitempty"`

	// Address
	// Example: 1/39 E Esplanade, Manly NSW 2095
	FormattedAddress string `json:"formattedAddress,omitempty"`

	// geometry
	Geometry *EnrichGeometry `json:"geometry,omitempty"`

	// Postal Code
	// Example: 2095
	PostalCode string `json:"postalCode,omitempty"`

	// Route Name
	// Example: E Esplanade
	Route string `json:"route,omitempty"`

	// Route Number
	// Example: 29
	RouteNo string `json:"routeNo,omitempty"`

	// State
	// Example: NSW
	State string `json:"state,omitempty"`

	// Suburb
	// Example: Manly
	Suburb string `json:"suburb,omitempty"`
}

// Validate validates this enrich location
func (m *EnrichLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGeometry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnrichLocation) validateGeometry(formats strfmt.Registry) error {

	if swag.IsZero(m.Geometry) { // not required
		return nil
	}

	if m.Geometry != nil {
		if err := m.Geometry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geometry")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnrichLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnrichLocation) UnmarshalBinary(b []byte) error {
	var res EnrichLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
