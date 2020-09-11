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

// JobData JobData
//
// Container object, containing job details.
//
// swagger:model JobData
type JobData struct {

	// The date time when the job was created.
	// Required: true
	Created *string `json:"created"`

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
	Updated *string `json:"updated"`

	// links
	Links *JobsLinks `json:"links,omitempty"`
}

// Validate validates this job data
func (m *JobData) Validate(formats strfmt.Registry) error {
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

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobData) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *JobData) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *JobData) validateSteps(formats strfmt.Registry) error {

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

func (m *JobData) validateUpdated(formats strfmt.Registry) error {

	if err := validate.Required("updated", "body", m.Updated); err != nil {
		return err
	}

	return nil
}

func (m *JobData) validateLinks(formats strfmt.Registry) error {

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
func (m *JobData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobData) UnmarshalBinary(b []byte) error {
	var res JobData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
