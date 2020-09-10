// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IrregularSource Irregular Income sources typically require at least 5 credits across a minimum 90 day time period
//
// swagger:model IrregularSource
type IrregularSource struct {

	// Duration irregular income (number days from first to last occurrence) returned as an integer with values zero or greater.
	// Required: true
	AgeDays *int64 `json:"ageDays"`

	// Mean of irregular income amount - calculated across all occurrences identified.
	// Required: true
	AmountAvg *string `json:"amountAvg"`

	// Average (mean) number of times per calendar month the credits in the series occur.
	// Required: true
	AvgMonthlyOccurence *string `json:"avgMonthlyOccurence"`

	// Each amount classified as income (repeated for each income credit and ordered by most recent)
	// Required: true
	ChangeHistory []*ChangeHistoryIncome `json:"changeHistory"`

	// Frequency is "irregular"
	// Required: true
	Frequency *string `json:"frequency"`

	// Number of instances of credits in the series.
	// Required: true
	NoOccurrences *int64 `json:"noOccurrences"`

	// Source irregular income (cleaned transaction description).
	// Required: true
	Source *string `json:"source"`

	// current
	// Required: true
	Current *CurrentIrregularSource `json:"current"`
}

// Validate validates this irregular source
func (m *IrregularSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgeDays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmountAvg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvgMonthlyOccurence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChangeHistory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNoOccurrences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IrregularSource) validateAgeDays(formats strfmt.Registry) error {

	if err := validate.Required("ageDays", "body", m.AgeDays); err != nil {
		return err
	}

	return nil
}

func (m *IrregularSource) validateAmountAvg(formats strfmt.Registry) error {

	if err := validate.Required("amountAvg", "body", m.AmountAvg); err != nil {
		return err
	}

	return nil
}

func (m *IrregularSource) validateAvgMonthlyOccurence(formats strfmt.Registry) error {

	if err := validate.Required("avgMonthlyOccurence", "body", m.AvgMonthlyOccurence); err != nil {
		return err
	}

	return nil
}

func (m *IrregularSource) validateChangeHistory(formats strfmt.Registry) error {

	if err := validate.Required("changeHistory", "body", m.ChangeHistory); err != nil {
		return err
	}

	for i := 0; i < len(m.ChangeHistory); i++ {
		if swag.IsZero(m.ChangeHistory[i]) { // not required
			continue
		}

		if m.ChangeHistory[i] != nil {
			if err := m.ChangeHistory[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("changeHistory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IrregularSource) validateFrequency(formats strfmt.Registry) error {

	if err := validate.Required("frequency", "body", m.Frequency); err != nil {
		return err
	}

	return nil
}

func (m *IrregularSource) validateNoOccurrences(formats strfmt.Registry) error {

	if err := validate.Required("noOccurrences", "body", m.NoOccurrences); err != nil {
		return err
	}

	return nil
}

func (m *IrregularSource) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *IrregularSource) validateCurrent(formats strfmt.Registry) error {

	if err := validate.Required("current", "body", m.Current); err != nil {
		return err
	}

	if m.Current != nil {
		if err := m.Current.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IrregularSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IrregularSource) UnmarshalBinary(b []byte) error {
	var res IrregularSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}