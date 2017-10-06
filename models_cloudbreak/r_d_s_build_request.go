// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RDSBuildRequest r d s build request
// swagger:model RDSBuildRequest

type RDSBuildRequest struct {

	// requested cluster name
	// Required: true
	ClusterName *string `json:"clusterName"`

	// rds config request
	// Required: true
	RdsConfigRequest *RDSConfig `json:"rdsConfigRequest"`
}

/* polymorph RDSBuildRequest clusterName false */

/* polymorph RDSBuildRequest rdsConfigRequest false */

// Validate validates this r d s build request
func (m *RDSBuildRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRdsConfigRequest(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RDSBuildRequest) validateClusterName(formats strfmt.Registry) error {

	if err := validate.Required("clusterName", "body", m.ClusterName); err != nil {
		return err
	}

	return nil
}

func (m *RDSBuildRequest) validateRdsConfigRequest(formats strfmt.Registry) error {

	if err := validate.Required("rdsConfigRequest", "body", m.RdsConfigRequest); err != nil {
		return err
	}

	if m.RdsConfigRequest != nil {

		if err := m.RdsConfigRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rdsConfigRequest")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RDSBuildRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RDSBuildRequest) UnmarshalBinary(b []byte) error {
	var res RDSBuildRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
