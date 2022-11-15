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

// ConnectorAuthorization connector authorization
//
// swagger:model ConnectorAuthorization
type ConnectorAuthorization struct {

	// meta
	Meta *ConnectorAuthorizationMeta `json:"meta,omitempty"`

	// Authorization type identifier
	// user AuthorizationUser  AuthorizationUser "User" authorization type identifier
	// token AuthorizationToken AuthorizationToken "Token" authorization type identifier
	// other AuthorizationOther  AuthorizationOther "Other" authorization type identifier
	// user-mfa AuthorizationUserMfa  AuthorizationUserMfa "UserMfa" authorization type identifier
	// user-mfa-intermittent AuthorizationUserMfaIntermittent  AuthorizationUserMfaIntermittent "UserMfaIntermittent" authorization type identifier
	// Example: user
	// Required: true
	// Enum: [user token other user-mfa user-mfa-intermittent]
	Type *string `json:"type"`
}

// Validate validates this connector authorization
func (m *ConnectorAuthorization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMeta(formats); err != nil {
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

func (m *ConnectorAuthorization) validateMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

var connectorAuthorizationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user","token","other","user-mfa","user-mfa-intermittent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectorAuthorizationTypeTypePropEnum = append(connectorAuthorizationTypeTypePropEnum, v)
	}
}

const (

	// ConnectorAuthorizationTypeUser captures enum value "user"
	ConnectorAuthorizationTypeUser string = "user"

	// ConnectorAuthorizationTypeToken captures enum value "token"
	ConnectorAuthorizationTypeToken string = "token"

	// ConnectorAuthorizationTypeOther captures enum value "other"
	ConnectorAuthorizationTypeOther string = "other"

	// ConnectorAuthorizationTypeUserDashMfa captures enum value "user-mfa"
	ConnectorAuthorizationTypeUserDashMfa string = "user-mfa"

	// ConnectorAuthorizationTypeUserDashMfaDashIntermittent captures enum value "user-mfa-intermittent"
	ConnectorAuthorizationTypeUserDashMfaDashIntermittent string = "user-mfa-intermittent"
)

// prop value enum
func (m *ConnectorAuthorization) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, connectorAuthorizationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConnectorAuthorization) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this connector authorization based on the context it is used
func (m *ConnectorAuthorization) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectorAuthorization) contextValidateMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.Meta != nil {
		if err := m.Meta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConnectorAuthorization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectorAuthorization) UnmarshalBinary(b []byte) error {
	var res ConnectorAuthorization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}