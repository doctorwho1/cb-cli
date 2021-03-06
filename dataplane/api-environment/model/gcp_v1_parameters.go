// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GcpV1Parameters gcp v1 parameters
// swagger:model GcpV1Parameters
type GcpV1Parameters struct {

	// json
	JSON *JSONV1Parameters `json:"json,omitempty"`

	// p12
	P12 *P12V1Parameters `json:"p12,omitempty"`
}

// Validate validates this gcp v1 parameters
func (m *GcpV1Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJSON(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateP12(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpV1Parameters) validateJSON(formats strfmt.Registry) error {

	if swag.IsZero(m.JSON) { // not required
		return nil
	}

	if m.JSON != nil {
		if err := m.JSON.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("json")
			}
			return err
		}
	}

	return nil
}

func (m *GcpV1Parameters) validateP12(formats strfmt.Registry) error {

	if swag.IsZero(m.P12) { // not required
		return nil
	}

	if m.P12 != nil {
		if err := m.P12.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("p12")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GcpV1Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpV1Parameters) UnmarshalBinary(b []byte) error {
	var res GcpV1Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
