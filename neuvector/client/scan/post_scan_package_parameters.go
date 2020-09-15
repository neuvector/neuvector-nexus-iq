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

	"github.com/neuvector/neuvector-nexus-iq/neuvector/models"
)

// NewPostScanPackageParams creates a new PostScanPackageParams object
// with the default values initialized.
func NewPostScanPackageParams() *PostScanPackageParams {
	var ()
	return &PostScanPackageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostScanPackageParamsWithTimeout creates a new PostScanPackageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostScanPackageParamsWithTimeout(timeout time.Duration) *PostScanPackageParams {
	var ()
	return &PostScanPackageParams{

		timeout: timeout,
	}
}

// NewPostScanPackageParamsWithContext creates a new PostScanPackageParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostScanPackageParamsWithContext(ctx context.Context) *PostScanPackageParams {
	var ()
	return &PostScanPackageParams{

		Context: ctx,
	}
}

// NewPostScanPackageParamsWithHTTPClient creates a new PostScanPackageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostScanPackageParamsWithHTTPClient(client *http.Client) *PostScanPackageParams {
	var ()
	return &PostScanPackageParams{
		HTTPClient: client,
	}
}

/*PostScanPackageParams contains all the parameters to send to the API endpoint
for the post scan package operation typically these are written to a http.Request
*/
type PostScanPackageParams struct {

	/*XAuthToken*/
	XAuthToken string
	/*Body
	  Request package scan data

	*/
	Body *models.RESTScanPackageReqData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post scan package params
func (o *PostScanPackageParams) WithTimeout(timeout time.Duration) *PostScanPackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post scan package params
func (o *PostScanPackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post scan package params
func (o *PostScanPackageParams) WithContext(ctx context.Context) *PostScanPackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post scan package params
func (o *PostScanPackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post scan package params
func (o *PostScanPackageParams) WithHTTPClient(client *http.Client) *PostScanPackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post scan package params
func (o *PostScanPackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the post scan package params
func (o *PostScanPackageParams) WithXAuthToken(xAuthToken string) *PostScanPackageParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the post scan package params
func (o *PostScanPackageParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithBody adds the body to the post scan package params
func (o *PostScanPackageParams) WithBody(body *models.RESTScanPackageReqData) *PostScanPackageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post scan package params
func (o *PostScanPackageParams) SetBody(body *models.RESTScanPackageReqData) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostScanPackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
