// Code generated by go-swagger; DO NOT EDIT.

package version_control

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

// NewGetRevertCommitParams creates a new GetRevertCommitParams object
// with the default values initialized.
func NewGetRevertCommitParams() *GetRevertCommitParams {
	var ()
	return &GetRevertCommitParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRevertCommitParamsWithTimeout creates a new GetRevertCommitParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRevertCommitParamsWithTimeout(timeout time.Duration) *GetRevertCommitParams {
	var ()
	return &GetRevertCommitParams{

		timeout: timeout,
	}
}

// NewGetRevertCommitParamsWithContext creates a new GetRevertCommitParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRevertCommitParamsWithContext(ctx context.Context) *GetRevertCommitParams {
	var ()
	return &GetRevertCommitParams{

		Context: ctx,
	}
}

// NewGetRevertCommitParamsWithHTTPClient creates a new GetRevertCommitParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRevertCommitParamsWithHTTPClient(client *http.Client) *GetRevertCommitParams {
	var ()
	return &GetRevertCommitParams{
		HTTPClient: client,
	}
}

/*GetRevertCommitParams contains all the parameters to send to the API endpoint
for the get revert commit operation typically these are written to a http.Request
*/
type GetRevertCommitParams struct {

	/*CommitFromID
	  The commit to start from (usually the latest commit)

	*/
	CommitFromID strfmt.UUID
	/*CommitToID
	  The commit you would like to mirror

	*/
	CommitToID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get revert commit params
func (o *GetRevertCommitParams) WithTimeout(timeout time.Duration) *GetRevertCommitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get revert commit params
func (o *GetRevertCommitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get revert commit params
func (o *GetRevertCommitParams) WithContext(ctx context.Context) *GetRevertCommitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get revert commit params
func (o *GetRevertCommitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get revert commit params
func (o *GetRevertCommitParams) WithHTTPClient(client *http.Client) *GetRevertCommitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get revert commit params
func (o *GetRevertCommitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommitFromID adds the commitFromID to the get revert commit params
func (o *GetRevertCommitParams) WithCommitFromID(commitFromID strfmt.UUID) *GetRevertCommitParams {
	o.SetCommitFromID(commitFromID)
	return o
}

// SetCommitFromID adds the commitFromId to the get revert commit params
func (o *GetRevertCommitParams) SetCommitFromID(commitFromID strfmt.UUID) {
	o.CommitFromID = commitFromID
}

// WithCommitToID adds the commitToID to the get revert commit params
func (o *GetRevertCommitParams) WithCommitToID(commitToID strfmt.UUID) *GetRevertCommitParams {
	o.SetCommitToID(commitToID)
	return o
}

// SetCommitToID adds the commitToId to the get revert commit params
func (o *GetRevertCommitParams) SetCommitToID(commitToID strfmt.UUID) {
	o.CommitToID = commitToID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRevertCommitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param commitFromID
	if err := r.SetPathParam("commitFromID", o.CommitFromID.String()); err != nil {
		return err
	}

	// path param commitToID
	if err := r.SetPathParam("commitToID", o.CommitToID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
