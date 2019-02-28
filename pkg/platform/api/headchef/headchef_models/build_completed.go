// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuildCompleted Build Completed
//
// A message indicating that a requested build has been completed.
// swagger:model buildCompleted
type BuildCompleted struct {

	// An array containing all the artifacts that make up this build. This is suitable for passing as part of a To Go Request.
	// Required: true
	Artifacts []*BuildCompletedArtifactsItems0 `json:"artifacts"`

	// An S3 URI containing the log for this build.
	// Format: uri
	LogURI strfmt.URI `json:"log_uri,omitempty"`

	// A user-facing message describing the build results.
	Message string `json:"message,omitempty"`
}

// Validate validates this build completed
func (m *BuildCompleted) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtifacts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildCompleted) validateArtifacts(formats strfmt.Registry) error {

	if err := validate.Required("artifacts", "body", m.Artifacts); err != nil {
		return err
	}

	for i := 0; i < len(m.Artifacts); i++ {
		if swag.IsZero(m.Artifacts[i]) { // not required
			continue
		}

		if m.Artifacts[i] != nil {
			if err := m.Artifacts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("artifacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BuildCompleted) validateLogURI(formats strfmt.Registry) error {

	if swag.IsZero(m.LogURI) { // not required
		return nil
	}

	if err := validate.FormatOf("log_uri", "body", "uri", m.LogURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildCompleted) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildCompleted) UnmarshalBinary(b []byte) error {
	var res BuildCompleted
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BuildCompletedArtifactsItems0 Artifact
//
// The result of building a single ingredient is an artifact, which contains the files created by the build.
// swagger:model BuildCompletedArtifactsItems0
type BuildCompletedArtifactsItems0 struct {

	// Artifact ID
	// Required: true
	// Format: uuid
	ArtifactID *strfmt.UUID `json:"artifact_id"`

	// Timestamp for when the artifact was created
	// Required: true
	// Format: date-time
	BuildTimestamp *strfmt.DateTime `json:"build_timestamp"`

	// dependency ids
	DependencyIds []strfmt.UUID `json:"dependency_ids"`

	// Ingredient Version ID
	//
	// Source Ingredient Version for the artifact
	// Format: uuid
	IngredientVersionID strfmt.UUID `json:"ingredient_version_id,omitempty"`

	// Platform ID for the artifact
	// Required: true
	// Format: uuid
	PlatformID *strfmt.UUID `json:"platform_id"`

	// URI for the storage location of the artifact
	// Required: true
	// Format: uri
	URI *strfmt.URI `json:"uri"`
}

// Validate validates this build completed artifacts items0
func (m *BuildCompletedArtifactsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtifactID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuildTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDependencyIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredientVersionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildCompletedArtifactsItems0) validateArtifactID(formats strfmt.Registry) error {

	if err := validate.Required("artifact_id", "body", m.ArtifactID); err != nil {
		return err
	}

	if err := validate.FormatOf("artifact_id", "body", "uuid", m.ArtifactID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildCompletedArtifactsItems0) validateBuildTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("build_timestamp", "body", m.BuildTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("build_timestamp", "body", "date-time", m.BuildTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildCompletedArtifactsItems0) validateDependencyIds(formats strfmt.Registry) error {

	if swag.IsZero(m.DependencyIds) { // not required
		return nil
	}

	for i := 0; i < len(m.DependencyIds); i++ {

		if err := validate.FormatOf("dependency_ids"+"."+strconv.Itoa(i), "body", "uuid", m.DependencyIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *BuildCompletedArtifactsItems0) validateIngredientVersionID(formats strfmt.Registry) error {

	if swag.IsZero(m.IngredientVersionID) { // not required
		return nil
	}

	if err := validate.FormatOf("ingredient_version_id", "body", "uuid", m.IngredientVersionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildCompletedArtifactsItems0) validatePlatformID(formats strfmt.Registry) error {

	if err := validate.Required("platform_id", "body", m.PlatformID); err != nil {
		return err
	}

	if err := validate.FormatOf("platform_id", "body", "uuid", m.PlatformID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildCompletedArtifactsItems0) validateURI(formats strfmt.Registry) error {

	if err := validate.Required("uri", "body", m.URI); err != nil {
		return err
	}

	if err := validate.FormatOf("uri", "body", "uri", m.URI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildCompletedArtifactsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildCompletedArtifactsItems0) UnmarshalBinary(b []byte) error {
	var res BuildCompletedArtifactsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}