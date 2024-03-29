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

// Previous6MonthsCreditLiabilities previous6 months credit liabilities
//
// swagger:model Previous6MonthsCreditLiabilities
type Previous6MonthsCreditLiabilities struct {

	// Value of cash advances in period
	// Example: -2053.50
	// Required: true
	CashAdvances *string `json:"cashAdvances"`
}

// Validate validates this previous6 months credit liabilities
func (m *Previous6MonthsCreditLiabilities) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCashAdvances(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Previous6MonthsCreditLiabilities) validateCashAdvances(formats strfmt.Registry) error {

	if err := validate.Required("cashAdvances", "body", m.CashAdvances); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this previous6 months credit liabilities based on context it is used
func (m *Previous6MonthsCreditLiabilities) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Previous6MonthsCreditLiabilities) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Previous6MonthsCreditLiabilities) UnmarshalBinary(b []byte) error {
	var res Previous6MonthsCreditLiabilities
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
