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

// SubCategoryExpenses Payments Sub-Category
//
// swagger:model SubCategoryExpenses
type SubCategoryExpenses struct {

	// change history
	// Required: true
	ChangeHistory []*ChangeHistoryExpensesClass `json:"changeHistory"`

	// Summary period "monthly".
	// Required: true
	Summary *string `json:"summary"`

	// category
	// Required: true
	Category *CategoryDataExpenses `json:"category"`
}

// Validate validates this sub category expenses
func (m *SubCategoryExpenses) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChangeHistory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubCategoryExpenses) validateChangeHistory(formats strfmt.Registry) error {

	if err := validate.Required("changeHistory", "body", m.ChangeHistory); err != nil {
		return err
	}

	for i := 0; i < len(m.ChangeHistory); i++ {
		if swag.IsZero(m.ChangeHistory[i]) { // not required
			continue
		}

		if m.ChangeHistory[i] != nil {
			if err := m.ChangeHistory[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("changeHistory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SubCategoryExpenses) validateSummary(formats strfmt.Registry) error {

	if err := validate.Required("summary", "body", m.Summary); err != nil {
		return err
	}

	return nil
}

func (m *SubCategoryExpenses) validateCategory(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *SubCategoryExpenses) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubCategoryExpenses) UnmarshalBinary(b []byte) error {
	var res SubCategoryExpenses
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
