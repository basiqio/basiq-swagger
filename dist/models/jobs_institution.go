// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JobsInstitution Institution
//
// Institution details.
//
// swagger:model JobsInstitution
type JobsInstitution struct {

	// A string that uniquely identifies institution.
	// Example: AU00000
	// Required: true
	ID *string `json:"id"`

	// links
	// Required: true
	Links *JobsLinks `json:"links"`

	// Always "institution".
	// Example: institution
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this jobs institution
func (m *JobsInstitution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
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

func (m *JobsInstitution) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *JobsInstitution) validateLinks(formats strfmt.Registry) error {

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

func (m *JobsInstitution) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobsInstitution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsInstitution) UnmarshalBinary(b []byte) error {
	var res JobsInstitution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
