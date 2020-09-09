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

// JobsData JobsData
//
// Container object, containing job details.
//
// swagger:model JobsData
type JobsData struct {

	// The date time when the job was created.
	// Required: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created"`

	// A string that uniquely identifies the job.
	// Required: true
	ID *string `json:"id"`

	// List of steps that need to be completed.
	// Required: true
	Steps []*Step `json:"steps"`

	// Value is "job".
	Type string `json:"type,omitempty"`

	// The date time when the job was last updated.
	// Required: true
	// Format: date-time
	Updated *strfmt.DateTime `json:"updated"`

	// institution
	// Required: true
	Institution *JobsInstitution `json:"institution"`

	// links
	Links *JobsLinks `json:"links,omitempty"`
}

// Validate validates this jobs data
func (m *JobsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstitution(formats); err != nil {
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

func (m *JobsData) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobsData) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *JobsData) validateSteps(formats strfmt.Registry) error {

	if err := validate.Required("steps", "body", m.Steps); err != nil {
		return err
	}

	for i := 0; i < len(m.Steps); i++ {
		if swag.IsZero(m.Steps[i]) { // not required
			continue
		}

		if m.Steps[i] != nil {
			if err := m.Steps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JobsData) validateUpdated(formats strfmt.Registry) error {

	if err := validate.Required("updated", "body", m.Updated); err != nil {
		return err
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobsData) validateInstitution(formats strfmt.Registry) error {

	if err := validate.Required("institution", "body", m.Institution); err != nil {
		return err
	}

	if m.Institution != nil {
		if err := m.Institution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("institution")
			}
			return err
		}
	}

	return nil
}

func (m *JobsData) validateLinks(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *JobsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsData) UnmarshalBinary(b []byte) error {
	var res JobsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
