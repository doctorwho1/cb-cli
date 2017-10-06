// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FlexUsageComponentJSON flex usage component Json
// swagger:model FlexUsageComponentJson

type FlexUsageComponentJSON struct {

	// component Id
	ComponentID string `json:"componentId,omitempty"`

	// instances
	Instances []*FlexUsageComponentInstanceJSON `json:"instances"`
}

/* polymorph FlexUsageComponentJson componentId false */

/* polymorph FlexUsageComponentJson instances false */

// Validate validates this flex usage component Json
func (m *FlexUsageComponentJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstances(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlexUsageComponentJSON) validateInstances(formats strfmt.Registry) error {

	if swag.IsZero(m.Instances) { // not required
		return nil
	}

	for i := 0; i < len(m.Instances); i++ {

		if swag.IsZero(m.Instances[i]) { // not required
			continue
		}

		if m.Instances[i] != nil {

			if err := m.Instances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlexUsageComponentJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlexUsageComponentJSON) UnmarshalBinary(b []byte) error {
	var res FlexUsageComponentJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
