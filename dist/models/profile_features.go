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

// ProfileFeatures ProfileFeatures describes set of institution profile features
//
// swagger:model ProfileFeatures
type ProfileFeatures struct {

	// EmailAddresses holds list of data source identifiers which are capable to fetch bank customer e-mail addresses.
	// Example: ["web","pdf","csv"]
	// Required: true
	EmailAddresses []SourceName `json:"emailAddresses"`

	// FirstName holds list of data source identifiers which are capable to fetch bank customer first name.
	// Example: ["web","pdf","csv"]
	// Required: true
	FirstName []SourceName `json:"firstName"`

	// FullName holds list of data source identifiers which are capable to fetch bank customer full name.
	// Example: ["web","pdf","csv"]
	// Required: true
	FullName []SourceName `json:"fullName"`

	// LastName holds list of data source identifiers which are capable to fetch bank customer last name.
	// Example: ["web","pdf","csv"]
	// Required: true
	LastName []SourceName `json:"lastName"`

	// MiddleName holds list of data source identifiers which are capable to fetch bank customer middle name.
	// Example: ["web","pdf","csv"]
	// Required: true
	MiddleName []SourceName `json:"middleName"`

	// PhoneNumbers holds list of data source identifiers which are capable to fetch bank customer phone number.
	// Example: ["web","pdf","csv"]
	// Required: true
	PhoneNumbers []SourceName `json:"phoneNumbers"`

	// PhysicalAddresses holds list of data source identifiers which are capable to fetch bank customer physical addresses.
	// Example: ["web","pdf","csv"]
	// Required: true
	PhysicalAddresses []SourceName `json:"physicalAddresses"`
}

// Validate validates this profile features
func (m *ProfileFeatures) Validate(formats strfmt.Registry) error {
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

func (m *ProfileFeatures) validateEmailAddresses(formats strfmt.Registry) error {

	if err := validate.Required("emailAddresses", "body", m.EmailAddresses); err != nil {
		return err
	}

	for i := 0; i < len(m.EmailAddresses); i++ {

		if err := m.EmailAddresses[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emailAddresses" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ProfileFeatures) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	for i := 0; i < len(m.FirstName); i++ {

		if err := m.FirstName[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firstName" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ProfileFeatures) validateFullName(formats strfmt.Registry) error {

	if err := validate.Required("fullName", "body", m.FullName); err != nil {
		return err
	}

	for i := 0; i < len(m.FullName); i++ {

		if err := m.FullName[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fullName" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ProfileFeatures) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", m.LastName); err != nil {
		return err
	}

	for i := 0; i < len(m.LastName); i++ {

		if err := m.LastName[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastName" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ProfileFeatures) validateMiddleName(formats strfmt.Registry) error {

	if err := validate.Required("middleName", "body", m.MiddleName); err != nil {
		return err
	}

	for i := 0; i < len(m.MiddleName); i++ {

		if err := m.MiddleName[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("middleName" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ProfileFeatures) validatePhoneNumbers(formats strfmt.Registry) error {

	if err := validate.Required("phoneNumbers", "body", m.PhoneNumbers); err != nil {
		return err
	}

	for i := 0; i < len(m.PhoneNumbers); i++ {

		if err := m.PhoneNumbers[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phoneNumbers" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ProfileFeatures) validatePhysicalAddresses(formats strfmt.Registry) error {

	if err := validate.Required("physicalAddresses", "body", m.PhysicalAddresses); err != nil {
		return err
	}

	for i := 0; i < len(m.PhysicalAddresses); i++ {

		if err := m.PhysicalAddresses[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("physicalAddresses" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this profile features based on the context it is used
func (m *ProfileFeatures) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmailAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirstName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFullName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMiddleName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhoneNumbers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhysicalAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProfileFeatures) contextValidateEmailAddresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EmailAddresses); i++ {

		if err := m.EmailAddresses[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emailAddresses" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ProfileFeatures) contextValidateFirstName(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FirstName); i++ {

		if err := m.FirstName[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firstName" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ProfileFeatures) contextValidateFullName(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FullName); i++ {

		if err := m.FullName[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fullName" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ProfileFeatures) contextValidateLastName(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LastName); i++ {

		if err := m.LastName[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastName" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ProfileFeatures) contextValidateMiddleName(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MiddleName); i++ {

		if err := m.MiddleName[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("middleName" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ProfileFeatures) contextValidatePhoneNumbers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PhoneNumbers); i++ {

		if err := m.PhoneNumbers[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phoneNumbers" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ProfileFeatures) contextValidatePhysicalAddresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PhysicalAddresses); i++ {

		if err := m.PhysicalAddresses[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("physicalAddresses" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProfileFeatures) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProfileFeatures) UnmarshalBinary(b []byte) error {
	var res ProfileFeatures
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
