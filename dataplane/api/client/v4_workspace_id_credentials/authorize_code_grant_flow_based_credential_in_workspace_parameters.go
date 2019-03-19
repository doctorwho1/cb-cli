// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_credentials

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

// NewAuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams creates a new AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams object
// with the default values initialized.
func NewAuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams() *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams {
	var ()
	return &AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAuthorizeCodeGrantFlowBasedCredentialInWorkspaceParamsWithTimeout creates a new AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAuthorizeCodeGrantFlowBasedCredentialInWorkspaceParamsWithTimeout(timeout time.Duration) *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams {
	var ()
	return &AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams{

		timeout: timeout,
	}
}

// NewAuthorizeCodeGrantFlowBasedCredentialInWorkspaceParamsWithContext creates a new AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewAuthorizeCodeGrantFlowBasedCredentialInWorkspaceParamsWithContext(ctx context.Context) *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams {
	var ()
	return &AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams{

		Context: ctx,
	}
}

// NewAuthorizeCodeGrantFlowBasedCredentialInWorkspaceParamsWithHTTPClient creates a new AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAuthorizeCodeGrantFlowBasedCredentialInWorkspaceParamsWithHTTPClient(client *http.Client) *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams {
	var ()
	return &AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams{
		HTTPClient: client,
	}
}

/*AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams contains all the parameters to send to the API endpoint
for the authorize code grant flow based credential in workspace operation typically these are written to a http.Request
*/
type AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams struct {

	/*CloudPlatform*/
	CloudPlatform string
	/*Code*/
	Code *string
	/*State*/
	State *string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the authorize code grant flow based credential in workspace params
func (o *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams) WithTimeout(timeout time.Duration) *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authorize code grant flow based credential in workspace params
func (o *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authorize code grant flow based credential in workspace params
func (o *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams) WithContext(ctx context.Context) *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authorize code grant flow based credential in workspace params
func (o *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authorize code grant flow based credential in workspace params
func (o *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams) WithHTTPClient(client *http.Client) *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authorize code grant flow based credential in workspace params
func (o *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudPlatform adds the cloudPlatform to the authorize code grant flow based credential in workspace params
func (o *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams) WithCloudPlatform(cloudPlatform string) *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams {
	o.SetCloudPlatform(cloudPlatform)
	return o
}

// SetCloudPlatform adds the cloudPlatform to the authorize code grant flow based credential in workspace params
func (o *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams) SetCloudPlatform(cloudPlatform string) {
	o.CloudPlatform = cloudPlatform
}

// WithCode adds the code to the authorize code grant flow based credential in workspace params
func (o *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams) WithCode(code *string) *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the authorize code grant flow based credential in workspace params
func (o *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams) SetCode(code *string) {
	o.Code = code
}

// WithState adds the state to the authorize code grant flow based credential in workspace params
func (o *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams) WithState(state *string) *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the authorize code grant flow based credential in workspace params
func (o *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams) SetState(state *string) {
	o.State = state
}

// WithWorkspaceID adds the workspaceID to the authorize code grant flow based credential in workspace params
func (o *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams) WithWorkspaceID(workspaceID int64) *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the authorize code grant flow based credential in workspace params
func (o *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *AuthorizeCodeGrantFlowBasedCredentialInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloudPlatform
	if err := r.SetPathParam("cloudPlatform", o.CloudPlatform); err != nil {
		return err
	}

	if o.Code != nil {

		// query param code
		var qrCode string
		if o.Code != nil {
			qrCode = *o.Code
		}
		qCode := qrCode
		if qCode != "" {
			if err := r.SetQueryParam("code", qCode); err != nil {
				return err
			}
		}

	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

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