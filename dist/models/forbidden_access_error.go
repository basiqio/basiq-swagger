// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ForbiddenAccessError forbidden access error
//
// swagger:model ForbiddenAccessError
type ForbiddenAccessError struct {

	// Unique identifier for this particular occurrence of the problem.
	// Required: true
	CorrelationID *string `json:"correlationId"`

	// Error data.
	// Required: true
	Data []*ForbiddenAccessErrorDataItems0 `json:"data"`

	// Always "list".
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this forbidden access error
func (m *ForbiddenAccessError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCorrelationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ForbiddenAccessError) validateCorrelationID(formats strfmt.Registry) error {

	if err := validate.Required("correlationId", "body", m.CorrelationID); err != nil {
		return err
	}

	return nil
}

func (m *ForbiddenAccessError) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ForbiddenAccessError) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ForbiddenAccessError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ForbiddenAccessError) UnmarshalBinary(b []byte) error {
	var res ForbiddenAccessError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ForbiddenAccessErrorDataItems0 forbidden access error data items0
//
// swagger:model ForbiddenAccessErrorDataItems0
type ForbiddenAccessErrorDataItems0 struct {

	// Application-specific error code, expressed as a string value.
	// Required: true
	// Enum: [forbidden-access no-production-access access-denied]
	Code *string `json:"code"`

	// Human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty"`

	// Title of the error
	Title string `json:"title,omitempty"`

	// Type of the response, always "error"
	// Required: true
	Type *string `json:"type"`

	// source
	// Required: true
	Source *Source `json:"source"`
}

// Validate validates this forbidden access error data items0
func (m *ForbiddenAccessErrorDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var forbiddenAccessErrorDataItems0TypeCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["forbidden-access","no-production-access","access-denied"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		forbiddenAccessErrorDataItems0TypeCodePropEnum = append(forbiddenAccessErrorDataItems0TypeCodePropEnum, v)
	}
}

const (

	// ForbiddenAccessErrorDataItems0CodeForbiddenAccess captures enum value "forbidden-access"
	ForbiddenAccessErrorDataItems0CodeForbiddenAccess string = "forbidden-access"

	// ForbiddenAccessErrorDataItems0CodeNoProductionAccess captures enum value "no-production-access"
	ForbiddenAccessErrorDataItems0CodeNoProductionAccess string = "no-production-access"

	// ForbiddenAccessErrorDataItems0CodeAccessDenied captures enum value "access-denied"
	ForbiddenAccessErrorDataItems0CodeAccessDenied string = "access-denied"
)

// prop value enum
func (m *ForbiddenAccessErrorDataItems0) validateCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, forbiddenAccessErrorDataItems0TypeCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ForbiddenAccessErrorDataItems0) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	// value enum
	if err := m.validateCodeEnum("code", "body", *m.Code); err != nil {
		return err
	}

	return nil
}

func (m *ForbiddenAccessErrorDataItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ForbiddenAccessErrorDataItems0) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ForbiddenAccessErrorDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ForbiddenAccessErrorDataItems0) UnmarshalBinary(b []byte) error {
	var res ForbiddenAccessErrorDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
