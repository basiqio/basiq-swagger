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

// AuthLinksResponseResource auth links response resource
//
// swagger:model AuthLinksResponseResource
type AuthLinksResponseResource struct {

	// The date time of auth link expiry.
	// Required: true
	// Format: date-time
	ExpiresAt *strfmt.DateTime `json:"expiresAt"`

	// Uniquely identifies the auth link.
	// Required: true
	ID *string `json:"id"`

	// A user's mobile phone, used as for authentication.
	// Required: true
	Mobile *string `json:"mobile"`

	// Type of the response, always "auth_link".
	// Required: true
	Type *string `json:"type"`

	// A string that uniquely identifies the user.
	// Required: true
	UserID *string `json:"userId"`

	// links
	Links *AuthLinkLinks `json:"links,omitempty"`
}

// Validate validates this auth links response resource
func (m *AuthLinksResponseResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMobile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
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

func (m *AuthLinksResponseResource) validateExpiresAt(formats strfmt.Registry) error {

	if err := validate.Required("expiresAt", "body", m.ExpiresAt); err != nil {
		return err
	}

	if err := validate.FormatOf("expiresAt", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuthLinksResponseResource) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *AuthLinksResponseResource) validateMobile(formats strfmt.Registry) error {

	if err := validate.Required("mobile", "body", m.Mobile); err != nil {
		return err
	}

	return nil
}

func (m *AuthLinksResponseResource) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *AuthLinksResponseResource) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

func (m *AuthLinksResponseResource) validateLinks(formats strfmt.Registry) error {

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
func (m *AuthLinksResponseResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthLinksResponseResource) UnmarshalBinary(b []byte) error {
	var res AuthLinksResponseResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
