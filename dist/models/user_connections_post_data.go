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

// UserConnectionsPostData user connections post data
//
// swagger:model UserConnectionsPostData
type UserConnectionsPostData struct {

	// The users institution login ID
	// Required: true
	LoginID *string `json:"loginId"`

	// The users institution password
	// Required: true
	Password *string `json:"password"`

	// User's institution secondary login id. Mandatory if required by institution's login process
	SecondaryLoginID string `json:"secondaryLoginId,omitempty"`

	// User's institution security code. Mandatory if required by institution's login process.
	SecurityCode string `json:"securityCode,omitempty"`

	// institution
	// Required: true
	Institution *InstitutionModel `json:"institution"`
}

// Validate validates this user connections post data
func (m *UserConnectionsPostData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoginID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
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

func (m *UserConnectionsPostData) validateLoginID(formats strfmt.Registry) error {

	if err := validate.Required("loginId", "body", m.LoginID); err != nil {
		return err
	}

	return nil
}

func (m *UserConnectionsPostData) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *UserConnectionsPostData) validateInstitution(formats strfmt.Registry) error {

	if err := validate.Required("institution", "body", m.Institution); err != nil {
		return err
	}

	if m.Institution != nil {
		if err := m.Institution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("institution")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserConnectionsPostData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserConnectionsPostData) UnmarshalBinary(b []byte) error {
	var res UserConnectionsPostData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}