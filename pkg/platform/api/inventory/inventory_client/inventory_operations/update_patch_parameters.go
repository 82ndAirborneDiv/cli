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

// NewUpdatePatchParams creates a new UpdatePatchParams object
// with the default values initialized.
func NewUpdatePatchParams() *UpdatePatchParams {
	var ()
	return &UpdatePatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePatchParamsWithTimeout creates a new UpdatePatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePatchParamsWithTimeout(timeout time.Duration) *UpdatePatchParams {
	var ()
	return &UpdatePatchParams{

		timeout: timeout,
	}
}

// NewUpdatePatchParamsWithContext creates a new UpdatePatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePatchParamsWithContext(ctx context.Context) *UpdatePatchParams {
	var ()
	return &UpdatePatchParams{

		Context: ctx,
	}
}

// NewUpdatePatchParamsWithHTTPClient creates a new UpdatePatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePatchParamsWithHTTPClient(client *http.Client) *UpdatePatchParams {
	var ()
	return &UpdatePatchParams{
		HTTPClient: client,
	}
}

/*UpdatePatchParams contains all the parameters to send to the API endpoint
for the update patch operation typically these are written to a http.Request
*/
type UpdatePatchParams struct {

	/*Patch*/
	Patch *inventory_models.V1PatchCore
	/*PatchID*/
	PatchID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update patch params
func (o *UpdatePatchParams) WithTimeout(timeout time.Duration) *UpdatePatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update patch params
func (o *UpdatePatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update patch params
func (o *UpdatePatchParams) WithContext(ctx context.Context) *UpdatePatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update patch params
func (o *UpdatePatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update patch params
func (o *UpdatePatchParams) WithHTTPClient(client *http.Client) *UpdatePatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update patch params
func (o *UpdatePatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatch adds the patch to the update patch params
func (o *UpdatePatchParams) WithPatch(patch *inventory_models.V1PatchCore) *UpdatePatchParams {
	o.SetPatch(patch)
	return o
}

// SetPatch adds the patch to the update patch params
func (o *UpdatePatchParams) SetPatch(patch *inventory_models.V1PatchCore) {
	o.Patch = patch
}

// WithPatchID adds the patchID to the update patch params
func (o *UpdatePatchParams) WithPatchID(patchID strfmt.UUID) *UpdatePatchParams {
	o.SetPatchID(patchID)
	return o
}

// SetPatchID adds the patchId to the update patch params
func (o *UpdatePatchParams) SetPatchID(patchID strfmt.UUID) {
	o.PatchID = patchID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Patch != nil {
		if err := r.SetBodyParam(o.Patch); err != nil {
			return err
		}
	}

	// path param patch_id
	if err := r.SetPathParam("patch_id", o.PatchID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
