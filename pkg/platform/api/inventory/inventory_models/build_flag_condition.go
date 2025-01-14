// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuildFlagCondition Build Flag Condition
//
// A build flag name and value that must be included in a recipe for the containing entity to apply. If nothing in the recipe matches this condition, the containing entity is disable/cannot be used.
//
// swagger:model buildFlagCondition
type BuildFlagCondition struct {

	// Whether the build flag value must match this condition's value (eq) in order to satisfy the condition or must not match (ne)
	// Required: true
	// Enum: [eq ne]
	Comparator *string `json:"comparator"`

	// The name of the build flag conditioned upon
	// Required: true
	Name *string `json:"name"`

	// The value of the build flag conditioned upon
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this build flag condition
func (m *BuildFlagCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComparator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var buildFlagConditionTypeComparatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["eq","ne"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buildFlagConditionTypeComparatorPropEnum = append(buildFlagConditionTypeComparatorPropEnum, v)
	}
}

const (

	// BuildFlagConditionComparatorEq captures enum value "eq"
	BuildFlagConditionComparatorEq string = "eq"

	// BuildFlagConditionComparatorNe captures enum value "ne"
	BuildFlagConditionComparatorNe string = "ne"
)

// prop value enum
func (m *BuildFlagCondition) validateComparatorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, buildFlagConditionTypeComparatorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BuildFlagCondition) validateComparator(formats strfmt.Registry) error {

	if err := validate.Required("comparator", "body", m.Comparator); err != nil {
		return err
	}

	// value enum
	if err := m.validateComparatorEnum("comparator", "body", *m.Comparator); err != nil {
		return err
	}

	return nil
}

func (m *BuildFlagCondition) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *BuildFlagCondition) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildFlagCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildFlagCondition) UnmarshalBinary(b []byte) error {
	var res BuildFlagCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
