// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStopSdxByCrnParams creates a new StopSdxByCrnParams object
// with the default values initialized.
func NewStopSdxByCrnParams() *StopSdxByCrnParams {
	var ()
	return &StopSdxByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopSdxByCrnParamsWithTimeout creates a new StopSdxByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopSdxByCrnParamsWithTimeout(timeout time.Duration) *StopSdxByCrnParams {
	var ()
	return &StopSdxByCrnParams{

		timeout: timeout,
	}
}

// NewStopSdxByCrnParamsWithContext creates a new StopSdxByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopSdxByCrnParamsWithContext(ctx context.Context) *StopSdxByCrnParams {
	var ()
	return &StopSdxByCrnParams{

		Context: ctx,
	}
}

// NewStopSdxByCrnParamsWithHTTPClient creates a new StopSdxByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopSdxByCrnParamsWithHTTPClient(client *http.Client) *StopSdxByCrnParams {
	var ()
	return &StopSdxByCrnParams{
		HTTPClient: client,
	}
}

/*StopSdxByCrnParams contains all the parameters to send to the API endpoint
for the stop sdx by crn operation typically these are written to a http.Request
*/
type StopSdxByCrnParams struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop sdx by crn params
func (o *StopSdxByCrnParams) WithTimeout(timeout time.Duration) *StopSdxByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop sdx by crn params
func (o *StopSdxByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop sdx by crn params
func (o *StopSdxByCrnParams) WithContext(ctx context.Context) *StopSdxByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop sdx by crn params
func (o *StopSdxByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop sdx by crn params
func (o *StopSdxByCrnParams) WithHTTPClient(client *http.Client) *StopSdxByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop sdx by crn params
func (o *StopSdxByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the stop sdx by crn params
func (o *StopSdxByCrnParams) WithCrn(crn string) *StopSdxByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the stop sdx by crn params
func (o *StopSdxByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *StopSdxByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}