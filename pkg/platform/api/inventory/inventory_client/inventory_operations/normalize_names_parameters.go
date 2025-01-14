// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

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

	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// NewNormalizeNamesParams creates a new NormalizeNamesParams object
// with the default values initialized.
func NewNormalizeNamesParams() *NormalizeNamesParams {
	var ()
	return &NormalizeNamesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNormalizeNamesParamsWithTimeout creates a new NormalizeNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNormalizeNamesParamsWithTimeout(timeout time.Duration) *NormalizeNamesParams {
	var ()
	return &NormalizeNamesParams{

		timeout: timeout,
	}
}

// NewNormalizeNamesParamsWithContext creates a new NormalizeNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewNormalizeNamesParamsWithContext(ctx context.Context) *NormalizeNamesParams {
	var ()
	return &NormalizeNamesParams{

		Context: ctx,
	}
}

// NewNormalizeNamesParamsWithHTTPClient creates a new NormalizeNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNormalizeNamesParamsWithHTTPClient(client *http.Client) *NormalizeNamesParams {
	var ()
	return &NormalizeNamesParams{
		HTTPClient: client,
	}
}

/*NormalizeNamesParams contains all the parameters to send to the API endpoint
for the normalize names operation typically these are written to a http.Request
*/
type NormalizeNamesParams struct {

	/*Names*/
	Names *inventory_models.UnnormalizedNames
	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the normalize names params
func (o *NormalizeNamesParams) WithTimeout(timeout time.Duration) *NormalizeNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the normalize names params
func (o *NormalizeNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the normalize names params
func (o *NormalizeNamesParams) WithContext(ctx context.Context) *NormalizeNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the normalize names params
func (o *NormalizeNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the normalize names params
func (o *NormalizeNamesParams) WithHTTPClient(client *http.Client) *NormalizeNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the normalize names params
func (o *NormalizeNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNames adds the names to the normalize names params
func (o *NormalizeNamesParams) WithNames(names *inventory_models.UnnormalizedNames) *NormalizeNamesParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the normalize names params
func (o *NormalizeNamesParams) SetNames(names *inventory_models.UnnormalizedNames) {
	o.Names = names
}

// WithNamespace adds the namespace to the normalize names params
func (o *NormalizeNamesParams) WithNamespace(namespace string) *NormalizeNamesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the normalize names params
func (o *NormalizeNamesParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *NormalizeNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Names != nil {
		if err := r.SetBodyParam(o.Names); err != nil {
			return err
		}
	}

	// query param namespace
	qrNamespace := o.Namespace
	qNamespace := qrNamespace
	if qNamespace != "" {
		if err := r.SetQueryParam("namespace", qNamespace); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
