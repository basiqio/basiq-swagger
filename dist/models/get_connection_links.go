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

// GetConnectionLinks GetConnectionLinks
//
// Object containing links to resources.
//
// swagger:model GetConnectionLinks
type GetConnectionLinks struct {

	// Accounts reference url.
	// Example: https://au-api.basiq.io/users/cd6fbd92/accounts?filter=institution.id.eq('AU00000')
	Accounts string `json:"accounts,omitempty"`

	// Connection self reference url.
	// Example: https://au-api.basiq.io/users/cd6fbd92-0b12-43ba-a3c1-286dd5f4f396/connections/29523951
	// Required: true
	Self *string `json:"self"`

	// Transactions reference url.
	// Example: https://au-api.basiq.io/users/cd6fbd92/transactions?filter=institution.id.eq('AU00000')
	Transactions string `json:"transactions,omitempty"`

	// User reference url.
	// Example: https://au-api.basiq.io/users/cd6fbd92
	// Required: true
	User *string `json:"user"`
}

// Validate validates this get connection links
func (m *GetConnectionLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetConnectionLinks) validateSelf(formats strfmt.Registry) error {

	if err := validate.Required("self", "body", m.Self); err != nil {
		return err
	}

	return nil
}

func (m *GetConnectionLinks) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get connection links based on context it is used
func (m *GetConnectionLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetConnectionLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetConnectionLinks) UnmarshalBinary(b []byte) error {
	var res GetConnectionLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
