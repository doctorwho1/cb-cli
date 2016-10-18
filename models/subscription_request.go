package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*SubscriptionRequest subscription request

swagger:model SubscriptionRequest
*/
type SubscriptionRequest struct {

	/* url of the endpoint

	Required: true
	Pattern: ^(https?:\/\/)((([\da-z\.-]+)\.([a-z]{2,6}))|localhost|[1-9][0-9]{0,2}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3})(:[1-9][0-9]{1,4})?\/([\/\w\.-]*)\/?$
	*/
	EndpointURL string `json:"endpointUrl"`
}

// Validate validates this subscription request
func (m *SubscriptionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpointURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionRequest) validateEndpointURL(formats strfmt.Registry) error {

	if err := validate.RequiredString("endpointUrl", "body", string(m.EndpointURL)); err != nil {
		return err
	}

	if err := validate.Pattern("endpointUrl", "body", string(m.EndpointURL), `^(https?:\/\/)((([\da-z\.-]+)\.([a-z]{2,6}))|localhost|[1-9][0-9]{0,2}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3})(:[1-9][0-9]{1,4})?\/([\/\w\.-]*)\/?$`); err != nil {
		return err
	}

	return nil
}
