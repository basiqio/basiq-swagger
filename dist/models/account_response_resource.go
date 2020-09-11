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

// AccountResponseResource AccountResponseResource
//
// Container object, containing account details.
//
// The account object represents an account held with a financial institution (e.g. a savings account). You can use this object to retrieve specific account details such as the account number, balance and available funds.
//
// swagger:model AccountResponseResource
type AccountResponseResource struct {

	// The name of the account holder as returned by the institution. No formatting is applied. Returns a string or null when not available.
	// Required: true
	AccountHolder *string `json:"accountHolder"`

	// Full account number.
	// Required: true
	AccountNo *string `json:"accountNo"`

	// Funds that are available to an account holder for withdrawal or other use. This may include funds from an overdraft facility or line of credit. As well as funds classified as the available balance, such as from cleared and existing deposits.
	// Required: true
	AvailableFunds *string `json:"availableFunds"`

	// Amount of funds in the account right now - excluding any pending transactions. For a credit card this would be zero or the minus amount spent.
	// Required: true
	Balance *string `json:"balance"`

	// The id of the connection resource that was used to retrieve the account.
	// Required: true
	Connection *string `json:"connection"`

	// The currency the funds are stored in, using ISO 4217 standard.
	// Required: true
	Currency *string `json:"currency"`

	// Uniquely identifies the account.
	// Required: true
	ID *string `json:"id"`

	// The id of the institution resource the account originated from.
	// Required: true
	Institution *string `json:"institution"`

	// Timestamp of last update, UTC, RFC 3339 format e.g. "2017-09-28T13:39:33Z"
	// Required: true
	LastUpdated *string `json:"lastUpdated"`

	// Account name as defined by institution or user.
	// Required: true
	Name *string `json:"name"`

	// Indicates the account status. Always set to 'available'. Field kept for backward compatibility. Possible values include:
	// <ul><li>available newest account data is available.</li></ul>
	// Required: true
	// Enum: [available unavailable]
	Status *string `json:"status"`

	// An array of date intervals indicating the coverage of the transaction data relating to the account.
	// Will return a single element for accounts sourced from a single bank connection.
	// Will return multiple elements in cases where there have been multiple PDF/CSV uploads for an account.
	// Required: true
	TransactionIntervals []*AccountTransactionInterval `json:"transactionIntervals"`

	// Always "account".
	// Required: true
	Type *string `json:"type"`

	// class
	// Required: true
	Class *ConnectionAccountClass `json:"class"`

	// links
	// Required: true
	Links *AccountLinks `json:"links"`
}

// Validate validates this account response resource
func (m *AccountResponseResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountHolder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountNo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailableFunds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstitution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionIntervals(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountResponseResource) validateAccountHolder(formats strfmt.Registry) error {

	if err := validate.Required("accountHolder", "body", m.AccountHolder); err != nil {
		return err
	}

	return nil
}

func (m *AccountResponseResource) validateAccountNo(formats strfmt.Registry) error {

	if err := validate.Required("accountNo", "body", m.AccountNo); err != nil {
		return err
	}

	return nil
}

func (m *AccountResponseResource) validateAvailableFunds(formats strfmt.Registry) error {

	if err := validate.Required("availableFunds", "body", m.AvailableFunds); err != nil {
		return err
	}

	return nil
}

func (m *AccountResponseResource) validateBalance(formats strfmt.Registry) error {

	if err := validate.Required("balance", "body", m.Balance); err != nil {
		return err
	}

	return nil
}

func (m *AccountResponseResource) validateConnection(formats strfmt.Registry) error {

	if err := validate.Required("connection", "body", m.Connection); err != nil {
		return err
	}

	return nil
}

func (m *AccountResponseResource) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *AccountResponseResource) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *AccountResponseResource) validateInstitution(formats strfmt.Registry) error {

	if err := validate.Required("institution", "body", m.Institution); err != nil {
		return err
	}

	return nil
}

func (m *AccountResponseResource) validateLastUpdated(formats strfmt.Registry) error {

	if err := validate.Required("lastUpdated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *AccountResponseResource) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var accountResponseResourceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["available","unavailable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountResponseResourceTypeStatusPropEnum = append(accountResponseResourceTypeStatusPropEnum, v)
	}
}

const (

	// AccountResponseResourceStatusAvailable captures enum value "available"
	AccountResponseResourceStatusAvailable string = "available"

	// AccountResponseResourceStatusUnavailable captures enum value "unavailable"
	AccountResponseResourceStatusUnavailable string = "unavailable"
)

// prop value enum
func (m *AccountResponseResource) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountResponseResourceTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccountResponseResource) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *AccountResponseResource) validateTransactionIntervals(formats strfmt.Registry) error {

	if err := validate.Required("transactionIntervals", "body", m.TransactionIntervals); err != nil {
		return err
	}

	for i := 0; i < len(m.TransactionIntervals); i++ {
		if swag.IsZero(m.TransactionIntervals[i]) { // not required
			continue
		}

		if m.TransactionIntervals[i] != nil {
			if err := m.TransactionIntervals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transactionIntervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountResponseResource) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *AccountResponseResource) validateClass(formats strfmt.Registry) error {

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

func (m *AccountResponseResource) validateLinks(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *AccountResponseResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountResponseResource) UnmarshalBinary(b []byte) error {
	var res AccountResponseResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
