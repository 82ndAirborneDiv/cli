// Code generated by go-swagger; DO NOT EDIT.

package mono_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Distro A fully resolved list of Source Packages and Languages that has been compiled for a specific Platform.
//
//
// swagger:model Distro
type Distro struct {

	// added
	// Format: date-time
	Added strfmt.DateTime `json:"added,omitempty"`

	// distro ID
	// Format: uuid
	DistroID strfmt.UUID `json:"distroID,omitempty"`

	// formats
	Formats []*Format `json:"formats"`

	// manifest
	Manifest []*IngredientVersion `json:"manifest"`

	// platform display name
	PlatformDisplayName *string `json:"platformDisplayName,omitempty"`

	// platform ID
	// Format: uuid
	PlatformID strfmt.UUID `json:"platformID,omitempty"`

	// release ID
	// Format: uuid
	ReleaseID strfmt.UUID `json:"releaseID,omitempty"`
}

// Validate validates this distro
func (m *Distro) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDistroID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManifest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Distro) validateAdded(formats strfmt.Registry) error {

	if swag.IsZero(m.Added) { // not required
		return nil
	}

	if err := validate.FormatOf("added", "body", "date-time", m.Added.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Distro) validateDistroID(formats strfmt.Registry) error {

	if swag.IsZero(m.DistroID) { // not required
		return nil
	}

	if err := validate.FormatOf("distroID", "body", "uuid", m.DistroID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Distro) validateFormats(formats strfmt.Registry) error {

	if swag.IsZero(m.Formats) { // not required
		return nil
	}

	for i := 0; i < len(m.Formats); i++ {
		if swag.IsZero(m.Formats[i]) { // not required
			continue
		}

		if m.Formats[i] != nil {
			if err := m.Formats[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("formats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Distro) validateManifest(formats strfmt.Registry) error {

	if swag.IsZero(m.Manifest) { // not required
		return nil
	}

	for i := 0; i < len(m.Manifest); i++ {
		if swag.IsZero(m.Manifest[i]) { // not required
			continue
		}

		if m.Manifest[i] != nil {
			if err := m.Manifest[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("manifest" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Distro) validatePlatformID(formats strfmt.Registry) error {

	if swag.IsZero(m.PlatformID) { // not required
		return nil
	}

	if err := validate.FormatOf("platformID", "body", "uuid", m.PlatformID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Distro) validateReleaseID(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseID) { // not required
		return nil
	}

	if err := validate.FormatOf("releaseID", "body", "uuid", m.ReleaseID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Distro) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Distro) UnmarshalBinary(b []byte) error {
	var res Distro
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
