// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BranchEditable branchEditable
// swagger:model BranchEditable
type BranchEditable struct {

	// The commit that this branch is currently pointing at
	// Format: uuid
	CommitID *strfmt.UUID `json:"commitID,omitempty"`

	// The tracking method used
	// Enum: [ignore auto_update notify]
	TrackingType *string `json:"tracking_type,omitempty"`

	// The branch_id of the branch that this branch tracks
	// Format: uuid
	Tracks *strfmt.UUID `json:"tracks,omitempty"`
}

// Validate validates this branch editable
func (m *BranchEditable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommitID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackingType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTracks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BranchEditable) validateCommitID(formats strfmt.Registry) error {

	if swag.IsZero(m.CommitID) { // not required
		return nil
	}

	if err := validate.FormatOf("commitID", "body", "uuid", m.CommitID.String(), formats); err != nil {
		return err
	}

	return nil
}

var branchEditableTypeTrackingTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ignore","auto_update","notify"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		branchEditableTypeTrackingTypePropEnum = append(branchEditableTypeTrackingTypePropEnum, v)
	}
}

const (

	// BranchEditableTrackingTypeIgnore captures enum value "ignore"
	BranchEditableTrackingTypeIgnore string = "ignore"

	// BranchEditableTrackingTypeAutoUpdate captures enum value "auto_update"
	BranchEditableTrackingTypeAutoUpdate string = "auto_update"

	// BranchEditableTrackingTypeNotify captures enum value "notify"
	BranchEditableTrackingTypeNotify string = "notify"
)

// prop value enum
func (m *BranchEditable) validateTrackingTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, branchEditableTypeTrackingTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BranchEditable) validateTrackingType(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackingType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTrackingTypeEnum("tracking_type", "body", *m.TrackingType); err != nil {
		return err
	}

	return nil
}

func (m *BranchEditable) validateTracks(formats strfmt.Registry) error {

	if swag.IsZero(m.Tracks) { // not required
		return nil
	}

	if err := validate.FormatOf("tracks", "body", "uuid", m.Tracks.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BranchEditable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BranchEditable) UnmarshalBinary(b []byte) error {
	var res BranchEditable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
