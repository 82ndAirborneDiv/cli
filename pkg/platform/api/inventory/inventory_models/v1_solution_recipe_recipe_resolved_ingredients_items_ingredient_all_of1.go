// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientAllOf1 Ingredient Core
//
// A unique ingredient that can be used in a recipe. These properties are shared by all ingredient models.
// swagger:model v1SolutionRecipeRecipeResolvedIngredientsItemsIngredientAllOf1
type V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientAllOf1 struct {

	// A concise summary of what this ingredient can be used for
	// Required: true
	Description *string `json:"description"`

	// The name of the ingredient (excluding any version information)
	// Required: true
	Name *string `json:"name"`

	// The UUID of the organization the ingredient belongs to, if it is private to a particular organization
	// Format: uuid
	OrganizationID strfmt.UUID `json:"organization_id,omitempty"`

	// The primary namespace to which this ingredient belongs
	// Required: true
	PrimaryNamespace *string `json:"primary_namespace"`

	// URL of the website about this ingredient (if any)
	// Format: uri
	Website strfmt.URI `json:"website,omitempty"`
}

// Validate validates this v1 solution recipe recipe resolved ingredients items ingredient all of1
func (m *V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebsite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientAllOf1) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientAllOf1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientAllOf1) validateOrganizationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganizationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organization_id", "body", "uuid", m.OrganizationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientAllOf1) validatePrimaryNamespace(formats strfmt.Registry) error {

	if err := validate.Required("primary_namespace", "body", m.PrimaryNamespace); err != nil {
		return err
	}

	return nil
}

func (m *V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientAllOf1) validateWebsite(formats strfmt.Registry) error {

	if swag.IsZero(m.Website) { // not required
		return nil
	}

	if err := validate.FormatOf("website", "body", "uri", m.Website.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientAllOf1) UnmarshalBinary(b []byte) error {
	var res V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
