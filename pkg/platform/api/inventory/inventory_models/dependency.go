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

// Dependency Dependency
//
// A single dependency for an ingredient version revision
//
// swagger:model dependency
type Dependency struct {

	// The features that must already be present in the recipe for this requirement to apply. For example, can be used to create requirements that only apply on specific operating systems.
	Conditions []*DependencyCondition `json:"conditions"`

	// The name of the feature this ingredient version is dependent on
	// Required: true
	Feature *string `json:"feature"`

	// The namespace the feature depended on is contained in
	// Required: true
	Namespace *string `json:"namespace"`

	// Whatever text or data structure we parsed to generate this dependency
	OriginalRequirement string `json:"original_requirement,omitempty"`

	// requirements
	// Required: true
	Requirements Requirements `json:"requirements"`
}

// Validate validates this dependency
func (m *Dependency) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeature(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequirements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Dependency) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Dependency) validateFeature(formats strfmt.Registry) error {

	if err := validate.Required("feature", "body", m.Feature); err != nil {
		return err
	}

	return nil
}

func (m *Dependency) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *Dependency) validateRequirements(formats strfmt.Registry) error {

	if err := validate.Required("requirements", "body", m.Requirements); err != nil {
		return err
	}

	if err := m.Requirements.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("requirements")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Dependency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Dependency) UnmarshalBinary(b []byte) error {
	var res Dependency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
