// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewGetSystemSummaryParams creates a new GetSystemSummaryParams object
// with the default values initialized.
func NewGetSystemSummaryParams() *GetSystemSummaryParams {
	var ()
	return &GetSystemSummaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemSummaryParamsWithTimeout creates a new GetSystemSummaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSystemSummaryParamsWithTimeout(timeout time.Duration) *GetSystemSummaryParams {
	var ()
	return &GetSystemSummaryParams{

		timeout: timeout,
	}
}

// NewGetSystemSummaryParamsWithContext creates a new GetSystemSummaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSystemSummaryParamsWithContext(ctx context.Context) *GetSystemSummaryParams {
	var ()
	return &GetSystemSummaryParams{

		Context: ctx,
	}
}

// NewGetSystemSummaryParamsWithHTTPClient creates a new GetSystemSummaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSystemSummaryParamsWithHTTPClient(client *http.Client) *GetSystemSummaryParams {
	var ()
	return &GetSystemSummaryParams{
		HTTPClient: client,
	}
}

/*GetSystemSummaryParams contains all the parameters to send to the API endpoint
for the get system summary operation typically these are written to a http.Request
*/
type GetSystemSummaryParams struct {

	/*XAuthToken*/
	XAuthToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get system summary params
func (o *GetSystemSummaryParams) WithTimeout(timeout time.Duration) *GetSystemSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system summary params
func (o *GetSystemSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system summary params
func (o *GetSystemSummaryParams) WithContext(ctx context.Context) *GetSystemSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system summary params
func (o *GetSystemSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system summary params
func (o *GetSystemSummaryParams) WithHTTPClient(client *http.Client) *GetSystemSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system summary params
func (o *GetSystemSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get system summary params
func (o *GetSystemSummaryParams) WithXAuthToken(xAuthToken string) *GetSystemSummaryParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get system summary params
func (o *GetSystemSummaryParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
