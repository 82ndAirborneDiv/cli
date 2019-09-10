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

// V1BuildScriptCore Build Script Core
//
// A short piece of scripting code that can be used to build an ingredient. This model only contains mutable fields of a build script.
// swagger:model v1BuildScriptCore
type V1BuildScriptCore struct {

	// The features that must already be present in the recipe for this build script to be used. For example, can be used to create build scripts that only work on specific operating systems.
	Conditions []*V1BuildScriptCoreConditionsItems0 `json:"conditions"`

	// The scripting language that the build script is written in
	// Required: true
	// Enum: [bash perl python]
	Language *string `json:"language"`

	// The build script itself
	// Required: true
	Script *string `json:"script"`
}

// Validate validates this v1 build script core
func (m *V1BuildScriptCore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScript(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1BuildScriptCore) validateConditions(formats strfmt.Registry) error {

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

var v1BuildScriptCoreTypeLanguagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bash","perl","python"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1BuildScriptCoreTypeLanguagePropEnum = append(v1BuildScriptCoreTypeLanguagePropEnum, v)
	}
}

const (

	// V1BuildScriptCoreLanguageBash captures enum value "bash"
	V1BuildScriptCoreLanguageBash string = "bash"

	// V1BuildScriptCoreLanguagePerl captures enum value "perl"
	V1BuildScriptCoreLanguagePerl string = "perl"

	// V1BuildScriptCoreLanguagePython captures enum value "python"
	V1BuildScriptCoreLanguagePython string = "python"
)

// prop value enum
func (m *V1BuildScriptCore) validateLanguageEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1BuildScriptCoreTypeLanguagePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1BuildScriptCore) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("language", "body", m.Language); err != nil {
		return err
	}

	// value enum
	if err := m.validateLanguageEnum("language", "body", *m.Language); err != nil {
		return err
	}

	return nil
}

func (m *V1BuildScriptCore) validateScript(formats strfmt.Registry) error {

	if err := validate.Required("script", "body", m.Script); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildScriptCore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildScriptCore) UnmarshalBinary(b []byte) error {
	var res V1BuildScriptCore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1BuildScriptCoreConditionsItems0 Condition Sub Schema
//
// A feature that must be present in a recipe for the containing entity to apply. If nothing in the recipe matches this condition, the containing entity is disable/cannot be used.
// swagger:model V1BuildScriptCoreConditionsItems0
type V1BuildScriptCoreConditionsItems0 struct {

	// What feature must be present for the containing entity to apply
	// Required: true
	Feature *string `json:"feature"`

	// The namespace the conditional feature is contained in
	// Required: true
	Namespace *string `json:"namespace"`

	// Requirements Sub Schema
	//
	// The version constraints that an ingredient version's requirement or condition puts on a feature
	// Required: true
	// Min Length: 1
	Requirements []*V1BuildScriptCoreConditionsItems0RequirementsItems0 `json:"requirements"`
}

// Validate validates this v1 build script core conditions items0
func (m *V1BuildScriptCoreConditionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *V1BuildScriptCoreConditionsItems0) validateFeature(formats strfmt.Registry) error {

	if err := validate.Required("feature", "body", m.Feature); err != nil {
		return err
	}

	return nil
}

func (m *V1BuildScriptCoreConditionsItems0) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *V1BuildScriptCoreConditionsItems0) validateRequirements(formats strfmt.Registry) error {

	if err := validate.Required("requirements", "body", m.Requirements); err != nil {
		return err
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

// MarshalBinary interface implementation
func (m *V1BuildScriptCoreConditionsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildScriptCoreConditionsItems0) UnmarshalBinary(b []byte) error {
	var res V1BuildScriptCoreConditionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1BuildScriptCoreConditionsItems0RequirementsItems0 v1 build script core conditions items0 requirements items0
// swagger:model V1BuildScriptCoreConditionsItems0RequirementsItems0
type V1BuildScriptCoreConditionsItems0RequirementsItems0 struct {

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

// Validate validates this v1 build script core conditions items0 requirements items0
func (m *V1BuildScriptCoreConditionsItems0RequirementsItems0) Validate(formats strfmt.Registry) error {
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

var v1BuildScriptCoreConditionsItems0RequirementsItems0TypeComparatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["eq","gt","gte","lt","lte","ne"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1BuildScriptCoreConditionsItems0RequirementsItems0TypeComparatorPropEnum = append(v1BuildScriptCoreConditionsItems0RequirementsItems0TypeComparatorPropEnum, v)
	}
}

const (

	// V1BuildScriptCoreConditionsItems0RequirementsItems0ComparatorEq captures enum value "eq"
	V1BuildScriptCoreConditionsItems0RequirementsItems0ComparatorEq string = "eq"

	// V1BuildScriptCoreConditionsItems0RequirementsItems0ComparatorGt captures enum value "gt"
	V1BuildScriptCoreConditionsItems0RequirementsItems0ComparatorGt string = "gt"

	// V1BuildScriptCoreConditionsItems0RequirementsItems0ComparatorGte captures enum value "gte"
	V1BuildScriptCoreConditionsItems0RequirementsItems0ComparatorGte string = "gte"

	// V1BuildScriptCoreConditionsItems0RequirementsItems0ComparatorLt captures enum value "lt"
	V1BuildScriptCoreConditionsItems0RequirementsItems0ComparatorLt string = "lt"

	// V1BuildScriptCoreConditionsItems0RequirementsItems0ComparatorLte captures enum value "lte"
	V1BuildScriptCoreConditionsItems0RequirementsItems0ComparatorLte string = "lte"

	// V1BuildScriptCoreConditionsItems0RequirementsItems0ComparatorNe captures enum value "ne"
	V1BuildScriptCoreConditionsItems0RequirementsItems0ComparatorNe string = "ne"
)

// prop value enum
func (m *V1BuildScriptCoreConditionsItems0RequirementsItems0) validateComparatorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1BuildScriptCoreConditionsItems0RequirementsItems0TypeComparatorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1BuildScriptCoreConditionsItems0RequirementsItems0) validateComparator(formats strfmt.Registry) error {

	if err := validate.Required("comparator", "body", m.Comparator); err != nil {
		return err
	}

	// value enum
	if err := m.validateComparatorEnum("comparator", "body", *m.Comparator); err != nil {
		return err
	}

	return nil
}

func (m *V1BuildScriptCoreConditionsItems0RequirementsItems0) validateSortableVersion(formats strfmt.Registry) error {

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

func (m *V1BuildScriptCoreConditionsItems0RequirementsItems0) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinLength("version", "body", string(*m.Version), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildScriptCoreConditionsItems0RequirementsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildScriptCoreConditionsItems0RequirementsItems0) UnmarshalBinary(b []byte) error {
	var res V1BuildScriptCoreConditionsItems0RequirementsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
