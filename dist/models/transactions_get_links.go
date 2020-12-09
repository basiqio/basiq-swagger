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

// TransactionsGetLinks transactions get links
//
// swagger:model TransactionsGetLinks
type TransactionsGetLinks struct {

	// Url to next result.
	// Example: https://au-api.basiq.io/users/6a52015e/transactions?next=bf1ec9d4
	Next string `json:"next,omitempty"`

	// Self reference url.
	// Example: https://au-api.basiq.io/users/ea3a81/transactions
	// Required: true
	Self *string `json:"self"`
}

// Validate validates this transactions get links
func (m *TransactionsGetLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionsGetLinks) validateSelf(formats strfmt.Registry) error {

	if err := validate.Required("self", "body", m.Self); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionsGetLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionsGetLinks) UnmarshalBinary(b []byte) error {
	var res TransactionsGetLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
