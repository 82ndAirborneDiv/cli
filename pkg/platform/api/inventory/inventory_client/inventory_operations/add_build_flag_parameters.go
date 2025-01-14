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

// NewAddBuildFlagParams creates a new AddBuildFlagParams object
// with the default values initialized.
func NewAddBuildFlagParams() *AddBuildFlagParams {
	var ()
	return &AddBuildFlagParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddBuildFlagParamsWithTimeout creates a new AddBuildFlagParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddBuildFlagParamsWithTimeout(timeout time.Duration) *AddBuildFlagParams {
	var ()
	return &AddBuildFlagParams{

		timeout: timeout,
	}
}

// NewAddBuildFlagParamsWithContext creates a new AddBuildFlagParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddBuildFlagParamsWithContext(ctx context.Context) *AddBuildFlagParams {
	var ()
	return &AddBuildFlagParams{

		Context: ctx,
	}
}

// NewAddBuildFlagParamsWithHTTPClient creates a new AddBuildFlagParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddBuildFlagParamsWithHTTPClient(client *http.Client) *AddBuildFlagParams {
	var ()
	return &AddBuildFlagParams{
		HTTPClient: client,
	}
}

/*AddBuildFlagParams contains all the parameters to send to the API endpoint
for the add build flag operation typically these are written to a http.Request
*/
type AddBuildFlagParams struct {

	/*BuildFlag*/
	BuildFlag *inventory_models.BuildFlagCore

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add build flag params
func (o *AddBuildFlagParams) WithTimeout(timeout time.Duration) *AddBuildFlagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add build flag params
func (o *AddBuildFlagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add build flag params
func (o *AddBuildFlagParams) WithContext(ctx context.Context) *AddBuildFlagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add build flag params
func (o *AddBuildFlagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add build flag params
func (o *AddBuildFlagParams) WithHTTPClient(client *http.Client) *AddBuildFlagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add build flag params
func (o *AddBuildFlagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildFlag adds the buildFlag to the add build flag params
func (o *AddBuildFlagParams) WithBuildFlag(buildFlag *inventory_models.BuildFlagCore) *AddBuildFlagParams {
	o.SetBuildFlag(buildFlag)
	return o
}

// SetBuildFlag adds the buildFlag to the add build flag params
func (o *AddBuildFlagParams) SetBuildFlag(buildFlag *inventory_models.BuildFlagCore) {
	o.BuildFlag = buildFlag
}

// WriteToRequest writes these params to a swagger request
func (o *AddBuildFlagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BuildFlag != nil {
		if err := r.SetBodyParam(o.BuildFlag); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
