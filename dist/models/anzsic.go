// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Anzsic anzsic
//
// swagger:model Anzsic
type Anzsic struct {

	// class
	Class *ClassDetails `json:"class,omitempty"`

	// division
	Division *DivisionDetails `json:"division,omitempty"`

	// group
	Group *GroupDetails `json:"group,omitempty"`

	// subdivision
	Subdivision *SubdivisionDetails `json:"subdivision,omitempty"`
}

// Validate validates this anzsic
func (m *Anzsic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubdivision(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Anzsic) validateClass(formats strfmt.Registry) error {

	if swag.IsZero(m.Class) { // not required
		return nil
	}

	if m.Class != nil {
		if err := m.Class.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("class")
			}
			return err
		}
	}

	return nil
}

func (m *Anzsic) validateDivision(formats strfmt.Registry) error {

	if swag.IsZero(m.Division) { // not required
		return nil
	}

	if m.Division != nil {
		if err := m.Division.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("division")
			}
			return err
		}
	}

	return nil
}

func (m *Anzsic) validateGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *Anzsic) validateSubdivision(formats strfmt.Registry) error {

	if swag.IsZero(m.Subdivision) { // not required
		return nil
	}

	if m.Subdivision != nil {
		if err := m.Subdivision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subdivision")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Anzsic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Anzsic) UnmarshalBinary(b []byte) error {
	var res Anzsic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}