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

// AccountTransactionInterval account transaction interval
//
// swagger:model AccountTransactionInterval
type AccountTransactionInterval struct {

	// Date of first transaction on this account
	// Required: true
	From *string `json:"from"`

	// Date of last transaction on this account
	// Required: true
	To *string `json:"to"`
}

// Validate validates this account transaction interval
func (m *AccountTransactionInterval) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountTransactionInterval) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	return nil
}

func (m *AccountTransactionInterval) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("to", "body", m.To); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountTransactionInterval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountTransactionInterval) UnmarshalBinary(b []byte) error {
	var res AccountTransactionInterval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}