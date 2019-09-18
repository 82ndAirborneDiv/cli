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

// V1OperatingSystemVersionPagedListOperatingSystemVersionsItemsAllOf0 v1 operating system version paged list operating system versions items all of0
// swagger:model v1OperatingSystemVersionPagedListOperatingSystemVersionsItemsAllOf0
type V1OperatingSystemVersionPagedListOperatingSystemVersionsItemsAllOf0 struct {

	// links
	// Required: true
	Links *V1OperatingSystemVersionPagedListOperatingSystemVersionsItemsAllOf0Links `json:"links"`

	// operating system version id
	// Required: true
	// Format: uuid
	OperatingSystemVersionID *strfmt.UUID `json:"operating_system_version_id"`
}

// Validate validates this v1 operating system version paged list operating system versions items all of0
func (m *V1OperatingSystemVersionPagedListOperatingSystemVersionsItemsAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatingSystemVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OperatingSystemVersionPagedListOperatingSystemVersionsItemsAllOf0) validateLinks(formats strfmt.Registry) error {

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

func (m *V1OperatingSystemVersionPagedListOperatingSystemVersionsItemsAllOf0) validateOperatingSystemVersionID(formats strfmt.Registry) error {

	if err := validate.Required("operating_system_version_id", "body", m.OperatingSystemVersionID); err != nil {
		return err
	}

	if err := validate.FormatOf("operating_system_version_id", "body", "uuid", m.OperatingSystemVersionID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OperatingSystemVersionPagedListOperatingSystemVersionsItemsAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OperatingSystemVersionPagedListOperatingSystemVersionsItemsAllOf0) UnmarshalBinary(b []byte) error {
	var res V1OperatingSystemVersionPagedListOperatingSystemVersionsItemsAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
