// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IngredientVersionRevisionPatchAllOf1 ingredient version revision patch all of1
//
// swagger:model ingredientVersionRevisionPatchAllOf1
type IngredientVersionRevisionPatchAllOf1 struct {

	// sequence number
	// Required: true
	// Minimum: 1
	SequenceNumber *int64 `json:"sequence_number"`
}

// Validate validates this ingredient version revision patch all of1
func (m *IngredientVersionRevisionPatchAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSequenceNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IngredientVersionRevisionPatchAllOf1) validateSequenceNumber(formats strfmt.Registry) error {

	if err := validate.Required("sequence_number", "body", m.SequenceNumber); err != nil {
		return err
	}

	if err := validate.MinimumInt("sequence_number", "body", int64(*m.SequenceNumber), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IngredientVersionRevisionPatchAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IngredientVersionRevisionPatchAllOf1) UnmarshalBinary(b []byte) error {
	var res IngredientVersionRevisionPatchAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
