// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WorkloadAnalyticsV4Response workload analytics v4 response
// swagger:model WorkloadAnalyticsV4Response
type WorkloadAnalyticsV4Response struct {

	// stack related telemetry component attributes
	Attributes *WorkloadAnalyticsAttributesHolder `json:"attributes,omitempty"`

	// telemetry - workload altus service (databus) endpoint url
	DatabusEndpoint string `json:"databusEndpoint,omitempty"`

	// enable telemetry component
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this workload analytics v4 response
func (m *WorkloadAnalyticsV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkloadAnalyticsV4Response) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkloadAnalyticsV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkloadAnalyticsV4Response) UnmarshalBinary(b []byte) error {
	var res WorkloadAnalyticsV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
