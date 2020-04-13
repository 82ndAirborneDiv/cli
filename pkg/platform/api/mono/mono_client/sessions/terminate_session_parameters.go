// Code generated by go-swagger; DO NOT EDIT.

package sessions

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

// NewTerminateSessionParams creates a new TerminateSessionParams object
// with the default values initialized.
func NewTerminateSessionParams() *TerminateSessionParams {
	var ()
	return &TerminateSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTerminateSessionParamsWithTimeout creates a new TerminateSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTerminateSessionParamsWithTimeout(timeout time.Duration) *TerminateSessionParams {
	var ()
	return &TerminateSessionParams{

		timeout: timeout,
	}
}

// NewTerminateSessionParamsWithContext creates a new TerminateSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewTerminateSessionParamsWithContext(ctx context.Context) *TerminateSessionParams {
	var ()
	return &TerminateSessionParams{

		Context: ctx,
	}
}

// NewTerminateSessionParamsWithHTTPClient creates a new TerminateSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTerminateSessionParamsWithHTTPClient(client *http.Client) *TerminateSessionParams {
	var ()
	return &TerminateSessionParams{
		HTTPClient: client,
	}
}

/*TerminateSessionParams contains all the parameters to send to the API endpoint
for the terminate session operation typically these are written to a http.Request
*/
type TerminateSessionParams struct {

	/*SessionID
	  Unique ID of session

	*/
	SessionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the terminate session params
func (o *TerminateSessionParams) WithTimeout(timeout time.Duration) *TerminateSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the terminate session params
func (o *TerminateSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the terminate session params
func (o *TerminateSessionParams) WithContext(ctx context.Context) *TerminateSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the terminate session params
func (o *TerminateSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the terminate session params
func (o *TerminateSessionParams) WithHTTPClient(client *http.Client) *TerminateSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the terminate session params
func (o *TerminateSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSessionID adds the sessionID to the terminate session params
func (o *TerminateSessionParams) WithSessionID(sessionID strfmt.UUID) *TerminateSessionParams {
	o.SetSessionID(sessionID)
	return o
}

// SetSessionID adds the sessionId to the terminate session params
func (o *TerminateSessionParams) SetSessionID(sessionID strfmt.UUID) {
	o.SessionID = sessionID
}

// WriteToRequest writes these params to a swagger request
func (o *TerminateSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sessionID
	if err := r.SetPathParam("sessionID", o.SessionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
