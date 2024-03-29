// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountsFeatures AccountsFeatures describes set of institution accounts features
//
// swagger:model AccountsFeatures
type AccountsFeatures struct {

	// AccountHolder holds list of data source identifiers which are capable to fetch account holder.
	// Example: ["web","pdf","csv"]
	// Required: true
	AccountHolder []SourceName `json:"accountHolder"`

	// AccountNumber holds list of data source identifiers which are capable to fetch account number.
	// Example: ["web","pdf","csv"]
	// Required: true
	AccountNo []SourceName `json:"accountNo"`

	// AvailableFunds holds list of data source identifiers which are capable to fetch available funds.
	// Example: ["web","pdf","csv"]
	// Required: true
	AvailableFunds []SourceName `json:"availableFunds"`

	// AccountBalance holds list of data source identifiers which are capable to fetch account balance.
	// Example: ["web","pdf","csv"]
	// Required: true
	Balance []SourceName `json:"balance"`

	// AccountCurrency holds list of data source identifiers which are capable to fetch account currency.
	// Example: ["web","pdf","csv"]
	// Required: true
	Currency []SourceName `json:"currency"`

	// LastUpdated holds list of data source identifiers which are capable to fetch account last updated date.
	// Example: ["web","pdf","csv"]
	// Required: true
	LastUpdated []SourceName `json:"lastUpdated"`

	// Meta holds list of data source identifiers which are capable to fetch account meta data (e.g. mortgage data).
	// Example: ["web","pdf","csv"]
	// Required: true
	Meta []SourceName `json:"meta"`

	// AccountName holds list of data source identifiers which are capable to fetch account name.
	// Example: ["web","pdf","csv"]
	// Required: true
	Name []SourceName `json:"name"`
}

// Validate validates this accounts features
func (m *AccountsFeatures) Validate(formats strfmt.Registry) error {
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

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountsFeatures) validateAccountHolder(formats strfmt.Registry) error {

	if err := validate.Required("accountHolder", "body", m.AccountHolder); err != nil {
		return err
	}

	for i := 0; i < len(m.AccountHolder); i++ {

		if err := m.AccountHolder[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accountHolder" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AccountsFeatures) validateAccountNo(formats strfmt.Registry) error {

	if err := validate.Required("accountNo", "body", m.AccountNo); err != nil {
		return err
	}

	for i := 0; i < len(m.AccountNo); i++ {

		if err := m.AccountNo[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accountNo" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AccountsFeatures) validateAvailableFunds(formats strfmt.Registry) error {

	if err := validate.Required("availableFunds", "body", m.AvailableFunds); err != nil {
		return err
	}

	for i := 0; i < len(m.AvailableFunds); i++ {

		if err := m.AvailableFunds[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("availableFunds" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AccountsFeatures) validateBalance(formats strfmt.Registry) error {

	if err := validate.Required("balance", "body", m.Balance); err != nil {
		return err
	}

	for i := 0; i < len(m.Balance); i++ {

		if err := m.Balance[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balance" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AccountsFeatures) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	for i := 0; i < len(m.Currency); i++ {

		if err := m.Currency[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currency" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AccountsFeatures) validateLastUpdated(formats strfmt.Registry) error {

	if err := validate.Required("lastUpdated", "body", m.LastUpdated); err != nil {
		return err
	}

	for i := 0; i < len(m.LastUpdated); i++ {

		if err := m.LastUpdated[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastUpdated" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AccountsFeatures) validateMeta(formats strfmt.Registry) error {

	if err := validate.Required("meta", "body", m.Meta); err != nil {
		return err
	}

	for i := 0; i < len(m.Meta); i++ {

		if err := m.Meta[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AccountsFeatures) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	for i := 0; i < len(m.Name); i++ {

		if err := m.Name[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this accounts features based on the context it is used
func (m *AccountsFeatures) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccountHolder(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAccountNo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAvailableFunds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBalance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCurrency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountsFeatures) contextValidateAccountHolder(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AccountHolder); i++ {

		if err := m.AccountHolder[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accountHolder" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AccountsFeatures) contextValidateAccountNo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AccountNo); i++ {

		if err := m.AccountNo[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accountNo" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AccountsFeatures) contextValidateAvailableFunds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AvailableFunds); i++ {

		if err := m.AvailableFunds[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("availableFunds" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AccountsFeatures) contextValidateBalance(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Balance); i++ {

		if err := m.Balance[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balance" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AccountsFeatures) contextValidateCurrency(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Currency); i++ {

		if err := m.Currency[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currency" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AccountsFeatures) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LastUpdated); i++ {

		if err := m.LastUpdated[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastUpdated" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AccountsFeatures) contextValidateMeta(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Meta); i++ {

		if err := m.Meta[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AccountsFeatures) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Name); i++ {

		if err := m.Name[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountsFeatures) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountsFeatures) UnmarshalBinary(b []byte) error {
	var res AccountsFeatures
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
