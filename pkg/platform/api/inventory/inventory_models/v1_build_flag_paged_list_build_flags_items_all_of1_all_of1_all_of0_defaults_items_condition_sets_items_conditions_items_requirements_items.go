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

// V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItems v1 build flag paged list build flags items all of1 all of1 all of0 defaults items condition sets items conditions items requirements items
// swagger:model v1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItems
type V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItems struct {

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

// Validate validates this v1 build flag paged list build flags items all of1 all of1 all of0 defaults items condition sets items conditions items requirements items
func (m *V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItems) Validate(formats strfmt.Registry) error {
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

var v1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItemsTypeComparatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["eq","gt","gte","lt","lte","ne"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItemsTypeComparatorPropEnum = append(v1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItemsTypeComparatorPropEnum, v)
	}
}

const (

	// V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItemsComparatorEq captures enum value "eq"
	V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItemsComparatorEq string = "eq"

	// V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItemsComparatorGt captures enum value "gt"
	V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItemsComparatorGt string = "gt"

	// V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItemsComparatorGte captures enum value "gte"
	V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItemsComparatorGte string = "gte"

	// V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItemsComparatorLt captures enum value "lt"
	V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItemsComparatorLt string = "lt"

	// V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItemsComparatorLte captures enum value "lte"
	V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItemsComparatorLte string = "lte"

	// V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItemsComparatorNe captures enum value "ne"
	V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItemsComparatorNe string = "ne"
)

// prop value enum
func (m *V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItems) validateComparatorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItemsTypeComparatorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItems) validateComparator(formats strfmt.Registry) error {

	if err := validate.Required("comparator", "body", m.Comparator); err != nil {
		return err
	}

	// value enum
	if err := m.validateComparatorEnum("comparator", "body", *m.Comparator); err != nil {
		return err
	}

	return nil
}

func (m *V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItems) validateSortableVersion(formats strfmt.Registry) error {

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

func (m *V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItems) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinLength("version", "body", string(*m.Version), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItems) UnmarshalBinary(b []byte) error {
	var res V1BuildFlagPagedListBuildFlagsItemsAllOf1AllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItemsRequirementsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
