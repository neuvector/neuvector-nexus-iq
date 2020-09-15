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

	"github.com/neuvector/neuvector-nexus-iq/neuvector/models"
)

// NewPatchSystemConfigParams creates a new PatchSystemConfigParams object
// with the default values initialized.
func NewPatchSystemConfigParams() *PatchSystemConfigParams {
	var ()
	return &PatchSystemConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchSystemConfigParamsWithTimeout creates a new PatchSystemConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchSystemConfigParamsWithTimeout(timeout time.Duration) *PatchSystemConfigParams {
	var ()
	return &PatchSystemConfigParams{

		timeout: timeout,
	}
}

// NewPatchSystemConfigParamsWithContext creates a new PatchSystemConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchSystemConfigParamsWithContext(ctx context.Context) *PatchSystemConfigParams {
	var ()
	return &PatchSystemConfigParams{

		Context: ctx,
	}
}

// NewPatchSystemConfigParamsWithHTTPClient creates a new PatchSystemConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchSystemConfigParamsWithHTTPClient(client *http.Client) *PatchSystemConfigParams {
	var ()
	return &PatchSystemConfigParams{
		HTTPClient: client,
	}
}

/*PatchSystemConfigParams contains all the parameters to send to the API endpoint
for the patch system config operation typically these are written to a http.Request
*/
type PatchSystemConfigParams struct {

	/*XAuthToken*/
	XAuthToken string
	/*Body
	  System configure data

	*/
	Body *models.RESTSystemConfigConfigData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch system config params
func (o *PatchSystemConfigParams) WithTimeout(timeout time.Duration) *PatchSystemConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch system config params
func (o *PatchSystemConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch system config params
func (o *PatchSystemConfigParams) WithContext(ctx context.Context) *PatchSystemConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch system config params
func (o *PatchSystemConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch system config params
func (o *PatchSystemConfigParams) WithHTTPClient(client *http.Client) *PatchSystemConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch system config params
func (o *PatchSystemConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the patch system config params
func (o *PatchSystemConfigParams) WithXAuthToken(xAuthToken string) *PatchSystemConfigParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the patch system config params
func (o *PatchSystemConfigParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithBody adds the body to the patch system config params
func (o *PatchSystemConfigParams) WithBody(body *models.RESTSystemConfigConfigData) *PatchSystemConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch system config params
func (o *PatchSystemConfigParams) SetBody(body *models.RESTSystemConfigConfigData) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchSystemConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Auth-Token
	if err := r.SetHeaderParam("X-Auth-Token", o.XAuthToken); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
