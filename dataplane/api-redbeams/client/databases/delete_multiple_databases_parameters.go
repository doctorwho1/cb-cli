// Code generated by go-swagger; DO NOT EDIT.

package databases

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

// NewDeleteMultipleDatabasesParams creates a new DeleteMultipleDatabasesParams object
// with the default values initialized.
func NewDeleteMultipleDatabasesParams() *DeleteMultipleDatabasesParams {
	var ()
	return &DeleteMultipleDatabasesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMultipleDatabasesParamsWithTimeout creates a new DeleteMultipleDatabasesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMultipleDatabasesParamsWithTimeout(timeout time.Duration) *DeleteMultipleDatabasesParams {
	var ()
	return &DeleteMultipleDatabasesParams{

		timeout: timeout,
	}
}

// NewDeleteMultipleDatabasesParamsWithContext creates a new DeleteMultipleDatabasesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMultipleDatabasesParamsWithContext(ctx context.Context) *DeleteMultipleDatabasesParams {
	var ()
	return &DeleteMultipleDatabasesParams{

		Context: ctx,
	}
}

// NewDeleteMultipleDatabasesParamsWithHTTPClient creates a new DeleteMultipleDatabasesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMultipleDatabasesParamsWithHTTPClient(client *http.Client) *DeleteMultipleDatabasesParams {
	var ()
	return &DeleteMultipleDatabasesParams{
		HTTPClient: client,
	}
}

/*DeleteMultipleDatabasesParams contains all the parameters to send to the API endpoint
for the delete multiple databases operation typically these are written to a http.Request
*/
type DeleteMultipleDatabasesParams struct {

	/*Body
	  Names of the databases

	*/
	Body []string
	/*EnvironmentCrn
	  CRN of the environment of the database(s)

	*/
	EnvironmentCrn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete multiple databases params
func (o *DeleteMultipleDatabasesParams) WithTimeout(timeout time.Duration) *DeleteMultipleDatabasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete multiple databases params
func (o *DeleteMultipleDatabasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete multiple databases params
func (o *DeleteMultipleDatabasesParams) WithContext(ctx context.Context) *DeleteMultipleDatabasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete multiple databases params
func (o *DeleteMultipleDatabasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete multiple databases params
func (o *DeleteMultipleDatabasesParams) WithHTTPClient(client *http.Client) *DeleteMultipleDatabasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete multiple databases params
func (o *DeleteMultipleDatabasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete multiple databases params
func (o *DeleteMultipleDatabasesParams) WithBody(body []string) *DeleteMultipleDatabasesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete multiple databases params
func (o *DeleteMultipleDatabasesParams) SetBody(body []string) {
	o.Body = body
}

// WithEnvironmentCrn adds the environmentCrn to the delete multiple databases params
func (o *DeleteMultipleDatabasesParams) WithEnvironmentCrn(environmentCrn string) *DeleteMultipleDatabasesParams {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the delete multiple databases params
func (o *DeleteMultipleDatabasesParams) SetEnvironmentCrn(environmentCrn string) {
	o.EnvironmentCrn = environmentCrn
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMultipleDatabasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param environmentCrn
	qrEnvironmentCrn := o.EnvironmentCrn
	qEnvironmentCrn := qrEnvironmentCrn
	if qEnvironmentCrn != "" {
		if err := r.SetQueryParam("environmentCrn", qEnvironmentCrn); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}