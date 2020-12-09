// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountHolder account holder
//
// swagger:model AccountHolder
type AccountHolder struct {

	// Identifies the Product as defined by institution
	// Example: Hooli Transaction
	// Required: true
	Product *string `json:"product"`

	// Identifies the Account type defined by institution
	// Example: transaction
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this account holder
func (m *AccountHolder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountHolder) validateProduct(formats strfmt.Registry) error {

	if err := validate.Required("product", "body", m.Product); err != nil {
		return err
	}

	return nil
}

func (m *AccountHolder) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountHolder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountHolder) UnmarshalBinary(b []byte) error {
	var res AccountHolder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
