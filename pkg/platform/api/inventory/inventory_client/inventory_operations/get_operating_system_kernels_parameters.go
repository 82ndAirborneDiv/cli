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
	"github.com/go-openapi/swag"
)

// NewGetOperatingSystemKernelsParams creates a new GetOperatingSystemKernelsParams object
// with the default values initialized.
func NewGetOperatingSystemKernelsParams() *GetOperatingSystemKernelsParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetOperatingSystemKernelsParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOperatingSystemKernelsParamsWithTimeout creates a new GetOperatingSystemKernelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOperatingSystemKernelsParamsWithTimeout(timeout time.Duration) *GetOperatingSystemKernelsParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetOperatingSystemKernelsParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,

		timeout: timeout,
	}
}

// NewGetOperatingSystemKernelsParamsWithContext creates a new GetOperatingSystemKernelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOperatingSystemKernelsParamsWithContext(ctx context.Context) *GetOperatingSystemKernelsParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetOperatingSystemKernelsParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,

		Context: ctx,
	}
}

// NewGetOperatingSystemKernelsParamsWithHTTPClient creates a new GetOperatingSystemKernelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOperatingSystemKernelsParamsWithHTTPClient(client *http.Client) *GetOperatingSystemKernelsParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetOperatingSystemKernelsParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,
		HTTPClient:    client,
	}
}

/*GetOperatingSystemKernelsParams contains all the parameters to send to the API endpoint
for the get operating system kernels operation typically these are written to a http.Request
*/
type GetOperatingSystemKernelsParams struct {

	/*AllowUnstable
	  Whether to show an unstable revision of a resource if there is an available unstable version newer than the newest available stable version

	*/
	AllowUnstable *bool
	/*Limit
	  The maximum number of items returned per page

	*/
	Limit *int64
	/*OperatingSystemID*/
	OperatingSystemID strfmt.UUID
	/*Page
	  The page number returned

	*/
	Page *int64
	/*StateAt
	  Show the state of a resource as it was at the specified timestamp. If omitted, shows the current state of the resource.

	*/
	StateAt *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get operating system kernels params
func (o *GetOperatingSystemKernelsParams) WithTimeout(timeout time.Duration) *GetOperatingSystemKernelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get operating system kernels params
func (o *GetOperatingSystemKernelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get operating system kernels params
func (o *GetOperatingSystemKernelsParams) WithContext(ctx context.Context) *GetOperatingSystemKernelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get operating system kernels params
func (o *GetOperatingSystemKernelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get operating system kernels params
func (o *GetOperatingSystemKernelsParams) WithHTTPClient(client *http.Client) *GetOperatingSystemKernelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get operating system kernels params
func (o *GetOperatingSystemKernelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowUnstable adds the allowUnstable to the get operating system kernels params
func (o *GetOperatingSystemKernelsParams) WithAllowUnstable(allowUnstable *bool) *GetOperatingSystemKernelsParams {
	o.SetAllowUnstable(allowUnstable)
	return o
}

// SetAllowUnstable adds the allowUnstable to the get operating system kernels params
func (o *GetOperatingSystemKernelsParams) SetAllowUnstable(allowUnstable *bool) {
	o.AllowUnstable = allowUnstable
}

// WithLimit adds the limit to the get operating system kernels params
func (o *GetOperatingSystemKernelsParams) WithLimit(limit *int64) *GetOperatingSystemKernelsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get operating system kernels params
func (o *GetOperatingSystemKernelsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOperatingSystemID adds the operatingSystemID to the get operating system kernels params
func (o *GetOperatingSystemKernelsParams) WithOperatingSystemID(operatingSystemID strfmt.UUID) *GetOperatingSystemKernelsParams {
	o.SetOperatingSystemID(operatingSystemID)
	return o
}

// SetOperatingSystemID adds the operatingSystemId to the get operating system kernels params
func (o *GetOperatingSystemKernelsParams) SetOperatingSystemID(operatingSystemID strfmt.UUID) {
	o.OperatingSystemID = operatingSystemID
}

// WithPage adds the page to the get operating system kernels params
func (o *GetOperatingSystemKernelsParams) WithPage(page *int64) *GetOperatingSystemKernelsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get operating system kernels params
func (o *GetOperatingSystemKernelsParams) SetPage(page *int64) {
	o.Page = page
}

// WithStateAt adds the stateAt to the get operating system kernels params
func (o *GetOperatingSystemKernelsParams) WithStateAt(stateAt *strfmt.DateTime) *GetOperatingSystemKernelsParams {
	o.SetStateAt(stateAt)
	return o
}

// SetStateAt adds the stateAt to the get operating system kernels params
func (o *GetOperatingSystemKernelsParams) SetStateAt(stateAt *strfmt.DateTime) {
	o.StateAt = stateAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetOperatingSystemKernelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	// path param operating_system_id
	if err := r.SetPathParam("operating_system_id", o.OperatingSystemID.String()); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

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
