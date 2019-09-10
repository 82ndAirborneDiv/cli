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

// V1IngredientPagedList Ingredient Paged List
//
// A paginated list of ingredients
// swagger:model v1IngredientPagedList
type V1IngredientPagedList struct {

	// A page of ingredients
	// Required: true
	Ingredients []*V1IngredientPagedListIngredientsItems0 `json:"ingredients"`

	// links
	// Required: true
	Links *V1IngredientPagedListLinks `json:"links"`

	// paging
	// Required: true
	Paging *V1IngredientPagedListPaging `json:"paging"`
}

// Validate validates this v1 ingredient paged list
func (m *V1IngredientPagedList) Validate(formats strfmt.Registry) error {
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

func (m *V1IngredientPagedList) validateIngredients(formats strfmt.Registry) error {

	if err := validate.Required("ingredients", "body", m.Ingredients); err != nil {
		return err
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

func (m *V1IngredientPagedList) validateLinks(formats strfmt.Registry) error {

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

func (m *V1IngredientPagedList) validatePaging(formats strfmt.Registry) error {

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
func (m *V1IngredientPagedList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IngredientPagedList) UnmarshalBinary(b []byte) error {
	var res V1IngredientPagedList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1IngredientPagedListIngredientsItems0 Ingredient
//
// A unique ingredient that can be used in a recipe. This model contains all ingredient properties and is returned from read requests.
// swagger:model V1IngredientPagedListIngredientsItems0
type V1IngredientPagedListIngredientsItems0 struct {

	// creation timestamp
	// Required: true
	// Format: date-time
	CreationTimestamp *strfmt.DateTime `json:"creation_timestamp"`

	// ingredient id
	// Required: true
	// Format: uuid
	IngredientID *strfmt.UUID `json:"ingredient_id"`

	// links
	// Required: true
	Links *V1IngredientPagedListIngredientsItems0AO0Links `json:"links"`

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

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1IngredientPagedListIngredientsItems0) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		CreationTimestamp *strfmt.DateTime `json:"creation_timestamp"`

		IngredientID *strfmt.UUID `json:"ingredient_id"`

		Links *V1IngredientPagedListIngredientsItems0AO0Links `json:"links"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.CreationTimestamp = dataAO0.CreationTimestamp

	m.IngredientID = dataAO0.IngredientID

	m.Links = dataAO0.Links

	// AO1
	var dataAO1 struct {
		Description *string `json:"description"`

		Name *string `json:"name"`

		OrganizationID strfmt.UUID `json:"organization_id,omitempty"`

		PrimaryNamespace *string `json:"primary_namespace"`

		Website strfmt.URI `json:"website,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Description = dataAO1.Description

	m.Name = dataAO1.Name

	m.OrganizationID = dataAO1.OrganizationID

	m.PrimaryNamespace = dataAO1.PrimaryNamespace

	m.Website = dataAO1.Website

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1IngredientPagedListIngredientsItems0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		CreationTimestamp *strfmt.DateTime `json:"creation_timestamp"`

		IngredientID *strfmt.UUID `json:"ingredient_id"`

		Links *V1IngredientPagedListIngredientsItems0AO0Links `json:"links"`
	}

	dataAO0.CreationTimestamp = m.CreationTimestamp

	dataAO0.IngredientID = m.IngredientID

	dataAO0.Links = m.Links

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	var dataAO1 struct {
		Description *string `json:"description"`

		Name *string `json:"name"`

		OrganizationID strfmt.UUID `json:"organization_id,omitempty"`

		PrimaryNamespace *string `json:"primary_namespace"`

		Website strfmt.URI `json:"website,omitempty"`
	}

	dataAO1.Description = m.Description

	dataAO1.Name = m.Name

	dataAO1.OrganizationID = m.OrganizationID

	dataAO1.PrimaryNamespace = m.PrimaryNamespace

	dataAO1.Website = m.Website

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 ingredient paged list ingredients items0
func (m *V1IngredientPagedListIngredientsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

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

func (m *V1IngredientPagedListIngredientsItems0) validateCreationTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("creation_timestamp", "body", m.CreationTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("creation_timestamp", "body", "date-time", m.CreationTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientPagedListIngredientsItems0) validateIngredientID(formats strfmt.Registry) error {

	if err := validate.Required("ingredient_id", "body", m.IngredientID); err != nil {
		return err
	}

	if err := validate.FormatOf("ingredient_id", "body", "uuid", m.IngredientID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientPagedListIngredientsItems0) validateLinks(formats strfmt.Registry) error {

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

func (m *V1IngredientPagedListIngredientsItems0) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientPagedListIngredientsItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientPagedListIngredientsItems0) validateOrganizationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganizationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organization_id", "body", "uuid", m.OrganizationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientPagedListIngredientsItems0) validatePrimaryNamespace(formats strfmt.Registry) error {

	if err := validate.Required("primary_namespace", "body", m.PrimaryNamespace); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientPagedListIngredientsItems0) validateWebsite(formats strfmt.Registry) error {

	if swag.IsZero(m.Website) { // not required
		return nil
	}

	if err := validate.FormatOf("website", "body", "uri", m.Website.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IngredientPagedListIngredientsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IngredientPagedListIngredientsItems0) UnmarshalBinary(b []byte) error {
	var res V1IngredientPagedListIngredientsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1IngredientPagedListIngredientsItems0AO0Links Self Link
//
// A self link
// swagger:model V1IngredientPagedListIngredientsItems0AO0Links
type V1IngredientPagedListIngredientsItems0AO0Links struct {

	// The URI of this resource
	// Required: true
	// Format: uri
	Self *strfmt.URI `json:"self"`
}

// Validate validates this v1 ingredient paged list ingredients items0 a o0 links
func (m *V1IngredientPagedListIngredientsItems0AO0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IngredientPagedListIngredientsItems0AO0Links) validateSelf(formats strfmt.Registry) error {

	if err := validate.Required("links"+"."+"self", "body", m.Self); err != nil {
		return err
	}

	if err := validate.FormatOf("links"+"."+"self", "body", "uri", m.Self.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IngredientPagedListIngredientsItems0AO0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IngredientPagedListIngredientsItems0AO0Links) UnmarshalBinary(b []byte) error {
	var res V1IngredientPagedListIngredientsItems0AO0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1IngredientPagedListLinks Paging Links
//
// Links for a model that include links for paging data
// swagger:model V1IngredientPagedListLinks
type V1IngredientPagedListLinks struct {

	// The URI of the first page
	// Required: true
	// Format: uri
	First *strfmt.URI `json:"first"`

	// The URI of last page
	// Required: true
	// Format: uri
	Last *strfmt.URI `json:"last"`

	// The URI of the next page
	// Format: uri
	Next strfmt.URI `json:"next,omitempty"`

	// The URI of the previous page
	// Format: uri
	Previous strfmt.URI `json:"previous,omitempty"`

	// The URI of this page
	// Required: true
	// Format: uri
	Self *strfmt.URI `json:"self"`
}

// Validate validates this v1 ingredient paged list links
func (m *V1IngredientPagedListLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLast(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IngredientPagedListLinks) validateFirst(formats strfmt.Registry) error {

	if err := validate.Required("links"+"."+"first", "body", m.First); err != nil {
		return err
	}

	if err := validate.FormatOf("links"+"."+"first", "body", "uri", m.First.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientPagedListLinks) validateLast(formats strfmt.Registry) error {

	if err := validate.Required("links"+"."+"last", "body", m.Last); err != nil {
		return err
	}

	if err := validate.FormatOf("links"+"."+"last", "body", "uri", m.Last.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientPagedListLinks) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("links"+"."+"next", "body", "uri", m.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientPagedListLinks) validatePrevious(formats strfmt.Registry) error {

	if swag.IsZero(m.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("links"+"."+"previous", "body", "uri", m.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientPagedListLinks) validateSelf(formats strfmt.Registry) error {

	if err := validate.Required("links"+"."+"self", "body", m.Self); err != nil {
		return err
	}

	if err := validate.FormatOf("links"+"."+"self", "body", "uri", m.Self.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IngredientPagedListLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IngredientPagedListLinks) UnmarshalBinary(b []byte) error {
	var res V1IngredientPagedListLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1IngredientPagedListPaging Paging
//
// Paging data
// swagger:model V1IngredientPagedListPaging
type V1IngredientPagedListPaging struct {

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

// Validate validates this v1 ingredient paged list paging
func (m *V1IngredientPagedListPaging) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *V1IngredientPagedListPaging) validateItemCount(formats strfmt.Registry) error {

	if err := validate.Required("paging"+"."+"item_count", "body", m.ItemCount); err != nil {
		return err
	}

	if err := validate.MinimumInt("paging"+"."+"item_count", "body", int64(*m.ItemCount), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientPagedListPaging) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("paging"+"."+"limit", "body", m.Limit); err != nil {
		return err
	}

	if err := validate.MinimumInt("paging"+"."+"limit", "body", int64(*m.Limit), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientPagedListPaging) validatePage(formats strfmt.Registry) error {

	if err := validate.Required("paging"+"."+"page", "body", m.Page); err != nil {
		return err
	}

	if err := validate.MinimumInt("paging"+"."+"page", "body", int64(*m.Page), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientPagedListPaging) validatePageCount(formats strfmt.Registry) error {

	if err := validate.Required("paging"+"."+"page_count", "body", m.PageCount); err != nil {
		return err
	}

	if err := validate.MinimumInt("paging"+"."+"page_count", "body", int64(*m.PageCount), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IngredientPagedListPaging) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IngredientPagedListPaging) UnmarshalBinary(b []byte) error {
	var res V1IngredientPagedListPaging
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
