// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1BuildScript Build Script
//
// A short piece of scripting code that can be used to build an ingredient. This model contains all build script properties and is returned from read requests
//
// swagger:model v1BuildScript
type V1BuildScript struct {
	V1BuildScriptAllOf0

	V1BuildScriptCore
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1BuildScript) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1BuildScriptAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1BuildScriptAllOf0 = aO0

	// AO1
	var aO1 V1BuildScriptCore
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1BuildScriptCore = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1BuildScript) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.V1BuildScriptAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1BuildScriptCore)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 build script
func (m *V1BuildScript) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1BuildScriptAllOf0
	if err := m.V1BuildScriptAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1BuildScriptCore
	if err := m.V1BuildScriptCore.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildScript) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildScript) UnmarshalBinary(b []byte) error {
	var res V1BuildScript
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
