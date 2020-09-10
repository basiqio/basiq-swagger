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

// TransactionData transaction data
//
// swagger:model TransactionData
type TransactionData struct {

	// The id of the account resource the transaction belongs to.
	// Required: true
	Account *string `json:"account"`

	// Transaction amount. Outgoing funds are expressed as negative values.
	// Required: true
	Amount *string `json:"amount"`

	// Value of the account balance at time the transaction was completed.
	// Required: true
	Balance *string `json:"balance"`

	// Describes the class(type) of transaction.
	// Required: true
	// Enum: [bank-fee payment cash-withdrawal transfer loan-interest refund direct-cedit interest loan-repayment]
	Class *string `json:"class"`

	// The id of the connection resource that was used to retrieve the transaction.
	// Required: true
	Connection *string `json:"connection"`

	// The transaction description as submitted by the institution..
	// Required: true
	Description *string `json:"description"`

	// Identifies if the transaction is of debit or credit type.
	// Required: true
	// Enum: [debit credit]
	Direction *string `json:"direction"`

	// Uniquely identifies the transaction for this connection. Note that when a connection is refreshed pending transactions will receive new id's, whilst posted transactions will receive the same id's as before the refresh.
	// Required: true
	ID *string `json:"id"`

	// The id of the institution resource the transaction originated from.
	// Required: true
	Institution *string `json:"institution"`

	// Date the transaction was posted as provided by the institution (this is the same date that appears on a bank statement). This value is null if the record is pending. e.g. "2017-11-10T21:46:44Z" or 2017-11-10T00:00:00Z.
	// Required: true
	PostDate *string `json:"postDate"`

	// Identifies if a transaction is pending or posted. A pending transaction is an approved debit or credit transaction that has not been fully processed yet (i.e. has not been posted). Find out more about pending transaction and how to deal with them within your app. Note that pending transactions are not available for all institutions.
	// Required: true
	// Enum: [pending posted]
	Status *string `json:"status"`

	// Date that the user executed the transaction as provided by the istitution. Note that not all transactions provide this value (varies by institution) e.g. "2017-11-10T00:00:00Z"
	// Required: true
	TransactionDate *string `json:"transactionDate"`

	// Value is "transaction".
	// Required: true
	Type *string `json:"type"`

	// links
	// Required: true
	Links *TransactionLinks `json:"links"`

	// sub class
	// Required: true
	SubClass *SubClass `json:"subClass"`
}

// Validate validates this transaction data
func (m *TransactionData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstitution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubClass(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionData) validateAccount(formats strfmt.Registry) error {

	if err := validate.Required("account", "body", m.Account); err != nil {
		return err
	}

	return nil
}

func (m *TransactionData) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *TransactionData) validateBalance(formats strfmt.Registry) error {

	if err := validate.Required("balance", "body", m.Balance); err != nil {
		return err
	}

	return nil
}

var transactionDataTypeClassPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bank-fee","payment","cash-withdrawal","transfer","loan-interest","refund","direct-cedit","interest","loan-repayment"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transactionDataTypeClassPropEnum = append(transactionDataTypeClassPropEnum, v)
	}
}

const (

	// TransactionDataClassBankFee captures enum value "bank-fee"
	TransactionDataClassBankFee string = "bank-fee"

	// TransactionDataClassPayment captures enum value "payment"
	TransactionDataClassPayment string = "payment"

	// TransactionDataClassCashWithdrawal captures enum value "cash-withdrawal"
	TransactionDataClassCashWithdrawal string = "cash-withdrawal"

	// TransactionDataClassTransfer captures enum value "transfer"
	TransactionDataClassTransfer string = "transfer"

	// TransactionDataClassLoanInterest captures enum value "loan-interest"
	TransactionDataClassLoanInterest string = "loan-interest"

	// TransactionDataClassRefund captures enum value "refund"
	TransactionDataClassRefund string = "refund"

	// TransactionDataClassDirectCedit captures enum value "direct-cedit"
	TransactionDataClassDirectCedit string = "direct-cedit"

	// TransactionDataClassInterest captures enum value "interest"
	TransactionDataClassInterest string = "interest"

	// TransactionDataClassLoanRepayment captures enum value "loan-repayment"
	TransactionDataClassLoanRepayment string = "loan-repayment"
)

// prop value enum
func (m *TransactionData) validateClassEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transactionDataTypeClassPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TransactionData) validateClass(formats strfmt.Registry) error {

	if err := validate.Required("class", "body", m.Class); err != nil {
		return err
	}

	// value enum
	if err := m.validateClassEnum("class", "body", *m.Class); err != nil {
		return err
	}

	return nil
}

func (m *TransactionData) validateConnection(formats strfmt.Registry) error {

	if err := validate.Required("connection", "body", m.Connection); err != nil {
		return err
	}

	return nil
}

func (m *TransactionData) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

var transactionDataTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["debit","credit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transactionDataTypeDirectionPropEnum = append(transactionDataTypeDirectionPropEnum, v)
	}
}

const (

	// TransactionDataDirectionDebit captures enum value "debit"
	TransactionDataDirectionDebit string = "debit"

	// TransactionDataDirectionCredit captures enum value "credit"
	TransactionDataDirectionCredit string = "credit"
)

// prop value enum
func (m *TransactionData) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transactionDataTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TransactionData) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("direction", "body", m.Direction); err != nil {
		return err
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", *m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *TransactionData) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *TransactionData) validateInstitution(formats strfmt.Registry) error {

	if err := validate.Required("institution", "body", m.Institution); err != nil {
		return err
	}

	return nil
}

func (m *TransactionData) validatePostDate(formats strfmt.Registry) error {

	if err := validate.Required("postDate", "body", m.PostDate); err != nil {
		return err
	}

	return nil
}

var transactionDataTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","posted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transactionDataTypeStatusPropEnum = append(transactionDataTypeStatusPropEnum, v)
	}
}

const (

	// TransactionDataStatusPending captures enum value "pending"
	TransactionDataStatusPending string = "pending"

	// TransactionDataStatusPosted captures enum value "posted"
	TransactionDataStatusPosted string = "posted"
)

// prop value enum
func (m *TransactionData) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transactionDataTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TransactionData) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *TransactionData) validateTransactionDate(formats strfmt.Registry) error {

	if err := validate.Required("transactionDate", "body", m.TransactionDate); err != nil {
		return err
	}

	return nil
}

func (m *TransactionData) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *TransactionData) validateLinks(formats strfmt.Registry) error {

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

func (m *TransactionData) validateSubClass(formats strfmt.Registry) error {

	if err := validate.Required("subClass", "body", m.SubClass); err != nil {
		return err
	}

	if m.SubClass != nil {
		if err := m.SubClass.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subClass")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionData) UnmarshalBinary(b []byte) error {
	var res TransactionData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}