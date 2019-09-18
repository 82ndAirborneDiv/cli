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

// NewGetNamespacesParams creates a new GetNamespacesParams object
// with the default values initialized.
func NewGetNamespacesParams() *GetNamespacesParams {
	var (
		limitDefault = int64(50)
		pageDefault  = int64(1)
	)
	return &GetNamespacesParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNamespacesParamsWithTimeout creates a new GetNamespacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNamespacesParamsWithTimeout(timeout time.Duration) *GetNamespacesParams {
	var (
		limitDefault = int64(50)
		pageDefault  = int64(1)
	)
	return &GetNamespacesParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: timeout,
	}
}

// NewGetNamespacesParamsWithContext creates a new GetNamespacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNamespacesParamsWithContext(ctx context.Context) *GetNamespacesParams {
	var (
		limitDefault = int64(50)
		pageDefault  = int64(1)
	)
	return &GetNamespacesParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		Context: ctx,
	}
}

// NewGetNamespacesParamsWithHTTPClient creates a new GetNamespacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNamespacesParamsWithHTTPClient(client *http.Client) *GetNamespacesParams {
	var (
		limitDefault = int64(50)
		pageDefault  = int64(1)
	)
	return &GetNamespacesParams{
		Limit:      &limitDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*GetNamespacesParams contains all the parameters to send to the API endpoint
for the get namespaces operation typically these are written to a http.Request
*/
type GetNamespacesParams struct {

	/*Limit
	  The maximum number of items returned per page

	*/
	Limit *int64
	/*Page
	  The page number returned

	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get namespaces params
func (o *GetNamespacesParams) WithTimeout(timeout time.Duration) *GetNamespacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get namespaces params
func (o *GetNamespacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get namespaces params
func (o *GetNamespacesParams) WithContext(ctx context.Context) *GetNamespacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get namespaces params
func (o *GetNamespacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get namespaces params
func (o *GetNamespacesParams) WithHTTPClient(client *http.Client) *GetNamespacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get namespaces params
func (o *GetNamespacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get namespaces params
func (o *GetNamespacesParams) WithLimit(limit *int64) *GetNamespacesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get namespaces params
func (o *GetNamespacesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get namespaces params
func (o *GetNamespacesParams) WithPage(page *int64) *GetNamespacesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get namespaces params
func (o *GetNamespacesParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetNamespacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
