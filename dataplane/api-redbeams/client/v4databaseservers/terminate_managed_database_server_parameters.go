// Code generated by go-swagger; DO NOT EDIT.

package v4databaseservers

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

// NewTerminateManagedDatabaseServerParams creates a new TerminateManagedDatabaseServerParams object
// with the default values initialized.
func NewTerminateManagedDatabaseServerParams() *TerminateManagedDatabaseServerParams {
	var ()
	return &TerminateManagedDatabaseServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTerminateManagedDatabaseServerParamsWithTimeout creates a new TerminateManagedDatabaseServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTerminateManagedDatabaseServerParamsWithTimeout(timeout time.Duration) *TerminateManagedDatabaseServerParams {
	var ()
	return &TerminateManagedDatabaseServerParams{

		timeout: timeout,
	}
}

// NewTerminateManagedDatabaseServerParamsWithContext creates a new TerminateManagedDatabaseServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewTerminateManagedDatabaseServerParamsWithContext(ctx context.Context) *TerminateManagedDatabaseServerParams {
	var ()
	return &TerminateManagedDatabaseServerParams{

		Context: ctx,
	}
}

// NewTerminateManagedDatabaseServerParamsWithHTTPClient creates a new TerminateManagedDatabaseServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTerminateManagedDatabaseServerParamsWithHTTPClient(client *http.Client) *TerminateManagedDatabaseServerParams {
	var ()
	return &TerminateManagedDatabaseServerParams{
		HTTPClient: client,
	}
}

/*TerminateManagedDatabaseServerParams contains all the parameters to send to the API endpoint
for the terminate managed database server operation typically these are written to a http.Request
*/
type TerminateManagedDatabaseServerParams struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the terminate managed database server params
func (o *TerminateManagedDatabaseServerParams) WithTimeout(timeout time.Duration) *TerminateManagedDatabaseServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the terminate managed database server params
func (o *TerminateManagedDatabaseServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the terminate managed database server params
func (o *TerminateManagedDatabaseServerParams) WithContext(ctx context.Context) *TerminateManagedDatabaseServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the terminate managed database server params
func (o *TerminateManagedDatabaseServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the terminate managed database server params
func (o *TerminateManagedDatabaseServerParams) WithHTTPClient(client *http.Client) *TerminateManagedDatabaseServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the terminate managed database server params
func (o *TerminateManagedDatabaseServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the terminate managed database server params
func (o *TerminateManagedDatabaseServerParams) WithCrn(crn string) *TerminateManagedDatabaseServerParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the terminate managed database server params
func (o *TerminateManagedDatabaseServerParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *TerminateManagedDatabaseServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
