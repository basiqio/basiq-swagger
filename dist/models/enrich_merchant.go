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

// EnrichMerchant enrich merchant
//
// swagger:model EnrichMerchant
type EnrichMerchant struct {

	// Merchant name
	// Example: Garfish Manly
	// Required: true
	BusinessName *string `json:"businessName"`

	// phone number
	PhoneNumber *EnrichPhoneNumber `json:"phoneNumber,omitempty"`

	// Merchant Website
	// Example: http://garfish.com.au/garfish-manly/
	// Required: true
	Website *string `json:"website"`
}

// Validate validates this enrich merchant
func (m *EnrichMerchant) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBusinessName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebsite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnrichMerchant) validateBusinessName(formats strfmt.Registry) error {

	if err := validate.Required("businessName", "body", m.BusinessName); err != nil {
		return err
	}

	return nil
}

func (m *EnrichMerchant) validatePhoneNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.PhoneNumber) { // not required
		return nil
	}

	if m.PhoneNumber != nil {
		if err := m.PhoneNumber.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phoneNumber")
			}
			return err
		}
	}

	return nil
}

func (m *EnrichMerchant) validateWebsite(formats strfmt.Registry) error {

	if err := validate.Required("website", "body", m.Website); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnrichMerchant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnrichMerchant) UnmarshalBinary(b []byte) error {
	var res EnrichMerchant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
