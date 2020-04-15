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

// NewAddTokenParams creates a new AddTokenParams object
// with the default values initialized.
func NewAddTokenParams() *AddTokenParams {
	var ()
	return &AddTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddTokenParamsWithTimeout creates a new AddTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddTokenParamsWithTimeout(timeout time.Duration) *AddTokenParams {
	var ()
	return &AddTokenParams{

		timeout: timeout,
	}
}

// NewAddTokenParamsWithContext creates a new AddTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddTokenParamsWithContext(ctx context.Context) *AddTokenParams {
	var ()
	return &AddTokenParams{

		Context: ctx,
	}
}

// NewAddTokenParamsWithHTTPClient creates a new AddTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddTokenParamsWithHTTPClient(client *http.Client) *AddTokenParams {
	var ()
	return &AddTokenParams{
		HTTPClient: client,
	}
}

/*AddTokenParams contains all the parameters to send to the API endpoint
for the add token operation typically these are written to a http.Request
*/
type AddTokenParams struct {

	/*TokenOptions*/
	TokenOptions *mono_models.TokenEditable

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add token params
func (o *AddTokenParams) WithTimeout(timeout time.Duration) *AddTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add token params
func (o *AddTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add token params
func (o *AddTokenParams) WithContext(ctx context.Context) *AddTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add token params
func (o *AddTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add token params
func (o *AddTokenParams) WithHTTPClient(client *http.Client) *AddTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add token params
func (o *AddTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTokenOptions adds the tokenOptions to the add token params
func (o *AddTokenParams) WithTokenOptions(tokenOptions *mono_models.TokenEditable) *AddTokenParams {
	o.SetTokenOptions(tokenOptions)
	return o
}

// SetTokenOptions adds the tokenOptions to the add token params
func (o *AddTokenParams) SetTokenOptions(tokenOptions *mono_models.TokenEditable) {
	o.TokenOptions = tokenOptions
}

// WriteToRequest writes these params to a swagger request
func (o *AddTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TokenOptions != nil {
		if err := r.SetBodyParam(o.TokenOptions); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
