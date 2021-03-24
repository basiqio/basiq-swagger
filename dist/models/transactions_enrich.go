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

// TransactionsEnrich transactions enrich
//
// swagger:model TransactionsEnrich
type TransactionsEnrich struct {

	// category
	// Required: true
	Category *TransactionsCategory `json:"category"`

	// location
	// Required: true
	Location *TransactionsLocation `json:"location"`

	// merchant
	// Required: true
	Merchant *TransactionsMerchant `json:"merchant"`
}

// Validate validates this transactions enrich
func (m *TransactionsEnrich) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMerchant(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionsEnrich) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	if m.Category != nil {
		if err := m.Category.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category")
			}
			return err
		}
	}

	return nil
}

func (m *TransactionsEnrich) validateLocation(formats strfmt.Registry) error {

	if err := validate.Required("location", "body", m.Location); err != nil {
		return err
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *TransactionsEnrich) validateMerchant(formats strfmt.Registry) error {

	if err := validate.Required("merchant", "body", m.Merchant); err != nil {
		return err
	}

	if m.Merchant != nil {
		if err := m.Merchant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("merchant")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionsEnrich) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionsEnrich) UnmarshalBinary(b []byte) error {
	var res TransactionsEnrich
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
