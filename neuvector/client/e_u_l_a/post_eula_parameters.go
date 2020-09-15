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

	"github.com/neuvector/neuvector-nexus-iq/neuvector/models"
)

// NewPostEulaParams creates a new PostEulaParams object
// with the default values initialized.
func NewPostEulaParams() *PostEulaParams {
	var ()
	return &PostEulaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostEulaParamsWithTimeout creates a new PostEulaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostEulaParamsWithTimeout(timeout time.Duration) *PostEulaParams {
	var ()
	return &PostEulaParams{

		timeout: timeout,
	}
}

// NewPostEulaParamsWithContext creates a new PostEulaParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostEulaParamsWithContext(ctx context.Context) *PostEulaParams {
	var ()
	return &PostEulaParams{

		Context: ctx,
	}
}

// NewPostEulaParamsWithHTTPClient creates a new PostEulaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostEulaParamsWithHTTPClient(client *http.Client) *PostEulaParams {
	var ()
	return &PostEulaParams{
		HTTPClient: client,
	}
}

/*PostEulaParams contains all the parameters to send to the API endpoint
for the post eula operation typically these are written to a http.Request
*/
type PostEulaParams struct {

	/*XAuthToken*/
	XAuthToken string
	/*Body
	  EULA data

	*/
	Body *models.RESTEULAData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post eula params
func (o *PostEulaParams) WithTimeout(timeout time.Duration) *PostEulaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post eula params
func (o *PostEulaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post eula params
func (o *PostEulaParams) WithContext(ctx context.Context) *PostEulaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post eula params
func (o *PostEulaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post eula params
func (o *PostEulaParams) WithHTTPClient(client *http.Client) *PostEulaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post eula params
func (o *PostEulaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the post eula params
func (o *PostEulaParams) WithXAuthToken(xAuthToken string) *PostEulaParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the post eula params
func (o *PostEulaParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithBody adds the body to the post eula params
func (o *PostEulaParams) WithBody(body *models.RESTEULAData) *PostEulaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post eula params
func (o *PostEulaParams) SetBody(body *models.RESTEULAData) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostEulaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
