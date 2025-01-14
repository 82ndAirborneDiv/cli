// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuildRequest Build Request
//
// A build request which is submitted to the Head Chef REST API.
// swagger:model buildRequest
type BuildRequest struct {

	// A list of additional command-line parameters to pass to setup-builds.pl. NOTE: this is a temporary feature to expose some camel features before build options are implemented.
	// Unique: true
	CamelFlags []string `json:"camel_flags"`

	// format
	// Required: true
	// Enum: [7zip dmg msi raw tarball zip]
	Format *string `json:"format"`

	// recipe
	// Required: true
	Recipe *BuildRequestRecipe `json:"recipe"`

	// requester
	Requester *Requester `json:"requester,omitempty"`
}

// Validate validates this build request
func (m *BuildRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCamelFlags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequester(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var buildRequestCamelFlagsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["--dynamic-core","--python-debug","--tcl-debug","--tcl-disable-threads"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buildRequestCamelFlagsItemsEnum = append(buildRequestCamelFlagsItemsEnum, v)
	}
}

func (m *BuildRequest) validateCamelFlagsItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, buildRequestCamelFlagsItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *BuildRequest) validateCamelFlags(formats strfmt.Registry) error {

	if swag.IsZero(m.CamelFlags) { // not required
		return nil
	}

	if err := validate.UniqueItems("camel_flags", "body", m.CamelFlags); err != nil {
		return err
	}

	for i := 0; i < len(m.CamelFlags); i++ {

		// value enum
		if err := m.validateCamelFlagsItemsEnum("camel_flags"+"."+strconv.Itoa(i), "body", m.CamelFlags[i]); err != nil {
			return err
		}

	}

	return nil
}

var buildRequestTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["7zip","dmg","msi","raw","tarball","zip"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buildRequestTypeFormatPropEnum = append(buildRequestTypeFormatPropEnum, v)
	}
}

const (

	// BuildRequestFormatNr7zip captures enum value "7zip"
	BuildRequestFormatNr7zip string = "7zip"

	// BuildRequestFormatDmg captures enum value "dmg"
	BuildRequestFormatDmg string = "dmg"

	// BuildRequestFormatMsi captures enum value "msi"
	BuildRequestFormatMsi string = "msi"

	// BuildRequestFormatRaw captures enum value "raw"
	BuildRequestFormatRaw string = "raw"

	// BuildRequestFormatTarball captures enum value "tarball"
	BuildRequestFormatTarball string = "tarball"

	// BuildRequestFormatZip captures enum value "zip"
	BuildRequestFormatZip string = "zip"
)

// prop value enum
func (m *BuildRequest) validateFormatEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, buildRequestTypeFormatPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BuildRequest) validateFormat(formats strfmt.Registry) error {

	if err := validate.Required("format", "body", m.Format); err != nil {
		return err
	}

	// value enum
	if err := m.validateFormatEnum("format", "body", *m.Format); err != nil {
		return err
	}

	return nil
}

func (m *BuildRequest) validateRecipe(formats strfmt.Registry) error {

	if err := validate.Required("recipe", "body", m.Recipe); err != nil {
		return err
	}

	if m.Recipe != nil {
		if err := m.Recipe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipe")
			}
			return err
		}
	}

	return nil
}

func (m *BuildRequest) validateRequester(formats strfmt.Registry) error {

	if swag.IsZero(m.Requester) { // not required
		return nil
	}

	if m.Requester != nil {
		if err := m.Requester.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requester")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildRequest) UnmarshalBinary(b []byte) error {
	var res BuildRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
