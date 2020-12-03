// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Image Image
//
// The full image data model
//
// swagger:model image
type Image struct {
	ImageAllOf0

	ImageCore

	RevisionedResource
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Image) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ImageAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ImageAllOf0 = aO0

	// AO1
	var aO1 ImageCore
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ImageCore = aO1

	// AO2
	var aO2 RevisionedResource
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.RevisionedResource = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Image) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.ImageAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ImageCore)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.RevisionedResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this image
func (m *Image) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ImageAllOf0
	if err := m.ImageAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ImageCore
	if err := m.ImageCore.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with RevisionedResource
	if err := m.RevisionedResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Image) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Image) UnmarshalBinary(b []byte) error {
	var res Image
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}