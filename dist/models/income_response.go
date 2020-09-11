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

// IncomeResponse IncomeResponse
//
// Container object, containing income details.
//
// The Income Summary object with income summarised and classified by type: regular, irregular or other sources (created/refreshed across up to 10 institutions)
//
// swagger:model IncomeResponse
type IncomeResponse struct {

	// Number of days covered by the report
	CoverageDays int64 `json:"coverageDays,omitempty"`

	// Start month for the period for which the Income summary is generated. The period of time relates to the account and transaction data used as input into the report.
	// Required: true
	FromMonth *string `json:"fromMonth"`

	// The identifier of the income resource to be retrieved.
	// Required: true
	ID *string `json:"id"`

	// Required true
	IrregularSources []*IrregularSource `json:"irregular"`

	// Required true
	RegularSources []*RegularSource `json:"regular"`

	// End month (usually the current month) for the period for which the Income summary is generated.
	// Required: true
	ToMonth *string `json:"toMonth"`

	// Always "income".
	// Required: true
	Type *string `json:"type"`

	// income summary
	// Required: true
	IncomeSummary *IncomeSummary `json:"incomeSummary"`

	// links
	Links *IncomeLinks `json:"links,omitempty"`

	// other credit
	OtherCredit *OtherCreditSource `json:"otherCredit,omitempty"`
}

// Validate validates this income response
func (m *IncomeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFromMonth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIrregularSources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegularSources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToMonth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncomeSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOtherCredit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IncomeResponse) validateFromMonth(formats strfmt.Registry) error {

	if err := validate.Required("fromMonth", "body", m.FromMonth); err != nil {
		return err
	}

	return nil
}

func (m *IncomeResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *IncomeResponse) validateIrregularSources(formats strfmt.Registry) error {

	if swag.IsZero(m.IrregularSources) { // not required
		return nil
	}

	for i := 0; i < len(m.IrregularSources); i++ {
		if swag.IsZero(m.IrregularSources[i]) { // not required
			continue
		}

		if m.IrregularSources[i] != nil {
			if err := m.IrregularSources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("irregular" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IncomeResponse) validateRegularSources(formats strfmt.Registry) error {

	if swag.IsZero(m.RegularSources) { // not required
		return nil
	}

	for i := 0; i < len(m.RegularSources); i++ {
		if swag.IsZero(m.RegularSources[i]) { // not required
			continue
		}

		if m.RegularSources[i] != nil {
			if err := m.RegularSources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("regular" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IncomeResponse) validateToMonth(formats strfmt.Registry) error {

	if err := validate.Required("toMonth", "body", m.ToMonth); err != nil {
		return err
	}

	return nil
}

func (m *IncomeResponse) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *IncomeResponse) validateIncomeSummary(formats strfmt.Registry) error {

	if err := validate.Required("incomeSummary", "body", m.IncomeSummary); err != nil {
		return err
	}

	if m.IncomeSummary != nil {
		if err := m.IncomeSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incomeSummary")
			}
			return err
		}
	}

	return nil
}

func (m *IncomeResponse) validateLinks(formats strfmt.Registry) error {

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

func (m *IncomeResponse) validateOtherCredit(formats strfmt.Registry) error {

	if swag.IsZero(m.OtherCredit) { // not required
		return nil
	}

	if m.OtherCredit != nil {
		if err := m.OtherCredit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("otherCredit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IncomeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IncomeResponse) UnmarshalBinary(b []byte) error {
	var res IncomeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
