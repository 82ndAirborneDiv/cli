// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1SubSchemaRevisionedResource Revisioned Resource
//
// Properties shared by all revisioned resources
//
// swagger:model v1SubSchemaRevisionedResource
type V1SubSchemaRevisionedResource struct {

	// The revision number of this revision of the resource. This number increases monotonically with each new revision.
	// Required: true
	// Minimum: 1
	Revision *int64 `json:"revision"`

	// The date and time at which this revision of the resource was created
	// Required: true
	// Format: date-time
	RevisionTimestamp *strfmt.DateTime `json:"revision_timestamp"`
}

// Validate validates this v1 sub schema revisioned resource
func (m *V1SubSchemaRevisionedResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevisionTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SubSchemaRevisionedResource) validateRevision(formats strfmt.Registry) error {

	if err := validate.Required("revision", "body", m.Revision); err != nil {
		return err
	}

	if err := validate.MinimumInt("revision", "body", int64(*m.Revision), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *V1SubSchemaRevisionedResource) validateRevisionTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("revision_timestamp", "body", m.RevisionTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("revision_timestamp", "body", "date-time", m.RevisionTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SubSchemaRevisionedResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SubSchemaRevisionedResource) UnmarshalBinary(b []byte) error {
	var res V1SubSchemaRevisionedResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
