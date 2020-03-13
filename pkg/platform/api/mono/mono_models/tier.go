// Code generated by go-swagger; DO NOT EDIT.

package mono_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Tier tier
// swagger:model Tier
type Tier struct {

	// description
	Description string `json:"description,omitempty"`

	// grants security and compliance
	GrantsSecurityAndCompliance bool `json:"grantsSecurityAndCompliance,omitempty"`

	// is default
	IsDefault bool `json:"isDefault,omitempty"`

	// is personal
	IsPersonal bool `json:"isPersonal,omitempty"`

	// limit minimum users
	LimitMinimumUsers int64 `json:"limitMinimumUsers,omitempty"`

	// limit node users threshold
	LimitNodeUsersThreshold int64 `json:"limitNodeUsersThreshold,omitempty"`

	// limit nodes per threshold
	LimitNodesPerThreshold int64 `json:"limitNodesPerThreshold,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// requires payment
	RequiresPayment bool `json:"requiresPayment,omitempty"`

	// user selectable
	UserSelectable bool `json:"userSelectable,omitempty"`
}

// Validate validates this tier
func (m *Tier) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Tier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tier) UnmarshalBinary(b []byte) error {
	var res Tier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
