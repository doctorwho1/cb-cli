// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_blueprints

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

// NewGetBlueprintRequestFromNameParams creates a new GetBlueprintRequestFromNameParams object
// with the default values initialized.
func NewGetBlueprintRequestFromNameParams() *GetBlueprintRequestFromNameParams {
	var ()
	return &GetBlueprintRequestFromNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBlueprintRequestFromNameParamsWithTimeout creates a new GetBlueprintRequestFromNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBlueprintRequestFromNameParamsWithTimeout(timeout time.Duration) *GetBlueprintRequestFromNameParams {
	var ()
	return &GetBlueprintRequestFromNameParams{

		timeout: timeout,
	}
}

// NewGetBlueprintRequestFromNameParamsWithContext creates a new GetBlueprintRequestFromNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBlueprintRequestFromNameParamsWithContext(ctx context.Context) *GetBlueprintRequestFromNameParams {
	var ()
	return &GetBlueprintRequestFromNameParams{

		Context: ctx,
	}
}

// NewGetBlueprintRequestFromNameParamsWithHTTPClient creates a new GetBlueprintRequestFromNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBlueprintRequestFromNameParamsWithHTTPClient(client *http.Client) *GetBlueprintRequestFromNameParams {
	var ()
	return &GetBlueprintRequestFromNameParams{
		HTTPClient: client,
	}
}

/*GetBlueprintRequestFromNameParams contains all the parameters to send to the API endpoint
for the get blueprint request from name operation typically these are written to a http.Request
*/
type GetBlueprintRequestFromNameParams struct {

	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get blueprint request from name params
func (o *GetBlueprintRequestFromNameParams) WithTimeout(timeout time.Duration) *GetBlueprintRequestFromNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get blueprint request from name params
func (o *GetBlueprintRequestFromNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get blueprint request from name params
func (o *GetBlueprintRequestFromNameParams) WithContext(ctx context.Context) *GetBlueprintRequestFromNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get blueprint request from name params
func (o *GetBlueprintRequestFromNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get blueprint request from name params
func (o *GetBlueprintRequestFromNameParams) WithHTTPClient(client *http.Client) *GetBlueprintRequestFromNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get blueprint request from name params
func (o *GetBlueprintRequestFromNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get blueprint request from name params
func (o *GetBlueprintRequestFromNameParams) WithName(name string) *GetBlueprintRequestFromNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get blueprint request from name params
func (o *GetBlueprintRequestFromNameParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the get blueprint request from name params
func (o *GetBlueprintRequestFromNameParams) WithWorkspaceID(workspaceID int64) *GetBlueprintRequestFromNameParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get blueprint request from name params
func (o *GetBlueprintRequestFromNameParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBlueprintRequestFromNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
