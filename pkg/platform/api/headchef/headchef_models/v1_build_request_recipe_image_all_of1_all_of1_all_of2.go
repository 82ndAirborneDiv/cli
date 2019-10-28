// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1BuildRequestRecipeImageAllOf1AllOf1AllOf2 v1 build request recipe image all of1 all of1 all of2
// swagger:model v1BuildRequestRecipeImageAllOf1AllOf1AllOf2
type V1BuildRequestRecipeImageAllOf1AllOf1AllOf2 struct {

	// conditions
	Conditions []*V1BuildRequestRecipeImageAllOf1AllOf1AllOf2ConditionsItems `json:"conditions"`
}

// Validate validates this v1 build request recipe image all of1 all of1 all of2
func (m *V1BuildRequestRecipeImageAllOf1AllOf1AllOf2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1BuildRequestRecipeImageAllOf1AllOf1AllOf2) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildRequestRecipeImageAllOf1AllOf1AllOf2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildRequestRecipeImageAllOf1AllOf1AllOf2) UnmarshalBinary(b []byte) error {
	var res V1BuildRequestRecipeImageAllOf1AllOf1AllOf2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
