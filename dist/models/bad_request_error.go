// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BadRequestError bad request error
//
// swagger:model BadRequestError
type BadRequestError struct {

	// Unique identifier for this particular occurrence of the problem.
	// Example: ac5ah5i
	// Required: true
	CorrelationID *string `json:"correlationId"`

	// Error data.
	// Required: true
	Data []*BadRequestErrorDataItems0 `json:"data"`

	// Always "list".
	// Example: list
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this bad request error
func (m *BadRequestError) Validate(formats strfmt.Registry) error {
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

func (m *BadRequestError) validateCorrelationID(formats strfmt.Registry) error {

	if err := validate.Required("correlationId", "body", m.CorrelationID); err != nil {
		return err
	}

	return nil
}

func (m *BadRequestError) validateData(formats strfmt.Registry) error {

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

func (m *BadRequestError) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this bad request error based on the context it is used
func (m *BadRequestError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BadRequestError) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Data); i++ {

		if m.Data[i] != nil {
			if err := m.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BadRequestError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BadRequestError) UnmarshalBinary(b []byte) error {
	var res BadRequestError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BadRequestErrorDataItems0 bad request error data items0
//
// swagger:model BadRequestErrorDataItems0
type BadRequestErrorDataItems0 struct {

	// Application-specific error code, expressed as a string value.
	// Example: parameter-not-valid
	// Required: true
	// Enum: [parameter-not-supplied parameter-not-valid unsupported-accept invalid-content institution-not-supported temporary-unavailable invalid-credentials]
	Code *string `json:"code"`

	// Human-readable explanation specific to this occurrence of the problem.
	// Example: ID value is not valid.
	Detail string `json:"detail,omitempty"`

	// source
	Source *Source `json:"source,omitempty"`

	// Title of the error
	// Example: Parameter not valid.
	Title string `json:"title,omitempty"`

	// Type of the response, always "error"
	// Example: error
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this bad request error data items0
func (m *BadRequestErrorDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
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

var badRequestErrorDataItems0TypeCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["parameter-not-supplied","parameter-not-valid","unsupported-accept","invalid-content","institution-not-supported","temporary-unavailable","invalid-credentials"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		badRequestErrorDataItems0TypeCodePropEnum = append(badRequestErrorDataItems0TypeCodePropEnum, v)
	}
}

const (

	// BadRequestErrorDataItems0CodeParameterDashNotDashSupplied captures enum value "parameter-not-supplied"
	BadRequestErrorDataItems0CodeParameterDashNotDashSupplied string = "parameter-not-supplied"

	// BadRequestErrorDataItems0CodeParameterDashNotDashValid captures enum value "parameter-not-valid"
	BadRequestErrorDataItems0CodeParameterDashNotDashValid string = "parameter-not-valid"

	// BadRequestErrorDataItems0CodeUnsupportedDashAccept captures enum value "unsupported-accept"
	BadRequestErrorDataItems0CodeUnsupportedDashAccept string = "unsupported-accept"

	// BadRequestErrorDataItems0CodeInvalidDashContent captures enum value "invalid-content"
	BadRequestErrorDataItems0CodeInvalidDashContent string = "invalid-content"

	// BadRequestErrorDataItems0CodeInstitutionDashNotDashSupported captures enum value "institution-not-supported"
	BadRequestErrorDataItems0CodeInstitutionDashNotDashSupported string = "institution-not-supported"

	// BadRequestErrorDataItems0CodeTemporaryDashUnavailable captures enum value "temporary-unavailable"
	BadRequestErrorDataItems0CodeTemporaryDashUnavailable string = "temporary-unavailable"

	// BadRequestErrorDataItems0CodeInvalidDashCredentials captures enum value "invalid-credentials"
	BadRequestErrorDataItems0CodeInvalidDashCredentials string = "invalid-credentials"
)

// prop value enum
func (m *BadRequestErrorDataItems0) validateCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, badRequestErrorDataItems0TypeCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BadRequestErrorDataItems0) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	// value enum
	if err := m.validateCodeEnum("code", "body", *m.Code); err != nil {
		return err
	}

	return nil
}

func (m *BadRequestErrorDataItems0) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
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

func (m *BadRequestErrorDataItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this bad request error data items0 based on the context it is used
func (m *BadRequestErrorDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BadRequestErrorDataItems0) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {
		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BadRequestErrorDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BadRequestErrorDataItems0) UnmarshalBinary(b []byte) error {
	var res BadRequestErrorDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
