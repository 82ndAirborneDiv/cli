// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// NewAddFormatParams creates a new AddFormatParams object
// with the default values initialized.
func NewAddFormatParams() *AddFormatParams {
	var ()
	return &AddFormatParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddFormatParamsWithTimeout creates a new AddFormatParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddFormatParamsWithTimeout(timeout time.Duration) *AddFormatParams {
	var ()
	return &AddFormatParams{

		timeout: timeout,
	}
}

// NewAddFormatParamsWithContext creates a new AddFormatParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddFormatParamsWithContext(ctx context.Context) *AddFormatParams {
	var ()
	return &AddFormatParams{

		Context: ctx,
	}
}

// NewAddFormatParamsWithHTTPClient creates a new AddFormatParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddFormatParamsWithHTTPClient(client *http.Client) *AddFormatParams {
	var ()
	return &AddFormatParams{
		HTTPClient: client,
	}
}

/*AddFormatParams contains all the parameters to send to the API endpoint
for the add format operation typically these are written to a http.Request
*/
type AddFormatParams struct {

	/*DistroID
	  desired distro

	*/
	DistroID strfmt.UUID
	/*Format
	  Format details

	*/
	Format *mono_models.Format
	/*OrganizationName
	  desired organization

	*/
	OrganizationName string
	/*ProjectName
	  desired project

	*/
	ProjectName string
	/*ReleaseID
	  desired release

	*/
	ReleaseID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add format params
func (o *AddFormatParams) WithTimeout(timeout time.Duration) *AddFormatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add format params
func (o *AddFormatParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add format params
func (o *AddFormatParams) WithContext(ctx context.Context) *AddFormatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add format params
func (o *AddFormatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add format params
func (o *AddFormatParams) WithHTTPClient(client *http.Client) *AddFormatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add format params
func (o *AddFormatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDistroID adds the distroID to the add format params
func (o *AddFormatParams) WithDistroID(distroID strfmt.UUID) *AddFormatParams {
	o.SetDistroID(distroID)
	return o
}

// SetDistroID adds the distroId to the add format params
func (o *AddFormatParams) SetDistroID(distroID strfmt.UUID) {
	o.DistroID = distroID
}

// WithFormat adds the format to the add format params
func (o *AddFormatParams) WithFormat(format *mono_models.Format) *AddFormatParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the add format params
func (o *AddFormatParams) SetFormat(format *mono_models.Format) {
	o.Format = format
}

// WithOrganizationName adds the organizationName to the add format params
func (o *AddFormatParams) WithOrganizationName(organizationName string) *AddFormatParams {
	o.SetOrganizationName(organizationName)
	return o
}

// SetOrganizationName adds the organizationName to the add format params
func (o *AddFormatParams) SetOrganizationName(organizationName string) {
	o.OrganizationName = organizationName
}

// WithProjectName adds the projectName to the add format params
func (o *AddFormatParams) WithProjectName(projectName string) *AddFormatParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the add format params
func (o *AddFormatParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithReleaseID adds the releaseID to the add format params
func (o *AddFormatParams) WithReleaseID(releaseID strfmt.UUID) *AddFormatParams {
	o.SetReleaseID(releaseID)
	return o
}

// SetReleaseID adds the releaseId to the add format params
func (o *AddFormatParams) SetReleaseID(releaseID strfmt.UUID) {
	o.ReleaseID = releaseID
}

// WriteToRequest writes these params to a swagger request
func (o *AddFormatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param distroID
	if err := r.SetPathParam("distroID", o.DistroID.String()); err != nil {
		return err
	}

	if o.Format != nil {
		if err := r.SetBodyParam(o.Format); err != nil {
			return err
		}
	}

	// path param organizationName
	if err := r.SetPathParam("organizationName", o.OrganizationName); err != nil {
		return err
	}

	// path param projectName
	if err := r.SetPathParam("projectName", o.ProjectName); err != nil {
		return err
	}

	// path param releaseID
	if err := r.SetPathParam("releaseID", o.ReleaseID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}