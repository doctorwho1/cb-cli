// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

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

// NewListRetryableFlowsDistroXV1Params creates a new ListRetryableFlowsDistroXV1Params object
// with the default values initialized.
func NewListRetryableFlowsDistroXV1Params() *ListRetryableFlowsDistroXV1Params {
	var ()
	return &ListRetryableFlowsDistroXV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListRetryableFlowsDistroXV1ParamsWithTimeout creates a new ListRetryableFlowsDistroXV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListRetryableFlowsDistroXV1ParamsWithTimeout(timeout time.Duration) *ListRetryableFlowsDistroXV1Params {
	var ()
	return &ListRetryableFlowsDistroXV1Params{

		timeout: timeout,
	}
}

// NewListRetryableFlowsDistroXV1ParamsWithContext creates a new ListRetryableFlowsDistroXV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewListRetryableFlowsDistroXV1ParamsWithContext(ctx context.Context) *ListRetryableFlowsDistroXV1Params {
	var ()
	return &ListRetryableFlowsDistroXV1Params{

		Context: ctx,
	}
}

// NewListRetryableFlowsDistroXV1ParamsWithHTTPClient creates a new ListRetryableFlowsDistroXV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListRetryableFlowsDistroXV1ParamsWithHTTPClient(client *http.Client) *ListRetryableFlowsDistroXV1Params {
	var ()
	return &ListRetryableFlowsDistroXV1Params{
		HTTPClient: client,
	}
}

/*ListRetryableFlowsDistroXV1Params contains all the parameters to send to the API endpoint
for the list retryable flows distro x v1 operation typically these are written to a http.Request
*/
type ListRetryableFlowsDistroXV1Params struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list retryable flows distro x v1 params
func (o *ListRetryableFlowsDistroXV1Params) WithTimeout(timeout time.Duration) *ListRetryableFlowsDistroXV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list retryable flows distro x v1 params
func (o *ListRetryableFlowsDistroXV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list retryable flows distro x v1 params
func (o *ListRetryableFlowsDistroXV1Params) WithContext(ctx context.Context) *ListRetryableFlowsDistroXV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list retryable flows distro x v1 params
func (o *ListRetryableFlowsDistroXV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list retryable flows distro x v1 params
func (o *ListRetryableFlowsDistroXV1Params) WithHTTPClient(client *http.Client) *ListRetryableFlowsDistroXV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list retryable flows distro x v1 params
func (o *ListRetryableFlowsDistroXV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the list retryable flows distro x v1 params
func (o *ListRetryableFlowsDistroXV1Params) WithName(name string) *ListRetryableFlowsDistroXV1Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the list retryable flows distro x v1 params
func (o *ListRetryableFlowsDistroXV1Params) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ListRetryableFlowsDistroXV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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