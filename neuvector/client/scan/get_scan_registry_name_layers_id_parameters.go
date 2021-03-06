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

// NewGetScanRegistryNameLayersIDParams creates a new GetScanRegistryNameLayersIDParams object
// with the default values initialized.
func NewGetScanRegistryNameLayersIDParams() *GetScanRegistryNameLayersIDParams {
	var ()
	return &GetScanRegistryNameLayersIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScanRegistryNameLayersIDParamsWithTimeout creates a new GetScanRegistryNameLayersIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScanRegistryNameLayersIDParamsWithTimeout(timeout time.Duration) *GetScanRegistryNameLayersIDParams {
	var ()
	return &GetScanRegistryNameLayersIDParams{

		timeout: timeout,
	}
}

// NewGetScanRegistryNameLayersIDParamsWithContext creates a new GetScanRegistryNameLayersIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScanRegistryNameLayersIDParamsWithContext(ctx context.Context) *GetScanRegistryNameLayersIDParams {
	var ()
	return &GetScanRegistryNameLayersIDParams{

		Context: ctx,
	}
}

// NewGetScanRegistryNameLayersIDParamsWithHTTPClient creates a new GetScanRegistryNameLayersIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScanRegistryNameLayersIDParamsWithHTTPClient(client *http.Client) *GetScanRegistryNameLayersIDParams {
	var ()
	return &GetScanRegistryNameLayersIDParams{
		HTTPClient: client,
	}
}

/*GetScanRegistryNameLayersIDParams contains all the parameters to send to the API endpoint
for the get scan registry name layers ID operation typically these are written to a http.Request
*/
type GetScanRegistryNameLayersIDParams struct {

	/*XAuthToken*/
	XAuthToken string
	/*ID
	  Layer ID

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

// WithTimeout adds the timeout to the get scan registry name layers ID params
func (o *GetScanRegistryNameLayersIDParams) WithTimeout(timeout time.Duration) *GetScanRegistryNameLayersIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scan registry name layers ID params
func (o *GetScanRegistryNameLayersIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scan registry name layers ID params
func (o *GetScanRegistryNameLayersIDParams) WithContext(ctx context.Context) *GetScanRegistryNameLayersIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scan registry name layers ID params
func (o *GetScanRegistryNameLayersIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scan registry name layers ID params
func (o *GetScanRegistryNameLayersIDParams) WithHTTPClient(client *http.Client) *GetScanRegistryNameLayersIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scan registry name layers ID params
func (o *GetScanRegistryNameLayersIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get scan registry name layers ID params
func (o *GetScanRegistryNameLayersIDParams) WithXAuthToken(xAuthToken string) *GetScanRegistryNameLayersIDParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get scan registry name layers ID params
func (o *GetScanRegistryNameLayersIDParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithID adds the id to the get scan registry name layers ID params
func (o *GetScanRegistryNameLayersIDParams) WithID(id string) *GetScanRegistryNameLayersIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get scan registry name layers ID params
func (o *GetScanRegistryNameLayersIDParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the get scan registry name layers ID params
func (o *GetScanRegistryNameLayersIDParams) WithName(name string) *GetScanRegistryNameLayersIDParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get scan registry name layers ID params
func (o *GetScanRegistryNameLayersIDParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetScanRegistryNameLayersIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
