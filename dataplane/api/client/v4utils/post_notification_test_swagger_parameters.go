// Code generated by go-swagger; DO NOT EDIT.

package v4utils

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

// NewPostNotificationTestParams creates a new PostNotificationTestParams object
// with the default values initialized.
func NewPostNotificationTestParams() *PostNotificationTestParams {

	return &PostNotificationTestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostNotificationTestParamsWithTimeout creates a new PostNotificationTestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostNotificationTestParamsWithTimeout(timeout time.Duration) *PostNotificationTestParams {

	return &PostNotificationTestParams{

		timeout: timeout,
	}
}

// NewPostNotificationTestParamsWithContext creates a new PostNotificationTestParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostNotificationTestParamsWithContext(ctx context.Context) *PostNotificationTestParams {

	return &PostNotificationTestParams{

		Context: ctx,
	}
}

// NewPostNotificationTestParamsWithHTTPClient creates a new PostNotificationTestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostNotificationTestParamsWithHTTPClient(client *http.Client) *PostNotificationTestParams {

	return &PostNotificationTestParams{
		HTTPClient: client,
	}
}

/*PostNotificationTestParams contains all the parameters to send to the API endpoint
for the post notification test operation typically these are written to a http.Request
*/
type PostNotificationTestParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post notification test params
func (o *PostNotificationTestParams) WithTimeout(timeout time.Duration) *PostNotificationTestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post notification test params
func (o *PostNotificationTestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post notification test params
func (o *PostNotificationTestParams) WithContext(ctx context.Context) *PostNotificationTestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post notification test params
func (o *PostNotificationTestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post notification test params
func (o *PostNotificationTestParams) WithHTTPClient(client *http.Client) *PostNotificationTestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post notification test params
func (o *PostNotificationTestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostNotificationTestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}