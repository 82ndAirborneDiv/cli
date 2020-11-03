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

// NewSearchIngredientsParams creates a new SearchIngredientsParams object
// with the default values initialized.
func NewSearchIngredientsParams() *SearchIngredientsParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		offsetDefault        = int64(0)
	)
	return &SearchIngredientsParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Offset:        &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchIngredientsParamsWithTimeout creates a new SearchIngredientsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchIngredientsParamsWithTimeout(timeout time.Duration) *SearchIngredientsParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		offsetDefault        = int64(0)
	)
	return &SearchIngredientsParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Offset:        &offsetDefault,

		timeout: timeout,
	}
}

// NewSearchIngredientsParamsWithContext creates a new SearchIngredientsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchIngredientsParamsWithContext(ctx context.Context) *SearchIngredientsParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		offsetDefault        = int64(0)
	)
	return &SearchIngredientsParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Offset:        &offsetDefault,

		Context: ctx,
	}
}

// NewSearchIngredientsParamsWithHTTPClient creates a new SearchIngredientsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchIngredientsParamsWithHTTPClient(client *http.Client) *SearchIngredientsParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		offsetDefault        = int64(0)
	)
	return &SearchIngredientsParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Offset:        &offsetDefault,
		HTTPClient:    client,
	}
}

/*SearchIngredientsParams contains all the parameters to send to the API endpoint
for the search ingredients operation typically these are written to a http.Request
*/
type SearchIngredientsParams struct {

	/*AllowUnstable
	  Whether to show an unstable revision of a resource if there is an available unstable version newer than the newest available stable version

	*/
	AllowUnstable *bool
	/*Limit
	  The maximum number of ingredients returned per page

	*/
	Limit *int64
	/*Namespaces*/
	Namespaces string
	/*Offset
	  The number of ingredients to skip

	*/
	Offset *int64
	/*Q
	  Return only ingredients whose names or features match the specified substring

	*/
	Q string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search ingredients params
func (o *SearchIngredientsParams) WithTimeout(timeout time.Duration) *SearchIngredientsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search ingredients params
func (o *SearchIngredientsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search ingredients params
func (o *SearchIngredientsParams) WithContext(ctx context.Context) *SearchIngredientsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search ingredients params
func (o *SearchIngredientsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search ingredients params
func (o *SearchIngredientsParams) WithHTTPClient(client *http.Client) *SearchIngredientsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search ingredients params
func (o *SearchIngredientsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowUnstable adds the allowUnstable to the search ingredients params
func (o *SearchIngredientsParams) WithAllowUnstable(allowUnstable *bool) *SearchIngredientsParams {
	o.SetAllowUnstable(allowUnstable)
	return o
}

// SetAllowUnstable adds the allowUnstable to the search ingredients params
func (o *SearchIngredientsParams) SetAllowUnstable(allowUnstable *bool) {
	o.AllowUnstable = allowUnstable
}

// WithLimit adds the limit to the search ingredients params
func (o *SearchIngredientsParams) WithLimit(limit *int64) *SearchIngredientsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search ingredients params
func (o *SearchIngredientsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNamespaces adds the namespaces to the search ingredients params
func (o *SearchIngredientsParams) WithNamespaces(namespaces string) *SearchIngredientsParams {
	o.SetNamespaces(namespaces)
	return o
}

// SetNamespaces adds the namespaces to the search ingredients params
func (o *SearchIngredientsParams) SetNamespaces(namespaces string) {
	o.Namespaces = namespaces
}

// WithOffset adds the offset to the search ingredients params
func (o *SearchIngredientsParams) WithOffset(offset *int64) *SearchIngredientsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search ingredients params
func (o *SearchIngredientsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the search ingredients params
func (o *SearchIngredientsParams) WithQ(q string) *SearchIngredientsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the search ingredients params
func (o *SearchIngredientsParams) SetQ(q string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *SearchIngredientsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param namespaces
	qrNamespaces := o.Namespaces
	qNamespaces := qrNamespaces
	if qNamespaces != "" {
		if err := r.SetQueryParam("namespaces", qNamespaces); err != nil {
			return err
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	// query param q
	qrQ := o.Q
	qQ := qrQ
	if qQ != "" {
		if err := r.SetQueryParam("q", qQ); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}