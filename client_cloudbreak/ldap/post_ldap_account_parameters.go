package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// NewPostLdapAccountParams creates a new PostLdapAccountParams object
// with the default values initialized.
func NewPostLdapAccountParams() *PostLdapAccountParams {
	var ()
	return &PostLdapAccountParams{}
}

/*PostLdapAccountParams contains all the parameters to send to the API endpoint
for the post ldap account operation typically these are written to a http.Request
*/
type PostLdapAccountParams struct {

	/*Body*/
	Body *models_cloudbreak.LdapConfigRequest
}

// WithBody adds the body to the post ldap account params
func (o *PostLdapAccountParams) WithBody(body *models_cloudbreak.LdapConfigRequest) *PostLdapAccountParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostLdapAccountParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.LdapConfigRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}