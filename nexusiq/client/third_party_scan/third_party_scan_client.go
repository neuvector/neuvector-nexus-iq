// Code generated by go-swagger; DO NOT EDIT.

package third_party_scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new third party scan API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for third party scan API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetScanStatus(params *GetScanStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetScanStatusOK, error)

	PostScan(params *PostScanParams, authInfo runtime.ClientAuthInfoWriter) (*PostScanAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetScanStatus gets scan status

  Get the status of a scan report submit

*/
func (a *Client) GetScanStatus(params *GetScanStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetScanStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScanStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScanStatus",
		Method:             "GET",
		PathPattern:        "/scan/applications/{applicationInternalId}/status/{statusId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetScanStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetScanStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getScanStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostScan submits a scan report

  Submit a scan report

*/
func (a *Client) PostScan(params *PostScanParams, authInfo runtime.ClientAuthInfoWriter) (*PostScanAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostScanParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postScan",
		Method:             "POST",
		PathPattern:        "/scan/applications/{applicationInternalId}/sources/{source}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostScanReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostScanAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postScan: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
