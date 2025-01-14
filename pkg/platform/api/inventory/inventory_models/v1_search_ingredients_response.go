// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1SearchIngredientsResponse Search ingredients response
// swagger:model v1SearchIngredientsResponse
type V1SearchIngredientsResponse struct {

	// A page of ingredients and version information
	Ingredients []*V1SearchIngredientsResponseIngredientsItems `json:"ingredients"`

	// links
	Links *V1SearchIngredientsResponseLinks `json:"links,omitempty"`

	// paging
	Paging *V1SearchIngredientsResponsePaging `json:"paging,omitempty"`
}

// Validate validates this v1 search ingredients response
func (m *V1SearchIngredientsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIngredients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaging(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SearchIngredientsResponse) validateIngredients(formats strfmt.Registry) error {

	if swag.IsZero(m.Ingredients) { // not required
		return nil
	}

	for i := 0; i < len(m.Ingredients); i++ {
		if swag.IsZero(m.Ingredients[i]) { // not required
			continue
		}

		if m.Ingredients[i] != nil {
			if err := m.Ingredients[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ingredients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SearchIngredientsResponse) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

func (m *V1SearchIngredientsResponse) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(m.Paging) { // not required
		return nil
	}

	if m.Paging != nil {
		if err := m.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("paging")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SearchIngredientsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SearchIngredientsResponse) UnmarshalBinary(b []byte) error {
	var res V1SearchIngredientsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
