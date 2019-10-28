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

// V1OperatingSystemVersionPagedList Operating System Version Paged List
//
// A paginated list of operating system versions
// swagger:model v1OperatingSystemVersionPagedList
type V1OperatingSystemVersionPagedList struct {

	// links
	// Required: true
	Links *V1OperatingSystemVersionPagedListLinks `json:"links"`

	// A page of operating system versions
	// Required: true
	OperatingSystemVersions []*V1OperatingSystemVersionPagedListOperatingSystemVersionsItems `json:"operating_system_versions"`

	// paging
	// Required: true
	Paging *V1OperatingSystemVersionPagedListPaging `json:"paging"`
}

// Validate validates this v1 operating system version paged list
func (m *V1OperatingSystemVersionPagedList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatingSystemVersions(formats); err != nil {
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

func (m *V1OperatingSystemVersionPagedList) validateLinks(formats strfmt.Registry) error {

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

func (m *V1OperatingSystemVersionPagedList) validateOperatingSystemVersions(formats strfmt.Registry) error {

	if err := validate.Required("operating_system_versions", "body", m.OperatingSystemVersions); err != nil {
		return err
	}

	for i := 0; i < len(m.OperatingSystemVersions); i++ {
		if swag.IsZero(m.OperatingSystemVersions[i]) { // not required
			continue
		}

		if m.OperatingSystemVersions[i] != nil {
			if err := m.OperatingSystemVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("operating_system_versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1OperatingSystemVersionPagedList) validatePaging(formats strfmt.Registry) error {

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
func (m *V1OperatingSystemVersionPagedList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OperatingSystemVersionPagedList) UnmarshalBinary(b []byte) error {
	var res V1OperatingSystemVersionPagedList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
