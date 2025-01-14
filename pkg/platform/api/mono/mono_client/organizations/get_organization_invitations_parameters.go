// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetOrganizationInvitationsParams creates a new GetOrganizationInvitationsParams object
// with the default values initialized.
func NewGetOrganizationInvitationsParams() *GetOrganizationInvitationsParams {
	var ()
	return &GetOrganizationInvitationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationInvitationsParamsWithTimeout creates a new GetOrganizationInvitationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrganizationInvitationsParamsWithTimeout(timeout time.Duration) *GetOrganizationInvitationsParams {
	var ()
	return &GetOrganizationInvitationsParams{

		timeout: timeout,
	}
}

// NewGetOrganizationInvitationsParamsWithContext creates a new GetOrganizationInvitationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrganizationInvitationsParamsWithContext(ctx context.Context) *GetOrganizationInvitationsParams {
	var ()
	return &GetOrganizationInvitationsParams{

		Context: ctx,
	}
}

// NewGetOrganizationInvitationsParamsWithHTTPClient creates a new GetOrganizationInvitationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrganizationInvitationsParamsWithHTTPClient(client *http.Client) *GetOrganizationInvitationsParams {
	var ()
	return &GetOrganizationInvitationsParams{
		HTTPClient: client,
	}
}

/*GetOrganizationInvitationsParams contains all the parameters to send to the API endpoint
for the get organization invitations operation typically these are written to a http.Request
*/
type GetOrganizationInvitationsParams struct {

	/*OrganizationName
	  organizationName of desired organization

	*/
	OrganizationName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get organization invitations params
func (o *GetOrganizationInvitationsParams) WithTimeout(timeout time.Duration) *GetOrganizationInvitationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization invitations params
func (o *GetOrganizationInvitationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization invitations params
func (o *GetOrganizationInvitationsParams) WithContext(ctx context.Context) *GetOrganizationInvitationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization invitations params
func (o *GetOrganizationInvitationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization invitations params
func (o *GetOrganizationInvitationsParams) WithHTTPClient(client *http.Client) *GetOrganizationInvitationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization invitations params
func (o *GetOrganizationInvitationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationName adds the organizationName to the get organization invitations params
func (o *GetOrganizationInvitationsParams) WithOrganizationName(organizationName string) *GetOrganizationInvitationsParams {
	o.SetOrganizationName(organizationName)
	return o
}

// SetOrganizationName adds the organizationName to the get organization invitations params
func (o *GetOrganizationInvitationsParams) SetOrganizationName(organizationName string) {
	o.OrganizationName = organizationName
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationInvitationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationName
	if err := r.SetPathParam("organizationName", o.OrganizationName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
