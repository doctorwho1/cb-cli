// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_environments

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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewRegisterDatalakePrerequisitesParams creates a new RegisterDatalakePrerequisitesParams object
// with the default values initialized.
func NewRegisterDatalakePrerequisitesParams() *RegisterDatalakePrerequisitesParams {
	var ()
	return &RegisterDatalakePrerequisitesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterDatalakePrerequisitesParamsWithTimeout creates a new RegisterDatalakePrerequisitesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRegisterDatalakePrerequisitesParamsWithTimeout(timeout time.Duration) *RegisterDatalakePrerequisitesParams {
	var ()
	return &RegisterDatalakePrerequisitesParams{

		timeout: timeout,
	}
}

// NewRegisterDatalakePrerequisitesParamsWithContext creates a new RegisterDatalakePrerequisitesParams object
// with the default values initialized, and the ability to set a context for a request
func NewRegisterDatalakePrerequisitesParamsWithContext(ctx context.Context) *RegisterDatalakePrerequisitesParams {
	var ()
	return &RegisterDatalakePrerequisitesParams{

		Context: ctx,
	}
}

// NewRegisterDatalakePrerequisitesParamsWithHTTPClient creates a new RegisterDatalakePrerequisitesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRegisterDatalakePrerequisitesParamsWithHTTPClient(client *http.Client) *RegisterDatalakePrerequisitesParams {
	var ()
	return &RegisterDatalakePrerequisitesParams{
		HTTPClient: client,
	}
}

/*RegisterDatalakePrerequisitesParams contains all the parameters to send to the API endpoint
for the register datalake prerequisites operation typically these are written to a http.Request
*/
type RegisterDatalakePrerequisitesParams struct {

	/*Body*/
	Body *model.DatalakePrerequisiteV4Request
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the register datalake prerequisites params
func (o *RegisterDatalakePrerequisitesParams) WithTimeout(timeout time.Duration) *RegisterDatalakePrerequisitesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register datalake prerequisites params
func (o *RegisterDatalakePrerequisitesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register datalake prerequisites params
func (o *RegisterDatalakePrerequisitesParams) WithContext(ctx context.Context) *RegisterDatalakePrerequisitesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register datalake prerequisites params
func (o *RegisterDatalakePrerequisitesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register datalake prerequisites params
func (o *RegisterDatalakePrerequisitesParams) WithHTTPClient(client *http.Client) *RegisterDatalakePrerequisitesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register datalake prerequisites params
func (o *RegisterDatalakePrerequisitesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the register datalake prerequisites params
func (o *RegisterDatalakePrerequisitesParams) WithBody(body *model.DatalakePrerequisiteV4Request) *RegisterDatalakePrerequisitesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the register datalake prerequisites params
func (o *RegisterDatalakePrerequisitesParams) SetBody(body *model.DatalakePrerequisiteV4Request) {
	o.Body = body
}

// WithName adds the name to the register datalake prerequisites params
func (o *RegisterDatalakePrerequisitesParams) WithName(name string) *RegisterDatalakePrerequisitesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the register datalake prerequisites params
func (o *RegisterDatalakePrerequisitesParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the register datalake prerequisites params
func (o *RegisterDatalakePrerequisitesParams) WithWorkspaceID(workspaceID int64) *RegisterDatalakePrerequisitesParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the register datalake prerequisites params
func (o *RegisterDatalakePrerequisitesParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterDatalakePrerequisitesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
