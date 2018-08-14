// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_smartsensesubscriptions

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

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// NewCreateSmartSenseSubscriptionInOrganizationParams creates a new CreateSmartSenseSubscriptionInOrganizationParams object
// with the default values initialized.
func NewCreateSmartSenseSubscriptionInOrganizationParams() *CreateSmartSenseSubscriptionInOrganizationParams {
	var ()
	return &CreateSmartSenseSubscriptionInOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSmartSenseSubscriptionInOrganizationParamsWithTimeout creates a new CreateSmartSenseSubscriptionInOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSmartSenseSubscriptionInOrganizationParamsWithTimeout(timeout time.Duration) *CreateSmartSenseSubscriptionInOrganizationParams {
	var ()
	return &CreateSmartSenseSubscriptionInOrganizationParams{

		timeout: timeout,
	}
}

// NewCreateSmartSenseSubscriptionInOrganizationParamsWithContext creates a new CreateSmartSenseSubscriptionInOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSmartSenseSubscriptionInOrganizationParamsWithContext(ctx context.Context) *CreateSmartSenseSubscriptionInOrganizationParams {
	var ()
	return &CreateSmartSenseSubscriptionInOrganizationParams{

		Context: ctx,
	}
}

// NewCreateSmartSenseSubscriptionInOrganizationParamsWithHTTPClient creates a new CreateSmartSenseSubscriptionInOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSmartSenseSubscriptionInOrganizationParamsWithHTTPClient(client *http.Client) *CreateSmartSenseSubscriptionInOrganizationParams {
	var ()
	return &CreateSmartSenseSubscriptionInOrganizationParams{
		HTTPClient: client,
	}
}

/*CreateSmartSenseSubscriptionInOrganizationParams contains all the parameters to send to the API endpoint
for the create smart sense subscription in organization operation typically these are written to a http.Request
*/
type CreateSmartSenseSubscriptionInOrganizationParams struct {

	/*Body*/
	Body *models_cloudbreak.SmartSenseSubscriptionJSON
	/*OrganizationID*/
	OrganizationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create smart sense subscription in organization params
func (o *CreateSmartSenseSubscriptionInOrganizationParams) WithTimeout(timeout time.Duration) *CreateSmartSenseSubscriptionInOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create smart sense subscription in organization params
func (o *CreateSmartSenseSubscriptionInOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create smart sense subscription in organization params
func (o *CreateSmartSenseSubscriptionInOrganizationParams) WithContext(ctx context.Context) *CreateSmartSenseSubscriptionInOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create smart sense subscription in organization params
func (o *CreateSmartSenseSubscriptionInOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create smart sense subscription in organization params
func (o *CreateSmartSenseSubscriptionInOrganizationParams) WithHTTPClient(client *http.Client) *CreateSmartSenseSubscriptionInOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create smart sense subscription in organization params
func (o *CreateSmartSenseSubscriptionInOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create smart sense subscription in organization params
func (o *CreateSmartSenseSubscriptionInOrganizationParams) WithBody(body *models_cloudbreak.SmartSenseSubscriptionJSON) *CreateSmartSenseSubscriptionInOrganizationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create smart sense subscription in organization params
func (o *CreateSmartSenseSubscriptionInOrganizationParams) SetBody(body *models_cloudbreak.SmartSenseSubscriptionJSON) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the create smart sense subscription in organization params
func (o *CreateSmartSenseSubscriptionInOrganizationParams) WithOrganizationID(organizationID int64) *CreateSmartSenseSubscriptionInOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create smart sense subscription in organization params
func (o *CreateSmartSenseSubscriptionInOrganizationParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSmartSenseSubscriptionInOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.SmartSenseSubscriptionJSON)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", swag.FormatInt64(o.OrganizationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}