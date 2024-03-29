// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AffordabilityTransactionLinks affordability transaction links
//
// swagger:model AffordabilityTransactionLinks
type AffordabilityTransactionLinks struct {

	// Url of the account.
	// Example: https://au-api.basiq.io/users/6a52015e/accounts/31eb30a0
	// Required: true
	Account *string `json:"account"`

	// Url of the institution.
	// Example: https://au-api.basiq.io/institutions/AU00000
	// Required: true
	Institution *string `json:"institution"`
}

// Validate validates this affordability transaction links
func (m *AffordabilityTransactionLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstitution(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AffordabilityTransactionLinks) validateAccount(formats strfmt.Registry) error {

	if err := validate.Required("account", "body", m.Account); err != nil {
		return err
	}

	return nil
}

func (m *AffordabilityTransactionLinks) validateInstitution(formats strfmt.Registry) error {

	if err := validate.Required("institution", "body", m.Institution); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this affordability transaction links based on context it is used
func (m *AffordabilityTransactionLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AffordabilityTransactionLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AffordabilityTransactionLinks) UnmarshalBinary(b []byte) error {
	var res AffordabilityTransactionLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
