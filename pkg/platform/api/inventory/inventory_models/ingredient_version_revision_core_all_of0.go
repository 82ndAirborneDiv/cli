// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IngredientVersionRevisionCoreAllOf0 ingredient version revision core all of0
//
// swagger:model ingredientVersionRevisionCoreAllOf0
type IngredientVersionRevisionCoreAllOf0 struct {

	// The SPDX license expression based on ActiveState's analysis of the package's licensing
	ActivestateLicenseExpression string `json:"activestate_license_expression,omitempty"`

	// Camel-specific metadata needed to build this ingredient version revision in camel, if there is any.
	CamelExtras interface{} `json:"camel_extras,omitempty"`

	// dependency sets
	DependencySets []*DependencySet `json:"dependency_sets"`

	// ingredient options
	IngredientOptions []*IngredientOption `json:"ingredient_options"`

	// Whether or not this revision is indemnified for customers paying for indemnification. If set to null, then this will use the is_indemnified value of the previous revision or false if this is the first revision.
	IsIndemnified *bool `json:"is_indemnified,omitempty"`

	// Whether or not this is a stable release of the package
	IsStableRelease *bool `json:"is_stable_release,omitempty"`

	// An S3 URI to a JSON manifest mapping files in the package to licenses for that file
	// Format: uri
	LicenseManifestURI *strfmt.URI `json:"license_manifest_uri,omitempty"`

	// S3 URL where the source distribution is stored for our platform
	// Format: uri
	PlatformSourceURI *strfmt.URI `json:"platform_source_uri,omitempty"`

	// The SPDX license expression based on running an automated scanner to determine the package's licensing
	ScannerLicenseExpression string `json:"scanner_license_expression,omitempty"`

	// A checksum of the source distribution. The actual type of the checksum (MD5, S3 Etag, etc.) is not specified. It's assumed that the system that populates and uses this data will know how to work with these checksums.
	SourceChecksum *string `json:"source_checksum,omitempty"`

	// The status of the revision. This can be one of stable, unstable, deleted, or deprecated.
	// Enum: [deleted deprecated stable unstable]
	Status string `json:"status,omitempty"`
}

// Validate validates this ingredient version revision core all of0
func (m *IngredientVersionRevisionCoreAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDependencySets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredientOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicenseManifestURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformSourceURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IngredientVersionRevisionCoreAllOf0) validateDependencySets(formats strfmt.Registry) error {

	if swag.IsZero(m.DependencySets) { // not required
		return nil
	}

	for i := 0; i < len(m.DependencySets); i++ {
		if swag.IsZero(m.DependencySets[i]) { // not required
			continue
		}

		if m.DependencySets[i] != nil {
			if err := m.DependencySets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependency_sets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IngredientVersionRevisionCoreAllOf0) validateIngredientOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.IngredientOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.IngredientOptions); i++ {
		if swag.IsZero(m.IngredientOptions[i]) { // not required
			continue
		}

		if m.IngredientOptions[i] != nil {
			if err := m.IngredientOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ingredient_options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IngredientVersionRevisionCoreAllOf0) validateLicenseManifestURI(formats strfmt.Registry) error {

	if swag.IsZero(m.LicenseManifestURI) { // not required
		return nil
	}

	if err := validate.FormatOf("license_manifest_uri", "body", "uri", m.LicenseManifestURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IngredientVersionRevisionCoreAllOf0) validatePlatformSourceURI(formats strfmt.Registry) error {

	if swag.IsZero(m.PlatformSourceURI) { // not required
		return nil
	}

	if err := validate.FormatOf("platform_source_uri", "body", "uri", m.PlatformSourceURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var ingredientVersionRevisionCoreAllOf0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["deleted","deprecated","stable","unstable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ingredientVersionRevisionCoreAllOf0TypeStatusPropEnum = append(ingredientVersionRevisionCoreAllOf0TypeStatusPropEnum, v)
	}
}

const (

	// IngredientVersionRevisionCoreAllOf0StatusDeleted captures enum value "deleted"
	IngredientVersionRevisionCoreAllOf0StatusDeleted string = "deleted"

	// IngredientVersionRevisionCoreAllOf0StatusDeprecated captures enum value "deprecated"
	IngredientVersionRevisionCoreAllOf0StatusDeprecated string = "deprecated"

	// IngredientVersionRevisionCoreAllOf0StatusStable captures enum value "stable"
	IngredientVersionRevisionCoreAllOf0StatusStable string = "stable"

	// IngredientVersionRevisionCoreAllOf0StatusUnstable captures enum value "unstable"
	IngredientVersionRevisionCoreAllOf0StatusUnstable string = "unstable"
)

// prop value enum
func (m *IngredientVersionRevisionCoreAllOf0) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ingredientVersionRevisionCoreAllOf0TypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IngredientVersionRevisionCoreAllOf0) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IngredientVersionRevisionCoreAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IngredientVersionRevisionCoreAllOf0) UnmarshalBinary(b []byte) error {
	var res IngredientVersionRevisionCoreAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
