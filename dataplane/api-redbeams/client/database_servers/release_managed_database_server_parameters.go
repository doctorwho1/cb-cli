// Code generated by go-swagger; DO NOT EDIT.

package database_servers

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

// NewReleaseManagedDatabaseServerParams creates a new ReleaseManagedDatabaseServerParams object
// with the default values initialized.
func NewReleaseManagedDatabaseServerParams() *ReleaseManagedDatabaseServerParams {
	var ()
	return &ReleaseManagedDatabaseServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReleaseManagedDatabaseServerParamsWithTimeout creates a new ReleaseManagedDatabaseServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReleaseManagedDatabaseServerParamsWithTimeout(timeout time.Duration) *ReleaseManagedDatabaseServerParams {
	var ()
	return &ReleaseManagedDatabaseServerParams{

		timeout: timeout,
	}
}

// NewReleaseManagedDatabaseServerParamsWithContext creates a new ReleaseManagedDatabaseServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewReleaseManagedDatabaseServerParamsWithContext(ctx context.Context) *ReleaseManagedDatabaseServerParams {
	var ()
	return &ReleaseManagedDatabaseServerParams{

		Context: ctx,
	}
}

// NewReleaseManagedDatabaseServerParamsWithHTTPClient creates a new ReleaseManagedDatabaseServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReleaseManagedDatabaseServerParamsWithHTTPClient(client *http.Client) *ReleaseManagedDatabaseServerParams {
	var ()
	return &ReleaseManagedDatabaseServerParams{
		HTTPClient: client,
	}
}

/*ReleaseManagedDatabaseServerParams contains all the parameters to send to the API endpoint
for the release managed database server operation typically these are written to a http.Request
*/
type ReleaseManagedDatabaseServerParams struct {

	/*Crn
	  CRN of the database server

	*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the release managed database server params
func (o *ReleaseManagedDatabaseServerParams) WithTimeout(timeout time.Duration) *ReleaseManagedDatabaseServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the release managed database server params
func (o *ReleaseManagedDatabaseServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the release managed database server params
func (o *ReleaseManagedDatabaseServerParams) WithContext(ctx context.Context) *ReleaseManagedDatabaseServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the release managed database server params
func (o *ReleaseManagedDatabaseServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the release managed database server params
func (o *ReleaseManagedDatabaseServerParams) WithHTTPClient(client *http.Client) *ReleaseManagedDatabaseServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the release managed database server params
func (o *ReleaseManagedDatabaseServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the release managed database server params
func (o *ReleaseManagedDatabaseServerParams) WithCrn(crn string) *ReleaseManagedDatabaseServerParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the release managed database server params
func (o *ReleaseManagedDatabaseServerParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *ReleaseManagedDatabaseServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
