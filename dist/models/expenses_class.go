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

// ExpensesClass Class represents one class of classification system
//
// swagger:model ExpensesClass
type ExpensesClass struct {

	// Classification code for HEC classification: 4 levels class, group, subdivision and division.
	// Required: true
	ClassCode *string `json:"classCode"`

	// Classification code for HEC classification: 4 levels class, group, subdivision and division.
	// Required: true
	ClassTitle *string `json:"classTitle"`

	// division code
	// Required: true
	DivisionCode *string `json:"divisionCode"`

	// division title
	// Required: true
	DivisionTitle *string `json:"divisionTitle"`
}

// Validate validates this expenses class
func (m *ExpensesClass) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClassCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClassTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivisionCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivisionTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExpensesClass) validateClassCode(formats strfmt.Registry) error {

	if err := validate.Required("classCode", "body", m.ClassCode); err != nil {
		return err
	}

	return nil
}

func (m *ExpensesClass) validateClassTitle(formats strfmt.Registry) error {

	if err := validate.Required("classTitle", "body", m.ClassTitle); err != nil {
		return err
	}

	return nil
}

func (m *ExpensesClass) validateDivisionCode(formats strfmt.Registry) error {

	if err := validate.Required("divisionCode", "body", m.DivisionCode); err != nil {
		return err
	}

	return nil
}

func (m *ExpensesClass) validateDivisionTitle(formats strfmt.Registry) error {

	if err := validate.Required("divisionTitle", "body", m.DivisionTitle); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExpensesClass) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExpensesClass) UnmarshalBinary(b []byte) error {
	var res ExpensesClass
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
