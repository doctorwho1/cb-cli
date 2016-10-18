package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*CredentialResponse credential response

swagger:model CredentialResponse
*/
type CredentialResponse struct {

	/* type of cloud provider

	Required: true
	*/
	CloudPlatform string `json:"cloudPlatform"`

	/* description of the resource

	Max Length: 1000
	Min Length: 0
	*/
	Description *string `json:"description,omitempty"`

	/* id of the resource
	 */
	ID *int64 `json:"id,omitempty"`

	/* authentication name for machines
	 */
	LoginUserName *string `json:"loginUserName,omitempty"`

	/* name of the resource

	Required: true
	Max Length: 100
	Min Length: 5
	Pattern: ([a-z][-a-z0-9]*[a-z0-9])
	*/
	Name string `json:"name"`

	/* cloud specific parameters for credential
	 */
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	/* resource is visible in account
	 */
	Public *bool `json:"public,omitempty"`

	/* public key for accessing instances

	Required: true
	*/
	PublicKey string `json:"publicKey"`

	/* id of the topology the resource belongs to
	 */
	TopologyID *int64 `json:"topologyId,omitempty"`
}

// Validate validates this credential response
func (m *CredentialResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudPlatform(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialResponse) validateCloudPlatform(formats strfmt.Registry) error {

	if err := validate.RequiredString("cloudPlatform", "body", string(m.CloudPlatform)); err != nil {
		return err
	}

	return nil
}

func (m *CredentialResponse) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *CredentialResponse) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 100); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(m.Name), `([a-z][-a-z0-9]*[a-z0-9])`); err != nil {
		return err
	}

	return nil
}

func (m *CredentialResponse) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.RequiredString("publicKey", "body", string(m.PublicKey)); err != nil {
		return err
	}

	return nil
}
