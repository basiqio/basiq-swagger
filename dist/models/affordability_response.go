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

// AffordabilityResponse AffordabilityResponse
//
// Container object, containing affordability details.
//
// The affordability object includes a summary of financial position, assets, liabilities, with links to an income object and an expenses object, for an individual user for account and transaction data stored against that user.
//
// swagger:model AffordabilityResponse
type AffordabilityResponse struct {

	// Assets
	// Required: true
	Assets []*AffordabilityAssetsData `json:"assets"`

	// Number of days covered by the report
	CoverageDays int64 `json:"coverageDays,omitempty"`

	// External
	// Required: true
	External []*ExternalLiabilityData `json:"external"`

	// Start month for the period for which the Affordability summary is generated. The period of time relates to the account and transaction data used as input into the report.
	// Required: true
	FromMonth *string `json:"fromMonth"`

	// Date the report was generated.
	// Required: true
	GeneratedDate *string `json:"generatedDate"`

	// Uniquely identifies the affordability report.
	// Required: true
	ID *string `json:"id"`

	// liabilities
	// Required: true
	Liabilities *LiabilitiesData `json:"liabilities"`

	// links
	// Required: true
	Links *GetAffordabilityLinks `json:"links"`

	// summary
	// Required: true
	Summary *AffordabilitySummary `json:"summary"`

	// End month (usually the current month) for the period for which the Affordability summary is generated.
	// Required: true
	ToMonth *string `json:"toMonth"`

	// Always "affordability".
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this affordability response
func (m *AffordabilityResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromMonth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeneratedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLiabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToMonth(formats); err != nil {
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

func (m *AffordabilityResponse) validateAssets(formats strfmt.Registry) error {

	if err := validate.Required("assets", "body", m.Assets); err != nil {
		return err
	}

	for i := 0; i < len(m.Assets); i++ {
		if swag.IsZero(m.Assets[i]) { // not required
			continue
		}

		if m.Assets[i] != nil {
			if err := m.Assets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AffordabilityResponse) validateExternal(formats strfmt.Registry) error {

	if err := validate.Required("external", "body", m.External); err != nil {
		return err
	}

	for i := 0; i < len(m.External); i++ {
		if swag.IsZero(m.External[i]) { // not required
			continue
		}

		if m.External[i] != nil {
			if err := m.External[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AffordabilityResponse) validateFromMonth(formats strfmt.Registry) error {

	if err := validate.Required("fromMonth", "body", m.FromMonth); err != nil {
		return err
	}

	return nil
}

func (m *AffordabilityResponse) validateGeneratedDate(formats strfmt.Registry) error {

	if err := validate.Required("generatedDate", "body", m.GeneratedDate); err != nil {
		return err
	}

	return nil
}

func (m *AffordabilityResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *AffordabilityResponse) validateLiabilities(formats strfmt.Registry) error {

	if err := validate.Required("liabilities", "body", m.Liabilities); err != nil {
		return err
	}

	if m.Liabilities != nil {
		if err := m.Liabilities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("liabilities")
			}
			return err
		}
	}

	return nil
}

func (m *AffordabilityResponse) validateLinks(formats strfmt.Registry) error {

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

func (m *AffordabilityResponse) validateSummary(formats strfmt.Registry) error {

	if err := validate.Required("summary", "body", m.Summary); err != nil {
		return err
	}

	if m.Summary != nil {
		if err := m.Summary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summary")
			}
			return err
		}
	}

	return nil
}

func (m *AffordabilityResponse) validateToMonth(formats strfmt.Registry) error {

	if err := validate.Required("toMonth", "body", m.ToMonth); err != nil {
		return err
	}

	return nil
}

func (m *AffordabilityResponse) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AffordabilityResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AffordabilityResponse) UnmarshalBinary(b []byte) error {
	var res AffordabilityResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
