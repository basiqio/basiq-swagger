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

// ExternalPayment external payment
//
// swagger:model ExternalPayment
type ExternalPayment struct {

	// Number of occurrences for same source (in this group)
	// Example: -50.50
	// Required: true
	AmountAvg *string `json:"amountAvg"`

	// Average monthly payment amount
	// Example: -20.00
	// Required: true
	AmountAvgMonthly *string `json:"amountAvgMonthly"`

	// Date of first payment
	// Example: 2019-03-15
	// Required: true
	// Format: date
	First *strfmt.Date `json:"first"`

	// Date of last payment
	// Example: 2020-03-15
	// Required: true
	// Format: date
	Last *strfmt.Date `json:"last"`

	// Number of occurrences for same source (in this group)
	// Example: 2
	// Required: true
	NoOccurrences *int64 `json:"noOccurrences"`

	// Amount of total payments identified for source in the affordability snapshot
	// Example: -146.50
	// Required: true
	Total *string `json:"total"`
}

// Validate validates this external payment
func (m *ExternalPayment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmountAvg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmountAvgMonthly(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLast(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNoOccurrences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalPayment) validateAmountAvg(formats strfmt.Registry) error {

	if err := validate.Required("amountAvg", "body", m.AmountAvg); err != nil {
		return err
	}

	return nil
}

func (m *ExternalPayment) validateAmountAvgMonthly(formats strfmt.Registry) error {

	if err := validate.Required("amountAvgMonthly", "body", m.AmountAvgMonthly); err != nil {
		return err
	}

	return nil
}

func (m *ExternalPayment) validateFirst(formats strfmt.Registry) error {

	if err := validate.Required("first", "body", m.First); err != nil {
		return err
	}

	if err := validate.FormatOf("first", "body", "date", m.First.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExternalPayment) validateLast(formats strfmt.Registry) error {

	if err := validate.Required("last", "body", m.Last); err != nil {
		return err
	}

	if err := validate.FormatOf("last", "body", "date", m.Last.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExternalPayment) validateNoOccurrences(formats strfmt.Registry) error {

	if err := validate.Required("noOccurrences", "body", m.NoOccurrences); err != nil {
		return err
	}

	return nil
}

func (m *ExternalPayment) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("total", "body", m.Total); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this external payment based on context it is used
func (m *ExternalPayment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExternalPayment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalPayment) UnmarshalBinary(b []byte) error {
	var res ExternalPayment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
