// Code generated by go-swagger; DO NOT EDIT.

package v1platform_resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPlatformNetworksForWorkspaceParams creates a new GetPlatformNetworksForWorkspaceParams object
// with the default values initialized.
func NewGetPlatformNetworksForWorkspaceParams() *GetPlatformNetworksForWorkspaceParams {
	var ()
	return &GetPlatformNetworksForWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlatformNetworksForWorkspaceParamsWithTimeout creates a new GetPlatformNetworksForWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPlatformNetworksForWorkspaceParamsWithTimeout(timeout time.Duration) *GetPlatformNetworksForWorkspaceParams {
	var ()
	return &GetPlatformNetworksForWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetPlatformNetworksForWorkspaceParamsWithContext creates a new GetPlatformNetworksForWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPlatformNetworksForWorkspaceParamsWithContext(ctx context.Context) *GetPlatformNetworksForWorkspaceParams {
	var ()
	return &GetPlatformNetworksForWorkspaceParams{

		Context: ctx,
	}
}

// NewGetPlatformNetworksForWorkspaceParamsWithHTTPClient creates a new GetPlatformNetworksForWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPlatformNetworksForWorkspaceParamsWithHTTPClient(client *http.Client) *GetPlatformNetworksForWorkspaceParams {
	var ()
	return &GetPlatformNetworksForWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetPlatformNetworksForWorkspaceParams contains all the parameters to send to the API endpoint
for the get platform networks for workspace operation typically these are written to a http.Request
*/
type GetPlatformNetworksForWorkspaceParams struct {

	/*AvailabilityZone*/
	AvailabilityZone *string
	/*CredentialName*/
	CredentialName *string
	/*PlatformVariant*/
	PlatformVariant *string
	/*Region*/
	Region *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) WithTimeout(timeout time.Duration) *GetPlatformNetworksForWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) WithContext(ctx context.Context) *GetPlatformNetworksForWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) WithHTTPClient(client *http.Client) *GetPlatformNetworksForWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityZone adds the availabilityZone to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) WithAvailabilityZone(availabilityZone *string) *GetPlatformNetworksForWorkspaceParams {
	o.SetAvailabilityZone(availabilityZone)
	return o
}

// SetAvailabilityZone adds the availabilityZone to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) SetAvailabilityZone(availabilityZone *string) {
	o.AvailabilityZone = availabilityZone
}

// WithCredentialName adds the credentialName to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) WithCredentialName(credentialName *string) *GetPlatformNetworksForWorkspaceParams {
	o.SetCredentialName(credentialName)
	return o
}

// SetCredentialName adds the credentialName to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) SetCredentialName(credentialName *string) {
	o.CredentialName = credentialName
}

// WithPlatformVariant adds the platformVariant to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) WithPlatformVariant(platformVariant *string) *GetPlatformNetworksForWorkspaceParams {
	o.SetPlatformVariant(platformVariant)
	return o
}

// SetPlatformVariant adds the platformVariant to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) SetPlatformVariant(platformVariant *string) {
	o.PlatformVariant = platformVariant
}

// WithRegion adds the region to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) WithRegion(region *string) *GetPlatformNetworksForWorkspaceParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) SetRegion(region *string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlatformNetworksForWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AvailabilityZone != nil {

		// query param availabilityZone
		var qrAvailabilityZone string
		if o.AvailabilityZone != nil {
			qrAvailabilityZone = *o.AvailabilityZone
		}
		qAvailabilityZone := qrAvailabilityZone
		if qAvailabilityZone != "" {
			if err := r.SetQueryParam("availabilityZone", qAvailabilityZone); err != nil {
				return err
			}
		}

	}

	if o.CredentialName != nil {

		// query param credentialName
		var qrCredentialName string
		if o.CredentialName != nil {
			qrCredentialName = *o.CredentialName
		}
		qCredentialName := qrCredentialName
		if qCredentialName != "" {
			if err := r.SetQueryParam("credentialName", qCredentialName); err != nil {
				return err
			}
		}

	}

	if o.PlatformVariant != nil {

		// query param platformVariant
		var qrPlatformVariant string
		if o.PlatformVariant != nil {
			qrPlatformVariant = *o.PlatformVariant
		}
		qPlatformVariant := qrPlatformVariant
		if qPlatformVariant != "" {
			if err := r.SetQueryParam("platformVariant", qPlatformVariant); err != nil {
				return err
			}
		}

	}

	if o.Region != nil {

		// query param region
		var qrRegion string
		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {
			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
