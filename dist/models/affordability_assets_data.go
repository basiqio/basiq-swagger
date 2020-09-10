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

// AffordabilityAssetsData affordability assets data
//
// swagger:model AffordabilityAssetsData
type AffordabilityAssetsData struct {

	// The available funds at the time of the query.
	// Required: true
	AvailableFunds *string `json:"availableFunds"`

	// The currency in which the account is recorded.
	// Required: true
	Balance *string `json:"balance"`

	// The currency in which the account is recorded.
	// Required: true
	Currency *string `json:"currency"`

	// The name of the financial institution with whom the account is held.
	// Required: true
	Institution *string `json:"institution"`

	// Type account
	// Required: true
	Type *string `json:"type"`

	// account
	// Required: true
	Account *AccountHolder `json:"account"`

	// previous6 months
	// Required: true
	Previous6Months *AssetsPrevious6MonthsData `json:"previous6Months"`
}

// Validate validates this affordability assets data
func (m *AffordabilityAssetsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableFunds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstitution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrevious6Months(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AffordabilityAssetsData) validateAvailableFunds(formats strfmt.Registry) error {

	if err := validate.Required("availableFunds", "body", m.AvailableFunds); err != nil {
		return err
	}

	return nil
}

func (m *AffordabilityAssetsData) validateBalance(formats strfmt.Registry) error {

	if err := validate.Required("balance", "body", m.Balance); err != nil {
		return err
	}

	return nil
}

func (m *AffordabilityAssetsData) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *AffordabilityAssetsData) validateInstitution(formats strfmt.Registry) error {

	if err := validate.Required("institution", "body", m.Institution); err != nil {
		return err
	}

	return nil
}

func (m *AffordabilityAssetsData) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *AffordabilityAssetsData) validateAccount(formats strfmt.Registry) error {

	if err := validate.Required("account", "body", m.Account); err != nil {
		return err
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

func (m *AffordabilityAssetsData) validatePrevious6Months(formats strfmt.Registry) error {

	if err := validate.Required("previous6Months", "body", m.Previous6Months); err != nil {
		return err
	}

	if m.Previous6Months != nil {
		if err := m.Previous6Months.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("previous6Months")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AffordabilityAssetsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AffordabilityAssetsData) UnmarshalBinary(b []byte) error {
	var res AffordabilityAssetsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}