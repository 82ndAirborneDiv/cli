// Code generated by go-swagger; DO NOT EDIT.

package invoices

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

// NewCalculateTaxParams creates a new CalculateTaxParams object
// with the default values initialized.
func NewCalculateTaxParams() *CalculateTaxParams {
	var ()
	return &CalculateTaxParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCalculateTaxParamsWithTimeout creates a new CalculateTaxParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCalculateTaxParamsWithTimeout(timeout time.Duration) *CalculateTaxParams {
	var ()
	return &CalculateTaxParams{

		timeout: timeout,
	}
}

// NewCalculateTaxParamsWithContext creates a new CalculateTaxParams object
// with the default values initialized, and the ability to set a context for a request
func NewCalculateTaxParamsWithContext(ctx context.Context) *CalculateTaxParams {
	var ()
	return &CalculateTaxParams{

		Context: ctx,
	}
}

// NewCalculateTaxParamsWithHTTPClient creates a new CalculateTaxParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCalculateTaxParamsWithHTTPClient(client *http.Client) *CalculateTaxParams {
	var ()
	return &CalculateTaxParams{
		HTTPClient: client,
	}
}

/*CalculateTaxParams contains all the parameters to send to the API endpoint
for the calculate tax operation typically these are written to a http.Request
*/
type CalculateTaxParams struct {

	/*TaxOptions*/
	TaxOptions *mono_models.TaxOptions

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the calculate tax params
func (o *CalculateTaxParams) WithTimeout(timeout time.Duration) *CalculateTaxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the calculate tax params
func (o *CalculateTaxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the calculate tax params
func (o *CalculateTaxParams) WithContext(ctx context.Context) *CalculateTaxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the calculate tax params
func (o *CalculateTaxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the calculate tax params
func (o *CalculateTaxParams) WithHTTPClient(client *http.Client) *CalculateTaxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the calculate tax params
func (o *CalculateTaxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaxOptions adds the taxOptions to the calculate tax params
func (o *CalculateTaxParams) WithTaxOptions(taxOptions *mono_models.TaxOptions) *CalculateTaxParams {
	o.SetTaxOptions(taxOptions)
	return o
}

// SetTaxOptions adds the taxOptions to the calculate tax params
func (o *CalculateTaxParams) SetTaxOptions(taxOptions *mono_models.TaxOptions) {
	o.TaxOptions = taxOptions
}

// WriteToRequest writes these params to a swagger request
func (o *CalculateTaxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TaxOptions != nil {
		if err := r.SetBodyParam(o.TaxOptions); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
