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

// GetUserLinks GetUserLinks
//
// Object containing links to resources.
//
// swagger:model GetUserLinks
type GetUserLinks struct {

	// Accounts reference url.
	// Required: true
	Accounts *string `json:"accounts"`

	// Auth link, possible null.
	// Required: true
	AuthLink *string `json:"auth_link"`

	// Connections reference url.
	// Required: true
	Connections *string `json:"connections"`

	// User self reference url.
	// Required: true
	Self *string `json:"self"`

	// Transactions reference url.
	// Required: true
	Transactions *string `json:"transactions"`
}

// Validate validates this get user links
func (m *GetUserLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUserLinks) validateAccounts(formats strfmt.Registry) error {

	if err := validate.Required("accounts", "body", m.Accounts); err != nil {
		return err
	}

	return nil
}

func (m *GetUserLinks) validateAuthLink(formats strfmt.Registry) error {

	if err := validate.Required("auth_link", "body", m.AuthLink); err != nil {
		return err
	}

	return nil
}

func (m *GetUserLinks) validateConnections(formats strfmt.Registry) error {

	if err := validate.Required("connections", "body", m.Connections); err != nil {
		return err
	}

	return nil
}

func (m *GetUserLinks) validateSelf(formats strfmt.Registry) error {

	if err := validate.Required("self", "body", m.Self); err != nil {
		return err
	}

	return nil
}

func (m *GetUserLinks) validateTransactions(formats strfmt.Registry) error {

	if err := validate.Required("transactions", "body", m.Transactions); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get user links based on context it is used
func (m *GetUserLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetUserLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUserLinks) UnmarshalBinary(b []byte) error {
	var res GetUserLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
