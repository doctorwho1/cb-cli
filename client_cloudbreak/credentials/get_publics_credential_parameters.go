package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetPublicsCredentialParams creates a new GetPublicsCredentialParams object
// with the default values initialized.
func NewGetPublicsCredentialParams() *GetPublicsCredentialParams {

	return &GetPublicsCredentialParams{}
}

/*GetPublicsCredentialParams contains all the parameters to send to the API endpoint
for the get publics credential operation typically these are written to a http.Request
*/
type GetPublicsCredentialParams struct {
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicsCredentialParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}