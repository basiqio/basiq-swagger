// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateUser create user
//
// swagger:model createUser
type CreateUser struct {

	// The end-users email address. Mandatory if mobile is not supplied.
	// Example: gavin@hooli.com
	Email string `json:"email,omitempty"`

	// The end-users mobile number, supplied in international format.
	// +[country-code][mobileno]. Mandatory if email is not supplied.
	// Example: +61410888999
	Mobile string `json:"mobile,omitempty"`
}

// Validate validates this create user
func (m *CreateUser) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create user based on context it is used
func (m *CreateUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateUser) UnmarshalBinary(b []byte) error {
	var res CreateUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
