// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AzureInstanceGroupV4Parameters azure instance group v4 parameters
// swagger:model AzureInstanceGroupV4Parameters
type AzureInstanceGroupV4Parameters struct {

	// availability set
	AvailabilitySet *AzureAvailabiltySetV4 `json:"availabilitySet,omitempty"`
}

// Validate validates this azure instance group v4 parameters
func (m *AzureInstanceGroupV4Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailabilitySet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureInstanceGroupV4Parameters) validateAvailabilitySet(formats strfmt.Registry) error {

	if swag.IsZero(m.AvailabilitySet) { // not required
		return nil
	}

	if m.AvailabilitySet != nil {
		if err := m.AvailabilitySet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("availabilitySet")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureInstanceGroupV4Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureInstanceGroupV4Parameters) UnmarshalBinary(b []byte) error {
	var res AzureInstanceGroupV4Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
