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

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// NewEditMemberParams creates a new EditMemberParams object
// with the default values initialized.
func NewEditMemberParams() *EditMemberParams {
	var ()
	return &EditMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEditMemberParamsWithTimeout creates a new EditMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEditMemberParamsWithTimeout(timeout time.Duration) *EditMemberParams {
	var ()
	return &EditMemberParams{

		timeout: timeout,
	}
}

// NewEditMemberParamsWithContext creates a new EditMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewEditMemberParamsWithContext(ctx context.Context) *EditMemberParams {
	var ()
	return &EditMemberParams{

		Context: ctx,
	}
}

// NewEditMemberParamsWithHTTPClient creates a new EditMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEditMemberParamsWithHTTPClient(client *http.Client) *EditMemberParams {
	var ()
	return &EditMemberParams{
		HTTPClient: client,
	}
}

/*EditMemberParams contains all the parameters to send to the API endpoint
for the edit member operation typically these are written to a http.Request
*/
type EditMemberParams struct {

	/*MemberAttrs
	  Member Attributes

	*/
	MemberAttrs *mono_models.MemberEditable
	/*OrganizationName
	  organizationID of desired organization

	*/
	OrganizationName string
	/*Username
	  username to join

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edit member params
func (o *EditMemberParams) WithTimeout(timeout time.Duration) *EditMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit member params
func (o *EditMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit member params
func (o *EditMemberParams) WithContext(ctx context.Context) *EditMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit member params
func (o *EditMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit member params
func (o *EditMemberParams) WithHTTPClient(client *http.Client) *EditMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit member params
func (o *EditMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMemberAttrs adds the memberAttrs to the edit member params
func (o *EditMemberParams) WithMemberAttrs(memberAttrs *mono_models.MemberEditable) *EditMemberParams {
	o.SetMemberAttrs(memberAttrs)
	return o
}

// SetMemberAttrs adds the memberAttrs to the edit member params
func (o *EditMemberParams) SetMemberAttrs(memberAttrs *mono_models.MemberEditable) {
	o.MemberAttrs = memberAttrs
}

// WithOrganizationName adds the organizationName to the edit member params
func (o *EditMemberParams) WithOrganizationName(organizationName string) *EditMemberParams {
	o.SetOrganizationName(organizationName)
	return o
}

// SetOrganizationName adds the organizationName to the edit member params
func (o *EditMemberParams) SetOrganizationName(organizationName string) {
	o.OrganizationName = organizationName
}

// WithUsername adds the username to the edit member params
func (o *EditMemberParams) WithUsername(username string) *EditMemberParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the edit member params
func (o *EditMemberParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *EditMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MemberAttrs != nil {
		if err := r.SetBodyParam(o.MemberAttrs); err != nil {
			return err
		}
	}

	// path param organizationName
	if err := r.SetPathParam("organizationName", o.OrganizationName); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
