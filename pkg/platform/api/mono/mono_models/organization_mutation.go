// Code generated by go-swagger; DO NOT EDIT.

package mono_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OrganizationMutation organization mutation
// swagger:model OrganizationMutation
type OrganizationMutation struct {

	// the date and time at which this mutation was created
	// Format: date-time
	Added strfmt.DateTime `json:"added,omitempty"`

	// days added or subtracted from the next billing date
	BillingDateDelta int64 `json:"billingDateDelta,omitempty"`

	// Invoiced customer ID newly associated with the org as of this mutation
	CustomerID string `json:"customerID,omitempty"`

	// the ID for the Invoiced event that caused this mutation, if any
	EventID *int64 `json:"eventID,omitempty"`

	// change in an external invoice identifier
	InvoiceID string `json:"invoiceID,omitempty"`

	// production nodes added or subtracted
	NodesDelta int64 `json:"nodesDelta,omitempty"`

	// free-form notes relating to this mutation
	Notes string `json:"notes,omitempty"`

	// identifier for the mutated organization
	// Format: uuid
	OrganizationID strfmt.UUID `json:"organizationID,omitempty"`

	// change in customer PO number
	PoNumber string `json:"poNumber,omitempty"`

	// change in identifier for a reseller
	ResellerID string `json:"resellerID,omitempty"`

	// indicated whether customer org is managed by a reseller
	ResellerManaged *bool `json:"resellerManaged,omitempty"`

	// Invoiced subscription ID newly associated with the org as of this mutation
	SubscriptionID string `json:"subscriptionID,omitempty"`

	// change in status of the organization's subscription
	SubscriptionStatus string `json:"subscriptionStatus,omitempty"`

	// change of tier by name
	TierName string `json:"tierName,omitempty"`

	// identifier for the mutating user
	UserName string `json:"userName,omitempty"`

	// users added or substracted
	UsersDelta int64 `json:"usersDelta,omitempty"`
}

// Validate validates this organization mutation
func (m *OrganizationMutation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationMutation) validateAdded(formats strfmt.Registry) error {

	if swag.IsZero(m.Added) { // not required
		return nil
	}

	if err := validate.FormatOf("added", "body", "date-time", m.Added.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OrganizationMutation) validateOrganizationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganizationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organizationID", "body", "uuid", m.OrganizationID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationMutation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationMutation) UnmarshalBinary(b []byte) error {
	var res OrganizationMutation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}