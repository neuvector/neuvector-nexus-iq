// Code generated by go-swagger; DO NOT EDIT.

package scan

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

// NewGetScanRegistryNameImageIDParams creates a new GetScanRegistryNameImageIDParams object
// with the default values initialized.
func NewGetScanRegistryNameImageIDParams() *GetScanRegistryNameImageIDParams {
	var ()
	return &GetScanRegistryNameImageIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScanRegistryNameImageIDParamsWithTimeout creates a new GetScanRegistryNameImageIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScanRegistryNameImageIDParamsWithTimeout(timeout time.Duration) *GetScanRegistryNameImageIDParams {
	var ()
	return &GetScanRegistryNameImageIDParams{

		timeout: timeout,
	}
}

// NewGetScanRegistryNameImageIDParamsWithContext creates a new GetScanRegistryNameImageIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScanRegistryNameImageIDParamsWithContext(ctx context.Context) *GetScanRegistryNameImageIDParams {
	var ()
	return &GetScanRegistryNameImageIDParams{

		Context: ctx,
	}
}

// NewGetScanRegistryNameImageIDParamsWithHTTPClient creates a new GetScanRegistryNameImageIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScanRegistryNameImageIDParamsWithHTTPClient(client *http.Client) *GetScanRegistryNameImageIDParams {
	var ()
	return &GetScanRegistryNameImageIDParams{
		HTTPClient: client,
	}
}

/*GetScanRegistryNameImageIDParams contains all the parameters to send to the API endpoint
for the get scan registry name image ID operation typically these are written to a http.Request
*/
type GetScanRegistryNameImageIDParams struct {

	/*XAuthToken*/
	XAuthToken string
	/*ID
	  Image ID

	*/
	ID string
	/*Name
	  Name of the registry

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get scan registry name image ID params
func (o *GetScanRegistryNameImageIDParams) WithTimeout(timeout time.Duration) *GetScanRegistryNameImageIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scan registry name image ID params
func (o *GetScanRegistryNameImageIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scan registry name image ID params
func (o *GetScanRegistryNameImageIDParams) WithContext(ctx context.Context) *GetScanRegistryNameImageIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scan registry name image ID params
func (o *GetScanRegistryNameImageIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scan registry name image ID params
func (o *GetScanRegistryNameImageIDParams) WithHTTPClient(client *http.Client) *GetScanRegistryNameImageIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scan registry name image ID params
func (o *GetScanRegistryNameImageIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get scan registry name image ID params
func (o *GetScanRegistryNameImageIDParams) WithXAuthToken(xAuthToken string) *GetScanRegistryNameImageIDParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get scan registry name image ID params
func (o *GetScanRegistryNameImageIDParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithID adds the id to the get scan registry name image ID params
func (o *GetScanRegistryNameImageIDParams) WithID(id string) *GetScanRegistryNameImageIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get scan registry name image ID params
func (o *GetScanRegistryNameImageIDParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the get scan registry name image ID params
func (o *GetScanRegistryNameImageIDParams) WithName(name string) *GetScanRegistryNameImageIDParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get scan registry name image ID params
func (o *GetScanRegistryNameImageIDParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetScanRegistryNameImageIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Auth-Token
	if err := r.SetHeaderParam("X-Auth-Token", o.XAuthToken); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
