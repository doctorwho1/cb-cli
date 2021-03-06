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

// NewCheckForUpgradeParams creates a new CheckForUpgradeParams object
// with the default values initialized.
func NewCheckForUpgradeParams() *CheckForUpgradeParams {
	var ()
	return &CheckForUpgradeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckForUpgradeParamsWithTimeout creates a new CheckForUpgradeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckForUpgradeParamsWithTimeout(timeout time.Duration) *CheckForUpgradeParams {
	var ()
	return &CheckForUpgradeParams{

		timeout: timeout,
	}
}

// NewCheckForUpgradeParamsWithContext creates a new CheckForUpgradeParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckForUpgradeParamsWithContext(ctx context.Context) *CheckForUpgradeParams {
	var ()
	return &CheckForUpgradeParams{

		Context: ctx,
	}
}

// NewCheckForUpgradeParamsWithHTTPClient creates a new CheckForUpgradeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckForUpgradeParamsWithHTTPClient(client *http.Client) *CheckForUpgradeParams {
	var ()
	return &CheckForUpgradeParams{
		HTTPClient: client,
	}
}

/*CheckForUpgradeParams contains all the parameters to send to the API endpoint
for the check for upgrade operation typically these are written to a http.Request
*/
type CheckForUpgradeParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the check for upgrade params
func (o *CheckForUpgradeParams) WithTimeout(timeout time.Duration) *CheckForUpgradeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check for upgrade params
func (o *CheckForUpgradeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check for upgrade params
func (o *CheckForUpgradeParams) WithContext(ctx context.Context) *CheckForUpgradeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check for upgrade params
func (o *CheckForUpgradeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check for upgrade params
func (o *CheckForUpgradeParams) WithHTTPClient(client *http.Client) *CheckForUpgradeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check for upgrade params
func (o *CheckForUpgradeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the check for upgrade params
func (o *CheckForUpgradeParams) WithName(name string) *CheckForUpgradeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the check for upgrade params
func (o *CheckForUpgradeParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *CheckForUpgradeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
