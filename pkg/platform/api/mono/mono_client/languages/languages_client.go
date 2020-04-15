// Code generated by go-swagger; DO NOT EDIT.

package languages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new languages API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for languages API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddLanguage(params *AddLanguageParams, authInfo runtime.ClientAuthInfoWriter) (*AddLanguageOK, error)

	ListLanguages(params *ListLanguagesParams, authInfo runtime.ClientAuthInfoWriter) (*ListLanguagesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddLanguage creates a new builder language

  Create a new language
*/
func (a *Client) AddLanguage(params *AddLanguageParams, authInfo runtime.ClientAuthInfoWriter) (*AddLanguageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddLanguageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addLanguage",
		Method:             "POST",
		PathPattern:        "/languages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddLanguageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddLanguageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addLanguage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListLanguages lists of builder languages

  Retrieve a list of all valid builder languages
*/
func (a *Client) ListLanguages(params *ListLanguagesParams, authInfo runtime.ClientAuthInfoWriter) (*ListLanguagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListLanguagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listLanguages",
		Method:             "GET",
		PathPattern:        "/languages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListLanguagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListLanguagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listLanguages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
