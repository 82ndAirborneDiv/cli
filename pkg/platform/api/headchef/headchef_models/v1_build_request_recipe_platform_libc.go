// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1BuildRequestRecipePlatformLibc Libc
//
// The full libc data model
// swagger:model v1BuildRequestRecipePlatformLibc
type V1BuildRequestRecipePlatformLibc struct {
	V1BuildRequestRecipePlatformLibcAllOf0

	V1BuildRequestRecipePlatformLibcAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1BuildRequestRecipePlatformLibc) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1BuildRequestRecipePlatformLibcAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1BuildRequestRecipePlatformLibcAllOf0 = aO0

	// AO1
	var aO1 V1BuildRequestRecipePlatformLibcAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1BuildRequestRecipePlatformLibcAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1BuildRequestRecipePlatformLibc) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.V1BuildRequestRecipePlatformLibcAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1BuildRequestRecipePlatformLibcAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 build request recipe platform libc
func (m *V1BuildRequestRecipePlatformLibc) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1BuildRequestRecipePlatformLibcAllOf0
	if err := m.V1BuildRequestRecipePlatformLibcAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1BuildRequestRecipePlatformLibcAllOf1
	if err := m.V1BuildRequestRecipePlatformLibcAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildRequestRecipePlatformLibc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildRequestRecipePlatformLibc) UnmarshalBinary(b []byte) error {
	var res V1BuildRequestRecipePlatformLibc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}