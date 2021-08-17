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

// ChangeHistoryExpensesClass Amount classified as spent that month (repeated each month of data)
//
// swagger:model ChangeHistoryExpensesClass
type ChangeHistoryExpensesClass struct {

	// Amount of expense that period
	// Example: -11.00
	// Required: true
	Amount *string `json:"amount"`

	// Month expense relates
	// Example: 2018-09
	// Required: true
	Date *string `json:"date"`
}

// Validate validates this change history expenses class
func (m *ChangeHistoryExpensesClass) Validate(formats strfmt.Registry) error {
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

func (m *ChangeHistoryExpensesClass) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *ChangeHistoryExpensesClass) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this change history expenses class based on context it is used
func (m *ChangeHistoryExpensesClass) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChangeHistoryExpensesClass) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangeHistoryExpensesClass) UnmarshalBinary(b []byte) error {
	var res ChangeHistoryExpensesClass
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
