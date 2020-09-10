// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExpensesResponse ExpensesResponse
//
// Container object, containing expenses details.
//
// The Expenses object returns an aggregated list of expenses by category of spend.
//
// swagger:model ExpensesResponse
type ExpensesResponse struct {

	// First 'month' occurrence of expenses categorised going back as far as 13 months.
	// Required: true
	FromMonth *string `json:"fromMonth"`

	// Uniquely identifies the expenses report.
	// Required: true
	ID *string `json:"id"`

	// payments
	// Required: true
	Payments []*PaymentsSummaryExpenses `json:"payments"`

	// Latest 'month' occurrence of expenses categorised.
	// Required: true
	ToMonth *string `json:"toMonth"`

	// Value of this resource is "expenses".
	// Required: true
	Type *string `json:"type"`

	// bank fees
	// Required: true
	BankFees *ClassResourceExpenses `json:"bankFees"`

	// cash withdrawals
	// Required: true
	CashWithdrawals *ClassResourceExpenses `json:"cashWithdrawals"`

	// external transfers
	// Required: true
	ExternalTransfers *ClassResourceExpenses `json:"externalTransfers"`

	// links
	Links *ExpensesLinks `json:"links,omitempty"`

	// loan interests
	// Required: true
	LoanInterests *ClassResourceExpenses `json:"loanInterests"`
}

// Validate validates this expenses response
func (m *ExpensesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFromMonth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToMonth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankFees(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCashWithdrawals(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalTransfers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoanInterests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExpensesResponse) validateFromMonth(formats strfmt.Registry) error {

	if err := validate.Required("fromMonth", "body", m.FromMonth); err != nil {
		return err
	}

	return nil
}

func (m *ExpensesResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ExpensesResponse) validatePayments(formats strfmt.Registry) error {

	if err := validate.Required("payments", "body", m.Payments); err != nil {
		return err
	}

	for i := 0; i < len(m.Payments); i++ {
		if swag.IsZero(m.Payments[i]) { // not required
			continue
		}

		if m.Payments[i] != nil {
			if err := m.Payments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExpensesResponse) validateToMonth(formats strfmt.Registry) error {

	if err := validate.Required("toMonth", "body", m.ToMonth); err != nil {
		return err
	}

	return nil
}

func (m *ExpensesResponse) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ExpensesResponse) validateBankFees(formats strfmt.Registry) error {

	if err := validate.Required("bankFees", "body", m.BankFees); err != nil {
		return err
	}

	if m.BankFees != nil {
		if err := m.BankFees.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bankFees")
			}
			return err
		}
	}

	return nil
}

func (m *ExpensesResponse) validateCashWithdrawals(formats strfmt.Registry) error {

	if err := validate.Required("cashWithdrawals", "body", m.CashWithdrawals); err != nil {
		return err
	}

	if m.CashWithdrawals != nil {
		if err := m.CashWithdrawals.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cashWithdrawals")
			}
			return err
		}
	}

	return nil
}

func (m *ExpensesResponse) validateExternalTransfers(formats strfmt.Registry) error {

	if err := validate.Required("externalTransfers", "body", m.ExternalTransfers); err != nil {
		return err
	}

	if m.ExternalTransfers != nil {
		if err := m.ExternalTransfers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalTransfers")
			}
			return err
		}
	}

	return nil
}

func (m *ExpensesResponse) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
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

func (m *ExpensesResponse) validateLoanInterests(formats strfmt.Registry) error {

	if err := validate.Required("loanInterests", "body", m.LoanInterests); err != nil {
		return err
	}

	if m.LoanInterests != nil {
		if err := m.LoanInterests.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loanInterests")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExpensesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExpensesResponse) UnmarshalBinary(b []byte) error {
	var res ExpensesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}