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

// IncomeLinks income links
//
// swagger:model IncomeLinks
type IncomeLinks struct {

	// Links of accounts
	// Required: true
	Accounts []string `json:"accounts"`

	// Link to the requested income resource
	// Required: true
	Self *string `json:"self"`
}

// Validate validates this income links
func (m *IncomeLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IncomeLinks) validateAccounts(formats strfmt.Registry) error {

	if err := validate.Required("accounts", "body", m.Accounts); err != nil {
		return err
	}

	return nil
}

func (m *IncomeLinks) validateSelf(formats strfmt.Registry) error {

	if err := validate.Required("self", "body", m.Self); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IncomeLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IncomeLinks) UnmarshalBinary(b []byte) error {
	var res IncomeLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
