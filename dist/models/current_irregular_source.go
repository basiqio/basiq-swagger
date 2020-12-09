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

// CurrentIrregularSource current irregular source
//
// swagger:model CurrentIrregularSource
type CurrentIrregularSource struct {

	// Most recent irregular income payment amount
	// Example: 62.00
	// Required: true
	Amount *string `json:"amount"`

	// Most recent irregular income payment date
	// Example: 2018-10-13T20:03:37
	// Required: true
	Date *string `json:"date"`
}

// Validate validates this current irregular source
func (m *CurrentIrregularSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CurrentIrregularSource) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *CurrentIrregularSource) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CurrentIrregularSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CurrentIrregularSource) UnmarshalBinary(b []byte) error {
	var res CurrentIrregularSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
