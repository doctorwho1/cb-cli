// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainKeystoneV3Parameters domain keystone v3 parameters
// swagger:model DomainKeystoneV3Parameters
type DomainKeystoneV3Parameters struct {

	// domain name
	// Required: true
	DomainName *string `json:"domainName"`

	// user domain
	// Required: true
	UserDomain *string `json:"userDomain"`
}

// Validate validates this domain keystone v3 parameters
func (m *DomainKeystoneV3Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomainName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserDomain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainKeystoneV3Parameters) validateDomainName(formats strfmt.Registry) error {

	if err := validate.Required("domainName", "body", m.DomainName); err != nil {
		return err
	}

	return nil
}

func (m *DomainKeystoneV3Parameters) validateUserDomain(formats strfmt.Registry) error {

	if err := validate.Required("userDomain", "body", m.UserDomain); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainKeystoneV3Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainKeystoneV3Parameters) UnmarshalBinary(b []byte) error {
	var res DomainKeystoneV3Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}