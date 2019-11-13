// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StatusMessageEnvelopeBody9 Package Build Failed
//
// A message indicating that a requested package build failed.
// swagger:model statusMessageEnvelopeBody9
type StatusMessageEnvelopeBody9 struct {

	// All of the errors from the failed build.
	// Required: true
	Errors []string `json:"errors"`

	// ingredient id
	// Required: true
	// Format: uuid
	IngredientID *strfmt.UUID `json:"ingredient_id"`

	// ingredient version
	// Required: true
	IngredientVersion *string `json:"ingredient_version"`

	// The percentage through the total number of packages to be built in the current tast
	// Maximum: 100
	// Minimum: 0
	Percent *int64 `json:"percent,omitempty"`

	// The current position in the list of packages to be built in the current task
	// Minimum: 0
	Sequence *int64 `json:"sequence,omitempty"`

	// The current total number of packages to be built in the current task
	// Minimum: 0
	Total *int64 `json:"total,omitempty"`
}

// Validate validates this status message envelope body9
func (m *StatusMessageEnvelopeBody9) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredientVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePercent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSequence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusMessageEnvelopeBody9) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("errors", "body", m.Errors); err != nil {
		return err
	}

	return nil
}

func (m *StatusMessageEnvelopeBody9) validateIngredientID(formats strfmt.Registry) error {

	if err := validate.Required("ingredient_id", "body", m.IngredientID); err != nil {
		return err
	}

	if err := validate.FormatOf("ingredient_id", "body", "uuid", m.IngredientID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StatusMessageEnvelopeBody9) validateIngredientVersion(formats strfmt.Registry) error {

	if err := validate.Required("ingredient_version", "body", m.IngredientVersion); err != nil {
		return err
	}

	return nil
}

func (m *StatusMessageEnvelopeBody9) validatePercent(formats strfmt.Registry) error {

	if swag.IsZero(m.Percent) { // not required
		return nil
	}

	if err := validate.MinimumInt("percent", "body", int64(*m.Percent), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("percent", "body", int64(*m.Percent), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *StatusMessageEnvelopeBody9) validateSequence(formats strfmt.Registry) error {

	if swag.IsZero(m.Sequence) { // not required
		return nil
	}

	if err := validate.MinimumInt("sequence", "body", int64(*m.Sequence), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *StatusMessageEnvelopeBody9) validateTotal(formats strfmt.Registry) error {

	if swag.IsZero(m.Total) { // not required
		return nil
	}

	if err := validate.MinimumInt("total", "body", int64(*m.Total), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatusMessageEnvelopeBody9) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusMessageEnvelopeBody9) UnmarshalBinary(b []byte) error {
	var res StatusMessageEnvelopeBody9
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}