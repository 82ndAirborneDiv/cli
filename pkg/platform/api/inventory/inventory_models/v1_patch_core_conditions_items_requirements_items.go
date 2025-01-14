// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1PatchCoreConditionsItemsRequirementsItems v1 patch core conditions items requirements items
// swagger:model v1PatchCoreConditionsItemsRequirementsItems
type V1PatchCoreConditionsItemsRequirementsItems struct {

	// The operator used to compare the sortable_version against a given provided feature to determine if it meets the requirement
	// Required: true
	// Enum: [eq gt gte lt lte ne]
	Comparator *string `json:"comparator"`

	// An array of decimal values representing all segments of a version, ordered from most to least significant. How a version string is rendered into a list of decimals will vary depending on the format of the source string and is therefore left up to the caller, but it must be done consistently across all versions of the same resource for sorting to work properly. This is represented as a string to avoid losing precision when converting to a floating point number.
	// Min Length: 1
	SortableVersion []string `json:"sortable_version"`

	// The required version in its original form.
	// Min Length: 1
	Version *string `json:"version,omitempty"`
}

// Validate validates this v1 patch core conditions items requirements items
func (m *V1PatchCoreConditionsItemsRequirementsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComparator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSortableVersion(formats); err != nil {
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

var v1PatchCoreConditionsItemsRequirementsItemsTypeComparatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["eq","gt","gte","lt","lte","ne"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1PatchCoreConditionsItemsRequirementsItemsTypeComparatorPropEnum = append(v1PatchCoreConditionsItemsRequirementsItemsTypeComparatorPropEnum, v)
	}
}

const (

	// V1PatchCoreConditionsItemsRequirementsItemsComparatorEq captures enum value "eq"
	V1PatchCoreConditionsItemsRequirementsItemsComparatorEq string = "eq"

	// V1PatchCoreConditionsItemsRequirementsItemsComparatorGt captures enum value "gt"
	V1PatchCoreConditionsItemsRequirementsItemsComparatorGt string = "gt"

	// V1PatchCoreConditionsItemsRequirementsItemsComparatorGte captures enum value "gte"
	V1PatchCoreConditionsItemsRequirementsItemsComparatorGte string = "gte"

	// V1PatchCoreConditionsItemsRequirementsItemsComparatorLt captures enum value "lt"
	V1PatchCoreConditionsItemsRequirementsItemsComparatorLt string = "lt"

	// V1PatchCoreConditionsItemsRequirementsItemsComparatorLte captures enum value "lte"
	V1PatchCoreConditionsItemsRequirementsItemsComparatorLte string = "lte"

	// V1PatchCoreConditionsItemsRequirementsItemsComparatorNe captures enum value "ne"
	V1PatchCoreConditionsItemsRequirementsItemsComparatorNe string = "ne"
)

// prop value enum
func (m *V1PatchCoreConditionsItemsRequirementsItems) validateComparatorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1PatchCoreConditionsItemsRequirementsItemsTypeComparatorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1PatchCoreConditionsItemsRequirementsItems) validateComparator(formats strfmt.Registry) error {

	if err := validate.Required("comparator", "body", m.Comparator); err != nil {
		return err
	}

	// value enum
	if err := m.validateComparatorEnum("comparator", "body", *m.Comparator); err != nil {
		return err
	}

	return nil
}

func (m *V1PatchCoreConditionsItemsRequirementsItems) validateSortableVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.SortableVersion) { // not required
		return nil
	}

	for i := 0; i < len(m.SortableVersion); i++ {

		if err := validate.MinLength("sortable_version"+"."+strconv.Itoa(i), "body", string(m.SortableVersion[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *V1PatchCoreConditionsItemsRequirementsItems) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinLength("version", "body", string(*m.Version), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PatchCoreConditionsItemsRequirementsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PatchCoreConditionsItemsRequirementsItems) UnmarshalBinary(b []byte) error {
	var res V1PatchCoreConditionsItemsRequirementsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
