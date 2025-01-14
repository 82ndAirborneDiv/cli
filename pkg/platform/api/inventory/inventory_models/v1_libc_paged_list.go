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

// V1LibcPagedList Libc Paged List
//
// A paginated list of libcs
// swagger:model v1LibcPagedList
type V1LibcPagedList struct {

	// A page of libcs
	// Required: true
	Libcs []*V1LibcPagedListLibcsItems `json:"libcs"`

	// links
	// Required: true
	Links *V1LibcPagedListLinks `json:"links"`

	// paging
	// Required: true
	Paging *V1LibcPagedListPaging `json:"paging"`
}

// Validate validates this v1 libc paged list
func (m *V1LibcPagedList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLibcs(formats); err != nil {
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

func (m *V1LibcPagedList) validateLibcs(formats strfmt.Registry) error {

	if err := validate.Required("libcs", "body", m.Libcs); err != nil {
		return err
	}

	for i := 0; i < len(m.Libcs); i++ {
		if swag.IsZero(m.Libcs[i]) { // not required
			continue
		}

		if m.Libcs[i] != nil {
			if err := m.Libcs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("libcs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1LibcPagedList) validateLinks(formats strfmt.Registry) error {

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

func (m *V1LibcPagedList) validatePaging(formats strfmt.Registry) error {

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
func (m *V1LibcPagedList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1LibcPagedList) UnmarshalBinary(b []byte) error {
	var res V1LibcPagedList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
