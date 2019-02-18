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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ActiveState/cli/pkg/platform/api/models"
)

// NewAddCommitParams creates a new AddCommitParams object
// with the default values initialized.
func NewAddCommitParams() *AddCommitParams {
	var ()
	return &AddCommitParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddCommitParamsWithTimeout creates a new AddCommitParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddCommitParamsWithTimeout(timeout time.Duration) *AddCommitParams {
	var ()
	return &AddCommitParams{

		timeout: timeout,
	}
}

// NewAddCommitParamsWithContext creates a new AddCommitParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddCommitParamsWithContext(ctx context.Context) *AddCommitParams {
	var ()
	return &AddCommitParams{

		Context: ctx,
	}
}

// NewAddCommitParamsWithHTTPClient creates a new AddCommitParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddCommitParamsWithHTTPClient(client *http.Client) *AddCommitParams {
	var ()
	return &AddCommitParams{
		HTTPClient: client,
	}
}

/*AddCommitParams contains all the parameters to send to the API endpoint
for the add commit operation typically these are written to a http.Request
*/
type AddCommitParams struct {

	/*Commit*/
	Commit *models.CommitEditable

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add commit params
func (o *AddCommitParams) WithTimeout(timeout time.Duration) *AddCommitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add commit params
func (o *AddCommitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add commit params
func (o *AddCommitParams) WithContext(ctx context.Context) *AddCommitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add commit params
func (o *AddCommitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add commit params
func (o *AddCommitParams) WithHTTPClient(client *http.Client) *AddCommitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add commit params
func (o *AddCommitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommit adds the commit to the add commit params
func (o *AddCommitParams) WithCommit(commit *models.CommitEditable) *AddCommitParams {
	o.SetCommit(commit)
	return o
}

// SetCommit adds the commit to the add commit params
func (o *AddCommitParams) SetCommit(commit *models.CommitEditable) {
	o.Commit = commit
}

// WriteToRequest writes these params to a swagger request
func (o *AddCommitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Commit != nil {
		if err := r.SetBodyParam(o.Commit); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
