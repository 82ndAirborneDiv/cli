// Code generated by go-swagger; DO NOT EDIT.

package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new invoices API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for invoices API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CalculateTax calculates the tax for the given address and options
*/
func (a *Client) CalculateTax(params *CalculateTaxParams, authInfo runtime.ClientAuthInfoWriter) (*CalculateTaxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCalculateTaxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "calculateTax",
		Method:             "POST",
		PathPattern:        "/taxes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CalculateTaxReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CalculateTaxOK), nil

}

/*
CreateInvoice creates new invoice

Creates a new invoice for the organization
*/
func (a *Client) CreateInvoice(params *CreateInvoiceParams, authInfo runtime.ClientAuthInfoWriter) (*CreateInvoiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInvoiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createInvoice",
		Method:             "POST",
		PathPattern:        "/organizations/{organizationIdentifier}/invoices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateInvoiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateInvoiceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
