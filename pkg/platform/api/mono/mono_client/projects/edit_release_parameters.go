// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// NewEditReleaseParams creates a new EditReleaseParams object
// with the default values initialized.
func NewEditReleaseParams() *EditReleaseParams {
	var ()
	return &EditReleaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEditReleaseParamsWithTimeout creates a new EditReleaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEditReleaseParamsWithTimeout(timeout time.Duration) *EditReleaseParams {
	var ()
	return &EditReleaseParams{

		timeout: timeout,
	}
}

// NewEditReleaseParamsWithContext creates a new EditReleaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewEditReleaseParamsWithContext(ctx context.Context) *EditReleaseParams {
	var ()
	return &EditReleaseParams{

		Context: ctx,
	}
}

// NewEditReleaseParamsWithHTTPClient creates a new EditReleaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEditReleaseParamsWithHTTPClient(client *http.Client) *EditReleaseParams {
	var ()
	return &EditReleaseParams{
		HTTPClient: client,
	}
}

/*EditReleaseParams contains all the parameters to send to the API endpoint
for the edit release operation typically these are written to a http.Request
*/
type EditReleaseParams struct {

	/*OrganizationName
	  organizationName of desired organization

	*/
	OrganizationName string
	/*ProjectName
	  projectName of desired project

	*/
	ProjectName string
	/*Release*/
	Release *mono_models.Release
	/*ReleaseID
	  release to fetch

	*/
	ReleaseID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edit release params
func (o *EditReleaseParams) WithTimeout(timeout time.Duration) *EditReleaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit release params
func (o *EditReleaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit release params
func (o *EditReleaseParams) WithContext(ctx context.Context) *EditReleaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit release params
func (o *EditReleaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit release params
func (o *EditReleaseParams) WithHTTPClient(client *http.Client) *EditReleaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit release params
func (o *EditReleaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationName adds the organizationName to the edit release params
func (o *EditReleaseParams) WithOrganizationName(organizationName string) *EditReleaseParams {
	o.SetOrganizationName(organizationName)
	return o
}

// SetOrganizationName adds the organizationName to the edit release params
func (o *EditReleaseParams) SetOrganizationName(organizationName string) {
	o.OrganizationName = organizationName
}

// WithProjectName adds the projectName to the edit release params
func (o *EditReleaseParams) WithProjectName(projectName string) *EditReleaseParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the edit release params
func (o *EditReleaseParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithRelease adds the release to the edit release params
func (o *EditReleaseParams) WithRelease(release *mono_models.Release) *EditReleaseParams {
	o.SetRelease(release)
	return o
}

// SetRelease adds the release to the edit release params
func (o *EditReleaseParams) SetRelease(release *mono_models.Release) {
	o.Release = release
}

// WithReleaseID adds the releaseID to the edit release params
func (o *EditReleaseParams) WithReleaseID(releaseID strfmt.UUID) *EditReleaseParams {
	o.SetReleaseID(releaseID)
	return o
}

// SetReleaseID adds the releaseId to the edit release params
func (o *EditReleaseParams) SetReleaseID(releaseID strfmt.UUID) {
	o.ReleaseID = releaseID
}

// WriteToRequest writes these params to a swagger request
func (o *EditReleaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationName
	if err := r.SetPathParam("organizationName", o.OrganizationName); err != nil {
		return err
	}

	// path param projectName
	if err := r.SetPathParam("projectName", o.ProjectName); err != nil {
		return err
	}

	if o.Release != nil {
		if err := r.SetBodyParam(o.Release); err != nil {
			return err
		}
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
