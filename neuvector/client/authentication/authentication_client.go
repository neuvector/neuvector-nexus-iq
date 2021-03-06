// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new authentication API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authentication API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAuth(params *DeleteAuthParams) (*DeleteAuthOK, error)

	PatchAuth(params *PatchAuthParams) (*PatchAuthOK, error)

	PostAuth(params *PostAuthParams) (*PostAuthOK, error)

	PostAuthServer(params *PostAuthServerParams) (*PostAuthServerOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteAuth logouts current logged in user
*/
func (a *Client) DeleteAuth(params *DeleteAuthParams) (*DeleteAuthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAuthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAuth",
		Method:             "DELETE",
		PathPattern:        "/auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAuthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAuthOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAuth: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchAuth keeps login session alive
*/
func (a *Client) PatchAuth(params *PatchAuthParams) (*PatchAuthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAuthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchAuth",
		Method:             "PATCH",
		PathPattern:        "/auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchAuthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchAuthOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAuth: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAuth logins authentication
*/
func (a *Client) PostAuth(params *PostAuthParams) (*PostAuthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAuthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAuth",
		Method:             "POST",
		PathPattern:        "/auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAuthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAuthOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAuth: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAuthServer authenticates to specified server
*/
func (a *Client) PostAuthServer(params *PostAuthServerParams) (*PostAuthServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAuthServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAuthServer",
		Method:             "POST",
		PathPattern:        "/auth/{server}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAuthServerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAuthServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAuthServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
