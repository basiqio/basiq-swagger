// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChangeHistoryAffordabilityData change history affordability data
//
// swagger:model ChangeHistoryAffordabilityData
type ChangeHistoryAffordabilityData struct {

	// Amount loan-interest or loan-repayment
	// Required: true
	Amount *string `json:"amount"`

	// Date
	// Required: true
	// Format: date
	Date *strfmt.Date `json:"date"`

	// Debit or Credit
	// Required: true
	// Enum: [debit credit]
	Direction *string `json:"direction"`

	// Cleaned transaction description
	// Required: true
	Source *string `json:"source"`
}

// Validate validates this change history affordability data
func (m *ChangeHistoryAffordabilityData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
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

func (m *ChangeHistoryAffordabilityData) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *ChangeHistoryAffordabilityData) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

var changeHistoryAffordabilityDataTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["debit","credit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeHistoryAffordabilityDataTypeDirectionPropEnum = append(changeHistoryAffordabilityDataTypeDirectionPropEnum, v)
	}
}

const (

	// ChangeHistoryAffordabilityDataDirectionDebit captures enum value "debit"
	ChangeHistoryAffordabilityDataDirectionDebit string = "debit"

	// ChangeHistoryAffordabilityDataDirectionCredit captures enum value "credit"
	ChangeHistoryAffordabilityDataDirectionCredit string = "credit"
)

// prop value enum
func (m *ChangeHistoryAffordabilityData) validateDirectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, changeHistoryAffordabilityDataTypeDirectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ChangeHistoryAffordabilityData) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("direction", "body", m.Direction); err != nil {
		return err
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", *m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *ChangeHistoryAffordabilityData) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChangeHistoryAffordabilityData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangeHistoryAffordabilityData) UnmarshalBinary(b []byte) error {
	var res ChangeHistoryAffordabilityData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
