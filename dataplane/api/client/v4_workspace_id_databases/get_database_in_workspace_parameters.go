// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_databases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDatabaseInWorkspaceParams creates a new GetDatabaseInWorkspaceParams object
// with the default values initialized.
func NewGetDatabaseInWorkspaceParams() *GetDatabaseInWorkspaceParams {
	var ()
	return &GetDatabaseInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatabaseInWorkspaceParamsWithTimeout creates a new GetDatabaseInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDatabaseInWorkspaceParamsWithTimeout(timeout time.Duration) *GetDatabaseInWorkspaceParams {
	var ()
	return &GetDatabaseInWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetDatabaseInWorkspaceParamsWithContext creates a new GetDatabaseInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDatabaseInWorkspaceParamsWithContext(ctx context.Context) *GetDatabaseInWorkspaceParams {
	var ()
	return &GetDatabaseInWorkspaceParams{

		Context: ctx,
	}
}

// NewGetDatabaseInWorkspaceParamsWithHTTPClient creates a new GetDatabaseInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDatabaseInWorkspaceParamsWithHTTPClient(client *http.Client) *GetDatabaseInWorkspaceParams {
	var ()
	return &GetDatabaseInWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetDatabaseInWorkspaceParams contains all the parameters to send to the API endpoint
for the get database in workspace operation typically these are written to a http.Request
*/
type GetDatabaseInWorkspaceParams struct {

	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get database in workspace params
func (o *GetDatabaseInWorkspaceParams) WithTimeout(timeout time.Duration) *GetDatabaseInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get database in workspace params
func (o *GetDatabaseInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get database in workspace params
func (o *GetDatabaseInWorkspaceParams) WithContext(ctx context.Context) *GetDatabaseInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get database in workspace params
func (o *GetDatabaseInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get database in workspace params
func (o *GetDatabaseInWorkspaceParams) WithHTTPClient(client *http.Client) *GetDatabaseInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get database in workspace params
func (o *GetDatabaseInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get database in workspace params
func (o *GetDatabaseInWorkspaceParams) WithName(name string) *GetDatabaseInWorkspaceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get database in workspace params
func (o *GetDatabaseInWorkspaceParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the get database in workspace params
func (o *GetDatabaseInWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetDatabaseInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get database in workspace params
func (o *GetDatabaseInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatabaseInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
