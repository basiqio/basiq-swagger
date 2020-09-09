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

// IncomeSummary Summary totals relating to income analysis
//
// swagger:model IncomeSummary
type IncomeSummary struct {

	// Total mean of irregular income monthly calculated across the whole time period for all irregular sources
	// Required: true
	IrregularIncomeAvg *string `json:"irregularIncomeAvg"`

	// Total median regular income monthly calculated over the past 3 months for all regular sources
	// Required: true
	RegularIncomeAvg *string `json:"regularIncomeAvg"`

	// Total regular income so far this financial year (year to date)
	// Required: true
	RegularIncomeYTD *string `json:"regularIncomeYTD"`

	// Total predicted regular income for this financial year year
	// Required: true
	RegularIncomeYear *string `json:"regularIncomeYear"`
}

// Validate validates this income summary
func (m *IncomeSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIrregularIncomeAvg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegularIncomeAvg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegularIncomeYTD(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegularIncomeYear(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IncomeSummary) validateIrregularIncomeAvg(formats strfmt.Registry) error {

	if err := validate.Required("irregularIncomeAvg", "body", m.IrregularIncomeAvg); err != nil {
		return err
	}

	return nil
}

func (m *IncomeSummary) validateRegularIncomeAvg(formats strfmt.Registry) error {

	if err := validate.Required("regularIncomeAvg", "body", m.RegularIncomeAvg); err != nil {
		return err
	}

	return nil
}

func (m *IncomeSummary) validateRegularIncomeYTD(formats strfmt.Registry) error {

	if err := validate.Required("regularIncomeYTD", "body", m.RegularIncomeYTD); err != nil {
		return err
	}

	return nil
}

func (m *IncomeSummary) validateRegularIncomeYear(formats strfmt.Registry) error {

	if err := validate.Required("regularIncomeYear", "body", m.RegularIncomeYear); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IncomeSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IncomeSummary) UnmarshalBinary(b []byte) error {
	var res IncomeSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
