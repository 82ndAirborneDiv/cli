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

// Project project
//
// swagger:model Project
type Project struct {

	// added
	// Format: date-time
	Added strfmt.DateTime `json:"added,omitempty"`

	// branches
	Branches Branches `json:"branches,omitempty"`

	// created by
	CreatedBy *string `json:"createdBy,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// forked from
	ForkedFrom *ProjectForkedFrom `json:"forkedFrom,omitempty"`

	// languages
	Languages []*ProjectLanguagesItems0 `json:"languages"`

	// last edited
	// Format: date-time
	LastEdited strfmt.DateTime `json:"lastEdited,omitempty"`

	// managed
	Managed bool `json:"managed,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization ID
	// Format: uuid
	OrganizationID strfmt.UUID `json:"organizationID,omitempty"`

	// platforms
	Platforms []*ProjectPlatformsItems0 `json:"platforms"`

	// private
	Private bool `json:"private,omitempty"`

	// project ID
	// Format: uuid
	ProjectID strfmt.UUID `json:"projectID,omitempty"`

	// repo Url
	RepoURL *string `json:"repoUrl,omitempty"`
}

// Validate validates this project
func (m *Project) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBranches(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForkedFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastEdited(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatforms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Project) validateAdded(formats strfmt.Registry) error {

	if swag.IsZero(m.Added) { // not required
		return nil
	}

	if err := validate.FormatOf("added", "body", "date-time", m.Added.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateBranches(formats strfmt.Registry) error {

	if swag.IsZero(m.Branches) { // not required
		return nil
	}

	if err := m.Branches.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("branches")
		}
		return err
	}

	return nil
}

func (m *Project) validateForkedFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.ForkedFrom) { // not required
		return nil
	}

	if m.ForkedFrom != nil {
		if err := m.ForkedFrom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forkedFrom")
			}
			return err
		}
	}

	return nil
}

func (m *Project) validateLanguages(formats strfmt.Registry) error {

	if swag.IsZero(m.Languages) { // not required
		return nil
	}

	for i := 0; i < len(m.Languages); i++ {
		if swag.IsZero(m.Languages[i]) { // not required
			continue
		}

		if m.Languages[i] != nil {
			if err := m.Languages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("languages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Project) validateLastEdited(formats strfmt.Registry) error {

	if swag.IsZero(m.LastEdited) { // not required
		return nil
	}

	if err := validate.FormatOf("lastEdited", "body", "date-time", m.LastEdited.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateOrganizationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganizationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organizationID", "body", "uuid", m.OrganizationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Project) validatePlatforms(formats strfmt.Registry) error {

	if swag.IsZero(m.Platforms) { // not required
		return nil
	}

	for i := 0; i < len(m.Platforms); i++ {
		if swag.IsZero(m.Platforms[i]) { // not required
			continue
		}

		if m.Platforms[i] != nil {
			if err := m.Platforms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("platforms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Project) validateProjectID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("projectID", "body", "uuid", m.ProjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Project) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Project) UnmarshalBinary(b []byte) error {
	var res Project
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ProjectForkedFrom project forked from
//
// swagger:model ProjectForkedFrom
type ProjectForkedFrom struct {

	// organization
	Organization string `json:"organization,omitempty"`

	// project
	Project string `json:"project,omitempty"`
}

// Validate validates this project forked from
func (m *ProjectForkedFrom) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectForkedFrom) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectForkedFrom) UnmarshalBinary(b []byte) error {
	var res ProjectForkedFrom
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ProjectLanguagesItems0 project languages items0
//
// swagger:model ProjectLanguagesItems0
type ProjectLanguagesItems0 struct {

	// ingredient ID
	// Format: uuid
	IngredientID strfmt.UUID `json:"ingredientID,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this project languages items0
func (m *ProjectLanguagesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIngredientID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectLanguagesItems0) validateIngredientID(formats strfmt.Registry) error {

	if swag.IsZero(m.IngredientID) { // not required
		return nil
	}

	if err := validate.FormatOf("ingredientID", "body", "uuid", m.IngredientID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectLanguagesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectLanguagesItems0) UnmarshalBinary(b []byte) error {
	var res ProjectLanguagesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ProjectPlatformsItems0 project platforms items0
//
// swagger:model ProjectPlatformsItems0
type ProjectPlatformsItems0 struct {

	// display name
	DisplayName *string `json:"displayName,omitempty"`

	// platform ID
	// Format: uuid
	PlatformID strfmt.UUID `json:"platformID,omitempty"`
}

// Validate validates this project platforms items0
func (m *ProjectPlatformsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlatformID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectPlatformsItems0) validatePlatformID(formats strfmt.Registry) error {

	if swag.IsZero(m.PlatformID) { // not required
		return nil
	}

	if err := validate.FormatOf("platformID", "body", "uuid", m.PlatformID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectPlatformsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectPlatformsItems0) UnmarshalBinary(b []byte) error {
	var res ProjectPlatformsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
