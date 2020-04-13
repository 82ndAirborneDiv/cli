// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewSetPreferredEmailParams creates a new SetPreferredEmailParams object
// with the default values initialized.
func NewSetPreferredEmailParams() *SetPreferredEmailParams {
	var ()
	return &SetPreferredEmailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetPreferredEmailParamsWithTimeout creates a new SetPreferredEmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetPreferredEmailParamsWithTimeout(timeout time.Duration) *SetPreferredEmailParams {
	var ()
	return &SetPreferredEmailParams{

		timeout: timeout,
	}
}

// NewSetPreferredEmailParamsWithContext creates a new SetPreferredEmailParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetPreferredEmailParamsWithContext(ctx context.Context) *SetPreferredEmailParams {
	var ()
	return &SetPreferredEmailParams{

		Context: ctx,
	}
}

// NewSetPreferredEmailParamsWithHTTPClient creates a new SetPreferredEmailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetPreferredEmailParamsWithHTTPClient(client *http.Client) *SetPreferredEmailParams {
	var ()
	return &SetPreferredEmailParams{
		HTTPClient: client,
	}
}

/*SetPreferredEmailParams contains all the parameters to send to the API endpoint
for the set preferred email operation typically these are written to a http.Request
*/
type SetPreferredEmailParams struct {

	/*Email
	  email address to change

	*/
	Email string
	/*Username
	  username of desired User

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set preferred email params
func (o *SetPreferredEmailParams) WithTimeout(timeout time.Duration) *SetPreferredEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set preferred email params
func (o *SetPreferredEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set preferred email params
func (o *SetPreferredEmailParams) WithContext(ctx context.Context) *SetPreferredEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set preferred email params
func (o *SetPreferredEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set preferred email params
func (o *SetPreferredEmailParams) WithHTTPClient(client *http.Client) *SetPreferredEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set preferred email params
func (o *SetPreferredEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the set preferred email params
func (o *SetPreferredEmailParams) WithEmail(email string) *SetPreferredEmailParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the set preferred email params
func (o *SetPreferredEmailParams) SetEmail(email string) {
	o.Email = email
}

// WithUsername adds the username to the set preferred email params
func (o *SetPreferredEmailParams) WithUsername(username string) *SetPreferredEmailParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the set preferred email params
func (o *SetPreferredEmailParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *SetPreferredEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param email
	if err := r.SetPathParam("email", o.Email); err != nil {
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
