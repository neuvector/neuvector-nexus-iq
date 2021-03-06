// Code generated by go-swagger; DO NOT EDIT.

package e_u_l_a

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

// NewGetEulaParams creates a new GetEulaParams object
// with the default values initialized.
func NewGetEulaParams() *GetEulaParams {
	var ()
	return &GetEulaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEulaParamsWithTimeout creates a new GetEulaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEulaParamsWithTimeout(timeout time.Duration) *GetEulaParams {
	var ()
	return &GetEulaParams{

		timeout: timeout,
	}
}

// NewGetEulaParamsWithContext creates a new GetEulaParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEulaParamsWithContext(ctx context.Context) *GetEulaParams {
	var ()
	return &GetEulaParams{

		Context: ctx,
	}
}

// NewGetEulaParamsWithHTTPClient creates a new GetEulaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEulaParamsWithHTTPClient(client *http.Client) *GetEulaParams {
	var ()
	return &GetEulaParams{
		HTTPClient: client,
	}
}

/*GetEulaParams contains all the parameters to send to the API endpoint
for the get eula operation typically these are written to a http.Request
*/
type GetEulaParams struct {

	/*XAuthToken*/
	XAuthToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get eula params
func (o *GetEulaParams) WithTimeout(timeout time.Duration) *GetEulaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get eula params
func (o *GetEulaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get eula params
func (o *GetEulaParams) WithContext(ctx context.Context) *GetEulaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get eula params
func (o *GetEulaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get eula params
func (o *GetEulaParams) WithHTTPClient(client *http.Client) *GetEulaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get eula params
func (o *GetEulaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get eula params
func (o *GetEulaParams) WithXAuthToken(xAuthToken string) *GetEulaParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get eula params
func (o *GetEulaParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetEulaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Auth-Token
	if err := r.SetHeaderParam("X-Auth-Token", o.XAuthToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
