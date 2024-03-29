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

// JobPostRequest job post request
//
// swagger:model JobPostRequest
type JobPostRequest struct {

	// One time password or answer to a security question/s e.g. ["1234"]
	// Example: ["1234"]
	// Required: true
	MfaResponse []string `json:"mfa-response"`
}

// Validate validates this job post request
func (m *JobPostRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMfaResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobPostRequest) validateMfaResponse(formats strfmt.Registry) error {

	if err := validate.Required("mfa-response", "body", m.MfaResponse); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this job post request based on context it is used
func (m *JobPostRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JobPostRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobPostRequest) UnmarshalBinary(b []byte) error {
	var res JobPostRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
