// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

// NewChangePasswordParams creates a new ChangePasswordParams object
// with the default values initialized.
func NewChangePasswordParams() *ChangePasswordParams {
	var ()
	return &ChangePasswordParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangePasswordParamsWithTimeout creates a new ChangePasswordParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangePasswordParamsWithTimeout(timeout time.Duration) *ChangePasswordParams {
	var ()
	return &ChangePasswordParams{

		timeout: timeout,
	}
}

// NewChangePasswordParamsWithContext creates a new ChangePasswordParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangePasswordParamsWithContext(ctx context.Context) *ChangePasswordParams {
	var ()
	return &ChangePasswordParams{

		Context: ctx,
	}
}

// NewChangePasswordParamsWithHTTPClient creates a new ChangePasswordParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangePasswordParamsWithHTTPClient(client *http.Client) *ChangePasswordParams {
	var ()
	return &ChangePasswordParams{
		HTTPClient: client,
	}
}

/*ChangePasswordParams contains all the parameters to send to the API endpoint
for the change password operation typically these are written to a http.Request
*/
type ChangePasswordParams struct {

	/*ChangeRequest
	  change Request

	*/
	ChangeRequest *mono_models.PasswordChange

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change password params
func (o *ChangePasswordParams) WithTimeout(timeout time.Duration) *ChangePasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change password params
func (o *ChangePasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change password params
func (o *ChangePasswordParams) WithContext(ctx context.Context) *ChangePasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change password params
func (o *ChangePasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change password params
func (o *ChangePasswordParams) WithHTTPClient(client *http.Client) *ChangePasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change password params
func (o *ChangePasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChangeRequest adds the changeRequest to the change password params
func (o *ChangePasswordParams) WithChangeRequest(changeRequest *mono_models.PasswordChange) *ChangePasswordParams {
	o.SetChangeRequest(changeRequest)
	return o
}

// SetChangeRequest adds the changeRequest to the change password params
func (o *ChangePasswordParams) SetChangeRequest(changeRequest *mono_models.PasswordChange) {
	o.ChangeRequest = changeRequest
}

// WriteToRequest writes these params to a swagger request
func (o *ChangePasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ChangeRequest != nil {
		if err := r.SetBodyParam(o.ChangeRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
