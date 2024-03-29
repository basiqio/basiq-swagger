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

// ConnectorInstitutionResource ConnectorInstitutionResource connector Institution  resource type
//
// swagger:model ConnectorInstitutionResource
type ConnectorInstitutionResource struct {

	// Institution country name
	// Example: Australia
	// Required: true
	Country *string `json:"country"`

	// logo
	// Required: true
	Logo *InstitutionLogoResource `json:"logo"`

	// Institution name
	// Example: Hooli Bank
	// Required: true
	Name *string `json:"name"`

	// Institution short name
	// Example: Hooli
	// Required: true
	ShortName *string `json:"shortName"`

	// Institution tier identifier
	// 1 TierOne  TierOne tier identifier for tier1 institution
	// 2 TierTwo  TierTwo tier identifier for tier2 institution
	// 3 TierThree  TierThree tier identifier for tier3 institution
	// 4 TierFour  TierFour tier identifier for tier4 institution
	// Example: 1
	// Required: true
	// Enum: [1 2 3 4]
	Tier *string `json:"tier"`

	// Institution type identifier
	// Bank BankInstitutionType  BankInstitutionType institution type identifier for Banks
	// Bank (Foreign) BankForeignInstitutionType  BankForeignInstitutionType institution type identifier for Foreign banks
	// Test Bank TestBankInstitutionType  TestBankInstitutionType institution type identifier for Test banks
	// Credit Union CreditUnionInstitutionType  CreditUnionInstitutionType institution type identifier for Credit union institutions
	// Financial Services FinancialServicesInstitutionType  FinancialServicesInstitutionType institution type identifier for Financial service institutions
	// Superannuation SuperannuationInstitutionType  SuperannuationInstitutionType institution type identifier for Superannuation institutions
	// Building Society BuildingSociety  BuildingSociety institution type identifier for Building Society institutions
	// Government Government  Government institution type identifier for Government institutions
	// Example: Bank
	// Required: true
	// Enum: [Bank Bank (Foreign) Test Bank Credit Union Financial Services Superannuation Building Society Government]
	Type *string `json:"type"`
}

// Validate validates this connector institution resource
func (m *ConnectorInstitutionResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShortName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTier(formats); err != nil {
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

func (m *ConnectorInstitutionResource) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

func (m *ConnectorInstitutionResource) validateLogo(formats strfmt.Registry) error {

	if err := validate.Required("logo", "body", m.Logo); err != nil {
		return err
	}

	if m.Logo != nil {
		if err := m.Logo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logo")
			}
			return err
		}
	}

	return nil
}

func (m *ConnectorInstitutionResource) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ConnectorInstitutionResource) validateShortName(formats strfmt.Registry) error {

	if err := validate.Required("shortName", "body", m.ShortName); err != nil {
		return err
	}

	return nil
}

var connectorInstitutionResourceTypeTierPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1","2","3","4"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectorInstitutionResourceTypeTierPropEnum = append(connectorInstitutionResourceTypeTierPropEnum, v)
	}
}

const (

	// ConnectorInstitutionResourceTierNr1 captures enum value "1"
	ConnectorInstitutionResourceTierNr1 string = "1"

	// ConnectorInstitutionResourceTierNr2 captures enum value "2"
	ConnectorInstitutionResourceTierNr2 string = "2"

	// ConnectorInstitutionResourceTierNr3 captures enum value "3"
	ConnectorInstitutionResourceTierNr3 string = "3"

	// ConnectorInstitutionResourceTierNr4 captures enum value "4"
	ConnectorInstitutionResourceTierNr4 string = "4"
)

// prop value enum
func (m *ConnectorInstitutionResource) validateTierEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, connectorInstitutionResourceTypeTierPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConnectorInstitutionResource) validateTier(formats strfmt.Registry) error {

	if err := validate.Required("tier", "body", m.Tier); err != nil {
		return err
	}

	// value enum
	if err := m.validateTierEnum("tier", "body", *m.Tier); err != nil {
		return err
	}

	return nil
}

var connectorInstitutionResourceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Bank","Bank (Foreign)","Test Bank","Credit Union","Financial Services","Superannuation","Building Society","Government"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectorInstitutionResourceTypeTypePropEnum = append(connectorInstitutionResourceTypeTypePropEnum, v)
	}
}

const (

	// ConnectorInstitutionResourceTypeBank captures enum value "Bank"
	ConnectorInstitutionResourceTypeBank string = "Bank"

	// ConnectorInstitutionResourceTypeBankForeign captures enum value "Bank (Foreign)"
	ConnectorInstitutionResourceTypeBankForeign string = "Bank (Foreign)"

	// ConnectorInstitutionResourceTypeTestBank captures enum value "Test Bank"
	ConnectorInstitutionResourceTypeTestBank string = "Test Bank"

	// ConnectorInstitutionResourceTypeCreditUnion captures enum value "Credit Union"
	ConnectorInstitutionResourceTypeCreditUnion string = "Credit Union"

	// ConnectorInstitutionResourceTypeFinancialServices captures enum value "Financial Services"
	ConnectorInstitutionResourceTypeFinancialServices string = "Financial Services"

	// ConnectorInstitutionResourceTypeSuperannuation captures enum value "Superannuation"
	ConnectorInstitutionResourceTypeSuperannuation string = "Superannuation"

	// ConnectorInstitutionResourceTypeBuildingSociety captures enum value "Building Society"
	ConnectorInstitutionResourceTypeBuildingSociety string = "Building Society"

	// ConnectorInstitutionResourceTypeGovernment captures enum value "Government"
	ConnectorInstitutionResourceTypeGovernment string = "Government"
)

// prop value enum
func (m *ConnectorInstitutionResource) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, connectorInstitutionResourceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConnectorInstitutionResource) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this connector institution resource based on the context it is used
func (m *ConnectorInstitutionResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectorInstitutionResource) contextValidateLogo(ctx context.Context, formats strfmt.Registry) error {

	if m.Logo != nil {
		if err := m.Logo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConnectorInstitutionResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectorInstitutionResource) UnmarshalBinary(b []byte) error {
	var res ConnectorInstitutionResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
