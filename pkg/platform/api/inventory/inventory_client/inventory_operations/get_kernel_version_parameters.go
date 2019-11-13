// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetKernelVersionParams creates a new GetKernelVersionParams object
// with the default values initialized.
func NewGetKernelVersionParams() *GetKernelVersionParams {
	var (
		allowUnstableDefault = bool(false)
	)
	return &GetKernelVersionParams{
		AllowUnstable: &allowUnstableDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKernelVersionParamsWithTimeout creates a new GetKernelVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKernelVersionParamsWithTimeout(timeout time.Duration) *GetKernelVersionParams {
	var (
		allowUnstableDefault = bool(false)
	)
	return &GetKernelVersionParams{
		AllowUnstable: &allowUnstableDefault,

		timeout: timeout,
	}
}

// NewGetKernelVersionParamsWithContext creates a new GetKernelVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKernelVersionParamsWithContext(ctx context.Context) *GetKernelVersionParams {
	var (
		allowUnstableDefault = bool(false)
	)
	return &GetKernelVersionParams{
		AllowUnstable: &allowUnstableDefault,

		Context: ctx,
	}
}

// NewGetKernelVersionParamsWithHTTPClient creates a new GetKernelVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKernelVersionParamsWithHTTPClient(client *http.Client) *GetKernelVersionParams {
	var (
		allowUnstableDefault = bool(false)
	)
	return &GetKernelVersionParams{
		AllowUnstable: &allowUnstableDefault,
		HTTPClient:    client,
	}
}

/*GetKernelVersionParams contains all the parameters to send to the API endpoint
for the get kernel version operation typically these are written to a http.Request
*/
type GetKernelVersionParams struct {

	/*AllowUnstable
	  Whether to show an unstable revision of a resource if there is an available unstable version newer than the newest available stable version

	*/
	AllowUnstable *bool
	/*KernelID*/
	KernelID strfmt.UUID
	/*KernelVersionID*/
	KernelVersionID strfmt.UUID
	/*StateAt
	  Show the state of a resource as it was at the specified timestamp. If omitted, shows the current state of the resource.

	*/
	StateAt *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get kernel version params
func (o *GetKernelVersionParams) WithTimeout(timeout time.Duration) *GetKernelVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kernel version params
func (o *GetKernelVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kernel version params
func (o *GetKernelVersionParams) WithContext(ctx context.Context) *GetKernelVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kernel version params
func (o *GetKernelVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kernel version params
func (o *GetKernelVersionParams) WithHTTPClient(client *http.Client) *GetKernelVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kernel version params
func (o *GetKernelVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowUnstable adds the allowUnstable to the get kernel version params
func (o *GetKernelVersionParams) WithAllowUnstable(allowUnstable *bool) *GetKernelVersionParams {
	o.SetAllowUnstable(allowUnstable)
	return o
}

// SetAllowUnstable adds the allowUnstable to the get kernel version params
func (o *GetKernelVersionParams) SetAllowUnstable(allowUnstable *bool) {
	o.AllowUnstable = allowUnstable
}

// WithKernelID adds the kernelID to the get kernel version params
func (o *GetKernelVersionParams) WithKernelID(kernelID strfmt.UUID) *GetKernelVersionParams {
	o.SetKernelID(kernelID)
	return o
}

// SetKernelID adds the kernelId to the get kernel version params
func (o *GetKernelVersionParams) SetKernelID(kernelID strfmt.UUID) {
	o.KernelID = kernelID
}

// WithKernelVersionID adds the kernelVersionID to the get kernel version params
func (o *GetKernelVersionParams) WithKernelVersionID(kernelVersionID strfmt.UUID) *GetKernelVersionParams {
	o.SetKernelVersionID(kernelVersionID)
	return o
}

// SetKernelVersionID adds the kernelVersionId to the get kernel version params
func (o *GetKernelVersionParams) SetKernelVersionID(kernelVersionID strfmt.UUID) {
	o.KernelVersionID = kernelVersionID
}

// WithStateAt adds the stateAt to the get kernel version params
func (o *GetKernelVersionParams) WithStateAt(stateAt *strfmt.DateTime) *GetKernelVersionParams {
	o.SetStateAt(stateAt)
	return o
}

// SetStateAt adds the stateAt to the get kernel version params
func (o *GetKernelVersionParams) SetStateAt(stateAt *strfmt.DateTime) {
	o.StateAt = stateAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetKernelVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowUnstable != nil {

		// query param allow_unstable
		var qrAllowUnstable bool
		if o.AllowUnstable != nil {
			qrAllowUnstable = *o.AllowUnstable
		}
		qAllowUnstable := swag.FormatBool(qrAllowUnstable)
		if qAllowUnstable != "" {
			if err := r.SetQueryParam("allow_unstable", qAllowUnstable); err != nil {
				return err
			}
		}

	}

	// path param kernel_id
	if err := r.SetPathParam("kernel_id", o.KernelID.String()); err != nil {
		return err
	}

	// path param kernel_version_id
	if err := r.SetPathParam("kernel_version_id", o.KernelVersionID.String()); err != nil {
		return err
	}

	if o.StateAt != nil {

		// query param state_at
		var qrStateAt strfmt.DateTime
		if o.StateAt != nil {
			qrStateAt = *o.StateAt
		}
		qStateAt := qrStateAt.String()
		if qStateAt != "" {
			if err := r.SetQueryParam("state_at", qStateAt); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}