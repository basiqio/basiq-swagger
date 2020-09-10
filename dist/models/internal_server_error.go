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

// InternalServerError internal server error
//
// swagger:model InternalServerError
type InternalServerError struct {

	// Unique identifier for this particular occurrence of the problem.
	// Required: true
	CorrelationID *string `json:"correlationId"`

	// Error data.
	// Required: true
	Data []*InternalServerErrorDataItems0 `json:"data"`

	// Always "list".
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this internal server error
func (m *InternalServerError) Validate(formats strfmt.Registry) error {
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

func (m *InternalServerError) validateCorrelationID(formats strfmt.Registry) error {

	if err := validate.Required("correlationId", "body", m.CorrelationID); err != nil {
		return err
	}

	return nil
}

func (m *InternalServerError) validateData(formats strfmt.Registry) error {

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

func (m *InternalServerError) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InternalServerError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InternalServerError) UnmarshalBinary(b []byte) error {
	var res InternalServerError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InternalServerErrorDataItems0 internal server error data items0
//
// swagger:model InternalServerErrorDataItems0
type InternalServerErrorDataItems0 struct {

	// Application-specific error code, expressed as a string value.
	// Required: true
	// Enum: [internal-server-error]
	Code *string `json:"code"`

	// Human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty"`

	// Title of the error
	Title string `json:"title,omitempty"`

	// Type of the response, always "error"
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this internal server error data items0
func (m *InternalServerErrorDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
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

var internalServerErrorDataItems0TypeCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["internal-server-error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		internalServerErrorDataItems0TypeCodePropEnum = append(internalServerErrorDataItems0TypeCodePropEnum, v)
	}
}

const (

	// InternalServerErrorDataItems0CodeInternalServerError captures enum value "internal-server-error"
	InternalServerErrorDataItems0CodeInternalServerError string = "internal-server-error"
)

// prop value enum
func (m *InternalServerErrorDataItems0) validateCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, internalServerErrorDataItems0TypeCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InternalServerErrorDataItems0) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	// value enum
	if err := m.validateCodeEnum("code", "body", *m.Code); err != nil {
		return err
	}

	return nil
}

func (m *InternalServerErrorDataItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InternalServerErrorDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InternalServerErrorDataItems0) UnmarshalBinary(b []byte) error {
	var res InternalServerErrorDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}