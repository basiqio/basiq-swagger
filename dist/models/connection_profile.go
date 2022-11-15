// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConnectionProfile connection profile
//
// swagger:model ConnectionProfile
type ConnectionProfile struct {

	// User email address
	// Example: ["gavin@hooli.com"]
	// Required: true
	EmailAddresses []string `json:"emailAddresses"`

	// User first name
	// Example: Gavin
	// Required: true
	FirstName *string `json:"firstName"`

	// User full name
	// Example: Gavin Belson
	// Required: true
	FullName *string `json:"fullName"`

	// User last name
	// Example: Belson
	// Required: true
	LastName *string `json:"lastName"`

	// User middle name
	// Required: true
	MiddleName *string `json:"middleName"`

	// User phone numbers
	// Example: ["XXXX 888 991"]
	// Required: true
	PhoneNumbers []string `json:"phoneNumbers"`

	// Physical user addresses
	// Required: true
	PhysicalAddresses []*PhysicalAddresses `json:"physicalAddresses"`
}

// Validate validates this connection profile
func (m *ConnectionProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmailAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMiddleName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneNumbers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysicalAddresses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectionProfile) validateEmailAddresses(formats strfmt.Registry) error {

	if err := validate.Required("emailAddresses", "body", m.EmailAddresses); err != nil {
		return err
	}

	return nil
}

func (m *ConnectionProfile) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *ConnectionProfile) validateFullName(formats strfmt.Registry) error {

	if err := validate.Required("fullName", "body", m.FullName); err != nil {
		return err
	}

	return nil
}

func (m *ConnectionProfile) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

func (m *ConnectionProfile) validateMiddleName(formats strfmt.Registry) error {

	if err := validate.Required("middleName", "body", m.MiddleName); err != nil {
		return err
	}

	return nil
}

func (m *ConnectionProfile) validatePhoneNumbers(formats strfmt.Registry) error {

	if err := validate.Required("phoneNumbers", "body", m.PhoneNumbers); err != nil {
		return err
	}

	return nil
}

func (m *ConnectionProfile) validatePhysicalAddresses(formats strfmt.Registry) error {

	if err := validate.Required("physicalAddresses", "body", m.PhysicalAddresses); err != nil {
		return err
	}

	for i := 0; i < len(m.PhysicalAddresses); i++ {
		if swag.IsZero(m.PhysicalAddresses[i]) { // not required
			continue
		}

		if m.PhysicalAddresses[i] != nil {
			if err := m.PhysicalAddresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("physicalAddresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("physicalAddresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this connection profile based on the context it is used
func (m *ConnectionProfile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePhysicalAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectionProfile) contextValidatePhysicalAddresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PhysicalAddresses); i++ {

		if m.PhysicalAddresses[i] != nil {
			if err := m.PhysicalAddresses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("physicalAddresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("physicalAddresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConnectionProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectionProfile) UnmarshalBinary(b []byte) error {
	var res ConnectionProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
