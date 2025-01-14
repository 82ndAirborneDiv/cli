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

// AuthorPagedList Author Paged List
//
// A page of authors from a paged author list
//
// swagger:model authorPagedList
type AuthorPagedList struct {

	// The authors
	// Required: true
	Authors []*Author `json:"authors"`

	// links
	// Required: true
	Links *PagingLinks `json:"links"`

	// paging
	// Required: true
	Paging *Paging `json:"paging"`
}

// Validate validates this author paged list
func (m *AuthorPagedList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthors(formats); err != nil {
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

func (m *AuthorPagedList) validateAuthors(formats strfmt.Registry) error {

	if err := validate.Required("authors", "body", m.Authors); err != nil {
		return err
	}

	for i := 0; i < len(m.Authors); i++ {
		if swag.IsZero(m.Authors[i]) { // not required
			continue
		}

		if m.Authors[i] != nil {
			if err := m.Authors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("authors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuthorPagedList) validateLinks(formats strfmt.Registry) error {

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

func (m *AuthorPagedList) validatePaging(formats strfmt.Registry) error {

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
func (m *AuthorPagedList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthorPagedList) UnmarshalBinary(b []byte) error {
	var res AuthorPagedList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
