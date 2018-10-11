// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterExposedServiceResponse cluster exposed service response
// swagger:model ClusterExposedServiceResponse

type ClusterExposedServiceResponse struct {

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// knox service
	KnoxService string `json:"knoxService,omitempty"`

	// mode
	Mode string `json:"mode,omitempty"`

	// open
	Open *bool `json:"open,omitempty"`

	// service name
	ServiceName string `json:"serviceName,omitempty"`

	// service Url
	ServiceURL string `json:"serviceUrl,omitempty"`
}

/* polymorph ClusterExposedServiceResponse displayName false */

/* polymorph ClusterExposedServiceResponse knoxService false */

/* polymorph ClusterExposedServiceResponse mode false */

/* polymorph ClusterExposedServiceResponse open false */

/* polymorph ClusterExposedServiceResponse serviceName false */

/* polymorph ClusterExposedServiceResponse serviceUrl false */

// Validate validates this cluster exposed service response
func (m *ClusterExposedServiceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var clusterExposedServiceResponseTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SSO_PROVIDER","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterExposedServiceResponseTypeModePropEnum = append(clusterExposedServiceResponseTypeModePropEnum, v)
	}
}

const (
	// ClusterExposedServiceResponseModeSSOPROVIDER captures enum value "SSO_PROVIDER"
	ClusterExposedServiceResponseModeSSOPROVIDER string = "SSO_PROVIDER"
	// ClusterExposedServiceResponseModeNONE captures enum value "NONE"
	ClusterExposedServiceResponseModeNONE string = "NONE"
)

// prop value enum
func (m *ClusterExposedServiceResponse) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterExposedServiceResponseTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterExposedServiceResponse) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterExposedServiceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterExposedServiceResponse) UnmarshalBinary(b []byte) error {
	var res ClusterExposedServiceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}