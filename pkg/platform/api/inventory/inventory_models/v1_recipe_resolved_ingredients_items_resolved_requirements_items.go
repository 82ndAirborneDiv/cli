// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1RecipeResolvedIngredientsItemsResolvedRequirementsItems Requirement Sub Schema
//
// An order requirement is a single package name and version specifier requested in an order.
// swagger:model v1RecipeResolvedIngredientsItemsResolvedRequirementsItems
type V1RecipeResolvedIngredientsItemsResolvedRequirementsItems struct {

	// The name of the required feature, If no ingredient ID is specified, the default provider of this feature will be chosen.
	// Required: true
	Feature *string `json:"feature"`

	// The ID of the ingredient version that should be used to fulfill this requirement. Can be used to override the default choice of provider for the specified feature. Must be an ingredient version that actually provides the specified feature.
	// Format: uuid
	IngredientVersionID strfmt.UUID `json:"ingredient_version_id,omitempty"`

	// The namespace for the required feature. For now, this can be empty as it is only used to request pre-platform installer ingredients.
	// Required: true
	Namespace *string `json:"namespace"`

	// The revision number of the ingredient version that should be used to fulfill this requirement.
	// Minimum: 1
	Revision int64 `json:"revision,omitempty"`

	// The requirements for the acceptable versions of this feature. This can be omitted, in which case any version is acceptable.
	// Min Items: 1
	VersionRequirements []*V1RecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems `json:"version_requirements"`
}

// Validate validates this v1 recipe resolved ingredients items resolved requirements items
func (m *V1RecipeResolvedIngredientsItemsResolvedRequirementsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeature(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredientVersionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionRequirements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RecipeResolvedIngredientsItemsResolvedRequirementsItems) validateFeature(formats strfmt.Registry) error {

	if err := validate.Required("feature", "body", m.Feature); err != nil {
		return err
	}

	return nil
}

func (m *V1RecipeResolvedIngredientsItemsResolvedRequirementsItems) validateIngredientVersionID(formats strfmt.Registry) error {

	if swag.IsZero(m.IngredientVersionID) { // not required
		return nil
	}

	if err := validate.FormatOf("ingredient_version_id", "body", "uuid", m.IngredientVersionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1RecipeResolvedIngredientsItemsResolvedRequirementsItems) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *V1RecipeResolvedIngredientsItemsResolvedRequirementsItems) validateRevision(formats strfmt.Registry) error {

	if swag.IsZero(m.Revision) { // not required
		return nil
	}

	if err := validate.MinimumInt("revision", "body", int64(m.Revision), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *V1RecipeResolvedIngredientsItemsResolvedRequirementsItems) validateVersionRequirements(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionRequirements) { // not required
		return nil
	}

	iVersionRequirementsSize := int64(len(m.VersionRequirements))

	if err := validate.MinItems("version_requirements", "body", iVersionRequirementsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.VersionRequirements); i++ {
		if swag.IsZero(m.VersionRequirements[i]) { // not required
			continue
		}

		if m.VersionRequirements[i] != nil {
			if err := m.VersionRequirements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("version_requirements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RecipeResolvedIngredientsItemsResolvedRequirementsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RecipeResolvedIngredientsItemsResolvedRequirementsItems) UnmarshalBinary(b []byte) error {
	var res V1RecipeResolvedIngredientsItemsResolvedRequirementsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
