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

// V1SolutionRecipeRecipe Recipe
//
// A recipe contains all required information (ingredient versions, platform, build options, etc.) to build a project for a single platform.
// swagger:model v1SolutionRecipeRecipe
type V1SolutionRecipeRecipe struct {

	// Camel-specific flags for controlling the build.
	CamelFlags []string `json:"camel_flags"`

	// image
	// Required: true
	Image *V1SolutionRecipeRecipeImage `json:"image"`

	// If all of the resolved ingredients resolved to indemnified versions, then this will be true.
	IsIndemnified *bool `json:"is_indemnified,omitempty"`

	// platform
	// Required: true
	Platform *V1SolutionRecipeRecipePlatform `json:"platform"`

	// This recipe's ID. Identical recipes will have the same ID.
	// Required: true
	// Format: uuid
	RecipeID *strfmt.UUID `json:"recipe_id"`

	// The resolved ingredients that comprise this recipe.
	// Required: true
	// Min Items: 1
	// Unique: true
	ResolvedIngredients []*V1SolutionRecipeRecipeResolvedIngredientsItems `json:"resolved_ingredients"`
}

// Validate validates this v1 solution recipe recipe
func (m *V1SolutionRecipeRecipe) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolvedIngredients(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SolutionRecipeRecipe) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *V1SolutionRecipeRecipe) validatePlatform(formats strfmt.Registry) error {

	if err := validate.Required("platform", "body", m.Platform); err != nil {
		return err
	}

	if m.Platform != nil {
		if err := m.Platform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

func (m *V1SolutionRecipeRecipe) validateRecipeID(formats strfmt.Registry) error {

	if err := validate.Required("recipe_id", "body", m.RecipeID); err != nil {
		return err
	}

	if err := validate.FormatOf("recipe_id", "body", "uuid", m.RecipeID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1SolutionRecipeRecipe) validateResolvedIngredients(formats strfmt.Registry) error {

	if err := validate.Required("resolved_ingredients", "body", m.ResolvedIngredients); err != nil {
		return err
	}

	iResolvedIngredientsSize := int64(len(m.ResolvedIngredients))

	if err := validate.MinItems("resolved_ingredients", "body", iResolvedIngredientsSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("resolved_ingredients", "body", m.ResolvedIngredients); err != nil {
		return err
	}

	for i := 0; i < len(m.ResolvedIngredients); i++ {
		if swag.IsZero(m.ResolvedIngredients[i]) { // not required
			continue
		}

		if m.ResolvedIngredients[i] != nil {
			if err := m.ResolvedIngredients[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resolved_ingredients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SolutionRecipeRecipe) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SolutionRecipeRecipe) UnmarshalBinary(b []byte) error {
	var res V1SolutionRecipeRecipe
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
