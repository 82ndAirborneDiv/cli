// Code generated by go-swagger; DO NOT EDIT.

package mono_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BillingInformation billing information
//
// swagger:model BillingInformation
type BillingInformation struct {

	// address
	Address *AddressInfo `json:"address,omitempty"`

	// billing date
	// Format: date
	BillingDate *strfmt.Date `json:"billingDate,omitempty"`

	// cc exp month
	CcExpMonth *string `json:"ccExpMonth,omitempty"`

	// cc exp year
	CcExpYear *string `json:"ccExpYear,omitempty"`

	// cc last4
	CcLast4 *string `json:"ccLast4,omitempty"`

	// customer ID
	CustomerID *string `json:"customerID,omitempty"`

	// email
	Email *string `json:"email,omitempty"`

	// external URL
	ExternalURL string `json:"externalURL,omitempty"`

	// is stripe customer
	IsStripeCustomer bool `json:"isStripeCustomer,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// reseller managed
	ResellerManaged bool `json:"resellerManaged,omitempty"`

	// self serve
	SelfServe bool `json:"selfServe,omitempty"`
}

// Validate validates this billing information
func (m *BillingInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingInformation) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *BillingInformation) validateBillingDate(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingDate) { // not required
		return nil
	}

	if err := validate.FormatOf("billingDate", "body", "date", m.BillingDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingInformation) UnmarshalBinary(b []byte) error {
	var res BillingInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
