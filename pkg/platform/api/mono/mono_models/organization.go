// Code generated by go-swagger; DO NOT EDIT.

package mono_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Organization organization
//
// swagger:model Organization
type Organization struct {

	// u r lname
	URLname string `json:"URLname,omitempty"`

	// add ons
	AddOns map[string]AddOn `json:"addOns,omitempty"`

	// added
	// Format: date-time
	Added strfmt.DateTime `json:"added,omitempty"`

	// billing date
	BillingDate *string `json:"billingDate,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// member count
	MemberCount int64 `json:"memberCount,omitempty"`

	// Deprecated; use displayName instead.
	Name string `json:"name,omitempty"`

	// organization ID
	// Format: uuid
	OrganizationID strfmt.UUID `json:"organizationID,omitempty"`

	// owner
	Owner bool `json:"owner,omitempty"`

	// personal
	Personal bool `json:"personal,omitempty"`

	// subscription status
	// Enum: [active cancelled expired]
	SubscriptionStatus *string `json:"subscriptionStatus,omitempty"`

	// tier
	Tier string `json:"tier,omitempty"`
}

// Validate validates this organization
func (m *Organization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddOns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Organization) validateAddOns(formats strfmt.Registry) error {

	if swag.IsZero(m.AddOns) { // not required
		return nil
	}

	for k := range m.AddOns {

		if err := validate.Required("addOns"+"."+k, "body", m.AddOns[k]); err != nil {
			return err
		}
		if val, ok := m.AddOns[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Organization) validateAdded(formats strfmt.Registry) error {

	if swag.IsZero(m.Added) { // not required
		return nil
	}

	if err := validate.FormatOf("added", "body", "date-time", m.Added.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateOrganizationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganizationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organizationID", "body", "uuid", m.OrganizationID.String(), formats); err != nil {
		return err
	}

	return nil
}

var organizationTypeSubscriptionStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","cancelled","expired"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		organizationTypeSubscriptionStatusPropEnum = append(organizationTypeSubscriptionStatusPropEnum, v)
	}
}

const (

	// OrganizationSubscriptionStatusActive captures enum value "active"
	OrganizationSubscriptionStatusActive string = "active"

	// OrganizationSubscriptionStatusCancelled captures enum value "cancelled"
	OrganizationSubscriptionStatusCancelled string = "cancelled"

	// OrganizationSubscriptionStatusExpired captures enum value "expired"
	OrganizationSubscriptionStatusExpired string = "expired"
)

// prop value enum
func (m *Organization) validateSubscriptionStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, organizationTypeSubscriptionStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Organization) validateSubscriptionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateSubscriptionStatusEnum("subscriptionStatus", "body", *m.SubscriptionStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Organization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Organization) UnmarshalBinary(b []byte) error {
	var res Organization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
