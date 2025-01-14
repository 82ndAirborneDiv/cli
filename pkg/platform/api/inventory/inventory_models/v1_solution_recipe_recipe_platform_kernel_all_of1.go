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

// V1SolutionRecipeRecipePlatformKernelAllOf1 Kernel Core
//
// The properties of a kernel needed to create a new one
// swagger:model v1SolutionRecipeRecipePlatformKernelAllOf1
type V1SolutionRecipeRecipePlatformKernelAllOf1 struct {

	// The name of the kernel (excluding any version information)
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this v1 solution recipe recipe platform kernel all of1
func (m *V1SolutionRecipeRecipePlatformKernelAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SolutionRecipeRecipePlatformKernelAllOf1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SolutionRecipeRecipePlatformKernelAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SolutionRecipeRecipePlatformKernelAllOf1) UnmarshalBinary(b []byte) error {
	var res V1SolutionRecipeRecipePlatformKernelAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
