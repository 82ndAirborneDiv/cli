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

	strfmt "github.com/go-openapi/strfmt"

	inventory_models "github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// NewAddOperatingSystemVersionRevisionParams creates a new AddOperatingSystemVersionRevisionParams object
// with the default values initialized.
func NewAddOperatingSystemVersionRevisionParams() *AddOperatingSystemVersionRevisionParams {
	var ()
	return &AddOperatingSystemVersionRevisionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddOperatingSystemVersionRevisionParamsWithTimeout creates a new AddOperatingSystemVersionRevisionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddOperatingSystemVersionRevisionParamsWithTimeout(timeout time.Duration) *AddOperatingSystemVersionRevisionParams {
	var ()
	return &AddOperatingSystemVersionRevisionParams{

		timeout: timeout,
	}
}

// NewAddOperatingSystemVersionRevisionParamsWithContext creates a new AddOperatingSystemVersionRevisionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddOperatingSystemVersionRevisionParamsWithContext(ctx context.Context) *AddOperatingSystemVersionRevisionParams {
	var ()
	return &AddOperatingSystemVersionRevisionParams{

		Context: ctx,
	}
}

// NewAddOperatingSystemVersionRevisionParamsWithHTTPClient creates a new AddOperatingSystemVersionRevisionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddOperatingSystemVersionRevisionParamsWithHTTPClient(client *http.Client) *AddOperatingSystemVersionRevisionParams {
	var ()
	return &AddOperatingSystemVersionRevisionParams{
		HTTPClient: client,
	}
}

/*AddOperatingSystemVersionRevisionParams contains all the parameters to send to the API endpoint
for the add operating system version revision operation typically these are written to a http.Request
*/
type AddOperatingSystemVersionRevisionParams struct {

	/*OperatingSystemID*/
	OperatingSystemID strfmt.UUID
	/*OperatingSystemVersionID*/
	OperatingSystemVersionID strfmt.UUID
	/*OperatingSystemVersionRevision*/
	OperatingSystemVersionRevision *inventory_models.V1Revision

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add operating system version revision params
func (o *AddOperatingSystemVersionRevisionParams) WithTimeout(timeout time.Duration) *AddOperatingSystemVersionRevisionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add operating system version revision params
func (o *AddOperatingSystemVersionRevisionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add operating system version revision params
func (o *AddOperatingSystemVersionRevisionParams) WithContext(ctx context.Context) *AddOperatingSystemVersionRevisionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add operating system version revision params
func (o *AddOperatingSystemVersionRevisionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add operating system version revision params
func (o *AddOperatingSystemVersionRevisionParams) WithHTTPClient(client *http.Client) *AddOperatingSystemVersionRevisionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add operating system version revision params
func (o *AddOperatingSystemVersionRevisionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOperatingSystemID adds the operatingSystemID to the add operating system version revision params
func (o *AddOperatingSystemVersionRevisionParams) WithOperatingSystemID(operatingSystemID strfmt.UUID) *AddOperatingSystemVersionRevisionParams {
	o.SetOperatingSystemID(operatingSystemID)
	return o
}

// SetOperatingSystemID adds the operatingSystemId to the add operating system version revision params
func (o *AddOperatingSystemVersionRevisionParams) SetOperatingSystemID(operatingSystemID strfmt.UUID) {
	o.OperatingSystemID = operatingSystemID
}

// WithOperatingSystemVersionID adds the operatingSystemVersionID to the add operating system version revision params
func (o *AddOperatingSystemVersionRevisionParams) WithOperatingSystemVersionID(operatingSystemVersionID strfmt.UUID) *AddOperatingSystemVersionRevisionParams {
	o.SetOperatingSystemVersionID(operatingSystemVersionID)
	return o
}

// SetOperatingSystemVersionID adds the operatingSystemVersionId to the add operating system version revision params
func (o *AddOperatingSystemVersionRevisionParams) SetOperatingSystemVersionID(operatingSystemVersionID strfmt.UUID) {
	o.OperatingSystemVersionID = operatingSystemVersionID
}

// WithOperatingSystemVersionRevision adds the operatingSystemVersionRevision to the add operating system version revision params
func (o *AddOperatingSystemVersionRevisionParams) WithOperatingSystemVersionRevision(operatingSystemVersionRevision *inventory_models.V1Revision) *AddOperatingSystemVersionRevisionParams {
	o.SetOperatingSystemVersionRevision(operatingSystemVersionRevision)
	return o
}

// SetOperatingSystemVersionRevision adds the operatingSystemVersionRevision to the add operating system version revision params
func (o *AddOperatingSystemVersionRevisionParams) SetOperatingSystemVersionRevision(operatingSystemVersionRevision *inventory_models.V1Revision) {
	o.OperatingSystemVersionRevision = operatingSystemVersionRevision
}

// WriteToRequest writes these params to a swagger request
func (o *AddOperatingSystemVersionRevisionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param operating_system_id
	if err := r.SetPathParam("operating_system_id", o.OperatingSystemID.String()); err != nil {
		return err
	}

	// path param operating_system_version_id
	if err := r.SetPathParam("operating_system_version_id", o.OperatingSystemVersionID.String()); err != nil {
		return err
	}

	if o.OperatingSystemVersionRevision != nil {
		if err := r.SetBodyParam(o.OperatingSystemVersionRevision); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
