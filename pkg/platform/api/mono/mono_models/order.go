// Code generated by go-swagger; DO NOT EDIT.

package mono_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Order order
//
// swagger:model Order
type Order struct {

	// Flags for camel
	CamelFlags []string `json:"camel_flags"`

	// Identifier for the order, currently arbitrary
	// Format: uuid
	OrderID strfmt.UUID `json:"order_id,omitempty"`

	// inventory API's UUIDs of requested platforms
	Platforms []strfmt.UUID `json:"platforms"`

	// requirements
	Requirements []*SolverRequirement `json:"requirements"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatforms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequirements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Order) validateOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderID) { // not required
		return nil
	}

	if err := validate.FormatOf("order_id", "body", "uuid", m.OrderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validatePlatforms(formats strfmt.Registry) error {

	if swag.IsZero(m.Platforms) { // not required
		return nil
	}

	for i := 0; i < len(m.Platforms); i++ {

		if err := validate.FormatOf("platforms"+"."+strconv.Itoa(i), "body", "uuid", m.Platforms[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *Order) validateRequirements(formats strfmt.Registry) error {

	if swag.IsZero(m.Requirements) { // not required
		return nil
	}

	for i := 0; i < len(m.Requirements); i++ {
		if swag.IsZero(m.Requirements[i]) { // not required
			continue
		}

		if m.Requirements[i] != nil {
			if err := m.Requirements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requirements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Order) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
