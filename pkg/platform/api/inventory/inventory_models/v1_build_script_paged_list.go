// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1BuildScriptPagedList Build Script Paged List
//
// A paginated list of build scripts
// swagger:model v1BuildScriptPagedList
type V1BuildScriptPagedList struct {

	// A page of build scripts
	// Required: true
	BuildScripts []*V1BuildScriptPagedListBuildScriptsItems `json:"build_scripts"`

	// links
	// Required: true
	Links *V1BuildScriptPagedListLinks `json:"links"`

	// paging
	// Required: true
	Paging *V1BuildScriptPagedListPaging `json:"paging"`
}

// Validate validates this v1 build script paged list
func (m *V1BuildScriptPagedList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildScripts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaging(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1BuildScriptPagedList) validateBuildScripts(formats strfmt.Registry) error {

	if err := validate.Required("build_scripts", "body", m.BuildScripts); err != nil {
		return err
	}

	for i := 0; i < len(m.BuildScripts); i++ {
		if swag.IsZero(m.BuildScripts[i]) { // not required
			continue
		}

		if m.BuildScripts[i] != nil {
			if err := m.BuildScripts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("build_scripts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1BuildScriptPagedList) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("links", "body", m.Links); err != nil {
		return err
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

func (m *V1BuildScriptPagedList) validatePaging(formats strfmt.Registry) error {

	if err := validate.Required("paging", "body", m.Paging); err != nil {
		return err
	}

	if m.Paging != nil {
		if err := m.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("paging")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildScriptPagedList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildScriptPagedList) UnmarshalBinary(b []byte) error {
	var res V1BuildScriptPagedList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
