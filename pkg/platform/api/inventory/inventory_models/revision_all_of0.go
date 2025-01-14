// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RevisionAllOf0 revision all of0
//
// swagger:model revisionAllOf0
type RevisionAllOf0 struct {

	// provided features
	// Required: true
	ProvidedFeatures []*ProvidedFeature `json:"provided_features"`
}

// Validate validates this revision all of0
func (m *RevisionAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvidedFeatures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RevisionAllOf0) validateProvidedFeatures(formats strfmt.Registry) error {

	if err := validate.Required("provided_features", "body", m.ProvidedFeatures); err != nil {
		return err
	}

	for i := 0; i < len(m.ProvidedFeatures); i++ {
		if swag.IsZero(m.ProvidedFeatures[i]) { // not required
			continue
		}

		if m.ProvidedFeatures[i] != nil {
			if err := m.ProvidedFeatures[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("provided_features" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RevisionAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RevisionAllOf0) UnmarshalBinary(b []byte) error {
	var res RevisionAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
