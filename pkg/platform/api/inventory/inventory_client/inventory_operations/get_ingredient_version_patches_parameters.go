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

// NewGetIngredientVersionPatchesParams creates a new GetIngredientVersionPatchesParams object
// with the default values initialized.
func NewGetIngredientVersionPatchesParams() *GetIngredientVersionPatchesParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetIngredientVersionPatchesParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIngredientVersionPatchesParamsWithTimeout creates a new GetIngredientVersionPatchesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIngredientVersionPatchesParamsWithTimeout(timeout time.Duration) *GetIngredientVersionPatchesParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetIngredientVersionPatchesParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,

		timeout: timeout,
	}
}

// NewGetIngredientVersionPatchesParamsWithContext creates a new GetIngredientVersionPatchesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIngredientVersionPatchesParamsWithContext(ctx context.Context) *GetIngredientVersionPatchesParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetIngredientVersionPatchesParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,

		Context: ctx,
	}
}

// NewGetIngredientVersionPatchesParamsWithHTTPClient creates a new GetIngredientVersionPatchesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIngredientVersionPatchesParamsWithHTTPClient(client *http.Client) *GetIngredientVersionPatchesParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetIngredientVersionPatchesParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,
		HTTPClient:    client,
	}
}

/*GetIngredientVersionPatchesParams contains all the parameters to send to the API endpoint
for the get ingredient version patches operation typically these are written to a http.Request
*/
type GetIngredientVersionPatchesParams struct {

	/*AllowUnstable
	  Whether to show an unstable revision of a resource if there is an available unstable version newer than the newest available stable version

	*/
	AllowUnstable *bool
	/*IngredientID*/
	IngredientID strfmt.UUID
	/*IngredientVersionID*/
	IngredientVersionID strfmt.UUID
	/*Limit
	  The maximum number of items returned per page

	*/
	Limit *int64
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

// WithTimeout adds the timeout to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) WithTimeout(timeout time.Duration) *GetIngredientVersionPatchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) WithContext(ctx context.Context) *GetIngredientVersionPatchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) WithHTTPClient(client *http.Client) *GetIngredientVersionPatchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowUnstable adds the allowUnstable to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) WithAllowUnstable(allowUnstable *bool) *GetIngredientVersionPatchesParams {
	o.SetAllowUnstable(allowUnstable)
	return o
}

// SetAllowUnstable adds the allowUnstable to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) SetAllowUnstable(allowUnstable *bool) {
	o.AllowUnstable = allowUnstable
}

// WithIngredientID adds the ingredientID to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) WithIngredientID(ingredientID strfmt.UUID) *GetIngredientVersionPatchesParams {
	o.SetIngredientID(ingredientID)
	return o
}

// SetIngredientID adds the ingredientId to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) SetIngredientID(ingredientID strfmt.UUID) {
	o.IngredientID = ingredientID
}

// WithIngredientVersionID adds the ingredientVersionID to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) WithIngredientVersionID(ingredientVersionID strfmt.UUID) *GetIngredientVersionPatchesParams {
	o.SetIngredientVersionID(ingredientVersionID)
	return o
}

// SetIngredientVersionID adds the ingredientVersionId to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) SetIngredientVersionID(ingredientVersionID strfmt.UUID) {
	o.IngredientVersionID = ingredientVersionID
}

// WithLimit adds the limit to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) WithLimit(limit *int64) *GetIngredientVersionPatchesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) WithPage(page *int64) *GetIngredientVersionPatchesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) SetPage(page *int64) {
	o.Page = page
}

// WithStateAt adds the stateAt to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) WithStateAt(stateAt *strfmt.DateTime) *GetIngredientVersionPatchesParams {
	o.SetStateAt(stateAt)
	return o
}

// SetStateAt adds the stateAt to the get ingredient version patches params
func (o *GetIngredientVersionPatchesParams) SetStateAt(stateAt *strfmt.DateTime) {
	o.StateAt = stateAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetIngredientVersionPatchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param ingredient_id
	if err := r.SetPathParam("ingredient_id", o.IngredientID.String()); err != nil {
		return err
	}

	// path param ingredient_version_id
	if err := r.SetPathParam("ingredient_version_id", o.IngredientVersionID.String()); err != nil {
		return err
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