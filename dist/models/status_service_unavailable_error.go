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

// StatusServiceUnavailableError status service unavailable error
//
// swagger:model StatusServiceUnavailableError
type StatusServiceUnavailableError struct {

	// Unique identifier for this particular occurrence of the problem.
	// Example: ac5ah5i
	// Required: true
	CorrelationID *string `json:"correlationId"`

	// Error data.
	// Required: true
	Data []*StatusServiceUnavailableErrorDataItems0 `json:"data"`

	// Always "list".
	// Example: list
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this status service unavailable error
func (m *StatusServiceUnavailableError) Validate(formats strfmt.Registry) error {
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

func (m *StatusServiceUnavailableError) validateCorrelationID(formats strfmt.Registry) error {

	if err := validate.Required("correlationId", "body", m.CorrelationID); err != nil {
		return err
	}

	return nil
}

func (m *StatusServiceUnavailableError) validateData(formats strfmt.Registry) error {

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

func (m *StatusServiceUnavailableError) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this status service unavailable error based on the context it is used
func (m *StatusServiceUnavailableError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusServiceUnavailableError) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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
func (m *StatusServiceUnavailableError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusServiceUnavailableError) UnmarshalBinary(b []byte) error {
	var res StatusServiceUnavailableError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StatusServiceUnavailableErrorDataItems0 status service unavailable error data items0
//
// swagger:model StatusServiceUnavailableErrorDataItems0
type StatusServiceUnavailableErrorDataItems0 struct {

	// Application-specific error code, expressed as a string value.
	// Example: service-unavailable
	// Required: true
	Code interface{} `json:"code"`

	// Human-readable explanation specific to this occurrence of the problem.
	// Example: Service Unavailable. Try again later.
	Detail string `json:"detail,omitempty"`

	// Title of the error
	// Example: Service Unavailable
	Title string `json:"title,omitempty"`

	// Type of the response, always "error"
	// Example: error
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this status service unavailable error data items0
func (m *StatusServiceUnavailableErrorDataItems0) Validate(formats strfmt.Registry) error {
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

func (m *StatusServiceUnavailableErrorDataItems0) validateCode(formats strfmt.Registry) error {

	if m.Code == nil {
		return errors.Required("code", "body", nil)
	}

	return nil
}

func (m *StatusServiceUnavailableErrorDataItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this status service unavailable error data items0 based on context it is used
func (m *StatusServiceUnavailableErrorDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StatusServiceUnavailableErrorDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusServiceUnavailableErrorDataItems0) UnmarshalBinary(b []byte) error {
	var res StatusServiceUnavailableErrorDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
