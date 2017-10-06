// Code generated by go-swagger; DO NOT EDIT.

package connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// NewGetGatewaysCredentialIDParams creates a new GetGatewaysCredentialIDParams object
// with the default values initialized.
func NewGetGatewaysCredentialIDParams() *GetGatewaysCredentialIDParams {
	var ()
	return &GetGatewaysCredentialIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGatewaysCredentialIDParamsWithTimeout creates a new GetGatewaysCredentialIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGatewaysCredentialIDParamsWithTimeout(timeout time.Duration) *GetGatewaysCredentialIDParams {
	var ()
	return &GetGatewaysCredentialIDParams{

		timeout: timeout,
	}
}

// NewGetGatewaysCredentialIDParamsWithContext creates a new GetGatewaysCredentialIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGatewaysCredentialIDParamsWithContext(ctx context.Context) *GetGatewaysCredentialIDParams {
	var ()
	return &GetGatewaysCredentialIDParams{

		Context: ctx,
	}
}

// NewGetGatewaysCredentialIDParamsWithHTTPClient creates a new GetGatewaysCredentialIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGatewaysCredentialIDParamsWithHTTPClient(client *http.Client) *GetGatewaysCredentialIDParams {
	var ()
	return &GetGatewaysCredentialIDParams{
		HTTPClient: client,
	}
}

/*GetGatewaysCredentialIDParams contains all the parameters to send to the API endpoint
for the get gateways credential Id operation typically these are written to a http.Request
*/
type GetGatewaysCredentialIDParams struct {

	/*Body*/
	Body *models_cloudbreak.PlatformResourceRequestJSON

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gateways credential Id params
func (o *GetGatewaysCredentialIDParams) WithTimeout(timeout time.Duration) *GetGatewaysCredentialIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gateways credential Id params
func (o *GetGatewaysCredentialIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gateways credential Id params
func (o *GetGatewaysCredentialIDParams) WithContext(ctx context.Context) *GetGatewaysCredentialIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gateways credential Id params
func (o *GetGatewaysCredentialIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gateways credential Id params
func (o *GetGatewaysCredentialIDParams) WithHTTPClient(client *http.Client) *GetGatewaysCredentialIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gateways credential Id params
func (o *GetGatewaysCredentialIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get gateways credential Id params
func (o *GetGatewaysCredentialIDParams) WithBody(body *models_cloudbreak.PlatformResourceRequestJSON) *GetGatewaysCredentialIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get gateways credential Id params
func (o *GetGatewaysCredentialIDParams) SetBody(body *models_cloudbreak.PlatformResourceRequestJSON) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetGatewaysCredentialIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.PlatformResourceRequestJSON)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
