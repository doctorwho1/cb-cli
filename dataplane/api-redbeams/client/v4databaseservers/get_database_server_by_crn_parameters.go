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

// NewGetDatabaseServerByCrnParams creates a new GetDatabaseServerByCrnParams object
// with the default values initialized.
func NewGetDatabaseServerByCrnParams() *GetDatabaseServerByCrnParams {
	var ()
	return &GetDatabaseServerByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatabaseServerByCrnParamsWithTimeout creates a new GetDatabaseServerByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDatabaseServerByCrnParamsWithTimeout(timeout time.Duration) *GetDatabaseServerByCrnParams {
	var ()
	return &GetDatabaseServerByCrnParams{

		timeout: timeout,
	}
}

// NewGetDatabaseServerByCrnParamsWithContext creates a new GetDatabaseServerByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDatabaseServerByCrnParamsWithContext(ctx context.Context) *GetDatabaseServerByCrnParams {
	var ()
	return &GetDatabaseServerByCrnParams{

		Context: ctx,
	}
}

// NewGetDatabaseServerByCrnParamsWithHTTPClient creates a new GetDatabaseServerByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDatabaseServerByCrnParamsWithHTTPClient(client *http.Client) *GetDatabaseServerByCrnParams {
	var ()
	return &GetDatabaseServerByCrnParams{
		HTTPClient: client,
	}
}

/*GetDatabaseServerByCrnParams contains all the parameters to send to the API endpoint
for the get database server by crn operation typically these are written to a http.Request
*/
type GetDatabaseServerByCrnParams struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get database server by crn params
func (o *GetDatabaseServerByCrnParams) WithTimeout(timeout time.Duration) *GetDatabaseServerByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get database server by crn params
func (o *GetDatabaseServerByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get database server by crn params
func (o *GetDatabaseServerByCrnParams) WithContext(ctx context.Context) *GetDatabaseServerByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get database server by crn params
func (o *GetDatabaseServerByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get database server by crn params
func (o *GetDatabaseServerByCrnParams) WithHTTPClient(client *http.Client) *GetDatabaseServerByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get database server by crn params
func (o *GetDatabaseServerByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the get database server by crn params
func (o *GetDatabaseServerByCrnParams) WithCrn(crn string) *GetDatabaseServerByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the get database server by crn params
func (o *GetDatabaseServerByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatabaseServerByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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