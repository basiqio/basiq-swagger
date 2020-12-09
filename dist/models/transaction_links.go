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

// TransactionLinks transaction links
//
// swagger:model TransactionLinks
type TransactionLinks struct {

	// Url of the account.
	// Example: https://au-api.basiq.io/users/6a52015e/accounts/31eb30a0
	// Required: true
	Account *string `json:"account"`

	// Url of the connection, always null.
	// Example: null
	// Required: true
	Connection *string `json:"connection"`

	// Url of the institution.
	// Example: https://au-api.basiq.io/institutions/AU00000
	// Required: true
	Institution *string `json:"institution"`

	// Transaction self reference.
	// Example: https://au-api.basiq.io/users/6a52015e/transactions/2082c765
	// Required: true
	Self *string `json:"self"`
}

// Validate validates this transaction links
func (m *TransactionLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstitution(formats); err != nil {
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

func (m *TransactionLinks) validateAccount(formats strfmt.Registry) error {

	if err := validate.Required("account", "body", m.Account); err != nil {
		return err
	}

	return nil
}

func (m *TransactionLinks) validateConnection(formats strfmt.Registry) error {

	if err := validate.Required("connection", "body", m.Connection); err != nil {
		return err
	}

	return nil
}

func (m *TransactionLinks) validateInstitution(formats strfmt.Registry) error {

	if err := validate.Required("institution", "body", m.Institution); err != nil {
		return err
	}

	return nil
}

func (m *TransactionLinks) validateSelf(formats strfmt.Registry) error {

	if err := validate.Required("self", "body", m.Self); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionLinks) UnmarshalBinary(b []byte) error {
	var res TransactionLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
