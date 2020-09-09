// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IncomePost income post
//
// swagger:model IncomePost
type IncomePost struct {

	// The list of accounts to be included in the income otherwise all the user accounts will be included in the report
	Accounts []string `json:"accounts"`

	// The first/start month to be included in the affordability output e.g. "fromMonth":"2019-05". Resulting income resource will be based on data between fromMonth and toMonth
	FromMonth string `json:"fromMonth,omitempty"`

	// The first/start month to be included in the affordability output e.g. "fromMonth":"2019-05". Resulting income resource will be based on data between fromMonth and toMonth
	ToMonth string `json:"toMonth,omitempty"`
}

// Validate validates this income post
func (m *IncomePost) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IncomePost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IncomePost) UnmarshalBinary(b []byte) error {
	var res IncomePost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
