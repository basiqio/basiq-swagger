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

// ExternalLiabilityData external liability data
//
// swagger:model ExternalLiabilityData
type ExternalLiabilityData struct {

	// Each transaction (repeated for each source ordered by most recent)
	// Required: true
	ChangeHistory []*ChangeHistoryExternal `json:"changeHistory"`

	// Aggregated attributes relating to payments for this source (identified as an external liability)
	// Required: true
	Payments []*ExternalPayment `json:"payments"`

	// Source of external liability (cleaned transaction description).
	// Example: afterpay
	// Required: true
	Source *string `json:"source"`
}

// Validate validates this external liability data
func (m *ExternalLiabilityData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChangeHistory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalLiabilityData) validateChangeHistory(formats strfmt.Registry) error {

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

func (m *ExternalLiabilityData) validatePayments(formats strfmt.Registry) error {

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

func (m *ExternalLiabilityData) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this external liability data based on the context it is used
func (m *ExternalLiabilityData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChangeHistory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePayments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalLiabilityData) contextValidateChangeHistory(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ChangeHistory); i++ {

		if m.ChangeHistory[i] != nil {
			if err := m.ChangeHistory[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("changeHistory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExternalLiabilityData) contextValidatePayments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Payments); i++ {

		if m.Payments[i] != nil {
			if err := m.Payments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExternalLiabilityData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalLiabilityData) UnmarshalBinary(b []byte) error {
	var res ExternalLiabilityData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
