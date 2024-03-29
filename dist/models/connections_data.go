// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConnectionsData connections data
//
// swagger:model ConnectionsData
type ConnectionsData struct {

	// Created date of the connection, available only for SERVER_SCOPE.
	// Example: 2019-07-29T07:34:09Z
	// Required: true
	CreatedDate *string `json:"createdDate"`

	// Connection identification.
	// Example: 61723
	// Required: true
	ID *string `json:"id"`

	// institution
	// Required: true
	Institution *ConnectionInstitution `json:"institution"`

	// Connection last used date, available only for SERVER_SCOPE.
	// Example: 2020-06-22T11:15:09Z
	LastUsed string `json:"lastUsed,omitempty"`

	// links
	// Required: true
	Links *GetConnectionsLinks `json:"links"`

	// Indicates whether MFA (multi factor authentication) is enabled for this connection. Where the value is true then expect an additional step in the Jobs response. Otherwise value is false.
	MfaEnabled bool `json:"mfaEnabled,omitempty"`

	// Connection status, available only for SERVER_SCOPE.
	// Example: active
	// Enum: [active pending invalid]
	Status string `json:"status,omitempty"`

	// Type, always "connection".
	// Example: connection
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this connections data
func (m *ConnectionsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstitution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *ConnectionsData) validateCreatedDate(formats strfmt.Registry) error {

	if err := validate.Required("createdDate", "body", m.CreatedDate); err != nil {
		return err
	}

	return nil
}

func (m *ConnectionsData) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ConnectionsData) validateInstitution(formats strfmt.Registry) error {

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

func (m *ConnectionsData) validateLinks(formats strfmt.Registry) error {

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

var connectionsDataTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","pending","invalid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectionsDataTypeStatusPropEnum = append(connectionsDataTypeStatusPropEnum, v)
	}
}

const (

	// ConnectionsDataStatusActive captures enum value "active"
	ConnectionsDataStatusActive string = "active"

	// ConnectionsDataStatusPending captures enum value "pending"
	ConnectionsDataStatusPending string = "pending"

	// ConnectionsDataStatusInvalid captures enum value "invalid"
	ConnectionsDataStatusInvalid string = "invalid"
)

// prop value enum
func (m *ConnectionsData) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, connectionsDataTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConnectionsData) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ConnectionsData) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this connections data based on the context it is used
func (m *ConnectionsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstitution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectionsData) contextValidateInstitution(ctx context.Context, formats strfmt.Registry) error {

	if m.Institution != nil {
		if err := m.Institution.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("institution")
			}
			return err
		}
	}

	return nil
}

func (m *ConnectionsData) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConnectionsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectionsData) UnmarshalBinary(b []byte) error {
	var res ConnectionsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
