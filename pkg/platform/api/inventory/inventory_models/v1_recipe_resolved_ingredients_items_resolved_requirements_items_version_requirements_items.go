// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems Version Requirement Sub Schema
//
// The version constraint for a feature
// swagger:model v1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems
type V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems struct {

	// The operator used to compare the version against a given provided feature to determine if it meets this requirement
	// Required: true
	// Enum: [eq gt gte lt lte ne]
	Comparator *string `json:"comparator"`

	// The required version in its original form
	// Required: true
	// Min Length: 1
	Version *string `json:"version"`
}

// Validate validates this v1 recipe resolved ingredients items resolved requirements items version requirements items
func (m *V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComparator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsTypeComparatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["eq","gt","gte","lt","lte","ne"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsTypeComparatorPropEnum = append(v1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsTypeComparatorPropEnum, v)
	}
}

const (

	// V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorEq captures enum value "eq"
	V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorEq string = "eq"

	// V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorGt captures enum value "gt"
	V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorGt string = "gt"

	// V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorGte captures enum value "gte"
	V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorGte string = "gte"

	// V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorLt captures enum value "lt"
	V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorLt string = "lt"

	// V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorLte captures enum value "lte"
	V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorLte string = "lte"

	// V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorNe captures enum value "ne"
	V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorNe string = "ne"
)

// prop value enum
func (m *V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems) validateComparatorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsTypeComparatorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems) validateComparator(formats strfmt.Registry) error {

	if err := validate.Required("comparator", "body", m.Comparator); err != nil {
		return err
	}

	// value enum
	if err := m.validateComparatorEnum("comparator", "body", *m.Comparator); err != nil {
		return err
	}

	return nil
}

func (m *V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinLength("version", "body", string(*m.Version), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems) UnmarshalBinary(b []byte) error {
	var res V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
