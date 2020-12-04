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

// AccountsData accounts data
//
// swagger:model AccountsData
type AccountsData struct {

	// Account number
	// Required: true
	AccountNo *string `json:"accountNo"`

	// Account available funds, nullable.
	// Required: true
	AvailableFunds *string `json:"availableFunds"`

	// Account balance, nullable.
	// Required: true
	Balance *string `json:"balance"`

	// class
	// Required: true
	Class *AccountClass `json:"class"`

	// Currency
	// Required: true
	Currency *string `json:"currency"`

	// Account identification.
	// Required: true
	ID *string `json:"id"`

	// Account last updated date and time.
	// Required: true
	LastUpdated *string `json:"lastUpdated"`

	// links
	// Required: true
	Links *ConnectionAccountLinks `json:"links"`

	// Account name.
	// Required: true
	Name *string `json:"name"`

	// Account status.
	// Required: true
	// Enum: [available unavailable]
	Status *string `json:"status"`

	// Type always "account".
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this accounts data
func (m *AccountsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailableFunds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountsData) validateAccountNo(formats strfmt.Registry) error {

	if err := validate.Required("accountNo", "body", m.AccountNo); err != nil {
		return err
	}

	return nil
}

func (m *AccountsData) validateAvailableFunds(formats strfmt.Registry) error {

	if err := validate.Required("availableFunds", "body", m.AvailableFunds); err != nil {
		return err
	}

	return nil
}

func (m *AccountsData) validateBalance(formats strfmt.Registry) error {

	if err := validate.Required("balance", "body", m.Balance); err != nil {
		return err
	}

	return nil
}

func (m *AccountsData) validateClass(formats strfmt.Registry) error {

	if err := validate.Required("class", "body", m.Class); err != nil {
		return err
	}

	if m.Class != nil {
		if err := m.Class.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("class")
			}
			return err
		}
	}

	return nil
}

func (m *AccountsData) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *AccountsData) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *AccountsData) validateLastUpdated(formats strfmt.Registry) error {

	if err := validate.Required("lastUpdated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *AccountsData) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("links", "body", m.Links); err != nil {
		return err
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

func (m *AccountsData) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var accountsDataTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["available","unavailable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountsDataTypeStatusPropEnum = append(accountsDataTypeStatusPropEnum, v)
	}
}

const (

	// AccountsDataStatusAvailable captures enum value "available"
	AccountsDataStatusAvailable string = "available"

	// AccountsDataStatusUnavailable captures enum value "unavailable"
	AccountsDataStatusUnavailable string = "unavailable"
)

// prop value enum
func (m *AccountsData) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountsDataTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccountsData) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *AccountsData) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountsData) UnmarshalBinary(b []byte) error {
	var res AccountsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
