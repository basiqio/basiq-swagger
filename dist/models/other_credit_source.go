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

// OtherCreditSource Other Credit Series are series where a pattern of credit is detected but the pattern does not meet the requirements for regular or irregular income, or the income series is too old e.g. former regular income or a potential future income
//
// swagger:model OtherCreditSource
type OtherCreditSource struct {

	// Duration other income (number days from first to last occurrence) returned as an integer with values zero or greater.
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

	// current
	// Required: true
	Current *CurrentOtherCreditSource `json:"current"`

	// Frequency is "other", "irregular" or a time period e.g. "bi-weekly"
	// Required: true
	// Enum: [daily weekly bi-weekly monthly bi-monthly quarterly half-year yearly other irregular]
	Frequency *string `json:"frequency"`

	// Number of instances of credits in the series.
	// Required: true
	NoOccurrences *int64 `json:"noOccurrences"`

	// Source Other Credit income (cleaned transaction description).
	// Required: true
	Source *string `json:"source"`
}

// Validate validates this other credit source
func (m *OtherCreditSource) Validate(formats strfmt.Registry) error {
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

	if err := m.validateCurrent(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OtherCreditSource) validateAgeDays(formats strfmt.Registry) error {

	if err := validate.Required("ageDays", "body", m.AgeDays); err != nil {
		return err
	}

	return nil
}

func (m *OtherCreditSource) validateAmountAvg(formats strfmt.Registry) error {

	if err := validate.Required("amountAvg", "body", m.AmountAvg); err != nil {
		return err
	}

	return nil
}

func (m *OtherCreditSource) validateAvgMonthlyOccurence(formats strfmt.Registry) error {

	if err := validate.Required("avgMonthlyOccurence", "body", m.AvgMonthlyOccurence); err != nil {
		return err
	}

	return nil
}

func (m *OtherCreditSource) validateChangeHistory(formats strfmt.Registry) error {

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

func (m *OtherCreditSource) validateCurrent(formats strfmt.Registry) error {

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

var otherCreditSourceTypeFrequencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["daily","weekly","bi-weekly","monthly","bi-monthly","quarterly","half-year","yearly","other","irregular"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		otherCreditSourceTypeFrequencyPropEnum = append(otherCreditSourceTypeFrequencyPropEnum, v)
	}
}

const (

	// OtherCreditSourceFrequencyDaily captures enum value "daily"
	OtherCreditSourceFrequencyDaily string = "daily"

	// OtherCreditSourceFrequencyWeekly captures enum value "weekly"
	OtherCreditSourceFrequencyWeekly string = "weekly"

	// OtherCreditSourceFrequencyBiWeekly captures enum value "bi-weekly"
	OtherCreditSourceFrequencyBiWeekly string = "bi-weekly"

	// OtherCreditSourceFrequencyMonthly captures enum value "monthly"
	OtherCreditSourceFrequencyMonthly string = "monthly"

	// OtherCreditSourceFrequencyBiMonthly captures enum value "bi-monthly"
	OtherCreditSourceFrequencyBiMonthly string = "bi-monthly"

	// OtherCreditSourceFrequencyQuarterly captures enum value "quarterly"
	OtherCreditSourceFrequencyQuarterly string = "quarterly"

	// OtherCreditSourceFrequencyHalfYear captures enum value "half-year"
	OtherCreditSourceFrequencyHalfYear string = "half-year"

	// OtherCreditSourceFrequencyYearly captures enum value "yearly"
	OtherCreditSourceFrequencyYearly string = "yearly"

	// OtherCreditSourceFrequencyOther captures enum value "other"
	OtherCreditSourceFrequencyOther string = "other"

	// OtherCreditSourceFrequencyIrregular captures enum value "irregular"
	OtherCreditSourceFrequencyIrregular string = "irregular"
)

// prop value enum
func (m *OtherCreditSource) validateFrequencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, otherCreditSourceTypeFrequencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OtherCreditSource) validateFrequency(formats strfmt.Registry) error {

	if err := validate.Required("frequency", "body", m.Frequency); err != nil {
		return err
	}

	// value enum
	if err := m.validateFrequencyEnum("frequency", "body", *m.Frequency); err != nil {
		return err
	}

	return nil
}

func (m *OtherCreditSource) validateNoOccurrences(formats strfmt.Registry) error {

	if err := validate.Required("noOccurrences", "body", m.NoOccurrences); err != nil {
		return err
	}

	return nil
}

func (m *OtherCreditSource) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OtherCreditSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OtherCreditSource) UnmarshalBinary(b []byte) error {
	var res OtherCreditSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
