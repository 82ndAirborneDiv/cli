// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new authentication API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authentication API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetLoginJwtToken logins with a valid j w t and redirect to a platform URL

Login with a valid JWT and redirect to a platform URL
*/
func (a *Client) GetLoginJwtToken(params *GetLoginJwtTokenParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoginJwtTokenParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLoginJwtToken",
		Method:             "GET",
		PathPattern:        "/login/jwt/{token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLoginJwtTokenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetLogout renews a valid j w t

Log out of the current session
*/
func (a *Client) GetLogout(params *GetLogoutParams) (*GetLogoutNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLogoutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLogout",
		Method:             "GET",
		PathPattern:        "/logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLogoutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLogoutNoContent), nil

}

/*
GetRenew renews a valid j w t

Renew your current JWT to forestall expiration
*/
func (a *Client) GetRenew(params *GetRenewParams) (*GetRenewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRenewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRenew",
		Method:             "GET",
		PathPattern:        "/renew",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRenewReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRenewOK), nil

}

/*
PostLogin trades your credentials for a j w t

Supply either username/password OR token
*/
func (a *Client) PostLogin(params *PostLoginParams) (*PostLoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLoginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLogin",
		Method:             "POST",
		PathPattern:        "/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostLoginReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLoginOK), nil

}

/*
AddToken generates an API token for current user

Produces an API token for use with automated API clients
*/
func (a *Client) AddToken(params *AddTokenParams, authInfo runtime.ClientAuthInfoWriter) (*AddTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addToken",
		Method:             "POST",
		PathPattern:        "/apikeys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddTokenOK), nil

}

/*
ChangePassword changes the current password

Prompts for current password which is used to change it to something new.
*/
func (a *Client) ChangePassword(params *ChangePasswordParams, authInfo runtime.ClientAuthInfoWriter) (*ChangePasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangePasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changePassword",
		Method:             "POST",
		PathPattern:        "/change-password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangePasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangePasswordOK), nil

}

/*
DeleteToken deletes an API token

Deletes the specified API Token
*/
func (a *Client) DeleteToken(params *DeleteTokenParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteToken",
		Method:             "DELETE",
		PathPattern:        "/apikeys/{tokenID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteTokenOK), nil

}

/*
DisableTOTP disables t o t p

Disable TOTP authentication
*/
func (a *Client) DisableTOTP(params *DisableTOTPParams, authInfo runtime.ClientAuthInfoWriter) (*DisableTOTPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisableTOTPParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "disableTOTP",
		Method:             "DELETE",
		PathPattern:        "/totp",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisableTOTPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DisableTOTPOK), nil

}

/*
EnableTOTP enables t o t p

Enable TOTP authentication by performing initial code validation
*/
func (a *Client) EnableTOTP(params *EnableTOTPParams, authInfo runtime.ClientAuthInfoWriter) (*EnableTOTPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnableTOTPParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "enableTOTP",
		Method:             "POST",
		PathPattern:        "/totp",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EnableTOTPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EnableTOTPOK), nil

}

/*
ListTokens lists current user s API tokens

List of all active API Tokens for current user
*/
func (a *Client) ListTokens(params *ListTokensParams, authInfo runtime.ClientAuthInfoWriter) (*ListTokensOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTokensParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listTokens",
		Method:             "GET",
		PathPattern:        "/apikeys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListTokensReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListTokensOK), nil

}

/*
LoginAs logins as given user requires you to be a superuser
*/
func (a *Client) LoginAs(params *LoginAsParams, authInfo runtime.ClientAuthInfoWriter) (*LoginAsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginAsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "loginAs",
		Method:             "POST",
		PathPattern:        "/login/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LoginAsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LoginAsOK), nil

}

/*
LoginWithGithub callbacks endpoint for github auth
*/
func (a *Client) LoginWithGithub(params *LoginWithGithubParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginWithGithubParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "loginWithGithub",
		Method:             "GET",
		PathPattern:        "/githubLogin",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LoginWithGithubReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
NewTOTP sets up a new t o t p key

Establish the private key for two-factor authentication
*/
func (a *Client) NewTOTP(params *NewTOTPParams, authInfo runtime.ClientAuthInfoWriter) (*NewTOTPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNewTOTPParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "newTOTP",
		Method:             "GET",
		PathPattern:        "/totp",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NewTOTPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NewTOTPOK), nil

}

/*
RequestReset requests a password recovery email

Sends a link which can be used to reset a forgotten password.
*/
func (a *Client) RequestReset(params *RequestResetParams) (*RequestResetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequestResetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "requestReset",
		Method:             "POST",
		PathPattern:        "/request-reset/{email}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RequestResetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RequestResetOK), nil

}

/*
ResetPassword resets a forgotten password

Sends a link which can be used to reset a forgotten password.
*/
func (a *Client) ResetPassword(params *ResetPasswordParams) (*ResetPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetPasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "resetPassword",
		Method:             "POST",
		PathPattern:        "/reset-password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ResetPasswordReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ResetPasswordOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
