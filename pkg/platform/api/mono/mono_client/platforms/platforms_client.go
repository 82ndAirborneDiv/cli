// Code generated by go-swagger; DO NOT EDIT.

package platforms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new platforms API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for platforms API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddPlatform(params *AddPlatformParams, authInfo runtime.ClientAuthInfoWriter) (*AddPlatformOK, error)

	ListPlatforms(params *ListPlatformsParams, authInfo runtime.ClientAuthInfoWriter) (*ListPlatformsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddPlatform creates a new builder platform

  Create a new platform
*/
func (a *Client) AddPlatform(params *AddPlatformParams, authInfo runtime.ClientAuthInfoWriter) (*AddPlatformOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddPlatformParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addPlatform",
		Method:             "POST",
		PathPattern:        "/platforms",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddPlatformReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddPlatformOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addPlatform: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListPlatforms lists of builder platforms

  Retrieve a list of all valid builder platforms
*/
func (a *Client) ListPlatforms(params *ListPlatformsParams, authInfo runtime.ClientAuthInfoWriter) (*ListPlatformsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPlatformsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listPlatforms",
		Method:             "GET",
		PathPattern:        "/platforms",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListPlatformsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPlatformsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listPlatforms: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
