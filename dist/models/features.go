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

// Features Features stucture that describes institution features per data source
//
// swagger:model Features
type Features struct {

	// accounts
	Accounts *AccountsFeatures `json:"accounts,omitempty"`

	// Login holds list of data source identifiers which are capable to do complete login step.
	// This feature is applicable only on web sources.
	// Example: ["web"]
	// Required: true
	Login []SourceName `json:"login"`

	// mfa challenge
	MfaChallenge []SourceName `json:"mfaChallenge"`

	// profile
	Profile *ProfileFeatures `json:"profile,omitempty"`

	// transactions
	Transactions *TransactionsFeatures `json:"transactions,omitempty"`
}

// Validate validates this features
func (m *Features) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMfaChallenge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Features) validateAccounts(formats strfmt.Registry) error {
	if swag.IsZero(m.Accounts) { // not required
		return nil
	}

	if m.Accounts != nil {
		if err := m.Accounts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accounts")
			}
			return err
		}
	}

	return nil
}

func (m *Features) validateLogin(formats strfmt.Registry) error {

	if err := validate.Required("login", "body", m.Login); err != nil {
		return err
	}

	for i := 0; i < len(m.Login); i++ {

		if err := m.Login[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("login" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Features) validateMfaChallenge(formats strfmt.Registry) error {
	if swag.IsZero(m.MfaChallenge) { // not required
		return nil
	}

	for i := 0; i < len(m.MfaChallenge); i++ {

		if err := m.MfaChallenge[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mfaChallenge" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Features) validateProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.Profile) { // not required
		return nil
	}

	if m.Profile != nil {
		if err := m.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

func (m *Features) validateTransactions(formats strfmt.Registry) error {
	if swag.IsZero(m.Transactions) { // not required
		return nil
	}

	if m.Transactions != nil {
		if err := m.Transactions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transactions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this features based on the context it is used
func (m *Features) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMfaChallenge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransactions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Features) contextValidateAccounts(ctx context.Context, formats strfmt.Registry) error {

	if m.Accounts != nil {
		if err := m.Accounts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accounts")
			}
			return err
		}
	}

	return nil
}

func (m *Features) contextValidateLogin(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Login); i++ {

		if err := m.Login[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("login" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Features) contextValidateMfaChallenge(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MfaChallenge); i++ {

		if err := m.MfaChallenge[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mfaChallenge" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Features) contextValidateProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.Profile != nil {
		if err := m.Profile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

func (m *Features) contextValidateTransactions(ctx context.Context, formats strfmt.Registry) error {

	if m.Transactions != nil {
		if err := m.Transactions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transactions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Features) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Features) UnmarshalBinary(b []byte) error {
	var res Features
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
