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

// V1IngredientVersionPagedListIngredientVersionsItemsAllOf0LinksAllOf1 v1 ingredient version paged list ingredient versions items all of0 links all of1
// swagger:model v1IngredientVersionPagedListIngredientVersionsItemsAllOf0LinksAllOf1
type V1IngredientVersionPagedListIngredientVersionsItemsAllOf0LinksAllOf1 struct {

	// The URI of the ingredient this is a version of
	// Required: true
	// Format: uri
	Ingredient *strfmt.URI `json:"ingredient"`
}

// Validate validates this v1 ingredient version paged list ingredient versions items all of0 links all of1
func (m *V1IngredientVersionPagedListIngredientVersionsItemsAllOf0LinksAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIngredient(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IngredientVersionPagedListIngredientVersionsItemsAllOf0LinksAllOf1) validateIngredient(formats strfmt.Registry) error {

	if err := validate.Required("ingredient", "body", m.Ingredient); err != nil {
		return err
	}

	if err := validate.FormatOf("ingredient", "body", "uri", m.Ingredient.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IngredientVersionPagedListIngredientVersionsItemsAllOf0LinksAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IngredientVersionPagedListIngredientVersionsItemsAllOf0LinksAllOf1) UnmarshalBinary(b []byte) error {
	var res V1IngredientVersionPagedListIngredientVersionsItemsAllOf0LinksAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
