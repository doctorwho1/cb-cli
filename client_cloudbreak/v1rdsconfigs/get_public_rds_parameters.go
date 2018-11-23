// Code generated by go-swagger; DO NOT EDIT.

package v1rdsconfigs

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
)

// NewGetPublicRdsParams creates a new GetPublicRdsParams object
// with the default values initialized.
func NewGetPublicRdsParams() *GetPublicRdsParams {
	var ()
	return &GetPublicRdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicRdsParamsWithTimeout creates a new GetPublicRdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicRdsParamsWithTimeout(timeout time.Duration) *GetPublicRdsParams {
	var ()
	return &GetPublicRdsParams{

		timeout: timeout,
	}
}

// NewGetPublicRdsParamsWithContext creates a new GetPublicRdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicRdsParamsWithContext(ctx context.Context) *GetPublicRdsParams {
	var ()
	return &GetPublicRdsParams{

		Context: ctx,
	}
}

// NewGetPublicRdsParamsWithHTTPClient creates a new GetPublicRdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicRdsParamsWithHTTPClient(client *http.Client) *GetPublicRdsParams {
	var ()
	return &GetPublicRdsParams{
		HTTPClient: client,
	}
}

/*GetPublicRdsParams contains all the parameters to send to the API endpoint
for the get public rds operation typically these are written to a http.Request
*/
type GetPublicRdsParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public rds params
func (o *GetPublicRdsParams) WithTimeout(timeout time.Duration) *GetPublicRdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public rds params
func (o *GetPublicRdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public rds params
func (o *GetPublicRdsParams) WithContext(ctx context.Context) *GetPublicRdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public rds params
func (o *GetPublicRdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public rds params
func (o *GetPublicRdsParams) WithHTTPClient(client *http.Client) *GetPublicRdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public rds params
func (o *GetPublicRdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get public rds params
func (o *GetPublicRdsParams) WithName(name string) *GetPublicRdsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get public rds params
func (o *GetPublicRdsParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicRdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}