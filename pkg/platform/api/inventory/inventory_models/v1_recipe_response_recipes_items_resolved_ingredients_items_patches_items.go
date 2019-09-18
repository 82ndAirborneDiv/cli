// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItems Patch
//
// A diff of changes that can be applied to an ingredient's source code. This model contains all patch properties and is returned from read requests
// swagger:model v1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItems
type V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItems struct {
	V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItemsAllOf0

	V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItemsAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItems) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItemsAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItemsAllOf0 = aO0

	// AO1
	var aO1 V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItemsAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItemsAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItems) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItemsAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItemsAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 recipe response recipes items resolved ingredients items patches items
func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItems) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItemsAllOf0
	if err := m.V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItemsAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItemsAllOf1
	if err := m.V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItemsAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItems) UnmarshalBinary(b []byte) error {
	var res V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
