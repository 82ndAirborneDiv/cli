// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IngredientVersionRevisionPagedList Ingredient Version Revision Paged List
//
// A paginated list of ingredient version revisions
//
// swagger:model ingredientVersionRevisionPagedList
type IngredientVersionRevisionPagedList struct {

	// A page of ingredient version revisions
	// Required: true
	IngredientVersionRevisions []*IngredientVersionRevision `json:"ingredient_version_revisions"`

	// links
	// Required: true
	Links *PagingLinks `json:"links"`

	// paging
	// Required: true
	Paging *Paging `json:"paging"`
}

// Validate validates this ingredient version revision paged list
func (m *IngredientVersionRevisionPagedList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIngredientVersionRevisions(formats); err != nil {
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

func (m *IngredientVersionRevisionPagedList) validateIngredientVersionRevisions(formats strfmt.Registry) error {

	if err := validate.Required("ingredient_version_revisions", "body", m.IngredientVersionRevisions); err != nil {
		return err
	}

	for i := 0; i < len(m.IngredientVersionRevisions); i++ {
		if swag.IsZero(m.IngredientVersionRevisions[i]) { // not required
			continue
		}

		if m.IngredientVersionRevisions[i] != nil {
			if err := m.IngredientVersionRevisions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ingredient_version_revisions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IngredientVersionRevisionPagedList) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("links", "body", m.Links); err != nil {
		return err
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

func (m *IngredientVersionRevisionPagedList) validatePaging(formats strfmt.Registry) error {

	if err := validate.Required("paging", "body", m.Paging); err != nil {
		return err
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
func (m *IngredientVersionRevisionPagedList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IngredientVersionRevisionPagedList) UnmarshalBinary(b []byte) error {
	var res IngredientVersionRevisionPagedList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
