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

// V1IngredientVersionPagedListPaging Paging
//
// Paging data
// swagger:model v1IngredientVersionPagedListPaging
type V1IngredientVersionPagedListPaging struct {

	// The total number of items available
	// Required: true
	// Minimum: 0
	AvailableCount *int64 `json:"available_count"`

	// The number of items on this page
	// Required: true
	// Minimum: 0
	ItemCount *int64 `json:"item_count"`

	// The maximum number of items that could be returned
	// Required: true
	// Minimum: 1
	Limit *int64 `json:"limit"`

	// The page number of this result set
	// Required: true
	// Minimum: 1
	Page *int64 `json:"page"`

	// The total number of pages
	// Required: true
	// Minimum: 1
	PageCount *int64 `json:"page_count"`
}

// Validate validates this v1 ingredient version paged list paging
func (m *V1IngredientVersionPagedListPaging) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IngredientVersionPagedListPaging) validateAvailableCount(formats strfmt.Registry) error {

	if err := validate.Required("available_count", "body", m.AvailableCount); err != nil {
		return err
	}

	if err := validate.MinimumInt("available_count", "body", int64(*m.AvailableCount), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientVersionPagedListPaging) validateItemCount(formats strfmt.Registry) error {

	if err := validate.Required("item_count", "body", m.ItemCount); err != nil {
		return err
	}

	if err := validate.MinimumInt("item_count", "body", int64(*m.ItemCount), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientVersionPagedListPaging) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	if err := validate.MinimumInt("limit", "body", int64(*m.Limit), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientVersionPagedListPaging) validatePage(formats strfmt.Registry) error {

	if err := validate.Required("page", "body", m.Page); err != nil {
		return err
	}

	if err := validate.MinimumInt("page", "body", int64(*m.Page), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientVersionPagedListPaging) validatePageCount(formats strfmt.Registry) error {

	if err := validate.Required("page_count", "body", m.PageCount); err != nil {
		return err
	}

	if err := validate.MinimumInt("page_count", "body", int64(*m.PageCount), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IngredientVersionPagedListPaging) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IngredientVersionPagedListPaging) UnmarshalBinary(b []byte) error {
	var res V1IngredientVersionPagedListPaging
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
